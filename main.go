package main

import (
	"backend/app/common/di"
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

type Env struct {
	Db *sql.DB
}

func main() {
	db, err := sql.Open("postgres", "user=gwp password=gwp dbname=go_blog_layerArchi sslmode=disable")
	if err != nil {
		panic(err)
	}

	if len(os.Args) > 1 {
		user := di.InitUser(db)
		switch os.Args[1] {
		case "showusers":
			err = user.GetAll()
		case "createsuperuser":
			err = user.Create()
		case "updateuser":
			err = user.Update()
		case "deleteuser":
			err = user.Delete()
		default:
			fmt.Printf("there is no such method: %s\n", os.Args[1])
		}
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		server := http.Server{
			Addr: "127.0.0.1:8080",
		}

		e := Env{Db: db}

		http.HandleFunc("/categories/", e.handleRequestCategory)
		http.HandleFunc("/sub-categories/", e.handleRequestSubCategory)
		http.HandleFunc("/posts/", e.handleRequestPost)

		server.ListenAndServe()
	}
}

func (e *Env) handleRequestCategory(w http.ResponseWriter, r *http.Request) {
	var err error
	category := di.InitCategory(e.Db)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	switch r.Method {
	case "GET":
		err = category.GetAll(w, r)
	case "POST":
		err = category.Create(w, r)
	case "PUT":
		err = category.Update(w, r)
	case "DELETE":
		err = category.Delete(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (e *Env) handleRequestSubCategory(w http.ResponseWriter, r *http.Request) {
	var err error
	subCategory := di.InitSubCategory(e.Db)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	switch r.Method {
	case "GET":
		err = subCategory.GetAll(w, r)
	case "POST":
		err = subCategory.Create(w, r)
	case "PUT":
		err = subCategory.Update(w, r)
	case "DELETE":
		err = subCategory.Delete(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (e *Env) handleRequestPost(w http.ResponseWriter, r *http.Request) {
	var err error
	post := di.InitPost(e.Db)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	switch r.Method {
	case "GET":
		err = post.GetAll(w, r)
	case "POST":
		err = post.Create(w, r)
	case "PUT":
		err = post.Update(w, r)
	case "DELETE":
		err = post.Delete(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
