package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	c := make(chan FreqMap) //make channel set to FreqMap type
	m := FreqMap{}          //variable m is a map

	for _, s := range l {
		go func(l string) {
			c <- Frequency(l) //call Frequency through a channel
		}(s)
	}
	for range l {
		for k, v := range <-c { //loop through channel/characters to map Frequency
			m[k] += v
		}
	}

	return m
}
