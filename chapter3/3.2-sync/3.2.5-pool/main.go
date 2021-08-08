package main

import (
	"fmt"
	"sync"
)

func main() {
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance")
			return struct{}{}
		},
	}

	// Getが呼び出されたらプール内に使用可能なインスタンスがあるか確認し、あればそれを返す
	// ない場合はNewを呼び出し、インスタンスを作成
	myPool.Get()
	instance := myPool.Get()
	// インスタンスをpoolに戻す
	myPool.Put(instance)
	myPool.Get() // Putでpoolに戻されたものを再利用するのでNew関数は呼び出されない
}
