package main

import (
	"fmt"
	"sync"
)

// countをラップ
type Counter struct {
	count int
	lock  sync.Mutex
}

func (c *Counter) Increment() {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.count++
	fmt.Printf("incrementing: %d\n", c.count)
}

func (c *Counter) Decrement() {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.count--
	fmt.Printf("decrementing: %d\n", c.count)
}

func main() {
	counter := Counter{count: 0, lock: sync.Mutex{}}

	var arithmetic sync.WaitGroup
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			counter.Increment()
		}()
	}

	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			counter.Decrement()
		}()
	}

	arithmetic.Wait()

	fmt.Println("Arithmetic complete")
}
