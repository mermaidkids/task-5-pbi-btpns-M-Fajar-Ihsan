package main

import (
	"fmt"
	"net/http"

	authcontroller "github.com/mermaidkids/project2/controllers"
)

func main() {
	http.HandleFunc("/", authcontroller.Index)
	http.HandleFunc("/login", authcontroller.Login)
	fmt.Println("server jalan di: http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
