package str

// Contains ...
func Contains(slices []string, comparizon string) bool {
	for _, a := range slices {
		if a == comparizon {
			return true
		}
	}

	return false
}

// ContainInts ...
func ContainInts(slices []int, comparizon int) bool {
	for _, a := range slices {
		if a == comparizon {
			return true
		}
	}

	return false
}
