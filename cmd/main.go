package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	calc "github.com/Rustam2202/PartyCostCalculatorGO/internal"
	person "github.com/Rustam2202/PartyCostCalculatorGO/internal/person"
	l "github.com/Rustam2202/PartyCostCalculatorGO/internal/language"
)

func main() {
	jsonInput, err := os.Open("LastNewYear.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonInput.Close()

	byteValue, _ := ioutil.ReadAll(jsonInput)
	var personsFromJSON person.Persons

	json.Unmarshal(byteValue, &personsFromJSON)
	result := calc.CalculateDebts(personsFromJSON, 1)
	result.ShowPayments(l.Language(l.ENG))
	//result.CheckCalculation(personsFromJSON)

	result.PrintToFile("result.txt", l.Language(l.RUS))

	//Test1()
	//Test2()
}

	result.PrintToFile("result.txt", pc.Language(pc.RUS))
}