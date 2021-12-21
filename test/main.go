package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"unicode/utf8"

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

	// range iterator
	// slice
	var numberSlices []int = []int{1, 1, 2, 3, 5, 7}
	for idx, number := range numberSlices {
		fmt.Printf("%d. %d\n", idx+1, number)
	}
	// map
	var countryCodes map[string]string = map[string]string{
		`vn`: "Vietnam",
		`gb`: "Great Britain",
		"us": "United States",
	}
	for ckey, cvalue := range countryCodes {
		fmt.Printf("%s - %s\n", ckey, cvalue)
	}
	// omit key/value/key+value
	for _, number := range numberSlices {
		fmt.Printf("%03d\n", number)
	}
	for c, _ := range countryCodes {
		fmt.Println(c)
	}
	count := 0
	for range countryCodes {
		count++
	}
	fmt.Println()

	var alphaCharStart = 'A'
	var alphaCharEnd = 'H'
	for alphaChar := alphaCharStart; alphaChar <= alphaCharEnd; alphaChar++ {
		fmt.Printf("%s\n", string(alphaChar))
	}

	var initialVotes int
	initialVotes = 2

	var counter *int
	fmt.Println(VoteCount(counter))
	counter = NewVoteCounter(initialVotes)
	fmt.Println(*counter == initialVotes) // true
	fmt.Println(VoteCount(counter))
	IncrementVoteCount(counter, 1)
	fmt.Println(VoteCount(counter))
	IncrementVoteCount(counter, 2)
	fmt.Println(VoteCount(counter))
	IncrementVoteCount(counter, 3)
	fmt.Println(VoteCount(counter))
	var electionResult *ElectionResult = NewElectionResult("John", 5)
	fmt.Println(DisplayResult(electionResult))
	var finalResults = map[string]int{
		"Mary": 10,
		"John": 51,
	}

	DecrementVotesOfCandidate(finalResults, "Mary")
	fmt.Println(finalResults)
	speed := 5
	batteryDrain := 2
	car := NewCar(speed, batteryDrain)
	fmt.Printf("car is now %v\n", car)
	car.DisplayDistance()
	car.Drive()
	car.DisplayBattery()
	fmt.Printf("car is now %v\n", car)
	car.DisplayDistance()
	trackDistance := 100
	fmt.Println(car.CanFinish(trackDistance))

	// Strings & Runes
	var myString string = "❗ABCDEF"
	for idx, char := range myString {
		fmt.Printf("%d. %c - %U - %d\n", idx+1, char, char, char)
	}

	// len vs rune
	stringLength := len(myString)
	numberOfRunes := utf8.RuneCountInString(myString)
	fmt.Printf("myString - Length: %d - Runes: %d\n", stringLength, numberOfRunes)

	// var applicationChars []ApplicationCharacter = []ApplicationCharacter{
	// 	{Application: "recommendation", Character: `❗`, UnicodeCodePoint: "U+2757"},
	// 	{Application: "search", Character: `🔍`, UnicodeCodePoint: "U+1F50D"},
	// 	{Application: "weather", Character: `☀`, UnicodeCodePoint: "U+2600"},
	// }
	fmt.Println(Application("❗ recommended search product 🔍"))
	log := "please replace '👎' with '👍'"

	fmt.Println(Replace(log, '👎', '👍'))

	// string conv

	fmt.Println(DescribeNumberBox(4.1))
	fmt.Println(DescribeNumberBox(4))
	fmt.Println(DescribeNumberBox(testNumberBox{4}))
	// fmt.Println(testNumberBox{4})
	// fmt.Println(testNumberBox{4}.Number())
	// var exampleObj NumberBox = testNumberBox{4}
	// if correctobj, ok := exampleObj.(interface{Number() string}); ok {
	//   value := correctobj.Number()
	// }
	//fmt.Printf("%s", reflect.ValueOf(exampleObj).MethodByName("Number"))
	fmt.Println(DescribeAnything(testNumberBox{4}))
}

type NumberBox interface {
}

type testNumberBox struct {
	number NumberBox
}

func (nb testNumberBox) Number() string {
	return fmt.Sprintf("%v", nb.number)
}

func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %1.1f", f)
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	strFloat64, isFloat64 := nb.(float64)
	if isFloat64 {
		return fmt.Sprintf("This is a box containing the number %1.1f", strFloat64)
	}
	strInt, isInt := nb.(int)
	if isInt {
		return fmt.Sprintf("This is a box containing the number %1.1f", float64(strInt))
	}
	if correctobj, ok := nb.(interface{Number() string}); ok { 
	  value := correctobj.Number()	  
	  return value
	}
	_, isString := nb.(string)
    if isString {
        return "Return to sender"
    }
	num, err := strconv.Atoi(fmt.Sprintf("%v", reflect.ValueOf(nb).MethodByName("Number").Call([]reflect.Value{})[0]))
    if err == nil {
        return DescribeNumberBox(num)
    }
	return "Return to sender"
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	fancyNumber, isFancyNumber := fnb.(FancyNumber)
    if isFancyNumber {
        num, err := strconv.Atoi(fancyNumber.Value())
        if err == nil {
            return num
        }
    }
	return 0
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	num := ExtractFancyNumber(fnb)
    if num > 0 {
        return fmt.Sprintf("This is a fancy box containing the number %1.1f", float64(num))
    }
	return fmt.Sprint("This is a fancy box containing the number 0.0")
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	switch v:= i.(type){
        case int:
    		return DescribeNumber(float64(v))
        case float64:
        	return DescribeNumber(v)
        case FancyNumberBox:
    		return DescribeFancyNumberBox(v)        
    	case NumberBox:
    		return DescribeNumberBox(v)
    	default:
    		return "Return to sender"
    }
}

