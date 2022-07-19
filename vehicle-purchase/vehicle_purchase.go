package purchase

import (
	"fmt"
	"strings"
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	result := strings.Compare(option1, option2)
	if result == -1 {
		return fmt.Sprintf("%s is clearly the better choice.", option1)
	}
	return fmt.Sprintf("%s is clearly the better choice.", option2)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var finalPrice float64
	if age >= 3 && age < 10 {
		finalPrice = originalPrice * .70
	}
	if age < 3 {
		finalPrice = originalPrice * .80
	}
	if age >= 10 {
		finalPrice = originalPrice * .50
	}
	return finalPrice
}
