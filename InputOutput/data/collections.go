package data

import "fmt"

var Codes = map[int]string{80: "http", 443: "https"}

var Countries [10]string

func init() {
	cities := []string{"Mumbai", "Delhi"}
	cities = append(cities, "Nagpur")
	fmt.Println(cities)
	Countries[0] = "India"
	Countries[2] = "Spain"
	Countries[4] = "Japan"
	Countries[5] = "Italy"
	Countries[7] = "Canada"
	Countries[9] = "Nepal"
}
