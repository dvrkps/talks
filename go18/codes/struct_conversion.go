package main

func main() {
	example()
}

// START OMIT
func example() {
	type T1 struct {
		X int `json:"foo"`
	}
	type T2 struct {
		X int `json:"bar"`
	}
	var v1 T1
	var v2 T2
	v1 = T1(v2) // now legal // HL
	_ = v1
}

// END OMIT
