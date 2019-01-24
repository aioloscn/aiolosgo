package main

import (
	real2 "aiolosgo/retriever/real"
	"fmt"
)

/**
 定义接口，实现者只要实现Get接口，并不需要Retriever名字一模一样，就像java中一个class implements一个接口
 */
type Retriever interface {
	Get(url string) string
}

type FakeRetriever interface {
	Get(url string) string
}

/**
 使用者
 */
func download(r Retriever) string {
	return r.Get("http://www.aiolosxhx.com")
}

func main() {
	//fmt.Println(download(mock.Retriever{"www.aiolosxhx.com"}))
	var r Retriever
	r = real2.Retriever1{}
	fmt.Println(download(r))
	/*str := real2.Retriever1{}.Get("http://www.aiolosxhx.com")			???????? 可以直接调用使用者创建的Get方法
	fmt.Println(str)*/
}
