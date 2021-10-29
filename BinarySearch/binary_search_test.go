package BinarySearch

import "testing"

func TestBinarySearch(t *testing.T) {
	arr := []int{1,2,3,4,5,6,7,8,9,10}
	var x int

	// TEST CASE 1:
	x = 1
	isIn, ind := BinarySearch(x, arr)
	if !isIn {
		t.Errorf("test1: `val:%v` is in testList, isIn:%v must be TRUE", x, isIn)
	}
	if ind != 0 {
		t.Errorf("test1: `val:%v` index is %v but must be 0", x, ind)
	}

	// TEST CASE 2:
	x = 7
	isIn, ind = BinarySearch(x, arr)
	if !isIn {
		t.Errorf("test2: `val:%v` is in testList, isIn:%v must be TRUE", x, isIn)
	}
	if ind != 6 {
		t.Errorf("test2: `val:%v` index is %v but must be 6", x, ind)
	}

	// TEST CASE 3:
	x = -5
	isIn, ind = BinarySearch(x, arr)
	if isIn {
		t.Errorf("test3: `val:%v` is not in testList, isIn:%v must be FALSE", x, isIn)
	}

	// TEST CASE 4:
	x = 15
	isIn, ind = BinarySearch(x, arr)
	if isIn {
		t.Errorf("test4: `val:%v` is not in testList, isIn:%v must be FALSE", x, isIn)
	}
}