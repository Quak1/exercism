package lasagnamaster

func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer == 0 {
		timePerLayer = 2
	}

	return len(layers) * timePerLayer
}

func Quantities(layers []string) (int, float64) {
	totalSauce := 0.0
	totalNoodles := 0

	for _, layer := range layers {
		switch layer {
		case "sauce":
			totalSauce += 0.2
		case "noodles":
			totalNoodles += 50
		default:
			continue
		}
	}

	return totalNoodles, totalSauce
}

func AddSecretIngredient(layers, myLayers []string) {
	myLayers[len(myLayers)-1] = layers[len(layers)-1]
}

func ScaleRecipe(amounts []float64, portions int) []float64 {
	newAmounts := []float64{}

	for _, amount := range amounts {
		newAmounts = append(newAmounts, amount*float64(portions)/2)
	}

	return newAmounts
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
