package main

import (
  "fmt"
  "log"
  "net/http"
  "github.com/julienschmidt/httprouter"
)

type Server struct {
  r *httprouter.Router
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json; charset=utf-8")
  s.r.ServeHTTP(w, r)
}

func ListTasks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  fmt.Fprint(w, "ListTasks\n")
}

func CreateTasks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  fmt.Fprint(w, "CreateTasks\n")
}

func ReadTasks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  fmt.Fprint(w, "ReadTasks " + ps.ByName("id") + "\n")
}

func UpdateTasks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  fmt.Fprint(w, "UpdateTasks " + ps.ByName("id") + "\n")
}

func DeleteTasks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  fmt.Fprint(w, "DeleteTasks " + ps.ByName("id") + "\n")
}

func main() {
  router := httprouter.New()
  router.GET("/", ListTasks)
  router.POST("/", CreateTasks)
  router.GET("/:id", ReadTasks)
  router.PUT("/:id", UpdateTasks)
  router.DELETE("/:id", DeleteTasks)

  log.Fatal(http.ListenAndServe(":8080", &Server{router}))
}
