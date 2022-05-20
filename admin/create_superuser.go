package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/ssh/terminal"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

func main() {
	db, err := sql.Open("postgres", "user=gwp password=gwp dbname=go_blog sslmode=disable")
	if err != nil {
		panic(err)
	}

	fmt.Printf("username: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	username := scanner.Text()
	fmt.Printf("password: ")
	password, err := terminal.ReadPassword(0)
	if err != nil {
		fmt.Println("Invalid password")
		return
	}
	fmt.Printf("username: %T, password: %T\n", username, password)
	hashedPassword, _ := bcrypt.GenerateFromPassword(password, 10)
	stringedHashedPassword := string(hashedPassword)

	user := User{
		Username: username,
		Password: stringedHashedPassword,
		IsAdmin:  true,
	}

	err = user.create(db)
	if err != nil {
		fmt.Printf("Error creating superuser: %s\n", err)
		return
	}
}

func (user *User) create(db *sql.DB) (err error) {
	err = db.QueryRow("insert into users (username, password, is_admin) values ($1, $2, $3) returning id",
		user.Username, user.Password, user.IsAdmin).Scan(&user.Id)
	return
}
