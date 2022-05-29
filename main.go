package main

import (
	"fmt"
	"time"
)

func main() {
	go printNumbers(1)
	go printNumbers(10)
	go printNumbers(100)

	time.Sleep(time.Second)

}

func printNumbers(start int32) {
	for i := start; i <= start+3; i++ {
		fmt.Println(i)
	}
}
