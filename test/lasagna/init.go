package lasagna

const NoodlePerLayer = 50 // grams
const SaucePerLayer = 0.2 // liters

func PreparationTime(layers []string, averageTime int) int {
	if averageTime == 0 {
		return 2
	}
	return len(layers) * averageTime
}

func Qualities(layers []string) (int, float64) {
	noodleQuality := NoodlePerLayer * len(layers)
	sauceQuality := SaucePerLayer * float64(len(layers))
	return noodleQuality, sauceQuality
}

func uniqueIncredientList(incredientList []string) []string {
	// value must be unique
	var values map[string]bool = make(map[string]bool)
	var newIncredientList = make([]string, len(values))
	for _, value := range incredientList {
		_, valueExisted := values[value]
		if !valueExisted {
			values[value] = true
			newIncredientList = append(newIncredientList, value)
		}
	}
	return newIncredientList
}

func AddSecretIncredient(friendList []string, myList []string) []string {
	mixedList := append(myList, friendList[:len(friendList)-1]...)
	mixedList = uniqueIncredientList(mixedList)
	return append(mixedList, friendList[len(friendList)-1])
}

func ScaleRecipe(quantities []float64, portions int) (scaledQuantities []float64) {
	for _, value := range quantities {
		scaledQuantities = append(scaledQuantities, value/2*float64(portions))
	}
	return
}

func NakedReturn(firstName, lastName string) (fullname string) {
	fullname = firstName + " " + lastName
	return
}
