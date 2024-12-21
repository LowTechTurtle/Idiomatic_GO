package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const data = `
{
	"id": "ABCD-123",
	"first_name": "Bob",
	"last_name": "Bobson",
	"title": "Senior Manager"
}
{
	"id": "XYZ-123",
	"first_name": "Mary",
	"last_name": "Maryson",
	"title": "Vice President"
}
{
	"id": "BOTX-263",
	"first_name": "",
	"last_name": "Garciason",
	"title": "Manager"
}
{
	"id": "HLXO-829",
	"first_name": "Pierre",
	"last_name": "",
	"title": "Intern"
}
{
	"id": "MOXW-821",
	"first_name": "Franklin",
	"last_name": "Watanabe",
	"title": ""
}
{
	"id": "",
	"first_name": "Shelly",
	"last_name": "Shellson",
	"title": "CEO"
}
{
	"id": "YDOD-324",
	"first_name": "",
	"last_name": "",
	"title": ""
}
`

var ErrInvalidID = errors.New("invalid ID")

type Employee struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Title     string `json:"title"`
}

var (
	validID = regexp.MustCompile(`\w{4}-\d{3}`)
)

type EmptyFieldErr struct {
	FieldName string
}

func (efe EmptyFieldErr) Error() string {
	return fmt.Sprintf("%s field is empty", efe.FieldName)
}

func (efe EmptyFieldErr) As(other *EmptyFieldErr) {
	other.FieldName = efe.FieldName
}

func ValidateEmployee(e Employee) error {
	var allErrors []error

	if len(e.ID) == 0 {
		allErrors = append(allErrors, EmptyFieldErr{FieldName: "ID"})
	} else if !validID.MatchString(e.ID) {
		allErrors = append(allErrors, ErrInvalidID)
	}

	if len(e.FirstName) == 0 {
		allErrors = append(allErrors, EmptyFieldErr{FieldName: "FirstName"})
	}

	if len(e.LastName) == 0 {
		allErrors = append(allErrors, EmptyFieldErr{FieldName: "LastName"})
	}

	if len(e.Title) == 0 {
		allErrors = append(allErrors, EmptyFieldErr{FieldName: "Title"})
	}

	switch len(allErrors) {
	case 0:
		return nil
	case 1:
		return allErrors[0]
	default:
		return errors.Join(allErrors...)
	}
}

func processError(err error, emp Employee) string {
	var fieldErr EmptyFieldErr
	if errors.Is(err, ErrInvalidID) {
		return fmt.Sprintf("invalid ID: %s", emp.ID)
	} else if errors.As(err, &fieldErr) {
		return fmt.Sprintf("empty field %s", fieldErr.FieldName)
	} else {
		return fmt.Sprintf("%v", err)
	}
}

func main() {
	d := json.NewDecoder(strings.NewReader(data))
	count := 0
	for d.More() {
		count++
		var emp Employee
		err := d.Decode(&emp)
		if err != nil {
			fmt.Printf("record %d: %v\n", count, err)
			continue
		}
		err = ValidateEmployee(emp)
		message := fmt.Sprintf("record %d: %+v", count, emp)
		switch err := err.(type) {
		case interface{ Unwrap() []error }:
			allErrors := err.Unwrap()
			var messages []string
			for _, e := range allErrors {
				messages = append(messages, processError(e, emp))
			}
			message = message + " allErrors: " + strings.Join(messages, ", ")
		default:
			message = message + " error: " + processError(err, emp)
		}

		if err != nil {
			var empfield EmptyFieldErr
			if errors.Is(err, ErrInvalidID) {
				fmt.Printf("record %d: %+v error: InvalidID, ID: %s, %v\n", count, emp, emp.ID, err)
				continue
			}
			if errors.As(err, &empfield) {
				fmt.Printf("record %d: %+v error: %s\n", count, emp, empfield.FieldName)
				continue
			} else {
				fmt.Printf("record %d: %+v error: %v\n", count, emp, err)
				continue
			}
		}
		fmt.Printf("record %d: %+v\n", count, emp)
	}
}
