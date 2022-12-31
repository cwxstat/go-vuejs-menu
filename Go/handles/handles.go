package handles

import (
	"fmt"
	"math/rand"
	"time"

	"encoding/json"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

type StatusData struct {
	Name string `json:"name"`
	P    int    `json:"p"`
	Code string `json:"code"`
	N    string `json:"n"`
}

type Status struct {
	M []StatusData `json:"data"`
	T string       `json:"type"`
}

func GenStatus() []byte {

	m := []StatusData{}

	for i := 0; i < 9; i++ {
		randomString := fmt.Sprintf("%f", rand.Float32()*100)
		randomN := fmt.Sprintf("n %f", rand.Float32()*100)
		m = append(m, StatusData{Name: "string", P: rand.Intn(100), Code: randomString, N: randomN})
	}
	n := &Status{M: m, T: "status"}
	resultJson, _ := json.Marshal(n)

	return resultJson

}

// https://play.golang.org/p/NjWxHFpgG6T
func GetStatus(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		// Read
		msg := ""
		err := websocket.Message.Receive(ws, &msg)
		if err != nil {
			c.Logger().Error(err)
		}
		fmt.Printf("%s\n", msg)
		for {

			// Write
			result := GenStatus()
			err = websocket.Message.Send(ws, string(result))
			if err != nil {
				c.Logger().Error(err)
				return

			}
			fmt.Printf("\n\nStatus out: %s\n", result)

			time.Sleep(2 * time.Second)
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

func GetType(f float32) string {
	if f > 30 {
		return "buy"
	}
	return "sell"
}

func CheckConn(ws *websocket.Conn) (<-chan string, <-chan bool) {
	// Adding in receive here..
	c := make(chan string)
	done := make(chan bool)
	go func() {
		for {
			msg := ""
			err := websocket.Message.Receive(ws, &msg)
			if err != nil {

				fmt.Printf("Error: %s\n", err)
				fmt.Printf("\n\n...closing connection\n")
				close(done)
				return

			}
			c <- msg
		}
	}()
	return c, done
}

func Hello(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		// Read
		msg := ""
		err := websocket.Message.Receive(ws, &msg)
		if err != nil {
			c.Logger().Error(err)
		}
		fmt.Printf("%s\n", msg)

		cM, done := CheckConn(ws)
		for {

			// Write
			result := fmt.Sprintf("{%q: [{%q: %4.2f}, {%q: %4.2f}], %q: %q}", "data", "p", rand.Float32()*100, "p", rand.Float32()*100, "type", GetType(rand.Float32()*100))
			err := websocket.Message.Send(ws, result)
			if err != nil {
				c.Logger().Error(err)
				fmt.Printf("Error: %s\n", err)
				fmt.Printf("\n\n...closing connection\n")
				return

			}

			select {
			case msg := <-cM:
				fmt.Println("Send operation on cM", msg)
			case <-done:
				fmt.Println("Closing connection!!!")
				return
			default:

			}

			fmt.Printf("out: %s\n", result)
			time.Sleep(2 * time.Second)
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}
