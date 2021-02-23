package variables

// START OMIT
func shadow() {
	var i int

	if i := 11; i > 0 {
		i++

		println(i) // 12 // HL
	}

	{
		i := 42
		i++

		println(i) // 43 // HL
	}

	println(i) // 1 // HL
}

// END OMIT
