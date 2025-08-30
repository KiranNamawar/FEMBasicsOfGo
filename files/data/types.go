package main

import (
    "fmt"
)

type distance float64

func (miles distance) ToKM() distance {
    return 1.60934 * miles
}

func main() {
	kiran := User {24, " Kiran"}
	kiran.printUser()
}

type User struct {
    id   int
    name string
}

func (u User) printUser() {
    fmt.Print(u.id, u.name)
}
