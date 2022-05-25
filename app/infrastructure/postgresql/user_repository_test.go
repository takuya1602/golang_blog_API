package postgresql

import (
	"backend/app/domain/entity"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"golang.org/x/crypto/bcrypt"
)

func TestUserRepositoryGetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "username", "password", "is_admin"}).
		AddRow(1, "testuser1", "testpass1", true).
		AddRow(2, "testuser2", "testpass2", true)

	mock.ExpectQuery(regexp.QuoteMeta("select * from users")).
		WillReturnRows(rows)

	r := NewUserRepository(db)

	users, err := r.GetAll()

	expectedUsers := []entity.User{
		{
			Id:       1,
			Name:     "testuser1",
			Password: "testpass1",
			IsAdmin:  true,
		},
		{
			Id:       2,
			Name:     "testuser2",
			Password: "testpass2",
			IsAdmin:  true,
		},
	}

	if !(reflect.DeepEqual(users, expectedUsers)) {
		t.Fatalf("Wrong content, was expecting %v, but got %v\n", expectedUsers, users)
	}
}

func TestUserRepositoryValidateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("testpass1"), 10)
	stringedHashedPassword := string(hashedPassword)

	rows := sqlmock.NewRows([]string{"id", "username", "password", "is_admin"}).
		AddRow(1, "testuser1", stringedHashedPassword, true)

	mock.ExpectQuery(regexp.QuoteMeta("select * from users where username = $1")).
		WithArgs("testuser1").
		WillReturnRows(rows)

	r := NewUserRepository(db)

	creds := entity.Credentials{
		Username: "testuser1",
		Password: "testpass1",
	}

	user, err := r.ValidateUser(creds)
	if err != nil {
		t.Fatal(err)
	}

	expectedUser := entity.User{
		Id:       1,
		Name:     "testuser1",
		Password: stringedHashedPassword,
		IsAdmin:  true,
	}

	if !(reflect.DeepEqual(user, expectedUser)) {
		t.Fatalf("Wrong content, was expecting %v, but got %v\n", expectedUser, user)
	}
}

func TestUserRepositoryCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta("insert into users (username, password) values ($1, $2)")).
		WithArgs("testuser1", "testpass1").
		WillReturnResult(sqlmock.NewResult(1, 3))

	r := NewUserRepository(db)

	user := entity.User{
		Name:     "testuser1",
		Password: "testpass1",
	}

	if err := r.Create(user); err != nil {
		t.Fatal(err)
	}
}

func TestUserRepositoryUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta("update users set username = $2, password = $3 where id = $1")).
		WithArgs(1, "testuser1", "testpass1").
		WillReturnResult(sqlmock.NewResult(1, 3))

	r := NewUserRepository(db)

	user := entity.User{
		Id:       1,
		Name:     "testuser1",
		Password: "testpass1",
	}

	if err := r.Update(user); err != nil {
		t.Fatal(err)
	}
}

func TestUserRepositoryDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta("delete from users where id = $1")).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 4))

	r := NewUserRepository(db)

	user := entity.User{
		Id:       1,
		Name:     "testuser1",
		Password: "testpass1",
	}

	if err := r.Delete(user); err != nil {
		t.Fatal(err)
	}
}

func TestUserRepositoryIsAdmin(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta("select is_admin from users where id = $1")).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"is_admin"}).AddRow(true))

	r := NewUserRepository(db)

	_, err = r.IsAdmin(1)
	if err != nil {
		t.Fatal(err)
	}
}
