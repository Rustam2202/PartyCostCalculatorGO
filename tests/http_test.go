package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nsf/jsondiff"
	"github.com/stretchr/testify/assert"

	"party-calc/internal"
	"party-calc/readers"
)

func TestPersonsHandler(t *testing.T) {
	router := gin.Default()
	router.GET("/", readers.PersonsHandler)
	req := httptest.NewRequest(http.MethodGet, "/", bytes.NewReader([]byte(case1.InputString)))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	//fmt.Println(w.Body.String())

	assert.JSONEq(t, case1.Want, w.Body.String())
}

func TestHandler1(t *testing.T) {
	marshalled, err := json.Marshal(case1.Input)
	if err != nil {
		t.Errorf("JSON marshalling error: %s", err)
	}
	r := httptest.NewRequest(http.MethodGet, "/", bytes.NewReader(marshalled))
	w := httptest.NewRecorder()
	readers.Handler(w, r)

	var result internal.PartyData
	err = json.Unmarshal(w.Body.Bytes(), &result)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(w.Body.String())
	opts := jsondiff.DefaultJSONOptions()
	diff, str := jsondiff.Compare(w.Body.Bytes(), []byte(case1.Want), &opts)
	fmt.Println(diff, str)

	if diff != jsondiff.FullMatch {
		t.Errorf("JSONs not matched")
	}

	//if w.Body.String() != case1.WantStruct {
	//	t.Errorf("Want %s, got %s", case1.Want,  w.Body.String())
	//}
}

func TestHandler(t *testing.T) {
	b := []byte(`{"persons": [{"name": "Рустам","spent": 4050}]}`)
	br := bytes.NewReader(b)
	r := httptest.NewRequest(http.MethodGet, "/", br)
	w := httptest.NewRecorder()
	readers.Handler(w, r)
	res := w.Result()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}
