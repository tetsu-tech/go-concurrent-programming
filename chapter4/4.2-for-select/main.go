package main

import "fmt"

func main() {
	sendToChannel := func(stringStream chan string, data []string) {
		defer close(stringStream)
		for _, s := range data {
			select {
			case stringStream <- s:
			}
		}
	}

	data := []string{"hello", "world"}
	stringStream := make(chan string)

	go sendToChannel(stringStream, data)

	for string := range stringStream {
		fmt.Println("str", string)
	}
}
