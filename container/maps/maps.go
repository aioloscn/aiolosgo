package main

import "fmt"

func main() {

	m := map[string]string {
		"a": "11",
		"b": "22",
		"c": "33",
	}
	m1 := make(map[string]int)	// m1 == empty map
	var m2 map[string]float32	// m2 == nil
	fmt.Println(m, m1, m2)
	for k, v := range m {
		fmt.Println(k, v)
	}
	if ret, ok := m2["b"]; ok {
		fmt.Println(ret, ok)
	}
	delete(m, "c")
	fmt.Println(m)
	m["a"] = "aiolos"
	fmt.Println(m)
}
