package algorithm

type Bubble struct{}

// Max returns the larger of x or y.
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func (Bubble) Sort(data []int) {

	sorting := true

	// loop though data until we iterate without having to swap any items
	for sorting == true {

		// consider sorting complete until we encounter values we need to swap
		sorting = false

		// loop through the each item in the array but skip the _last_ element
		for i := 1; i < len(data); i++ {

			// if index-1 is < index, move on to the next index as they are in order.
			if data[i-1] < data[i] {
				continue
			}

			// swap the values of index and index - 1 and set the "sorting" flag
			temp := data[i]
			data[i] = data[i-1]
			data[i-1] = temp
			sorting = true
		}
	}
}