package main

import (
	"fmt"
	//	"go.dev/io/data"
)

func calculateTax(price float32) (cgst float32, sgst float32) {
	cgst = price * .09
	sgst = price * .09
	return 0, 0
}

func birthday(age *int) {
	fmt.Printf("Address: %v, Value: %d\n", age, *age)
	if *age < 10 {
		panic("WTF")
	}
	*age = *age + 1
}

func main() {
	fmt.Println("Hello from Go")
	//	fmt.Println(len(data.Countries), data.Countries)
	//	SGST, CGST := calculateTax(100)
	//	fmt.Println(SGST + CGST)

	defer fmt.Println("GoodBye")

	age := 45
	birthday(&age)
	fmt.Println(age)
}
