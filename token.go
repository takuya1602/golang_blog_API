package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int
	Email    string
	Password string
}

type Credentials struct {
	Username string `json:"email"`
	Password string `json:"password"`
}

var validUser = User{
	Id:       10,
	Email:    "test@example.com",
	Password: "$2a$12$TBZJBBs0TfWdXHeujpGBn.TTwJq5V7Ra4yu.w9VV/Xgp9R3XS2YCq",
}

func (e *Env) handleRequestAdmin(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var creds Credentials
	json.Unmarshal(body, &creds)

	// Warning: anyone can get token with any email and password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(creds.Password), 10)

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(creds.Password))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	claims := jwt.MapClaims{
		"admin":   true,
		"user_id": validUser.Id,
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

	fmt.Printf("jwt-token: %s\n", tokenString)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(tokenString))
	return
}

func validateToken(authToken string) (isAdmin bool) {
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
		isAdmin = claims["admin"].(bool)
	} else {
		fmt.Println(err)
	}
	return
}
