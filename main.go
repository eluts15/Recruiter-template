package main

import (
	"fmt"
	"html/template"
	"os"
	"time"
)

type Contact struct {
	ID        int
	LastName  string
	FirstName string
}

type Recuiter struct {
	EmployeeID   int
	LastName     string
	FirstName    string
	EmailAddress string
	Phone        string
}

type Position struct {
	Description string
}

type Statement struct {
	Date     time.Time
	Contact  Contact
	Position Position
}

func main() {
	// Generate a template with New(), run it through the functions, parse the template.
	t := template.Must(template.New("recruit.tmpl").ParseFiles("recruit.tmpl"))
	err := t.Execute(os.Stdout, createMockStatement())
	if err != nil {
		panic(err)
	}
}

func formatDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d%d%d", day, month, year)
}

func createMockStatement() Statement {
	return Statement{
		Date: time.Date(2017, 31, 10, 0, 0, 0, 0, time.UTC),
		Contact: Contact{
			ID:        1,
			LastName:  "Luts",
			FirstName: "Ethan",
		},
		Position: Position{
			Description: "Junior Developer",
		},
	}
}
