package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"
)

// Benchmark the performance of a map with int keys
func benchmarkIntMap(elements int) {
	fmt.Printf("Benchmarking map[int]int with %d elements\n", elements)
	// Create a large map
	m := make(map[int]int, elements)
	for i := 0; i < elements; i++ {
		m[i] = i
	}

	// Measure lookup performance
	start := time.Now()
	for i := 0; i < elements; i++ {
		_ = m[i%elements]
	}
	fmt.Printf("Lookup time: %v\n", time.Since(start))

	// Measure insertion performance
	start = time.Now()
	for i := elements; i < elements*2; i++ {
		m[i] = i
	}
	fmt.Printf("Insertion time: %v\n", time.Since(start))

	// Measure deletion performance
	start = time.Now()
	for i := 0; i < elements; i++ {
		delete(m, i)
	}
	fmt.Printf("Deletion time: %v\n", time.Since(start))
}

// Benchmark the performance of a map with string keys
func benchmarkStringMap(elements int) {
	fmt.Printf("Benchmarking map[string]string with %d elements\n", elements)
	// Create a large map
	m := make(map[string]string, elements)
	for i := 0; i < elements; i++ {
		key := strconv.Itoa(i)
		m[key] = key
	}

	// Measure lookup performance
	start := time.Now()
	for i := 0; i < elements; i++ {
		key := strconv.Itoa(i % elements)
		_ = m[key]
	}
	fmt.Printf("Lookup time: %v\n", time.Since(start))

	// Measure insertion performance
	start = time.Now()
	for i := elements; i < elements*2; i++ {
		key := strconv.Itoa(i)
		m[key] = key
	}
	fmt.Printf("Insertion time: %v\n", time.Since(start))

	// Measure deletion performance
	start = time.Now()
	for i := 0; i < elements; i++ {
		key := strconv.Itoa(i)
		delete(m, key)
	}
	fmt.Printf("Deletion time: %v\n", time.Since(start))
}

func main() {
	elements := flag.Int("n", 10_00_000, "Number of elements in the map")
	mType := flag.String("type", "int", "Type of data structure to use, int or string")
	flag.Parse()
	if *mType == "int" {
		benchmarkIntMap(*elements)
	} else if *mType == "string" {
		benchmarkStringMap(*elements)
	}
}
