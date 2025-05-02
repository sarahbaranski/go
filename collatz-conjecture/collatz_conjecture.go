package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("input must be a positive integer")
	}
	counter := 0
	for n > 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = (n * 3) + 1
		}
		counter += 1
	}
	return counter, nil
}
