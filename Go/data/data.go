package data

import "fmt"

var d *Data

func init() {
	d = &Data{}
}

func GenOptions() string {
	return d.genOptions()
}

type Data struct {
}

// Format [{%q: %d, %q: %q}]
func (d *Data) genOptions() string {
	s := ""
	comma := ""
	for i := 0; i < 10; i++ {
		s += fmt.Sprintf("%s{%q: %d, %q: %q}", comma, "id", i, "text", fmt.Sprintf("option %d", i))
		comma = ","
	}
	s = fmt.Sprintf("[%s]", s)
	return s
}
