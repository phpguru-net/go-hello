package main

import (
	"fmt"
	"unsafe"
)


type DataTypes struct {
	aBool bool
	aString string
	anIntegerNumber int
	aFloatNumber float64
}
// Generic type
func displayVariableTypeAndSize(param DataTypes) {
	fmt.Printf("'%v' is %T type. And its size is %v bytes\n", param.aBool, param.aBool, unsafe.Sizeof(param.aBool))
	fmt.Printf("'%v' is %T type. And its size is %v bytes\n", param.aString, param.aString, unsafe.Sizeof(param.aString))
	fmt.Printf("'%v' is %T type. And its size is %v bytes\n", param.anIntegerNumber, param.anIntegerNumber, unsafe.Sizeof(param.anIntegerNumber))
	fmt.Printf("'%v' is %T type. And its size is %v bytes\n", param.aFloatNumber, param.aFloatNumber, unsafe.Sizeof(param.aFloatNumber))

}

func main() {
	fmt.Println("Day 1 Examples")
	// list all primitive types and its size
	
	var aBool bool = true	
	var aString string = "Golang is a general purpose language"	
	var anIntegerNumber int = 1e9 // 1,000,000,000
	var aFloatNumber = 0.123456789

	var dTypes DataTypes
	dTypes.aBool = aBool
	dTypes.aString = aString
	dTypes.anIntegerNumber = anIntegerNumber
	dTypes.aFloatNumber = aFloatNumber
	
	displayVariableTypeAndSize(dTypes)

}
