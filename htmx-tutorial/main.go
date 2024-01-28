package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/massanaRoger/go/htmx-tutorial/components"
	user_model "github.com/massanaRoger/go/htmx-tutorial/models"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	fileserver := http.FileServer(http.Dir("public"))
	r.Handle("/", http.StripPrefix("/", fileserver))

	r.Get("/users", getUsers)

	fmt.Println("Server running on port 3000")
	http.ListenAndServe(":3000", r)

}

func getUsers(w http.ResponseWriter, r *http.Request) {
	var users []user_model.User = []user_model.User{
		{Id: 1, Name: "John"},
		{Id: 2, Name: "Jane"},
		{Id: 3, Name: "Billy"},
	}

	templ.Handler(components.User(users)).ServeHTTP(w, r)

}
