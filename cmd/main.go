package main

import (
	"encoding/json"
	"net/http"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type User struct{
	gorm.Model
	Username string
}

type Blogs struct {
	gorm.Model
	Title    string `gorm:"unique;not null;type:varchar(100)"`
	Author   string
	Content  string `gorm:"not null"`
	Likes    uint
	Dislikes uint
}

var db, _ = gorm.Open(sqlite.Open("blogs.db"), &gorm.Config{})

func main() {
	db.AutoMigrate(&Blogs{})

	app := http.NewServeMux()

	app.HandleFunc("POST /", CreateBlog)
	app.HandleFunc("GET /", GetBlogs)
	app.HandleFunc("GET /{id}", GetABlog)
	app.HandleFunc("PATCH /{id}", EditABlog)
	app.HandleFunc("DELETE /{id}", DeleteABlog)

	http.ListenAndServe(":8080", app)
}

// CreateBlog creates a blog
func CreateBlog(w http.ResponseWriter, r *http.Request) {
	var newBlog Blogs

	err := json.NewDecoder(r.Body).Decode(&newBlog)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if res := db.Create(&newBlog); res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(http.StatusCreated)
		return
	}
}

// GetBlogs gets all blogs in an array
func GetBlogs(w http.ResponseWriter, r *http.Request) {
	var allBlogs []Blogs

	if res := db.Find(&allBlogs); res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		j, err := json.Marshal(allBlogs)
		if err != nil {
			http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		}
		w.Write(j)
		return
	}
}

// GetABlog gets a single blog by ID
func GetABlog(w http.ResponseWriter, r *http.Request) {
	var blog Blogs
	id := r.PathValue("id")

	if res := db.First(&blog, "id = ?", id); res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		j, err := json.Marshal(blog)
		if err != nil {
			http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		}
		w.Write(j)

		return
	}
}

// EditABlog allows the editing of blogs
func EditABlog(w http.ResponseWriter, r *http.Request) {
	// Grabs JSON from r
	var blog Blogs

	err := json.NewDecoder(r.Body).Decode(&blog)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Search DB using GORM
	id := r.PathValue("id")
	if res := db.Where("id = ?", id).Updates(&blog); 
	res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		j, err := json.Marshal(blog)
		if err != nil {
			http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		}
		w.Write(j)

		return
	}
}

// DeleteABlog allows the deletion of blogs
func DeleteABlog(w http.ResponseWriter, r *http.Request) {
	// Grabs JSON from r
	var blog Blogs

	err := json.NewDecoder(r.Body).Decode(&blog)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Search DB using GORM
	id := r.PathValue("id")
	if res := db.Delete(&blog, "id = ?", id); res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		return
	}
}
