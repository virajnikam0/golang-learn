package goroutines

import (
	"fmt"
	"time"
)

// concurrent
// parellel

func F1() {

	fmt.Println("line-2")
	time.Sleep(3 * time.Second)

}

func F2() {

	fmt.Println("line-3")
}
