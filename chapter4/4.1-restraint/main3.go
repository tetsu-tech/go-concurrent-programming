package main

import (
	"bytes"
	"fmt"
	"sync"
)

// 並行安全ではないデータ構造を使った拘束の例
func main() {
	printData := func(wg *sync.WaitGroup, data []byte) {
		defer wg.Done()

		var buff bytes.Buffer
		for _, b := range data {
			fmt.Fprintf(&buff, "%c", b)
		}
		fmt.Println(buff.String())
	}

	var wg sync.WaitGroup
	wg.Add(2)
	// TODO: 配列は並行安全ではないということかな？
	data := []byte("golang")
	// スライスを渡すコードでデータへのアクセスを制限してる
	go printData(&wg, data[:3])
	go printData(&wg, data[3:])

	wg.Wait()
}
