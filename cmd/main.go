package main

import (
	"fmt"
	"os"

	c "github.com/Rustam2202/PartyCostCalculatorGO/internal"
	l "github.com/Rustam2202/PartyCostCalculatorGO/internal/language"
	j "github.com/Rustam2202/PartyCostCalculatorGO/json"
)

func main() {
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
