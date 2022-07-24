package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer == 0 {
		timePerLayer = 2
	}
	numberOfLayers := len(layers)
	totalPrepTime := numberOfLayers * timePerLayer
	return totalPrepTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	var noodles int
	var sauce float64
	for i := 0; i < len(layers); i++ {
		if layers[i] == "sauce" {
			sauce += 0.2
		} else if layers[i] == "noodles" {
			noodles += 50
		}
	}
	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) []string {
	myList = myList[:len(myList)-1]
	newList := append(myList, friendsList[len(friendsList)-1])
	return newList
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(twoPortions []float64, portionsCooked int) []float64 {
	var finalPorts []float64
	for i := 0; i < len(twoPortions); i++ {
		finalPorts = append(finalPorts, (twoPortions[i]/2.0)*float64(portionsCooked))
	}
	return finalPorts
}
