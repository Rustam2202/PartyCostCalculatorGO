package json

import (
	"encoding/json"
	"io/ioutil"
	"os"

	p "party-calc/internal/person"
)

func ReadJSON(jsonInput *os.File) p.Persons {
	byteValue, _ := ioutil.ReadAll(jsonInput)
	var personsFromJSON p.Persons
	json.Unmarshal(byteValue, &personsFromJSON)
	return personsFromJSON
}
