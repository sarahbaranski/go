package robotname

import (
	"math/rand"
	"strconv"
	"time"
)

// Define the Robot type here.
type Robot struct {
	name string
}

const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)

}

func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}
	newRobot := RandStringBytes(2)
	rand.Seed(time.Now().UnixNano())
	min := 100
	max := 999
	randomInteger := rand.Intn(max-min+1) + min
	intString := strconv.Itoa(randomInteger)
	newName := newRobot + intString
	r.name = newName

	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}
