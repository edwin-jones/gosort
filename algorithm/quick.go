package algorithm

type Quick struct{}

func (Quick) Sort(data []int) {
	sort(data, 0, len(data)-1)
}

func partition(data []int, start int, end int) int {

	pivot := end
	pivotValue := data[pivot]

	var i int = start
	var j int = end

	for {
		for data[i] < pivotValue {
			i++
		}
		for data[j] > pivotValue {
			j--
		}

		if data[i] <= data[j] {
			swap(&data[i], &data[j])
		}

		i++
		j--

		if i != j {
			continue
		}

		pivot = i
		return pivot
	}
}

func sort(data []int, start int, end int) {
	if end < start {
		pivot := partition(data, start, end)
		sort(data, pivot+1, end)
		sort(data, start, pivot-1)
	}
}

func swap(a, b *int) {
	*a, *b = *b, *a
}
