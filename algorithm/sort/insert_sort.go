// Recording insert sort algorithm.
// time complexity: n^2, but not need other memory
package sort

func InsertSort(data []int) bool {
	if len(data) == 0 {
		return false
	}

	for j := 0; j < len(data); j++ {
		if data[j] < data[j-1] {
			i := j - 1
			cur := data[j]
			for i >= 0 && data[i] > cur {
				data[i+1] = data[i]
				i--
			}
			data[i+1] = cur
		}
	}
	return true
}
