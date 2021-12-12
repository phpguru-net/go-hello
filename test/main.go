package main

import (
	"fmt"
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
	// compare two string
	if(option1 < option2) {
		return option1 + " is clearly the better choice."
	}
	return option2 + " is clearly the better choice"
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if(age < 3) {
		return originalPrice * 0.8
	}
	if(age >= 10) {
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


func main(){
	fmt.Printf("%v need NeedsLicense? %v\n", "car", NeedsLicense("car"))
	fmt.Printf("%v vs %v ? %v\n", "Bugatti Veyron", "Ford Pinto", ChooseVehicle("Bugatti Veyron","Ford Pinto"))
	originalPrice := float64(1000)
	fmt.Printf("originalPrice = %v\n", originalPrice)
	fmt.Printf("%v years car's price is %1.2f\n", 2, CalculateResellPrice(originalPrice,2))
	fmt.Printf("%v years car's price is %1.2f\n", 5, CalculateResellPrice(originalPrice,5))
	fmt.Printf("%v years car's price is %1.2f\n", 11, CalculateResellPrice(originalPrice,11))

	fmt.Println(AssignTable("Christiane", 27, "Frank", "on the left", 23.7834298))
}