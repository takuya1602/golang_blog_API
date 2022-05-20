package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"path"
	"time"

	_ "github.com/lib/pq"
)

type Env struct {
	Db *sql.DB
}

type Category struct {
	Id           int    `json:"id"`
	CategoryName string `json:"category_name"`
	Slug         string `json:"slug"`
}

type SubCategory struct {
	Id               int    `json:"id"`
	CategoryName     string `json:"category_name"`
	Slug             string `json:"slug"`
	ParentCategoryId int    `json:"parent_category_id"`
}

type Post struct {
	Id              int       `json:"id"`
	CategoryId      int       `json:"category_id"`
	SubCategoryId   int       `json:"sub_category_id"`
	Title           string    `json:"title"`
	Slug            string    `json:"slug"`
	EyeCatchingImg  string    `json:"eye_catching_img"`
	Content         string    `json:"content"`
	MetaDescription string    `json:"meta_description"`
	IsPublic        bool      `json:"is_public"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func main() {
	db, err := sql.Open("postgres", "user=gwp password=gwp dbname=go_blog sslmode=disable")
	if err != nil {
		panic(err)
	}
	e := Env{Db: db}
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/categories/", e.handleRequestCategory)
	http.HandleFunc("/sub-categories/", e.handleRequestSubCategory)
	http.HandleFunc("/posts/", e.handleRequestPosts)
	http.HandleFunc("/admin/", e.handleRequestAdmin)
	http.Handle("/media/", http.StripPrefix("/media/", http.FileServer(http.Dir("media"))))

	server.ListenAndServe()
}

func (e *Env) handleRequestCategory(w http.ResponseWriter, r *http.Request) {
	var err error
	authToken := r.Header.Get("Authorization")
	isAdmin := validateToken(e.Db, authToken)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	switch r.Method {
	case "GET":
		err = handleGetCategory(w, r, e.Db)
	case "POST":
		if isAdmin {
			err = handlePostCategory(w, r, e.Db)
		} else {
			http.Error(w, "You don't have permission", http.StatusUnauthorized)
		}
	case "PUT":
		err = handlePutCategory(w, r, e.Db)
	case "DELETE":
		err = handleDeleteCategory(w, r, e.Db)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleGetCategory(w http.ResponseWriter, r *http.Request, db *sql.DB) (err error) {
	category, err := retrieveCategories(db)
	if err != nil {
		return
	}
	output, err := json.MarshalIndent(&category, "", "\t")
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

func handlePostCategory(w http.ResponseWriter, r *http.Request, db *sql.DB) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	category := Category{}
	json.Unmarshal(body, &category)
	err = category.create(db)
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handlePutCategory(w http.ResponseWriter, r *http.Request, db *sql.DB) (err error) {
	slug := path.Base(r.URL.Path)
	if err != nil {
		return
	}
	category, err := retrieveCategory(db, slug)
	if err != nil {
		return
	}
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, &category)
	err = category.update(db)
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handleDeleteCategory(w http.ResponseWriter, r *http.Request, db *sql.DB) (err error) {
	slug := path.Base(r.URL.Path)
	if err != nil {
		return
	}
	category, err := retrieveCategory(db, slug)
	if err != nil {
		return
	}
	err = category.delete(db)
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func (e *Env) handleRequestSubCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var err error
	switch r.Method {
	case "GET":
		err = handleGetSubCategory(w, r, e.Db)
	case "POST":
		err = handlePostSubCategory(w, r, e.Db)
	case "PUT":
		err = handlePutSubCategory(w, r, e.Db)
	case "DELETE":
		err = handleDeleteSubCategory(w, r, e.Db)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleGetSubCategory(w http.ResponseWriter, r *http.Request, db *sql.DB) (err error) {
	subCategories, err := retrieveSubCategories(db, r.URL.Query())
	if err != nil {
		return
	}
	output, err := json.MarshalIndent(&subCategories, "", "\t")
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

func handlePostSubCategory(w http.ResponseWriter, r *http.Request, db *sql.DB) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	subCategory := SubCategory{}
	json.Unmarshal(body, &subCategory)
	err = subCategory.create(db)
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handlePutSubCategory(w http.ResponseWriter, r *http.Request, db *sql.DB) (err error) {
	slug := path.Base(r.URL.Path)
	if err != nil {
		return
	}
	subCategory, err := retrieveSubCategory(db, slug)
	if err != nil {
		return
	}
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, &subCategory)
	err = subCategory.update(db)
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handleDeleteSubCategory(w http.ResponseWriter, r *http.Request, db *sql.DB) (err error) {
	slug := path.Base(r.URL.Path)
	if err != nil {
		return
	}
	subCategory, err := retrieveSubCategory(db, slug)
	if err != nil {
		return
	}
	err = subCategory.delete(db)
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func (e *Env) handleRequestPosts(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	switch r.Method {
	case "GET":
		slug := path.Base(r.URL.Path)
		if slug == "posts" {
			err = handleGetPosts(w, r, e.Db)
		} else {
			err = handleGetPost(w, r, e.Db, slug)
		}
	case "POST":
		err = handlePostPost(w, r, e.Db)
	case "PUT":
		err = handlePutPost(w, r, e.Db)
	case "DELETE":
		err = handleDeletePost(w, r, e.Db)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleGetPosts(w http.ResponseWriter, r *http.Request, db *sql.DB) (err error) {
	posts, err := retrievePosts(db, r.URL.Query())
	if err != nil {
		return err
	}
	output, err := json.MarshalIndent(&posts, "", "\t")
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

func handleGetPost(w http.ResponseWriter, r *http.Request, db *sql.DB, slug string) (err error) {
	post, err := retrievePost(db, slug)
	if err != nil {
		return
	}
	output, err := json.MarshalIndent(&post, "", "\t")
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

func handlePostPost(w http.ResponseWriter, r *http.Request, db *sql.DB) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	post := Post{}
	json.Unmarshal(body, &post)
	err = post.create(db)
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handlePutPost(w http.ResponseWriter, r *http.Request, db *sql.DB) (err error) {
	slug := path.Base(r.URL.Path)
	if err != nil {
		return
	}
	post, err := retrievePost(db, slug)
	if err != nil {
		return
	}
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, &post)
	err = post.update(db)
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handleDeletePost(w http.ResponseWriter, r *http.Request, db *sql.DB) (err error) {
	slug := path.Base(r.URL.Path)
	if err != nil {
		return
	}
	post, err := retrievePost(db, slug)
	if err != nil {
		return
	}
	err = post.delete(db)
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}
