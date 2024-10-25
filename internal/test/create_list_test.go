package main

import (
	"encoding/json"
	"github.com/Redvin-dt/Lists-for-all-life-cases.git/internal/entities"
	"github.com/Redvin-dt/Lists-for-all-life-cases.git/internal/setup"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateList(t *testing.T) {
	router := setup.SetupRouters()

	w := httptest.NewRecorder()

	exampleList := entities.List{
		Values: []string{"first item", "second item"},
	}

	listJson, _ := json.Marshal(exampleList)
	req, _ := http.NewRequest("POST", "/list/create", strings.NewReader(string(listJson)))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code, w.Body)

	w = httptest.NewRecorder()

	req, _ = http.NewRequest("POST", "/register", strings.NewReader(""))
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code, w.Body)
}
