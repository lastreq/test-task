package models

type Book struct {
	BookID        int64  `json:"book_id,omitempty"`
	Title         string `json:"title,omitempty"`
	Author        string `json:"author,omitempty"`
	PublisherYear string `json:"publisher_year,omitempty"`
}

type RequestedBookID struct {
	BookID int64 `json:"book_id,omitempty"`
}
