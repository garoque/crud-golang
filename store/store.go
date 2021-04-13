package store

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-crud/model"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func openDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:12345678@/crudgo")

	return db, err
}

//SetUser cria um novo usuário
func SetUser(w http.ResponseWriter, r *http.Request) {
	db, err := openDB()
	if err != nil {
		fmt.Println("error SetUser openDB:", err)
	}
	defer db.Close()

	var user model.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("error SetUser Decode:", err)
	}

	user.CreatedAt = time.Now()

	stmt, err := db.Prepare("INSERT INTO usuarios (nome, idade, created_at) VALUES (?, ?, ?)")
	if err != nil {
		fmt.Println("error SetUser db.Prepare:", err)
	}

	_, err = stmt.Exec(user.Nome, user.Idade, user.CreatedAt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Algum erro ocorreu!"))
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Usuário cadastrado com sucesso!"))
}

//GetUser retorna um usuário específico pelo ID
func GetUser(w http.ResponseWriter, id int) {
	db, err := openDB()
	if err != nil {
		fmt.Println("error GetUser openDB:", err)
	}
	defer db.Close()

	row := db.QueryRow(`SELECT * FROM usuarios WHERE id = ?`, id)

	var user model.User
	row.Scan(&user.ID, &user.Nome, &user.Idade, &user.CreatedAt)

	if user.ID == 0 {
		w.Write([]byte("Não foi encontrado nenhum usuário!"))
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}

//GetAllUsers retorna todos os usuários
func GetAllUsers(w http.ResponseWriter) {
	db, err := openDB()
	if err != nil {
		fmt.Println("error GetAllUsers openDB:", err)
	}
	defer db.Close()

	rows, _ := db.Query("SELECT * FROM usuarios")

	var users []model.User
	for rows.Next() {
		var user model.User
		rows.Scan(&user.ID, &user.Nome, &user.Idade, &user.CreatedAt)

		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

//EditUser edita um usuário específico pelo ID
func EditUser(w http.ResponseWriter, r *http.Request, id int) {
	db, err := openDB()
	if err != nil {
		fmt.Println("error EditUser openDB:", err)
	}
	defer db.Close()

	var user model.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("error EditUser Decode:", err)
	}

	stmt, err := db.Prepare("UPDATE usuarios SET nome = ?, idade = ? WHERE id = ?")
	if err != nil {
		fmt.Println("error EditUser stmt.Exec:", err)
	}

	stmt.Exec(user.Nome, user.Idade, id)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Usuário atualizado com sucesso!"))
}

//DeleteUser remove um usuário específico pelo ID
func DeleteUser(w http.ResponseWriter, id int) {
	db, err := openDB()
	if err != nil {
		fmt.Println("error EditUser openDB:", err)
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM usuarios WHERE id = ?")
	if err != nil {
		fmt.Println("error EditUser stmt.Exec:", err)
	}

	stmt.Exec(id)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Usuário removido com sucesso!"))
}
