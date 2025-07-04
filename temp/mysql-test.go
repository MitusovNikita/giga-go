package main

// packages
import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Подключение к БД: username:password@tcp(host:port)/dbname
	dsn := "myuser:mypassword@tcp(127.0.0.1:3306)/mydb"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Ошибка подключения:", err)
	}
	defer db.Close()

	// Проверим соединение
	if err := db.Ping(); err != nil {
		log.Fatal("База не отвечает:", err)
	}

	fmt.Println("✅ Успешное подключение к базе")

	// Выполним SELECT-запрос
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		log.Fatal("Ошибка запроса:", err)
	}
	defer rows.Close()

	// Обход результатов
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal("Ошибка чтения строки:", err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}

	// Проверим ошибки при итерации
	if err := rows.Err(); err != nil {
		log.Fatal("Ошибка при чтении результатов:", err)
	}
}
