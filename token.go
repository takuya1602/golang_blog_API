package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int
	Username string
	Password string
	IsAdmin  bool
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ReturnToken struct {
	Token string `json:"token"`
}

func (e *Env) handleRequestAdmin(w http.ResponseWriter, r *http.Request) {
	db := e.Db
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var creds Credentials
	json.Unmarshal(body, &creds)

	var user User
	err := db.QueryRow("select * from users where username = $1", creds.Username).
		Scan(&user.Id, &user.Username, &user.Password, &user.IsAdmin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	claims := jwt.MapClaims{
		"user_id": user.Id,
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	err = godotenv.Load(fmt.Sprint(".env", os.Getenv("GO_ENV")))
	if err != nil {
		panic(err)
	}

	secret := os.Getenv("SECRET_KEY")

	tokenString, _ := token.SignedString([]byte(secret))

	returnToken := ReturnToken{
		Token: tokenString,
	}

	output, err := json.MarshalIndent(&returnToken, "", "\t")
	if err != nil {
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	w.WriteHeader(200)
	return
}

func validateToken(db *sql.DB, authToken string) (isAdmin bool) {
	err := godotenv.Load(fmt.Sprint(".env", os.Getenv("GO_ENV")))
	if err != nil {
		panic(err)
	}

	secret := os.Getenv("SECRET_KEY")

	token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if token == nil {
		isAdmin = false
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := int64(claims["user_id"].(float64))
		err := db.QueryRow("select is_admin from users where id = $1", userId).Scan(&isAdmin)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
	return
}
