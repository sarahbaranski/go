package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	message string
	details int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("%s%d cows", e.message, e.details)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	negative := "negative fodder"
	division := "division by zero"
	amount, err := weightFodder.FodderAmount()
	if err == nil && amount >= 0 && cows > 0 {
		return amount / float64(cows), nil
	}
	if err == ErrScaleMalfunction && amount > 0 && cows > 0 {
		return amount / float64(cows) * 2, nil
	}
	if amount < 0 {
		return 0, errors.New(negative)
	}
	if cows == 0 {
		return 0, errors.New(division)
	}
	if cows < 0 {
		return 0, &SillyNephewError{
			message: "silly nephew, there cannot be ",
			details: cows,
		}
	} else {
		return 0, errors.New("non-scale error")
	}
}
