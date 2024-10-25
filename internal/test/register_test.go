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

func TestPostUser(t *testing.T) {
	router := setup.SetupRouters()

	w := httptest.NewRecorder()

	exampleUser := entities.User{
		Login:          "test_name",
		HashedPassword: "test_pass",
	}

	userJson, _ := json.Marshal(exampleUser)
	req, _ := http.NewRequest("POST", "/register", strings.NewReader(string(userJson)))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code, w.Body)

	w = httptest.NewRecorder()

	req, _ = http.NewRequest("POST", "/register", strings.NewReader("{login:test}"))
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code, w.Body)

	w = httptest.NewRecorder()

	req, _ = http.NewRequest("POST", "/register", strings.NewReader(""))
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code, w.Body)
}

func TestLoginUser(t *testing.T) {
	router := setup.SetupRouters()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/login", strings.NewReader("{login:test}"))
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code, w.Body)

	w = httptest.NewRecorder()

	req, _ = http.NewRequest("POST", "/login", strings.NewReader(""))
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code, w.Body)
}
