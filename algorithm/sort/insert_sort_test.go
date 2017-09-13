package sort

import "testing"

func testInsertSort(t *testing.T) {

	testdata := []int{23, 21, 76, 16, 52, 43}
	if result := InsertSort(testdata); !result {
		t.Fatal("insert sort failed.")
	}
}
