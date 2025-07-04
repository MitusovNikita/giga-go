// docker compose up -d --build
package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World")
}

func main() {
	http.HandleFunc("/", helloHandler)        // Регистрируем обработчик по пути "/"
	fmt.Println("Server is running on :8080") // Просто сообщение в консоль
	http.ListenAndServe(":8080", nil)         // Запускаем сервер на порту 8080
}
