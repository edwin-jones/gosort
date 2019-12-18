package main

import (
	"flag"
	"fmt"
	"github.com/edwin-jones/gosort/algorithm"
	"log"
	"strconv"
	"strings"
)

type Sorter interface {
	Sort(data []int)
}

func main() {

	var selectedAlgorithm string
	var rawData string
	var sorter Sorter

	flag.StringVar(&selectedAlgorithm, "algorithm", "bubble", "the sorting algorithm to use")
	flag.StringVar(&rawData, "data", "1,2,3", "the data to sort, represented as a list of integers delimited by a comma")
	flag.Parse()

	switch selectedAlgorithm {
	case "bubble":
		sorter = algorithm.Bubble{}
	case "insertion":
		sorter = algorithm.Insertion{}
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
