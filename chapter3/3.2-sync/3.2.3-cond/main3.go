package main

import (
	"fmt"
	"sync"
)

func main() {
	type Button struct {
		Clicked *sync.Cond
	}
	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}

	subscribe := func(c *sync.Cond, fn func()) {
		// TODO: ここでwaitgroupを使うのがよくわからん、すぐDoneにしてるからいらないかな？とか思ったけどなくすとDeadlockになる
		var gorutineRunning sync.WaitGroup
		gorutineRunning.Add(1)
		go func() {
			gorutineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		gorutineRunning.Wait()
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
