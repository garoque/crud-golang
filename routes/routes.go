package routes

import (
	"go-crud/store"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func StartServer() {
	http.HandleFunc("/users/", UserHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	stringId := strings.TrimPrefix(r.URL.Path, "/users/")
	id, _ := strconv.Atoi(stringId)

	switch {
	case r.Method == "GET" && id > 0:
		store.GetUser(w, id)
	case r.Method == "GET":
		store.GetAllUsers(w)
	case r.Method == "POST":
		store.SetUser(w, r)
	case r.Method == "PUT" && id > 0:
		store.EditUser(w, r, id)
	case r.Method == "DELETE":
		store.DeleteUser(w, id)
	}
}
