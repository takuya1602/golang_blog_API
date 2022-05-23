package postgresql

import (
	"backend/app/domain/entity"
	"backend/app/domain/repository"
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	*sql.DB
}

func NewUserRepository(db *sql.DB) (userRepository repository.IUserRepository) {
	userRepository = &UserRepository{db}
	return
}

func (r *UserRepository) GetAll() (users []entity.User, err error) {
	rows, err := r.Query("select * from users")
	if err != nil {
		return
	}
	for rows.Next() {
		var user entity.User
		rows.Scan(&user.Id, &user.Name, &user.Password, &user.IsAdmin)
		users = append(users, user)
	}
	return
}

func (r *UserRepository) ValidateUser(creds entity.Credentials) (user entity.User, err error) {
	err = r.QueryRow("select * from users where username = $1", creds.Username).
		Scan(&user.Id, &user.Name, &user.Password, &user.IsAdmin)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), creds.Password)
	if err != nil {
		panic(err)
	}
	return
}

func (r *UserRepository) Create(user entity.User) (err error) {
	_, err = r.Exec("insert into users (username, password) values ($1, $2)", user.Name, user.Password)
	return
}

func (r *UserRepository) Update(user entity.User) (err error) {
	_, err = r.Exec("update users set username = $2, password = $3 where id = $1",
		user.Id, user.Name, user.Password)
	return
}

func (r *UserRepository) Delete(user entity.User) (err error) {
	_, err = r.Exec("delete from users where id = $1", user.Id)
	return
}
