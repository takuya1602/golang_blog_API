package main

import (
	"database/sql"
)

type User struct {
	Id       int
	Username string
	Password string
	IsAdmin  bool
}

func (user *User) create(db *sql.DB) (err error) {
	err = db.QueryRow("insert into users (username, password, is_admin) values ($1, $2, $3) returning id",
		user.Username, user.Password, user.IsAdmin).Scan(&user.Id)
	return
}

func (user *User) update(db *sql.DB) (err error) {
	_, err = db.Exec("update users set username = $2, password = $3 where id = $1",
		user.Id, user.Username, user.Password)
	return
}

func (user *User) delete(db *sql.DB) (err error) {
	_, err = db.Exec("delete from users where id = $1", user.Id)
	return
}
