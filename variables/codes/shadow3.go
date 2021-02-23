package variables

import "fmt"

// START OMIT
func shadow() error {
	err := doOne()
	if err != nil {
		return fmt.Errorf("one: %v", err)
	}

	err = doTwo()
	if err != nil {
		return fmt.Errorf("two: %v", err)
	}

	err = doThree()
	if err != nil {
		return fmt.Errorf("three: %v", err)
	}

	return nil
}

// END OMIT
