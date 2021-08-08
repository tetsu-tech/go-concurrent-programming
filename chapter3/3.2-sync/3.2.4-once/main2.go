package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	increment := func() { count++ }
	decrement := func() { count-- }

	var once sync.Once
	once.Do(increment)
	once.Do(decrement)

	// 出力は「Count is 1」になる
	// sync.OnceはDoが呼びされた回数だけ数えていてDoに渡された一意な関数が呼び出された回数を数えているわけではない
	fmt.Printf("Count is %d\n", count)
}
