package main

import (
	"fmt"
	"math/bits"
)

// START OMIT
func ExampleLeadingZeros() {
	fmt.Printf("%08b\n", 42)
	got := bits.LeadingZeros8(42) // HL
	fmt.Println(got)
	// Output:
	// 00101010
	// 2
}

// END OMIT
