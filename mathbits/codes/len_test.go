package main

import (
	"fmt"
	"math/bits"
)

// START OMIT
func ExampleLen() {
	fmt.Printf("%b\n", 42)
	got := bits.Len(42) // HL
	fmt.Println(got)
	// Output:
	// 101010
	// 6
}

// END OMIT
