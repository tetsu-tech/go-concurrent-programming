package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	type Button struct {
		Clicked *sync.Cond
	}
	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}

	subscribe := func(c *sync.Cond, fn func()) {
		go func() {
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		// sleepを入れたら動いたwaitGroupを使っていたのはgorutineがスケジューリングされるのを待つため？な気がする
		time.Sleep(1 * time.Second)
	}

	var clieckRegistered sync.WaitGroup
	clieckRegistered.Add(3)
	subscribe(button.Clicked, func() {
		fmt.Println("Maximazing window.")
		clieckRegistered.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Displaying annoying dialog box!")
		clieckRegistered.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Mouse clicked.")
		clieckRegistered.Done()
	})

	button.Clicked.Broadcast()
	clieckRegistered.Wait()
}
