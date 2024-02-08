package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"log"
	"os"
)

const json = `{"aaa":{"bbb":{"ccc":{"ddd":{"eee":{"fff":{"ggg":{"hhh":"end", "hhh1":"end1"}}}}}}}}`

func main() {
	// Read the entire contents of a file into a byte slice
	data, err := os.ReadFile("bigjson.json") //size: 2.3 MB
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	// Convert byte slice to string
	content := string(data)

	value := gjson.Get(content, "d.results.125.url")
	fmt.Println(value.String())

	arrValue := gjson.Get(content, "d.results.#.url")
	fmt.Println(arrValue.String())

	fmt.Println("Hello, 世界")
}
