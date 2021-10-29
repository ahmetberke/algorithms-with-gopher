package LinearSearch

func LinearSearchInt(sVal int, list []int) (bool, int) {

	for i, v := range list {
		if sVal == v {
			return true, i
		}
	}

	return false, -1
}

func LinearSearchStr(sVal string, list []string) (bool, int) {

	for i, v := range list {
		if sVal == v {
			return true, i
		}
	}

	return false, -1
}