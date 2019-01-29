package main

import (
	"aiolosgo/pipeline"
	"fmt"
	"os"
)

func main() {
	//p := pipeline.Merge(pipeline.InMemSort(pipeline.ArraySource(3, 2, 6, 7, 4)), pipeline.InMemSort(pipeline.ArraySource(9, 13, 7, 5, 0, 20, 8)))
	//for v := range p {
	//	fmt.Println(v)
	//}

	file, err := os.Create("small.in")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipeline.RandomSource(50)
	pipeline.WriterSink(file, p)

	file, err = os.Open("small.in")
	if err != nil {
		panic(err)
	}

	source := pipeline.ReaderSource(file)
	p = pipeline.Merge(pipeline.InMemSort(source), pipeline.InMemSort(source))
	for v := range p {
		fmt.Println(v)
	}
}
