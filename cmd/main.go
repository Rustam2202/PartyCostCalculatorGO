package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	pc "github.com/Rustam2202/PartyCostCalculatorGO/internal/partycalc"
)

func main() {
	jsonInput, err := os.Open("LastNewYear.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonInput.Close()

	byteValue, _ := ioutil.ReadAll(jsonInput)
	var personsFromJSON pc.Persons

	json.Unmarshal(byteValue, &personsFromJSON)
	result := pc.CalculateDebts(personsFromJSON, 1)
	result.ShowPayments(pc.Language(pc.ENG))
	//result.CheckCalculation(personsFromJSON)

	result.PrintToFile("result.txt", pc.Language(pc.RUS))
}