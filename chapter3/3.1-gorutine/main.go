package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// MEMO: 全てのgorutineで "good day" が表示される
			fmt.Println(salutation)
		}()
	}
	wg.Wait()
}
