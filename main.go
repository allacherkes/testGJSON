package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"log"
	"os"
	"time"
)

const json = `{"aaa":{"bbb":{"ccc":{"ddd":{"eee":{"fff":{"ggg":{"hhh":"end", "hhh1":"end1"}}}}}}}}`

func main() {
	measurePerformance(json, "*.*.*.*.*")

	// Read the entire contents of a file into a byte slice
	data, err := os.ReadFile("1kb.json")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	// Convert byte slice to string
	content := string(data)

	measurePerformance(content, "property5.nested5.nested5_1.nested5_1_1.nested5_1_1_1.nested5_1_1_1_1.nested5_1_1_1_1_1.nested5_1_1_1_1_1_1.nested5_1_1_1_1_1_1_1.nested5_1_1_1_1_1_1_1_1")
	measurePerformance(content, "property5.nested5.nested5_1.nested5_1_1.nested5_1_1_1.nested5_1_1_1_1.nested5_1_1_1_1_1")
	measurePerformance(content, "property5.nested5.nested5_1.nested5_1_1.nested5_1_1_1")
	measurePerformance(content, "*.*.*.*.*.*.*.*.*.*")
	measurePerformance(content, "*.*.*.*.*.*.*")
	measurePerformance(content, "*.*.*.*.*")

	asdata, err := os.ReadFile("smalljson.json") //size: 1.76 KB
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	// Convert byte slice to string
	acontent := string(asdata)
	measurePerformance(acontent, "d.results.3.url")
	measurePerformance(acontent, "d.results.#.url")

	abdata, err := os.ReadFile("bigjson.json") //size: 2.3 MB
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	// Convert byte slice to string
	acontent = string(abdata)
	measurePerformance(acontent, "d.results.3.url")
	measurePerformance(acontent, "d.results.#.url")
	fmt.Println("Hello")

}

func measurePerformance(input string, path string) {
	sizeInKB := float64(len(input)) / 1024
	fmt.Printf("Size of iput: %.2f KB\n", sizeInKB)
	fmt.Printf("The path is: %s \n", path)

	var sumDuration time.Duration
	for i := 0; i < 10; i++ {
		start := time.Now()
		if !gjson.Valid(input) {
			return
		}
		_ = gjson.Get(input, path)
		//fmt.Println(value.String())
		duration := time.Since(start)
		sumDuration = sumDuration + duration
	}

	fmt.Println("Function execution duration:", sumDuration.Nanoseconds()/10, "nanoseconds \n")
}
