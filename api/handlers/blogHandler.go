package handlers

import (
	"encoding/json"
	"net/http"

	models "github.com/wikixen/blogapp/database/models"
)

// CreateBlog creates a blog
func CreateBlog(w http.ResponseWriter, r *http.Request) {
	dbInit()
	var newBlog models.Blogs

	err := json.NewDecoder(r.Body).Decode(&newBlog)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if res := db.Create(&newBlog); res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

// GetBlogs gets all blogs in an array
func GetBlogs(w http.ResponseWriter, r *http.Request) {
	var allBlogs []models.Blogs

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
	var blog models.Blogs
	
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
func EditBlog(w http.ResponseWriter, r *http.Request) {
	// Grabs JSON from r
	var blog models.Blogs

	err := json.NewDecoder(r.Body).Decode(&blog)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Search DB using GORM
	id := r.PathValue("id")
	if res := db.Where("id = ?", id).Updates(&blog); res.Error != nil {
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
func DeleteBlog(w http.ResponseWriter, r *http.Request) {
	// Grabs JSON from r
	var blog models.Blogs

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
