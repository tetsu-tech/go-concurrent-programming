package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var stdouBuff bytes.Buffer
	defer stdouBuff.WriteTo(os.Stdout)

	intStream := make(chan int, 4)
	go func() {
		defer close(intStream)
		defer fmt.Fprintln(&stdouBuff, "Producer Done.")
		for i := 0; i < 5; i++ {
			fmt.Fprintf(&stdouBuff, "Sending %d\n", i)
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Fprintf(&stdouBuff, "Received %v.\n", integer)
	}
}
