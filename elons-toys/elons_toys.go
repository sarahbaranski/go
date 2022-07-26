package elon

import (
	"fmt"
	"math"
)

// TODO: define the 'Drive()' method
func (car *Car) Drive() {
	if car.battery > car.batteryDrain {
		car.distance += car.speed
		car.battery -= car.batteryDrain
	}
}

// TODO: define the 'DisplayDistance() string' method
func (c Car) DisplayDistance() string {
	distance := c.distance
	return fmt.Sprintf("Driven %d meters", distance)
}

// TODO: define the 'DisplayBattery() string' method
func (c Car) DisplayBattery() string {
	battery := c.battery
	str := "%"
	return fmt.Sprintf("Battery at %d%s", battery, str)
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (c Car) CanFinish(trackDistance int) bool {
	carTimes := int(math.Ceil(float64(trackDistance) / float64(c.speed)))
	carDrain := c.batteryDrain * carTimes
	return carDrain <= c.battery
}
