package algorithm_test

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/edwin-jones/gosort/algorithm"
)

func generateLargeSetOfRandomNumbers() []int {
	len := 1000
	data := make([]int, len)
	for i := 0; i <= len-1; i++ {
		data[i] = rand.Intn(len)
	}
	return data
}

var testData = [][]int{
	{0, 1, 2, 3},
	{3, 2, 1, 0},
	{4, 5, 6, 7, 1, 3, 10},
	{0, -2, 999, 4, 1, 1, 1},
	{1111, 10, 22, 3, 4, 5, 2, 5, 5, 5, 33, -100, -232, 56, 44},
}

var bubble = algorithm.Bubble{}
var insertion = algorithm.Insertion{}
var quick = algorithm.Quick{}

type Sorter interface {
	Sort(data []int)
}

var algorithms = []Sorter{bubble, insertion, quick}

func TestSortingAlgorithms(t *testing.T) {
	for _, sorter := range algorithms {
		for _, input := range testData {

			output := make([]int, len(input))
			copy(output, input)

			sorter.Sort(output)

			if !sort.IntsAreSorted(output) || len(output) != len(input) {
				t.Errorf("Sorting failed with %T, input: %v, output: %v.", sorter, input, output)
			}
		}
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bubble.Sort(generateLargeSetOfRandomNumbers())
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertion.Sort(generateLargeSetOfRandomNumbers())
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		quick.Sort(generateLargeSetOfRandomNumbers())
	}
}
