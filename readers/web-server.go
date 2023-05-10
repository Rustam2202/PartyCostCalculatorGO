package readers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"party-calc/internal"
	"party-calc/internal/person"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var pers person.Persons
	err := json.NewDecoder(r.Body).Decode(&pers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// curl -s -XPOST -d'{ "persons": [{"name": "Рустам","spent": 4050}]}' http://localhost:8080/

	result := internal.CalculateDebts(pers, 1)

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func NewWebServer() {
	fmt.Println("Server started")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
