package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"party-calc/readers"
)

type caseType struct {
	input string
	want  string
}

var case3 = caseType{"", ""}

var cases []caseType

func TestJsonHandler(t *testing.T) {

	router := gin.Default()
	router.GET("/", readers.JsonHandler)
	req := httptest.NewRequest(http.MethodGet, "/", bytes.NewReader([]byte(case1.InputString)))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.JSONEq(t, case1.Want, w.Body.String())
}
