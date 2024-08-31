package main

import (
	"encoding/json"
	"net/http"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Blogs struct {
	gorm.Model
	Title    string
	Author   string
	Content  string
	Likes    uint
	Dislikes uint
}

var db, err = gorm.Open(sqlite.Open("blogs.db"), &gorm.Config{})

func main() {
	db.AutoMigrate(&Blogs{})

	app := http.NewServeMux()

	app.HandleFunc("POST /", CreateBlog)
	app.HandleFunc("GET /", GetBlogs)
	app.HandleFunc("GET /{id}", GetABlog)

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

	if newBlog.Title == "" {
		http.Error(w, "Title can't be empty", http.StatusBadRequest)
	} else if newBlog.Content == "" {
		http.Error(w, "Content can't be empty", http.StatusBadRequest)
	}

	if res := db.Create(newBlog); res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(http.StatusCreated)
		return
	}
}

func GetBlogs(w http.ResponseWriter, r *http.Request) {
	var allBlogs []Blogs
	if res := db.Find(&allBlogs); res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		j, err := json.Marshal(allBlogs)
		if err != nil {
			http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		}
		w.Write(j)
	}
}

func GetABlog(w http.ResponseWriter, r *http.Request) {
	var blog Blogs
	id := r.PathValue("id")
	if res := db.First(&blog, "id = ?", id); res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		j, err := json.Marshal(blog)
		if err != nil {
			http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		}
		w.Write(j)
	}
}
