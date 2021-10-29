package BinarySearch

import "testing"

func TestBinarySearch(t *testing.T) {
	testList := []int{1,2,3,4,5,6,7,8,9,10}
	var testVal int

	// TEST CASE 1:
	testVal = 1
	isIn, ind := BinarySearch(testVal, testList)
	if !isIn {
		t.Errorf("test1: `val:%v` is in testList, isIn:%v must be TRUE", testVal, isIn)
	}
	if ind != 0 {
		t.Errorf("test1: `val:%v` index is %v but must be 0", testVal, ind)
	}

	// TEST CASE 2:
	testVal = 7
	isIn, ind = BinarySearch(testVal, testList)
	if !isIn {
		t.Errorf("test2: `val:%v` is in testList, isIn:%v must be TRUE", testVal, isIn)
	}
	if ind != 6 {
		t.Errorf("test2: `val:%v` index is %v but must be 6", testVal, ind)
	}

	// TEST CASE 3:
	testVal = -5
	isIn, ind = BinarySearch(testVal, testList)
	if isIn {
		t.Errorf("test3: `val:%v` is not in testList, isIn:%v must be FALSE", testVal, isIn)
	}

	// TEST CASE 4:
	testVal = 15
	isIn, ind = BinarySearch(testVal, testList)
	if isIn {
		t.Errorf("test4: `val:%v` is not in testList, isIn:%v must be FALSE", testVal, isIn)
	}




}