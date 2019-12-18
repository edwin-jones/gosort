
package algorithm

type Bubble struct {}

// Max returns the larger of x or y.
func max(x, y int) int {
    if x < y {
        return y
    }
    return x
}

func (Bubble) Sort(data []int) () {
	sorting := true

	for sorting == true {

		sorting = false;
		size := max(len(data)-1, 0)

		for i := 0;  i<size; i++ {

			if (data[i+1] > data[i]){
				continue;
			}

			temp := data[i]
			data[i] = data[i+1]
			data[i+1] = temp
			sorting = true
			
		}	
	}
}