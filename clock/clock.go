package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
	h int
	m int
}

func New(h, m int) Clock {
	c := new(Clock)
	c.m = Remap(m, 0, 60)
	h += MinToHour(m)
	c.h = Remap(h, 0, 24)

	return *c
}

func (c Clock) Add(m int) Clock {
	total := (c.h * 60) + c.m + m
	if total > DAILY_MINUTES {
		total %= DAILY_MINUTES
	}
	c.h = Remap(total/60, 0, 24)
	c.m = Remap(total%60, 0, 60)

	return c
}

func (c Clock) Subtract(m int) Clock {
	if m > DAILY_MINUTES {
		m %= DAILY_MINUTES
	}
	total := (c.h * 60) + c.m
	total -= m
	if total < 0 {
		total += DAILY_MINUTES
	}
	c.h = Remap(total/60, 0, 24)
	c.m = Remap(total%60, 0, 60)

	return c
}

func (c Clock) String() string {
	minute := fmt.Sprintf("%d", c.m)
	if c.m < 10 {
		minute = "0" + minute
	}
	hour := fmt.Sprintf("%d", c.h)
	if c.h < 10 {
		hour = "0" + hour
	}

	return fmt.Sprintf("%s:%s", hour, minute)
}

// UTILITIES
const DAILY_MINUTES = 1440 // 24 * 60

func Remap(value, min, max int) int {
	if value < min {
		if (value * -1) > max {
			return max + value%max
		}
		return max + value
	}
	if value > max {
		return value % max
	}
	if value == max {
		return 0
	}

	return value
}

func MinToHour(min int) int {
	if min >= 60 || min%60 == 0 {
		return min / 60
	} else if min < 0 {
		return ((min / 60) - 1)
	}

	return 0
}
