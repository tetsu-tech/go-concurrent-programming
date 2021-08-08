package main

import (
	"fmt"
	"sync"
)

type Counter {
	count int
	lock sync.Mutex
}

func(c *Counter) increment() {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.count++
}

func (c *Counter) decrement() {
	c.lock.Lock()
	defer c.lock.Unlock()
	count--
}


func main() {
	// mutexを使って変数のアクセスを同期する
	var count int
	var lock sync.Mutex

	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("incrementing: %d\n", count)
	}

	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("decrementing: %d\n", count)
	}

	var arithmetic sync.WaitGroup
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}

	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}

	arithmetic.Wait()

	fmt.Println("Arithmetic complete")
}
