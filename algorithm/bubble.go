package algorithm

type Bubble struct{}

func (Bubble) Sort(data []int) {

	sorting := true

	// loop though data until we iterate without having to swap any items
	for sorting {

		// consider sorting complete until we encounter values we need to swap
		sorting = false

		// loop through the each item in the set starting from the second element
		for i := 1; i < len(data); i++ {

			// if index-1 is <= index, move on as they are in order already
			if data[i-1] <= data[i] {
				continue
			}

			// swap the values of index and index - 1, then set the "sorting" flag
			temp := data[i]
			data[i] = data[i-1]
			data[i-1] = temp
			sorting = true
		}
	}
}
