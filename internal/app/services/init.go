package services

import (
	"github.com/lastreq/test-task/internal/app/models"
)

type Service struct {
	repo repository
}

type repository interface {
	GetBooks() (books []models.Book, err error)
	DeleteBook(requestedBookID models.RequestedBookID) (err error)
}

func New(repo repository) *Service {
	svc := Service{repo: repo}

	return &svc
}
