package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				//fmt.Printf("aiolosxhx goroutine %d\n", i)			Printf  io操作是抢占式的
				a[i]++
				runtime.Gosched()		// 不交出控制权这里会死循环
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
