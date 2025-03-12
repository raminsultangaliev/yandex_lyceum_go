package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// глобальные переменные для всего
var db *sql.DB
var logger *log.Logger

// структура книги
type Book struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Available bool   `json:"available"`
}

// структура для внешнего API
type Quote struct {
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	// инициализация логгера
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Ошибка открытия файла лога:", err)
	}
	logger = log.New(file, "BOOKSERVICE: ", log.Ldate|log.Ltime|log.Lshortfile)

	// подключение к базе данных
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "5432"
	}
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "postgres"
	}
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "postgres"
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "bookstore"
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	database, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Fatal("Ошибка подключения к БД:", err)
	}
	db = database

	// создание таблицы, если её нет
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS books (
            id SERIAL PRIMARY KEY,
            title VARCHAR(255) NOT NULL,
            author VARCHAR(255) NOT NULL,
            available BOOLEAN DEFAULT true
        )
    `)
	if err != nil {
		logger.Fatal("Ошибка создания таблицы:", err)
	}

	// роутинг
	r := mux.NewRouter()
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books", addBook).Methods("POST")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")
	r.HandleFunc("/books/{id}/borrow", borrowBook).Methods("PUT")
	r.HandleFunc("/books/{id}/return", returnBook).Methods("PUT")
	r.HandleFunc("/random-quote", getRandomQuote).Methods("GET")

	// запуск сервера
	logger.Println("Сервер запущен на порту 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// получение всех книг
func getBooks(w http.ResponseWriter, r *http.Request) {
	logger.Println("Запрос на получение списка книг")

	rows, err := db.Query("SELECT id, title, author, available FROM books")
	if err != nil {
		logger.Println("Ошибка получения книг:", err)
		http.Error(w, "Ошибка получения книг", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var b Book
		err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.Available)
		if err != nil {
			logger.Println("Ошибка сканирования строки:", err)
			continue
		}
		books = append(books, b)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// добавление книги
func addBook(w http.ResponseWriter, r *http.Request) {
	logger.Println("Запрос на добавление книги")

	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		logger.Println("Ошибка декодирования запроса:", err)
		http.Error(w, "Неверный формат данных", http.StatusBadRequest)
		return
	}

	err = db.QueryRow(
		"INSERT INTO books (title, author, available) VALUES ($1, $2, $3) RETURNING id",
		book.Title, book.Author, true,
	).Scan(&book.ID)

	if err != nil {
		logger.Println("Ошибка добавления книги:", err)
		http.Error(w, "Ошибка добавления книги", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

// удаление книги
func deleteBook(w http.ResponseWriter, r *http.Request) {
	logger.Println("Запрос на удаление книги")

	vars := mux.Vars(r)
	id := vars["id"]

	result, err := db.Exec("DELETE FROM books WHERE id = $1", id)
	if err != nil {
		logger.Println("Ошибка удаления книги:", err)
		http.Error(w, "Ошибка удаления книги", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Println("Ошибка получения количества затронутых строк:", err)
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Книга не найдена", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// взять книгу
func borrowBook(w http.ResponseWriter, r *http.Request) {
	logger.Println("Запрос на получение книги")

	vars := mux.Vars(r)
	id := vars["id"]

	var available bool
	err := db.QueryRow("SELECT available FROM books WHERE id = $1", id).Scan(&available)
	if err == sql.ErrNoRows {
		http.Error(w, "Книга не найдена", http.StatusNotFound)
		return
	}
	if err != nil {
		logger.Println("Ошибка проверки доступности книги:", err)
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	if !available {
		http.Error(w, "Книга уже взята", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("UPDATE books SET available = false WHERE id = $1", id)
	if err != nil {
		logger.Println("Ошибка обновления статуса книги:", err)
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// вернуть книгу
func returnBook(w http.ResponseWriter, r *http.Request) {
	logger.Println("Запрос на возврат книги")

	vars := mux.Vars(r)
	id := vars["id"]

	var available bool
	err := db.QueryRow("SELECT available FROM books WHERE id = $1", id).Scan(&available)
	if err == sql.ErrNoRows {
		http.Error(w, "Книга не найдена", http.StatusNotFound)
		return
	}
	if err != nil {
		logger.Println("Ошибка проверки доступности книги:", err)
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	if available {
		http.Error(w, "Книга уже в библиотеке", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("UPDATE books SET available = true WHERE id = $1", id)
	if err != nil {
		logger.Println("Ошибка обновления статуса книги:", err)
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// получение случайной цитаты из внешнего API
func getRandomQuote(w http.ResponseWriter, r *http.Request) {
	logger.Println("Запрос на получение случайной цитаты")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get("https://api.quotable.io/random")
	if err != nil {
		logger.Println("Ошибка получения цитаты:", err)
		http.Error(w, "Ошибка получения цитаты", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var quote Quote
	if err := json.NewDecoder(resp.Body).Decode(&quote); err != nil {
		logger.Println("Ошибка декодирования цитаты:", err)
		http.Error(w, "Ошибка обработки цитаты", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quote)
}
