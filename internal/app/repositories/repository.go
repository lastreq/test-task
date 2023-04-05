package repositories

import (
	"fmt"

	"github.com/lastreq/test-task/internal/app/models"
)

func (repo *BooksRepository) GetBooks() (books []models.Book, err error) {
	return repo.Books, nil
}

func (repo *BooksRepository) DeleteBook(requestedBookID models.RequestedBookID) (err error) {
	for i := range repo.Books {
		if repo.Books[i].BookID == requestedBookID.BookID {
			repo.Books = remove(repo.Books, i)

			return nil
		}
	}

	err = fmt.Errorf("delete error, book with id:%v not found", requestedBookID)

	return err
}

func remove[T any](slice []T, s int) []T {
	return append(slice[:s], slice[s+1:]...)
}
