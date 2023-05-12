package main

import "party-calc/readers"


func main() {
	readers.NewGinServer()

	// tests.TestHandler()
	//	readers.NewWebServer()
	/*
		jsonInput, err := os.Open("../LastNewYear.json")
		if err != nil {
			fmt.Println(err)
		}
		defer jsonInput.Close()

		persons := j.ReadJSON(jsonInput)

		result := c.CalculateDebts(persons, 1)
		result.ShowPayments(l.Language(l.ENG))
		result.PrintToFile("result.txt", l.Language(l.RUS))
	*/
}
