package mydata

import (
	"testing"
)

func Test_Sort(t *testing.T) {
	tests := []struct{
		in []int
		expected []int
	}{
		{[]int{1,2}, []int{1,2}},
		{[]int{2,1}, []int{1,2}},
		{[]int{1,2,3}, []int{1,2,3}},
		{[]int{1,3,2}, []int{1,2,3}},
		{[]int{1,3,2,5,4}, []int{1,2,3,4,5}},
		{[]int{5,4,3,2,1}, []int{1,2,3,4,5}},
		{[]int{1,5,3,2,4}, []int{1,2,3,4,5}},
		{[]int{3,5,1,3,2,4}, []int{1,2,3,3,4,5}},
	}

	// t.Run("bubble", func(t *testing.T) {
	// 	for _, ts := range tests {
	// 		Bubble(ts.in)
	// 		assert.Equal(t, ts.expected, ts.in)
	// 	}
	// })

	// t.Run("insert", func(t *testing.T) {
	// 	for _, ts := range tests {
	// 		Insert(ts.in)
	// 		assert.Equal(t, ts.expected, ts.in)
	// 	}
	// })

	// t.Run("merge", func(t *testing.T) {
	// 	for _, ts := range tests {
	// 		// fmt.Printf("%+v\n", ts.in)
	// 		actual := MergeSort(ts.in)
	// 		assert.Equal(t, ts.expected, actual)
	// 	}
	// })

	// t.Run("quick", func(t *testing.T) {
	// 	for _, ts := range tests {
	// 		// fmt.Printf("%+v\n", ts.in)
	// 		QuickSort(ts.in)
	// 		assert.Equal(t, ts.expected, ts.in)
	// 	}
	// })
	t.Run("heap", func(t *testing.T) {
		for _, ts := range tests {
			HeapSort(ts.in)
		}
	})
}