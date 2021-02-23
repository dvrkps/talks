package variables

import "os"

// START OMIT
func short(infile, outfile string) {
	f, err := os.Open(infile)

	// ...

	f, err := os.Create(outfile) // no new variables on left side of := // HL

	// ...
}

// END OMIT
