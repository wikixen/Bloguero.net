package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/wikixen/blogapp/api/middleware"
	models "github.com/wikixen/blogapp/database/models"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var allUsers []models.User

	if res := db.Find(&allUsers); res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		j, err := json.Marshal(allUsers)
		if err != nil {
			http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		}
		w.Write(j)
		return
	}
}

func GetAUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	id := r.PathValue("id")
	if res := db.First(&user, "id = ?", id); res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		j, err := json.Marshal(user)
		if err != nil {
			http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		}
		w.Write(j)

		return
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if len(newUser.Password) < 8 && len(newUser.Password) > 128 {
		http.Error(w, "Invalid password length", http.StatusBadRequest)
	} else {
		newUser.Password, _ = middleware.HashPW(newUser.Password)
	
		if res := db.Create(&newUser); res.Error != nil {
			http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusCreated)
		}
	}
}

func EditUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	user.Password, _ = middleware.HashPW(user.Password)

	id := r.PathValue("id")
	if res := db.Where("id = ?", id).Updates(&user); res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		j, err := json.Marshal(user)
		if err != nil {
			http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		}
		w.Write(j)

		return
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	id := r.PathValue("id")
	if res := db.Delete(&user, "id = ?", id); res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		return
	}
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var input models.User

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	id := r.PathValue("id")
	if res := db.First(&user, "id = ?", id); res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusBadRequest)
	} else {
		match, err := middleware.AuthenticatePW(input.Password, user.Password)
		if err != nil {
			log.Fatal(err)
			// http.Error(w, "Incorrect username or password provided",http.StatusBadRequest)
		}
		if !match && input.Username == user.Username {
			w.WriteHeader(http.StatusOK)
			tokenStr, err := middleware.GenToken(user.Username)
			if err != nil {
				http.Error(w, "Error finding user", http.StatusInternalServerError)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(tokenStr))
			return
		} else {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		}
	}
}