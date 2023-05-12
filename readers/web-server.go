package readers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"party-calc/internal"
	"party-calc/internal/language"
	"party-calc/internal/person"

	"github.com/gin-gonic/gin"
)

func PersonsHandler(ctx *gin.Context) {
	var pers person.Persons
	err := ctx.ShouldBindJSON(&pers)
	if err != nil {
		fmt.Println(err)
	}
//	fmt.Println(pers)

	result := internal.CalculateDebts(pers, 1)
	ctx.JSON(http.StatusOK, result)
}

func NewGinServer() {
	router := gin.Default()
	router.GET("/", PersonsHandler)
	router.Run(":8080")
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var pers person.Persons
	err := json.NewDecoder(r.Body).Decode(&pers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// curl -X POST -d'{ "persons": [{"name": "Рустам","spent": 4050}]}' http://localhost:8080/

	result := internal.CalculateDebts(pers, 1)
	result.PrintPayments(language.ENG)

	encoder := json.NewEncoder(w)
	//	encoder.SetIndent("\t", "")

	err = encoder.Encode(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func NewWebServer() {
	fmt.Println("Server started")
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
