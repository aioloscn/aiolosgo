package main

import (
	"aiolosgo/retriever/mock"
	real2 "aiolosgo/retriever/real"
	"fmt"
	"time"
)

/**
 定义接口，实现者只要实现Get接口，并不需要Retriever名字一模一样，就像java中一个class implements一个接口
 */
type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

/**
 实现者只实现方法，所有不同的接口都有Get方法都能实现，根据入参和返回类型区分
 */
type FakeRetriever interface {
	Get(url string) string
}

const url = "http://www.aiolosxhx.com"

/**
 使用者
 */
func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url,
		map[string]string {"aaa": "111", "bbb": "222"})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string {
		"contents": "update Contents",
	})
	return s.Get(url)
}

func main() {
	fmt.Printf("%T %v\n", mock.Retriever{"www.aiolosxhx.com"}, mock.Retriever{"www.aiolosxhx.com"})
	var r Retriever
	r = &real2.Retriever1{
		UserAgent: "aiolos",
		TimeOut: time.Minute,
	}
	fmt.Printf("%T %v\n", r, r)
	//fmt.Println(download(r))
	//var str real2.Retriever1
	//fmt.Println(str.Get(url))					// ???????? 可以直接调用实现者创建的Get方法
	//str := real2.Retriever1{}.Get(url)			// ???????? 可以直接调用实现者创建的Get方法 Get参数为指针接受者时该行不可用↑
	//fmt.Println(str)

	// Type assertion
	realRetriever := r.(*real2.Retriever1)
	fmt.Println(realRetriever.TimeOut)

	retriever := &mock.Retriever{"Contents"}
	fmt.Printf("%T %v\n", retriever, retriever)
	fmt.Println(session(retriever))
}
