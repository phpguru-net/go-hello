package main

import (
	"fmt"

	"go-hello.phpguru.net/test/lasagna"
	"go-hello.phpguru.net/test/timer"
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
	// compare two string
	if option1 < option2 {
		return option1 + " is clearly the better choice."
	}
	return option2 + " is clearly the better choice"
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		return originalPrice * 0.8
	}
	if age >= 10 {
		return originalPrice * 0.5
	}
	return originalPrice * 0.7
}

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!\n", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!\n", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	output := Welcome(name)
	output += fmt.Sprintf("You have been assigned to table %03d. Your table is %s, exactly %0.1f meters from here.\n", table, direction, distance)
	output += fmt.Sprintf("You will be sitting next to %s.", neighbor)
	return output
}

// GetItem retrieves an item from a slice at given position. The second return value indicates whether
// the given index exists in the slice or not.
func GetItem(slice []int, index int) (int, bool) {
	if index < 0 {
		return 0, false
	}
	if index >= len(slice) {
		return 0, false
	}
	// otherwise
	return slice[index], true
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if index < 0 || index >= len(slice) {
		slice = append(slice, value)
		return slice
	}
	slice[index] = value
	return slice
}

// PrefilledSlice creates a slice of given length and prefills it with the given value.
func PrefilledSlice(value, length int) []int {
	var slices []int = []int{}
	if length > 0 {
		for i := 0; i < length; i++ {
			slices = append(slices, value)
		}
	}
	return slices
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

// - Stand (S)
// - Hit (H)
// - Split (P)
// - Automatically win (W)

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var cardValues map[string]int = map[string]int{
		"ace":   11,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"ten":   10,
		"jack":  10,
		"queen": 10,
		"king":  10,
		"other": 0,
	}
	value, ok := cardValues[card]
	if ok {
		return value
	}
	return 0
}

// IsBlackjack returns true if the player has a blackjack, false otherwise.
// If 2 cards have the combined value of 21 ==> Blackjack
func IsBlackjack(card1, card2 string) bool {
	card1Value := ParseCard(card1)
	card2Value := ParseCard(card2)
	return card1Value+card2Value == 21
}

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string {
	if !isBlackjack {
		return "P"
	}
	if isBlackjack && dealerScore < 10 {
		return "W"
	}
	return "S"
}

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore, dealerScore int) string {
	switch {
	case handScore >= 17:
		return "S"
	case handScore <= 11:
		return "H"
	case handScore >= 12 && handScore <= 16 && dealerScore <= 6:
		return "S"
	case handScore >= 12 && handScore <= 16 && dealerScore >= 7:
		return "H"
	}
	return "S"
}

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var sum int = 0
	for i := 0; i < len(birdsPerDay); i++ {
		sum += birdsPerDay[i]
	}
	return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	// 1 week = 7 days
	sum := 0
	start := 7 * (week - 1)
	for i := start; i < start+7 && i < len(birdsPerDay); i++ {
		sum += birdsPerDay[i]
	}
	return sum
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i] += 1
	}
	return birdsPerDay
}

func main() {
	fmt.Printf("%v need NeedsLicense? %v\n", "car", NeedsLicense("car"))
	fmt.Printf("%v vs %v ? %v\n", "Bugatti Veyron", "Ford Pinto", ChooseVehicle("Bugatti Veyron", "Ford Pinto"))
	originalPrice := float64(1000)
	fmt.Printf("originalPrice = %v\n", originalPrice)
	fmt.Printf("%v years car's price is %1.2f\n", 2, CalculateResellPrice(originalPrice, 2))
	fmt.Printf("%v years car's price is %1.2f\n", 5, CalculateResellPrice(originalPrice, 5))
	fmt.Printf("%v years car's price is %1.2f\n", 11, CalculateResellPrice(originalPrice, 11))

	fmt.Println(AssignTable("Christiane", 27, "Frank", "on the left", 23.7834298))

	var slice []int = []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	fmt.Println(GetItem(slice, 2))
	fmt.Println(GetItem(slice, 100))
	fmt.Println(GetItem(slice, -100))
	SetItem(slice, 2, 100)
	fmt.Println(slice)
	var newSlice []int
	newSlice = PrefilledSlice(1, 99)
	fmt.Println(newSlice)
	slice = RemoveItem(slice, 3)
	fmt.Println(slice)
	for i := 0; i < count; i++ {

	}

	fmt.Printf("%s = %v\n", "queen", ParseCard("queen"))
	fmt.Printf("%s = %v\n", "ace", ParseCard("ace"))
	fmt.Printf("%s = %v\n", "jack", ParseCard("jack"))

	// carSpeed := 5
	// batteryDrain := 2
	// car := speed.NewCar(carSpeed, batteryDrain)
	// distance := 100
	// fmt.Println(car)
	// raceTrack := speed.NewTrack(distance)
	// fmt.Println(raceTrack)
	// speed.Drive(car)
	// fmt.Println(car)
	friendsList := []string{"noodles", "sauce", "mozzarella", "kampot pepper"}
	myList := []string{"noodles", "meat", "sauce", "mozzarella"}

	mixedList := lasagna.AddSecretIncredient(friendsList, myList)
	fmt.Println(myList)
	fmt.Println(friendsList)
	fmt.Println(mixedList)
	quantities := []float64{1.2, 3.6, 10.5}
	scaledQuantities := lasagna.ScaleRecipe(quantities, 4)
	fmt.Println(scaledQuantities)

	fmt.Println(timer.Schedule("7/25/2019 13:45:00"))
	fmt.Println(timer.Description("7/25/2019 13:45:00"))
	fmt.Printf("July 25, 2019 13:45:00 is passed? %v\n", timer.HasPassed("July 25, 2019 13:45:00"))
	fmt.Printf("December 9, 2112 11:59:59 is passed? %v\n", timer.HasPassed("December 9, 2112 11:59:59"))
	fmt.Printf("Thursday, July 25, 2019 13:45:00? %v\n", timer.IsAfternoonAppointment("Thursday, July 25, 2019 13:45:00"))
	fmt.Println(timer.AnniversaryDate())
	fmt.Println(timer.UTCToLocal("16/12/2021 03:38:00"))
	timer.DayLightSaving()
}
