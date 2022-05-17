package main

import (
	"encoding/json"
	"net/http"
	"path"
	"time"

	_ "github.com/lib/pq"
)

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
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/categories/", handleRequestCategory)
	http.HandleFunc("/sub-categories/", handleRequestSubCategory)
	http.HandleFunc("/posts/", handleRequestPosts)
	http.Handle("/media/", http.StripPrefix("/media/", http.FileServer(http.Dir("media"))))

	server.ListenAndServe()
}

func handleRequestCategory(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	switch r.Method {
	case "GET":
		err = handleGetCategory(w, r)
	case "POST":
		err = handlePostCategory(w, r)
	case "PUT":
		err = handlePutCategory(w, r)
	case "DELETE":
		err = handleDeleteCategory(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleGetCategory(w http.ResponseWriter, r *http.Request) (err error) {
	category, err := retrieveCategories()
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

func handlePostCategory(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	category := Category{}
	json.Unmarshal(body, &category)
	err = category.create()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handlePutCategory(w http.ResponseWriter, r *http.Request) (err error) {
	slug := path.Base(r.URL.Path)
	if err != nil {
		return
	}
	category, err := retrieveCategory(slug)
	if err != nil {
		return
	}
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, &category)
	err = category.update()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handleDeleteCategory(w http.ResponseWriter, r *http.Request) (err error) {
	slug := path.Base(r.URL.Path)
	if err != nil {
		return
	}
	category, err := retrieveCategory(slug)
	if err != nil {
		return
	}
	err = category.delete()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handleRequestSubCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var err error
	switch r.Method {
	case "GET":
		err = handleGetSubCategory(w, r)
	case "POST":
		err = handlePostSubCategory(w, r)
	case "PUT":
		err = handlePutSubCategory(w, r)
	case "DELETE":
		err = handleDeleteSubCategory(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleGetSubCategory(w http.ResponseWriter, r *http.Request) (err error) {
	subCategories, err := retrieveSubCategories(r.URL.Query())
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

func handlePostSubCategory(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	subCategory := SubCategory{}
	json.Unmarshal(body, &subCategory)
	err = subCategory.create()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handlePutSubCategory(w http.ResponseWriter, r *http.Request) (err error) {
	slug := path.Base(r.URL.Path)
	if err != nil {
		return
	}
	subCategory, err := retrieveSubCategory(slug)
	if err != nil {
		return
	}
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, &subCategory)
	err = subCategory.update()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handleDeleteSubCategory(w http.ResponseWriter, r *http.Request) (err error) {
	slug := path.Base(r.URL.Path)
	if err != nil {
		return
	}
	subCategory, err := retrieveSubCategory(slug)
	if err != nil {
		return
	}
	err = subCategory.delete()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handleRequestPosts(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	switch r.Method {
	case "GET":
		slug := path.Base(r.URL.Path)
		if slug == "posts" {
			err = handleGetPosts(w, r)
		} else {
			err = handleGetPost(w, r, slug)
		}
	case "POST":
		err = handlePostPost(w, r)
	case "PUT":
		err = handlePutPost(w, r)
	case "DELETE":
		err = handleDeletePost(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleGetPosts(w http.ResponseWriter, r *http.Request) (err error) {
	posts, err := retrievePosts(r.URL.Query())
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

func handleGetPost(w http.ResponseWriter, r *http.Request, slug string) (err error) {
	post, err := retrievePost(slug)
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

func handlePostPost(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	post := Post{}
	json.Unmarshal(body, &post)
	err = post.create()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handlePutPost(w http.ResponseWriter, r *http.Request) (err error) {
	slug := path.Base(r.URL.Path)
	if err != nil {
		return
	}
	post, err := retrievePost(slug)
	if err != nil {
		return
	}
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, &post)
	err = post.update()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handleDeletePost(w http.ResponseWriter, r *http.Request) (err error) {
	slug := path.Base(r.URL.Path)
	if err != nil {
		return
	}
	post, err := retrievePost(slug)
	if err != nil {
		return
	}
	err = post.delete()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}
