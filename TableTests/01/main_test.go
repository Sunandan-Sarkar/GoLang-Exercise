package main

import "testing"

func TestMyTot(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}
	tests := []test{
		test{[]int{44, 44}, 88},
		test{[]int{21, 44}, 65},
		test{[]int{-2, 4}, 2},
		test{[]int{-11, 66, 0}, 55},
	}

	for _, v := range tests {
		x := myTot(v.data...)
		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		}
	}

}
