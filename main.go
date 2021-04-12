package main

import (
	"encoding/json"
	"fmt"
	"go-crud/model"
	"go-crud/routes"
	"net/http"
)

// func exec(db *sql.DB, sql string) sql.Result {
// 	result, err := db.Exec(sql)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return result
// }

func UserHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(user)

	// user := model.User{
	// 	ID:    1,
	// 	Nome:  "Gabriel",
	// 	Idade: 21,
	// }

	// data, err := json.Marshal(user)
	// if err != nil {
	// 	fmt.Println("error:", err)
	// }

	// w.Header().Set("Content-Type", "application/json")
	// w.Write(data)
}

func main() {
	// for {
	routes.StartServer()
	// }
	// servidor
	// http.HandleFunc("/users", UserHandler)
	// log.Fatal(http.ListenAndServe(":3000", nil))

	// conexão com o banco
	// db, err := sql.Open("mysql", "root:12345678@/")
	// if err != nil {
	// 	panic(err)
	// }

	// defer db.Close()

	// exec(db, "create database if not exists crudgo")
	// exec(db, "use crudgo")
	// exec(db, "drop table if exists usuarios")
	// exec(db, `create table usuarios (
	// 	id integer auto_increment,
	// 	nome varchar(80),
	// 	idade integer,
	// 	created_at datetime,
	// 	PRIMARY KEY (id)
	// )`)
}
