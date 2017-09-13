package sort

import "testing"

func TestQuickSort(t *testing.T) {
	testdata := []int{24, 52, 11, 94, 28, 36, 14, 80}
	if result := QuickSort(testdata); !result {
		t.Fatal("quick sort failed.")
	}
}
