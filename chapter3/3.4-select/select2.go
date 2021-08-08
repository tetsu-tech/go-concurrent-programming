package main

func main() {
	c1 := make(chan interface{})
	c2 := make(chan interface{})

	var c1Count, c2Count int
	for i := 1000; i >= 0; i-- {
		select {
		case <-c1:
			c1Count++
		case <-c2:
			c2Count++
		}
	}
}
