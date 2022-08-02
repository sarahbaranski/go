package grains

import (
	"errors"
)

const maxUint64 = (1 << 64) - 1

func Square(number int) (uint64, error) {
	if number <= 0 || number > 64 {
		return 0, errors.New("there are only 64 squares on a chessboard")
	}

	return uint64(1 << uint(number-1)), nil
}

func Total() uint64 {
	return uint64(maxUint64)

}
