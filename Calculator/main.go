package main

import "fmt"

func main() {
	var operation string
	var num1, num2 int

	fmt.Println("Go Calculator")
	fmt.Println("<===========>")
	fmt.Println("Operation to perform (add, subtract, multiply, divide):")
	fmt.Scanf("%s \n", &operation)
	fmt.Println("Enter two numbers seperated by space:")
	fmt.Scanln(&num1, &num2)
	switch operation {
	case "add":
		fmt.Println(num1 + num2)
	case "subtract":
		fmt.Println(num1 - num2)
	case "multiply":
		fmt.Println(num1 * num2)
	case "divide":
		fmt.Println(num1 / num2)
	}
}
