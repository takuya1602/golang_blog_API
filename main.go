package main

import (
	"encoding/json"
	"net/http"
	"path"
	"strconv"

	_ "github.com/lib/pq"
)

type ParentCategory struct {
	Id           int    `json:"id"`
	CategoryName string `json:"category_name"`
	Slug         string `json:"slug"`
}

type Category struct {
	Id               int    `json:"id"`
	CategoryName     string `json:"category_name"`
	Slug             string `json:"slug"`
	ParentCategoryId int    `json:"parent_category_id"`
}

type Post struct {
	Id             int    `json:"id"`
	Slug           string `json:"slug"`
	CategoryId     int    `json:"category_id"`
	Title          string `json:"title"`
	Content        string `json:"content"`
	EyeCatchingImg string `json:"eye_catching_img"`
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/parent-categories/", handleRequestParentCategory)
	http.HandleFunc("/categories/", handleRequestCategory)
	http.HandleFunc("/posts/", handleRequestPosts)
	http.Handle("/media/", http.StripPrefix("/media/", http.FileServer(http.Dir("media"))))

	server.ListenAndServe()
}

func handleRequestParentCategory(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	switch r.Method {
	case "GET":
		err = handleGetParentCategory(w, r)
	case "POST":
		err = handlePostParentCategory(w, r)
	case "PUT":
		err = handlePutParentCategory(w, r)
	case "DELETE":
		err = handleDeleteParentCategory(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleGetParentCategory(w http.ResponseWriter, r *http.Request) (err error) {
	parentCategory, err := retrieveParentCategories()
	if err != nil {
		return
	}
	output, err := json.MarshalIndent(&parentCategory, "", "\t")
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

func handlePostParentCategory(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	parentCategory := ParentCategory{}
	json.Unmarshal(body, &parentCategory)
	err = parentCategory.create()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handlePutParentCategory(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	parentCategory, err := retrieveParentCategory(id)
	if err != nil {
		return
	}
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, &parentCategory)
	err = parentCategory.update()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handleDeleteParentCategory(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	parentCategory, err := retrieveParentCategory(id)
	if err != nil {
		return
	}
	err = parentCategory.delete()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handleRequestCategory(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	switch r.Method {
	case "GET":
		slug := path.Base(r.URL.Path)
		if slug == "categories" {
			err = handleGetCategories(w, r)
		} else {
			err = handleGetCategoryPosts(w, r, slug)
		}
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

func handleGetCategories(w http.ResponseWriter, r *http.Request) (err error) {
	categories, err := retrieveCategories()
	if err != nil {
		return
	}
	output, err := json.MarshalIndent(&categories, "", "\t")
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

func handleGetCategoryPosts(w http.ResponseWriter, r *http.Request, slug string) (err error) {
	posts, err := retrieveCategoryPosts(slug)
	if err != nil {
		return
	}
	output, err := json.MarshalIndent(&posts, "", "\t")
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
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	category, err := retrieveCategory(id)
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
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	category, err := retrieveCategory(id)
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
	posts, err := retrievePosts()
	if err != nil {
		return
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
