package sort_t

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{name: "empty", in: []int{}, want: []int{}},
		{name: "single", in: []int{1}, want: []int{1}},
		{name: "ordered", in: []int{1, 2, 3, 4}, want: []int{1, 2, 3, 4}},
		{name: "reverse", in: []int{5, 4, 3, 2, 1}, want: []int{1, 2, 3, 4, 5}},
		{name: "duplicate", in: []int{4, 2, 4, 1, 2}, want: []int{1, 2, 2, 4, 4}},
		{name: "negative", in: []int{3, -1, 0, -5, 2}, want: []int{-5, -1, 0, 2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeSort(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("MergeSort(%v) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestMergeSortDoesNotChangeInput(t *testing.T) {
	nums := []int{3, 1, 2}
	wantNums := []int{3, 1, 2}

	_ = MergeSort(nums)

	if !reflect.DeepEqual(nums, wantNums) {
		t.Fatalf("MergeSort changed input: got %v, want %v", nums, wantNums)
	}
}
