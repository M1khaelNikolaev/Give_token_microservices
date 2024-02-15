package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("supersecret")

func generateToken(w http.ResponseWriter, r *http.Request) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Fprintln(w, err.Error())
	}

	fmt.Fprintln(w, tokenString)
}

func handleToken(w http.ResponseWriter, r *http.Request) {
	tokenString := "" // Здесь должен быть ваш сгенерированный токен

	// Отправка токена в ответе
	fmt.Fprintln(w, tokenString)
}

func main() {
	http.HandleFunc("/generateToken", generateToken)
	http.HandleFunc("/getToken", handleToken) // Добавляем новый обработчик для получения токена
	http.ListenAndServe(":9001", nil)         // Сервер слушает на порту 9001
}
