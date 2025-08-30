package data

import "fmt"

type Instructor struct {
	Id        int
	FirstName string
	LastName  string
	Score     int
}

func (i Instructor) String() string {
	return fmt.Sprintf("[%s %s](%d)", i.FirstName, i.LastName, i.Id)
}