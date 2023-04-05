package services

import (
	"github.com/lastreq/test-task/internal/app/models"
)

func (svc *Service) GetBooks() (books []models.Book, err error) {
	books, err = svc.repo.GetBooks()
	if err != nil {
		return books, err
	}

	return books, nil
}

func (svc *Service) DeleteBook(requestedBookID models.RequestedBookID) (err error) {
	err = svc.repo.DeleteBook(requestedBookID)
	if err != nil {
		return err
	}

	return nil
}
