package main

import (
	"fmt"

	complexpb "github.com/sunil206b/Protobuff3_Go/src/complex"
	enumpb "github.com/sunil206b/Protobuff3_Go/src/enums"
	"github.com/sunil206b/Protobuff3_Go/src/simple"
)

func main() {
	person := doSimple()
	readAndWriteToFile(person)
	fmt.Println("--------------------------------------------------------")
	readAndWriteTOJSON(person)
	fmt.Println("--------------------------------------------------------")
	em := doEnum()
	fmt.Println(em)
	fmt.Println("--------------------------------------------------------")
	cm := doComplex()
	fmt.Println(cm)
}

func doSimple() *simple.Person {
	person := simple.Person{
		Age:               31,
		FirstName:         "Kohli",
		LastName:          "Virat",
		Height:            5.9,
		IsProfileVerified: true,
		PoneNumbers:       []string{"5234534254", "7545634545"},
	}
	return &person
}

func doEnum() *enumpb.EnumMessage {
	em := enumpb.EnumMessage{
		Id:        42,
		DayOfWeek: enumpb.DayOfWeek_SUNDAY,
	}
	return &em
}

func doComplex() *complexpb.Building {
	cm := complexpb.Building{
		Name:   "vaikuntapuram",
		Number: "535",
		Street: &complexpb.Street{
			Name: "Jubli 123",
			City: &complexpb.City{
				Name:    "Hyderabad",
				Country: "India",
				ZipCode: "500001",
			},
		},
	}
	return &cm
}
