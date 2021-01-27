package main

import (
  "github.com/julienschmidt/httprouter"
  "net/http"
  "net/http/httptest"
  "testing"
  "io/ioutil"
  "strings"
)

func TestListTasksHandler(t *testing.T) {
  router := httprouter.New()
  router.GET("/", ListTasks)
  r, _ := http.NewRequest("GET", "/", nil)
  w := httptest.NewRecorder()
  s := Server{router}
  s.ServeHTTP(w, r)

  // test content type
  expectedContentTypeHeader := "application/json; charset=utf-8"
  if w.Header().Get("Content-Type") != expectedContentTypeHeader {
    t.Errorf("handler returned unexpected Content-Type header: got '%v', but want '%v'", w.Header().Get("Content-Type"), expectedContentTypeHeader)
  } else {
    t.Logf("Pass test content type")
  }

  // test status code
  expectedCode := http.StatusOK
  if w.Code != expectedCode {
    t.Errorf("handler returned unexpected status code: got '%v', but want '%v'", w.Code, expectedCode)
  } else {
    t.Logf("Pass test status code")
  }

  t.Log("Passes all tests")
}

func TestDeleteTasksHandler(t *testing.T) {
  router := httprouter.New()
  router.DELETE("/:id", DeleteTasks)

  r, _ := http.NewRequest("DELETE", "/mytask", nil)
  w := httptest.NewRecorder()
  s := Server{router}
  s.ServeHTTP(w, r)

  // test content type
  expectedContentTypeHeader := "application/json; charset=utf-8"
  if w.Header().Get("Content-Type") != expectedContentTypeHeader {
    t.Errorf("handler returned unexpected Content-Type header: got '%v', but want '%v'", w.Header().Get("Content-Type"), expectedContentTypeHeader)
  } else {
    t.Logf("Pass test content type")
  }

  expectedContentData := "DeleteTasks mytask"
  resp := w.Result()
  respBody, _ := ioutil.ReadAll(resp.Body)
  body := string(respBody)
  body = strings.Replace(body, "\n", "", -1)
  if body != expectedContentData {
    t.Errorf("handler returned unexpected body: got '%v', but want '%v'",
             body,
             expectedContentTypeHeader)
  } else {
    t.Log("Pass test body content")
  }

  t.Log("Passes all tests")
}
