package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/ssh/terminal"
)

func retrieveUser(db *sql.DB) (user User, err error) {
	fmt.Printf("username: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	username := scanner.Text()
	fmt.Printf("password: ")
	password, err := terminal.ReadPassword(0)
	fmt.Printf("\n")
	if err != nil {
		fmt.Println("Invalid password")
		return
	}
	err = db.QueryRow("select * from users where username = $1", username).
		Scan(&user.Id, &user.Username, &user.Password, &user.IsAdmin)
	if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return
	}
	return
}

func createUser(db *sql.DB) (err error) {
	fmt.Printf("username: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	username := scanner.Text()
	fmt.Printf("password: ")
	password, err := terminal.ReadPassword(0)
	fmt.Printf("\n")
	if err != nil {
		fmt.Println("Invalid password")
		return
	}

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
	fmt.Println("new user has been created")
	return
}

func changeUserName(db *sql.DB) (err error) {
	user, err := retrieveUser(db)
	if err != nil {
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("new username: ")
	scanner.Scan()
	newUsername := scanner.Text()
	user.Username = newUsername
	err = user.update(db)
	if err != nil {
		return
	}
	fmt.Printf("usesrname has been changed: new username is %s\n", newUsername)
	return
}

func changePassword(db *sql.DB) (err error) {
	user, err := retrieveUser(db)
	if err != nil {
		return
	}

	fmt.Printf("new password: ")
	newPassword, err := terminal.ReadPassword(0)
	fmt.Printf("\n")
	if err != nil {
		fmt.Println("Invalid password")
		return
	}

	hashedNewPassword, _ := bcrypt.GenerateFromPassword(newPassword, 10)
	stringedHashedNewPassword := string(hashedNewPassword)

	user.Password = stringedHashedNewPassword
	err = user.update(db)
	if err != nil {
		return
	}
	fmt.Printf("%s's password has been changed\n", user.Username)
	return
}

func deleteUser(db *sql.DB) (err error) {
	user, err := retrieveUser(db)
	if err != nil {
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Are you sure? (y/n): ")
	scanner.Scan()
	switch scanner.Text() {
	case "y":
		err = user.delete(db)
		if err != nil {
			return
		}
		fmt.Println("The user has been deleted")
	case "n":
		fmt.Println("To delete user was canceled")
	}
	return
}
