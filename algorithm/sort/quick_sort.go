// Recording quick sort algorithm.
// Though time complexity: n(lgn), but not need other memory
package sort

// QuickSort is sort data by quick sort of algorithm.
// Thinking:
// 	The data set is divided into two parts, the first part of the data
// 	than the second part of the data smaller.
// 	Then, follow these methods to quickly sort the two parts.
// 	Usually in the data set to select a partition, put a smaller than the
// 	partition in front of it, put the larger than partition behind it.
func QuickSort(data []int) bool {
	if len(data) == 0 {
		return false
	}
	recurSort(data, 0, len(data)-1)
	return true
}

func recurSort(data []int, start, end int) {
	if start < end {
		pivot := partition(data, start, end)
		recurSort(data, start, pivot-1)
		recurSort(data, pivot+1, end)
	}
}

func partition(data []int, start, end int) int {
	for start < end {
		for start < end && data[start] <= data[end] {
			end--
		}
		if start < end {
			data[start], data[end] = data[end], data[start]
			start++
		}

		for start < end && data[start] <= data[end] {
			start++
		}
		if start < end {
			data[start], data[end] = data[end], data[start]
			end--
		}
	}
	return start
}
