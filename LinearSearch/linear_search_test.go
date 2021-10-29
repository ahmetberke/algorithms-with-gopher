package LinearSearch

import "testing"

func TestLinearSearchInt(t *testing.T) {
	// Test Case 1:
	arr := []int{76,2,33,12,5,1,3,5,2,6}
	x := 5
	isIn, ind := LinearSearchInt(x, arr)
	if !isIn {
		t.Errorf("%v is in arr, isIn must be True", x)
	}
	if ind != 4 {
		t.Errorf("index of %v is %v, must be 4", x, ind)
	}

	// Test Case 2:
	x = -1245
	isIn, ind = LinearSearchInt(x, arr)
	if isIn {
		t.Errorf("%v is not in arr, isIn must be False", x)
	}

}

func TestLinearSearchString(t *testing.T) {
	// Test Case 1:
	arr := []string{"a","b","","asd","asfasgasg","test","hello","world"}
	x := "test"
	isIn, ind := LinearSearchStr(x, arr)
	if !isIn {
		t.Errorf("%v is in arr, isIn must be True", x)
	}
	if ind != 5 {
		t.Errorf("index of %v is %v, must be 5", x, ind)
	}

	// Test Case 2:
	x = "blabla"
	isIn, ind = LinearSearchStr(x, arr)
	if isIn {
		t.Errorf("%v is not in arr, isIn must be False", x)
	}
}