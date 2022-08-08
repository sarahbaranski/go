package beer

import (
	"errors"
	"fmt"
)

const part1 = "bottles of beer on the wall,"
const part2 = "bottles of beer.\nTake one down and pass it around,"
const part3 = "bottles of beer on the wall.\n"
const verse02 = "2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.\n"
const verse01 = "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n"
const verse00 = "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n"

func Song() string {
	s, _ := Verses(99, 0)
	return s
}

func Verses(start, stop int) (string, error) {
	if start > 99 || stop < 0 || start <= stop {
		return "", errors.New("expected error")
	}
	total := ""
	for i := start; i >= stop; i-- {
		s, e := Verse(i)
		if e != nil {
			return "", e
		}
		total += s + "\n"
	}
	return total, nil
}

func Verse(n int) (string, error) {
	if n > 99 {
		return "", errors.New("expected error")
	}
	if n == 2 {
		return verse02, nil
	}
	if n == 1 {
		return verse01, nil
	}
	if n == 0 {
		return verse00, nil
	}
	next := n - 1
	return fmt.Sprintf("%d %s %d %s %d %s", n, part1, n, part2, next, part3), nil
}
