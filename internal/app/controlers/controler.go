package controlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/lastreq/test-task/internal/app/models"
)

func (ctr *Controller) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := ctr.svc.GetBooks()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	booksJSONBytes, err := json.Marshal(books)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(booksJSONBytes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)

		return
	}

	return
}

func (ctr *Controller) DeleteBook(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	var requestedBookID models.RequestedBookID

	err = json.Unmarshal(bodyBytes, &requestedBookID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	err = ctr.svc.DeleteBook(requestedBookID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	return
}
