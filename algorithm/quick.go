package algorithm

type Quick struct{}

func (Quick) Sort(data []int) {
	sort(data, 0, len(data)-1)
}

func sort(data []int, start int, end int) {

	if len(data) < 1 || start >= end {
		return
	}

	left := start
	right := end
	pivot := data[end]

	for left <= right {

		for data[left] < pivot {
			left++
		}

		for data[right] > pivot {
			right--
		}

		if left <= right {
			data[left], data[right] = data[right], data[left]
			left++
			right--
		}

		sort(data, start, right)
		sort(data, left, end)
	}
}

func swap(a, b *int) {
	*a, *b = *b, *a
}
