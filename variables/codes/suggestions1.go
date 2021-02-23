package variables

// START OMIT
func zero(isCelsius bool) float64 {
	var z float64 // HL
	if isCelsius {
		z = -273.15
	}

	return z
}

// END OMIT
