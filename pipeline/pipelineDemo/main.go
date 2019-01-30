package main

import (
	"aiolosgo/pipeline"
	"bufio"
	"fmt"
	"os"
)

func main() {
	//p := pipeline.Merge(pipeline.InMemSort(pipeline.ArraySource(3, 2, 6, 7, 4)), pipeline.InMemSort(pipeline.ArraySource(9, 13, 7, 5, 0, 20, 8)))
	//for v := range p {
	//	fmt.Println(v)
	//}

	const filename = "small.in"
	const count = 64

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipeline.RandomSource(count)
	writer := bufio.NewWriter(file)
	pipeline.WriterSink(writer, p)
	writer.Flush()

	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}

	p = pipeline.ReaderSource(bufio.NewReader(file), -1)
	for v := range p {
		fmt.Println(v)
	}
}
