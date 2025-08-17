package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	request, error := ioutil.ReadAll(r.Body)
	if error != nil {
		log.Fatal(error)
	}

	var user models.User
	if error := json.Unmarshal(request, &user); error != nil {
		log.Fatal(error)
	}

	db, error := database.Connect()
	if error != nil {
		log.Fatal(error)
	}

	repositories := repositories.NewRepositoryUsers(db)
	repositories.Create(user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listando usuários..."))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Obtendo usuário..."))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuário..."))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuário..."))
}
