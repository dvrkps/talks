package variables

import "fmt"

// START OMIT
func shadow() error {
	err := doOne()
	if err != nil {
		return fmt.Errorf("one: %v", err)
	}

	if err := doTwo(); err != nil {
		return fmt.Errorf("two: %v", err)
	}

	if err := doThree(); err != nil {
		return fmt.Errorf("three: %v", err)
	}

	return nil
}

// END OMIT
