package repositories

import (
	"github.com/lastreq/test-task/internal/app/models"
)

type BooksRepository struct {
	Books []models.Book
}

func New() *BooksRepository {
	books := []models.Book{
		{BookID: 1, Title: "До самого рая", Author: "Ханья Янагихара", PublisherYear: "2023"},
		{BookID: 2, Title: "Прыжок", Author: "Сергей Лукьяненко", PublisherYear: "2023"},
		{BookID: 3, Title: "Люди удачи", Author: "Надифа Мохамед", PublisherYear: "2023"},
	}

	return &BooksRepository{
		Books: books,
	}
}
