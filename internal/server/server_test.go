package server

import (
	"bytes"
	"encoding/json"
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

		assert.JSONEq(t, string(want), w.Body.String(), tt.testName)
	}
}
