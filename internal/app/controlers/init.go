package controlers

import (
	"github.com/lastreq/test-task/internal/app/models"
)

type Controller struct {
	svc service
}
type service interface {
	GetBooks() (books []models.Book, err error)
	DeleteBook(requestedBookID models.RequestedBookID) (err error)
}

func New(svc service) *Controller {
	ctr := &Controller{svc: svc}

	return ctr
}
