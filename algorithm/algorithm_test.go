package algorithm_test

import (
	"sort"
	"testing"

	"github.com/edwin-jones/gosort/algorithm"
)

var testData = [][]int{
	{0, 1, 2, 3},
	{3, 2, 1, 0},
	{4, 5, 6, 7, 1, 3, 10},
	{0, -2, 999, 4, 1, 1, 1},
	{1111, 10, 22, 3, 4, 5, 2, 5, 5, 5, 33, -100, -232, 56, 44},
}

var bubble = algorithm.Bubble{}
var insertion = algorithm.Insertion{}

func TestBubbleSort(t *testing.T) {
	for _, input := range testData {

		output := []int{}
		copy(output, input)

		bubble.Sort(output)

		if !sort.IntsAreSorted(output) {
			t.Errorf("Sorting failed, input: %v, output: %v.", input, output)
		}
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, input := range testData {

			output := []int{}
			copy(output, input)

			bubble.Sort(output)
		}
	}
}

func TestInsertionSort(t *testing.T) {
	for _, input := range testData {

		output := []int{}
		copy(output, input)

		insertion.Sort(output)

		if !sort.IntsAreSorted(output) {
			t.Errorf("Sorting failed, input: %v, output: %v.", input, output)
		}
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, input := range testData {

			output := []int{}
			copy(output, input)

			insertion.Sort(output)
		}
	}
}
