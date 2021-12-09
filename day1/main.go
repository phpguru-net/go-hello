package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

const ONE_BILLION = 1e9
const A_BYTE = 1
const A_MEGA_BYTE = 1 << 10

type DataTypes struct {
	aBool           bool
	aString         string
	anIntegerNumber int
	aFloatNumber    float64
}

type Season int64

const (
	Spring Season = iota // = 0
	Summer               // = 1
	Autumn               // = 2
	Winter               // = 3
)

type TaskStatus string

const (
	Pending    TaskStatus = "pending"
	Inprogress TaskStatus = "inprogress"
	Closed     TaskStatus = "closed"
	Rejected   TaskStatus = "rejected"
)

type Task struct {
	name       string     `default:"Undefined Task" maxLength:"100"`
	estimation float64    `default:"0"`
	startDate  string     `default:"yyyy-mm-dd"`
	endDate    string     `default:"yyyy-mm-dd"`
	status     TaskStatus `default:"Pending"`
}

// define print method for struct
func (task *Task) print() {
	var output string = ""
	output += fmt.Sprintf("Name: %s\n", task.name)
	output += fmt.Sprintf("Estimation: %0.1f\n", task.estimation)
	output += fmt.Sprintf("Start Date: %s\n", task.startDate)
	output += fmt.Sprintf("End Date: %s\n", task.endDate)
	output += fmt.Sprintf("Status: %s\n", task.status)
	fmt.Println(output)
}

func printFieldDefaultValue(typ reflect.Type, name string) {
	var field reflect.StructField

	field, _ = typ.FieldByName(name)
	fmt.Printf("%v 's default value is %v\n", field.Name, field.Tag.Get("default"))
	fmt.Printf("%v 's maxLength is %v\n", field.Name, field.Tag.Get("maxLength"))
}

func (task *Task) setDefaultValues() {
	// check empty and set default value
	typ := reflect.TypeOf(*task)
	// print default value of all field
	printFieldDefaultValue(typ, "name")
	printFieldDefaultValue(typ, "estimation")
	printFieldDefaultValue(typ, "startDate")
	printFieldDefaultValue(typ, "endDate")
	printFieldDefaultValue(typ, "status")
}

// func printTask(task Task) {
// 	var output string = ""
// 	output += fmt.Sprintf("Name: %s\n", task.name)
// 	output += fmt.Sprintf("Estimation: %0.1f\n", task.estimation)
// 	output += fmt.Sprintf("Start Date: %s\n", task.startDate)
// 	output += fmt.Sprintf("End Date: %s\n", task.endDate)
// 	output += fmt.Sprintf("Status: %s\n", task.status)
// 	fmt.Println(output)
// }

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
	var anIntegerNumber int = ONE_BILLION // 1,000,000,000
	var aFloatNumber = 0.123456789
	var dTypes DataTypes
	dTypes.aBool = aBool
	dTypes.aString = aString
	dTypes.anIntegerNumber = anIntegerNumber
	dTypes.aFloatNumber = aFloatNumber
	displayVariableTypeAndSize(dTypes)

	fmt.Printf("A_BYTE = %v bits\n", A_BYTE<<3)
	fmt.Printf("A_MEGA_BYTE = %v bytes\n", A_MEGA_BYTE)

	// user define types
	// custom enum with constant
	var season Season
	season = Winter

	fmt.Printf("season == 1 vs season == Winter ||||| %v vs %v \n", season == 1, season == Winter)

	// A map
	cities := make(map[string]string)
	cities["hcm"] = "Ho Chi Minh"
	cities["hn"] = "Ha Noi"

	for key, value := range cities {
		fmt.Printf("%v - %v\n", key, value)
	}

	// struct
	var taskA Task = Task{name: "Task A"}
	taskA.startDate = "2021-12-09"
	taskA.status = Pending
	// define a method for struct

	// default value for struct
	// 1. define a new method and set default value
	// 2. use tags to define default value

	var taskB Task = taskA
	taskB.name = "Task B"

	// pointer
	var taskACopy *Task = &taskA
	taskACopy.name = "Task A ( New name )"

	taskA.setDefaultValues()
	taskB.setDefaultValues()

	taskA.print()
	(*taskACopy).print()
	taskB.print()
}
