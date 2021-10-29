package BinarySearch

func BinarySearch(sVal int, list []int) (bool, int) {

	var smallInd int = 0
	var bigInd int = len(list) - 1
	var midInd int

	for bigInd >= smallInd {
		midInd = (bigInd + smallInd) / 2

		if sVal == list[midInd] {
			return true, midInd
		}

		if sVal < list[midInd] {
			bigInd = midInd - 1
		}

		if sVal > list[midInd] {
			smallInd = midInd + 1
		}
	}

	return false, 0
}