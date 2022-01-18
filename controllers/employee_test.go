package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"happyhr/models"
	"happyhr/router"
	"net/http"
	"net/http/httptest"
	"testing"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/assert"
)

func TestEmployeeGet(t *testing.T) {
	router := router.Route

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/employees", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestEmployeeGPost(t *testing.T) {
	var emp models.Employee
	router := router.Route
	f := fuzz.New()
	f.Fuzz(&emp)
	emp_json, err := json.Marshal(emp)
	if err != nil {
		assert.Error(t, err)
	}
	fmt.Println(string(emp_json))
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/employees", bytes.NewReader(emp_json))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	fmt.Println(w.Body)
}
