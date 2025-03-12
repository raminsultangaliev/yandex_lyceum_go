package booksservice

import (
	"awesomeProject34/internal/books/model"
)

type Repository interface {
	Get() ([]model.Book, error)
	Add(req model.Book) error
}

type Whether interface {
	GetWhether()
}

type Service struct {
	repo    Repository
	whether Whether
}

func NewService(repo Repository, whether Whether) *Service {
	return &Service{repo: repo, whether: whether}
}

func (s *Service) Get() ([]model.Book, error) {
	s.whether.GetWhether()
	return s.repo.Get()
}

func (s *Service) Add(req model.Book) error {
	return s.repo.Add(req)
}
