package booksrepository

import (
	"awesomeProject34/internal/books/model"
	"database/sql"
	"fmt"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Get() ([]model.Book, error) {
	rows, err := r.db.Query("SELECT id, title, author, available FROM books")
	if err != nil {
		return nil, fmt.Errorf("booksRepository.Get: %w", err)
	}
	defer rows.Close()

	var books []model.Book
	for rows.Next() {
		var b model.Book
		err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.Available)
		if err != nil {
			continue
		}
		books = append(books, b)
	}

	return books, nil
}

// добавление книги
func (r *Repository) Add(req model.Book) error {
	_, err := r.db.Exec(
		"INSERT INTO books (title, author, available) VALUES ($1, $2, $3)",
		req.Title, req.Author, true,
	)
	if err != nil {
		return fmt.Errorf("booksRepository.Add: %w", err)
	}

	return nil
}
