package main

import (
	"fmt"
	"net/http"
)

func main() {
	err := errTest.Unauthorized.New("ある認証の処理内で返されたエラー")
	fmt.Println(statusCode(err)) // 401
}

func statusCode(err error) int {
	switch GetType(err) {
	case errTest.ConnectionFailed:
		return http.StatusInternalServerError // 500
	case Unauthorized:
		return http.StatusUnauthorized // 401
	default:
		return http.StatusBadRequest // 400
	}
}
