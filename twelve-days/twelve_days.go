package twelve

import "fmt"

var gifts = [12]string{"a Partridge in a Pear Tree", "two Turtle Doves", "three French Hens", "four Calling Birds", "five Gold Rings", "six Geese-a-Laying", "seven Swans-a-Swimming", "eight Maids-a-Milking", "nine Ladies Dancing", "ten Lords-a-Leaping", "eleven Pipers Piping", "twelve Drummers Drumming"}

var days = [12]string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}

func Verse(i int) string {
	verse := fmt.Sprintf("On the %s day of Christmas my true love gave to me:", days[i-1])
	for j := i - 1; j >= 0; j-- {
		v := fmt.Sprintf(" %s", gifts[j])
		if j > 0 {
			v += ","
		} else {
			v += "."
			if i > 1 {
				v = " and" + v
			}
		}
		verse += v
	}

	return verse
}

func Song() string {
	song := ""
	for i := 0; i < 12; i++ {
		song += Verse(i + 1)
		if i < 11 {
			song += "\n"
		}
	}
	return song
}
