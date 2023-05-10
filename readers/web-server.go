package readers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	//"github.com/gorilla/mux"

	c "party-calc/internal"
	"party-calc/internal/person"
)

func handlePost(w http.ResponseWriter, r *http.Request) {
	var pers person.Persons
	err := json.NewDecoder(r.Body).Decode(&pers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := c.CalculateDebts(pers, 1)

	err = json.NewEncoder(w).Encode(result)

}

func handleGet(w http.ResponseWriter, r *http.Request) {
	//err := json.NewEncoder(w).Encode()
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusBadRequest)
	//	return
	//}
}

func NewWebServer() {
	//r := mux.NewRouter()
	//r.HandleFunc("/", handlePost).Methods("POST")
	fmt.Println("Server started")
	http.HandleFunc("/", handlePost)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
