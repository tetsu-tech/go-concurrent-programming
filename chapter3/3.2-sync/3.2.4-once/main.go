package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	increment := func() {
		count++
	}

	var once sync.Once

	var increments sync.WaitGroup
	increments.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer increments.Done()
			// Doの引数に渡された関数が一度しか呼び出されないようにする
			once.Do(increment)
		}()
	}

	increments.Wait()
	fmt.Printf("Count is %d\n", count)
}
