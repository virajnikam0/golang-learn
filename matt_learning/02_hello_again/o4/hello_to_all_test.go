package o4

import (
	"fmt"
	"os"
	"testing"
)

func testHelloToAll() {
	names := os.Args

	if len(names) > 0 {
		for _, name := range names {
			fmt.Println("Hello ", name)
		}
	}
}

func TestZ1(t *testing.T) {
	arr := []struct {
		name string
		age  int
	}{
		{"John", 12},
		{"nice", 45},
		{"Hey", 55},
	}

	for _, data := range arr {
		fmt.Println(data)
	}
}
