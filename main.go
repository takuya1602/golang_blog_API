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
	Id             int             `json:"id"`
	CategoryName   string          `json:"category_name"`
	Slug           string          `json:"slug"`
	ParentCategory *ParentCategory `json:"parent_category"`
}

type Post struct {
	Id             int       `json:"id"`
	Slug           string    `json:"slug"`
	Category       *Category `json:"category"`
	Title          string    `json:"title"`
	Content        string    `json:"content"`
	EyeCatchingImg string    `json:"eye_catching_img"`
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/posts/", handleRequest)
	http.Handle("/media/", http.StripPrefix("/media/", http.FileServer(http.Dir("media"))))

	server.ListenAndServe()
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	case "PUT":
		err = handlePut(w, r)
	case "DELETE":
		err = handleDelete(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post, err := retrievePost(id)
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

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
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

func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post, err := retrievePost(id)
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

func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post, err := retrievePost(id)
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
