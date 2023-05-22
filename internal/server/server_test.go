package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var cases = []testStruct{
	onePerson, twoPersons, threePersons,
}

func TestHandler(t *testing.T) {
	for _, tt := range cases {
		input, err := json.Marshal(tt.input)
		if err != nil {
			t.Errorf("Not correct input-JSON: %s", err)
		}
		want, err := json.Marshal(tt.want)
		if err != nil {
			t.Errorf("Not correct want-JSON: %s", err)
		}

		router := gin.Default()
		router.GET("/", JsonHandler)
		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewReader(input))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		fmt.Println(w.Body.String())
		assert.JSONEq(t, string(want), w.Body.String(), tt.testName)
	}
}

func TestAny(t *testing.T) {
	// chose between test cases ('onePersons', 'twoPersons' ...)
	var test = threePersons
	input, err := json.Marshal(test.input)
	if err != nil {
		t.Errorf("Not correct input-JSON: %s", err)
	}
	want, err := json.Marshal(test.want)
	if err != nil {
		t.Errorf("Not correct want-JSON: %s", err)
	}

	router := gin.Default()
	router.GET("/", JsonHandler)
	req := httptest.NewRequest(http.MethodGet, "/", bytes.NewReader(input))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	// fmt.Println(string(want))
	// fmt.Println(w.Body.String())
	assert.JSONEq(t, string(want), w.Body.String(), test.testName) // ?? sensitive to persons-blocks order
}
