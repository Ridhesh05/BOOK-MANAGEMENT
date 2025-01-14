package main

import (
	"encoding/json"
	"io/ioutil"
)
func getBooks() ([]Book, error) {
	books := []Book{}
	booksByte, err := ioutil.ReadFile("./books.json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(booksByte, &books)
	if err != nil {
		return nil, err
	}
	return books, nil
}
func getBookById(id string) (Book, int, error) {
	books, err := getBooks()
	var requestedBook Book
	var requestedBookIndex int

	if err != nil {
		return Book{}, 0, err
	}

	for i, book := range books {
		if book.Id == id {
			requestedBook = book
			requestedBookIndex = i
		}
	}

	return requestedBook, requestedBookIndex, nil
}
func saveBooks(books []Book) error {

	booksBytes, err := json.Marshal(books)

	checkError(err)

	err = ioutil.WriteFile("./books.json", booksBytes, 0644)

	return err

}
