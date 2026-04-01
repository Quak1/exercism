package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(c FodderCalculator, cows int) (float64, error) {
	totalFodder, err := c.FodderAmount(cows)
	if err != nil {
		return 0, err
	}

	factor, err := c.FatteningFactor()
	if err != nil {
		return 0, err
	}

	return totalFodder / float64(cows) * factor, nil
}

func ValidateInputAndDivideFood(c FodderCalculator, cows int) (float64, error) {
	if cows <= 0 {
		return 0, errors.New("invalid number of cows")
	}

	return DivideFood(c, cows)
}

type InvalidCowsError struct {
	cows int
	msg  string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.msg)
}

func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{
			cows: cows,
			msg:  "there are no negative cows",
		}
	} else if cows == 0 {
		return &InvalidCowsError{
			cows: cows,
			msg:  "no cows don't need food",
		}
	} else {
		return nil
	}
}
