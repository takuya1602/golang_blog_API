package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=gwp password=gwp dbname=go_blog sslmode=disable")
	if err != nil {
		panic(err)
	}

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "createsuperuser":
			err = createUser(db)
		case "changeusername":
			err = changeUserName(db)
		case "changePassword":
			err = changePassword(db)
		case "deleteuser":
			err = deleteUser(db)
		default:
			fmt.Printf("there is no such method: %s\n", os.Args[1])
		}
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
