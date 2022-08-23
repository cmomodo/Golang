package main

import "fmt"

func main() {
	fmt.Println("Hello")
	var operator string
	var input1, input2 int16

	fmt.Println("Please enter your first number: ")
	fmt.Scan(&input1)

	fmt.Println("Please enter your second number: ")
	fmt.Scan(&input2)

	fmt.Println("Select Your Operator? +,-,/,*")
	fmt.Scan(&operator)
	switch operator {
	case "+":
		fmt.Println(input1 + input2)
	case "-":
		fmt.Println(input1 - input2)
	case "/":
		fmt.Println(input1 / input2)
	case "*":
		fmt.Println(input1 * input2)
	default:
		fmt.Println("Please enter an input!")
	}

}
