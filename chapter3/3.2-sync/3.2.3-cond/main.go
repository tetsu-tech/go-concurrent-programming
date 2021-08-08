package main

import "sync"

func main() {
	// ここは本に書いてなかったけど追記、よくわかってないw
	condition := false
	conditionTrue := func() bool {
		return condition
	}

	// NewCondはsync.Lockerを満たすものを引数として受け取る
	c := sync.NewCond(&sync.Mutex{})
	c.L.Lock()
	for conditionTrue() == false {
		c.Wait()
	}
	c.L.Unlock()
}
