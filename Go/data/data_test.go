package data

import (
	"fmt"
	"testing"
)

func TestData_genOptions(t *testing.T) {
	r := GenOptions()
	fmt.Println(r)
}
