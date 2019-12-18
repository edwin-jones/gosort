package algorithm

type Insertion struct{}

func (Insertion) Sort(data []int) {

	// loop through every item in the set
	for i := 1; i < len(data); i++ {

		// set the default insertion point as the last known sorted value (index 0)
		j := i - 1

		// keep track of the value at the current index while we swap things around
		current := data[i]

		// loop backwards from the current index to the start of the set, looking for the insertion point of the swap
		// We move each value one to the right so we have space to insert when we need to.
		for j >= 0 && data[j] > current {
			data[j+1] = data[j]
			j--
		}

		// insert the swapper value at the insertion point
		data[j+1] = current
	}
}
