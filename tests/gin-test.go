package tests

import (
	"fmt"
	"party-calc/internal/person"

	"github.com/gin-gonic/gin"
)

func TestGin() {
	router := gin.Default()
	router.GET("/")
}

func postPerson(ctx *gin.Context) {
	var per person.Person
	err := ctx.ShouldBindJSON(&per)
	if err != nil {
		fmt.Println(err)
	}
}
