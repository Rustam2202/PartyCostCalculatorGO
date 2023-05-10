package main

import (
	"fmt"
	"os"

	c "party-calc/internal"
	l "party-calc/internal/language"
	j "party-calc/json"
	"party-calc/readers"
)

func main() {
	readers.NewWebServer()


	jsonInput, err := os.Open("../LastNewYear.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonInput.Close()

	persons := j.ReadJSON(jsonInput)

	result := c.CalculateDebts(persons, 1)
	result.ShowPayments(l.Language(l.ENG))
	result.PrintToFile("result.txt", l.Language(l.RUS))
}
