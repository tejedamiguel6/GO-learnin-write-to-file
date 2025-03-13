package main

import (
	"bufio"
	"fmt"
	"os"

	"example.com/firsproject/fileOperations"
)



type Employee struct {
	firstName string
	placeOfWork string
	timeOfEmployment float64
}


func main () {

	fileOperations.ReadFromFile()

	var  employee Employee

	// var employee = Employee{firstName , placeOfWork, timeOfEmployment}

	fmt.Println("hello first go app!")
	fmt.Println("++++====================++++")

	fmt.Println("what is your name?")
	fmt.Scanln(&employee.firstName)
	fmt.Println("Hello, ", employee.firstName)

	fmt.Println("where do you work?")
	reader := bufio.NewReader(os.Stdin)
	placeOfWork, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	employee.placeOfWork = placeOfWork[:len(placeOfWork)-1]

	fmt.Println("I see, you work at", employee.placeOfWork)
	fmt.Println("How many years at?", employee.placeOfWork, "?")
	fmt.Scanln(&employee.timeOfEmployment)

	if(employee.timeOfEmployment <= 1) {
		fmt.Println("I see, less than one year, you are a new employee")
		fileOperations.WriteOutputToFile(employee.firstName, placeOfWork, employee.timeOfEmployment)

	} else if( employee.timeOfEmployment >= 2 && employee.timeOfEmployment <= 5) {
		fmt.Println("You are a mid level employee")
		fileOperations.WriteOutputToFile(employee.firstName, placeOfWork, employee.timeOfEmployment)
	}

}