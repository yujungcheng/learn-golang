package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListTasksHandler(t *testing.T) {
	router := httprouter.New()
	router.GET("/", ListTasks)
	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	// test content type
	expectedContentTypeHeader := "application/json; charset=utf-8"
	if w.Header().Get("Content-Type") != expectedContentTypeHeader {
		t.Errorf("handler returned unexpected Content-Type header: got '%v', but want '%v'", w.Header().Get("Content-Type"), expectedContentTypeHeader)
	}

	// test status code
	expectedCode := http.StatusOK
	if w.Code != expectedCode {
		t.Errorf("handler returned unexpected status code: got '%v', but want '%v'", w.Code, expectedCode)
	}

}
