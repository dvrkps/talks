package variables

// START OMIT
func isValid(v int) bool {
	var ( // HL
		min = 0   // HL
		max = 100 // HL
	) // HL

	if v < min {
		return false
	}

	if v > max {
		return false
	}

	return true
}

// END OMIT