// This file contains types used in the exercise but should not be modified.
type SillyNephewError struct {
	cows int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %v cows", e.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodderAmount, err := weightFodder.FodderAmount()
	if cows == 0 {
		return 0.0, errors.New("Division by zero")
	}
	if cows < 0 {
		return 0.0, &SillyNephewError{cows: cows}
	}
	if err == ErrScaleMalfunction && fodderAmount > 0 {
		if cows > 0 {
			return float64(fodderAmount) * 2 / float64(cows), nil
		}

	}
	if fodderAmount < 0 {
		return 0.0, errors.New("Negative fodder")
	}
	if err != nil {
		return 0.0, err
	}
	return float64(fodderAmount) / float64(cows), nil
}

// WeightFodder returns the amount of available fodder.
type WeightFodder interface {
	FodderAmount() (float64, error)
}

// ErrScaleMalfunction indicates an error with the scale.
var ErrScaleMalfunction = errors.New("sensor error")

type ApplicationCharacter struct {
	Application      string
	Character        rune
	UnicodeCodePoint string
}

func findApplication(applicationChars []ApplicationCharacter, char rune) (string, bool) {

	for _, appChar := range applicationChars {

		if appChar.Character == char {
			return appChar.Application, true
		}
	}
	return "", false
}

func Application(line string) string {
	var applicationChars []ApplicationCharacter = []ApplicationCharacter{
		{Application: "recommendation", Character: '❗', UnicodeCodePoint: "U+2757"},
		{Application: "search", Character: '🔍', UnicodeCodePoint: "U+1F50D"},
		{Application: "weather", Character: '☀', UnicodeCodePoint: "U+2600"},
	}
	for _, char := range line {
		if application, ok := findApplication(applicationChars, char); ok {
			return application
		}
	}
	return "default"
}

func Replace(log string, oldRune, newRune rune) string {
	var newLog []rune = []rune{}
	// var lastIndex int = 0
	var runes []rune = []rune(log)
	for _, char := range runes {
		if char == oldRune {
			newLog = append(newLog, newRune)
		} else {
			newLog = append(newLog, char)
		}
	}
	return string(newLog)
}

func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}

type Car struct {
	speed        int
	batteryDrain int
	battery      int
	distance     int
}

func NewCar(speed, batteryDrain int) Car {
	return Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
	}
}

func (car *Car) Drive() {
	if car.battery-car.batteryDrain >= 0 {
		car.battery -= car.batteryDrain
		car.distance += car.speed
	}
}

func (car Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %v meters", car.distance)
}

func (car Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %v", int(float64(car.battery)/100*100)) + "%"
}

func (car Car) CanFinish(trackDistance int) bool {
	// speed - batteryDrain
	// maximumDistance - battery
	maxiumDistance := float64(car.battery) * float64(car.speed) / float64(car.batteryDrain)
	return maxiumDistance >= float64(trackDistance)
}

func VoteCount(counter *int) int {
	if counter == nil {
		return 0
	}
	return *counter
}

func NewVoteCounter(initialVotes int) *int {
	return &initialVotes
}

// IncrementVoteCount increments the value in a vote counter
func IncrementVoteCount(counter *int, increment int) {
	if counter == nil {
		*counter = 0
		return
	}
	*counter += increment
}

type ElectionResult struct {
	// Candidate's name
	Name string
	// Number of votes the candidate has
	Votes int
}

// NewElectionResult creates a new election result
func NewElectionResult(candidateName string, votes int) *ElectionResult {
	var electionResult *ElectionResult = &ElectionResult{
		Name:  candidateName,
		Votes: votes,
	}
	return electionResult
}

// DisplayResult creates a message with the result to be displayed
func DisplayResult(result *ElectionResult) string {
	return fmt.Sprintf("%s (%d)", result.Name, result.Votes)
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	for c, _ := range results {
		if c == candidate {
			results[c] -= 1
			break
		}
	}
}

// ZERO VALUES

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	return &Resident{
		Name:    name,
		Age:     age,
		Address: address,
	}
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
	if r.Name == "" {
		return false
	}
	if r.Address == nil || len(r.Address) == 0 {
		return false
	}

	for _, value := range r.Address {
		if value == "" {
			return false
		}
	}
	return true
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	r.Name = ""
	r.Address = nil
	r.Age = 0
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	count := 0
	for _, resident := range residents {
		if resident.HasRequiredInfo() {
			count++
		}
	}
	return count
}
