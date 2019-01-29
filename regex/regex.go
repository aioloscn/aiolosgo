package main

import (
	"fmt"
	"regexp"
)

const text = `
email0 is asjfd@163.com
email1 is asj232@ldj.com.cn
email2 is 456789@qq.com
`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	submatch := re.FindAllStringSubmatch(text, -1)
	for _, s := range submatch {
		fmt.Println(s)
	}
}
