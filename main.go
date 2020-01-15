package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/edwin-jones/gosort/algorithm"
)

type Sorter interface {
	Sort(data []int)
}

func main() {

	var selectedAlgorithm string
	var rawData string
	var sorter Sorter

	flag.StringVar(&selectedAlgorithm, "algorithm", "bubble", "the sorting algorithm to use")
	flag.StringVar(&selectedAlgorithm, "a", "bubble", "the sorting algorithm to use (shorthand)")
	flag.StringVar(&rawData, "data", "1,2,3", "the data to sort, represented as a list of integers delimited by a comma")
	flag.StringVar(&rawData, "d", "1,2,3", "the data to sort, represented as a list of integers delimited by a comma (shorthand)")
	flag.Parse()

	switch selectedAlgorithm {
	case "bubble":
		sorter = algorithm.Bubble{}
	case "insertion":
		sorter = algorithm.Insertion{}
	case "quick":
		sorter = algorithm.Quick{}
	default:
		log.Fatal("Invalid sorting algorithm selected")
		return
	}

	data := []int{}
	numbers := strings.Split(rawData, ",")

	for _, number := range numbers {
		thing, _ := strconv.Atoi(number)
		data = append(data, thing)
	}

	sorter.Sort(data)

	fmt.Printf("sorted results: %v\n", data)
}
