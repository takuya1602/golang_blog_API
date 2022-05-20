package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"golang.org/x/crypto/bcrypt"
)

func TestHandleRequestAdmin(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	e := Env{Db: db}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("testpass1"), 10)
	stringedHashedPassword := string(hashedPassword)

	row := sqlmock.NewRows([]string{"id", "username", "password", "is_admin"}).
		AddRow(1, "testuser1", stringedHashedPassword, true)

	mock.ExpectQuery(regexp.QuoteMeta("select * from users where username = $1")).
		WithArgs("testuser1").
		WillReturnRows(row)

	requestJson := strings.NewReader(`{
		"username": "testuser1",
		"password": "testpass1"
	}`)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/", e.handleRequestAdmin)

	writer := httptest.NewRecorder()

	request, err := http.NewRequest("POST", "/admin/", requestJson)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Fatalf("Response code is %v\n", writer.Code)
	}

	// delete later
	returnToken := ReturnToken{}
	json.Unmarshal(writer.Body.Bytes(), &returnToken)
}
