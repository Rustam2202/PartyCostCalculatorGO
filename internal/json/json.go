package json

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	p "party-calc/internal/person"
)

func ReadJSON(jsonInput *os.File) p.Persons {
	jsonInput, err := os.Open("") // path/filename.json

	if err != nil {
		fmt.Println(err)
		return p.Persons{}
	}
	defer jsonInput.Close()

	byteValue, _ := ioutil.ReadAll(jsonInput)
	var personsFromJSON p.Persons
	json.Unmarshal(byteValue, &personsFromJSON)
	return personsFromJSON
}
