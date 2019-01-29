package main

import (
	"aiolosgo/crawler/engine"
	"aiolosgo/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}