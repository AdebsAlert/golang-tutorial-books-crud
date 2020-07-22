package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book Model
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author model
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}

// init books var as Book struct
var books []Book

func apiStatus(w http.ResponseWriter, r *http.Request) {
	println("App is running")

}

// Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	// set header
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)

}

func getBook(w http.ResponseWriter, r *http.Request) {
	// set header
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get any params in the request URL

	// Loop through books and find by ID
	for _, book := range books {
		if book.ID == params["book"] {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})

}

func createBook(w http.ResponseWriter, r *http.Request) {
	// set header
	w.Header().Set("Content-Type", "application/json")

	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000)) // mock an ID
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {

}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	// set header
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) // get any params in the request URL

	// Loop through books and find by ID
	for index, book := range books {
		if book.ID == params["book"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)

}

func main() {
	// Init the router
	r := mux.NewRouter()

	// mock data
	books = append(books, Book{ID: "1", Isbn: "9085884", Title: "Book One", Author: &Author{Firstname: "Adebayo", Lastname: "Mustafa", Email: "adebsalert@gmail.com"}})
	books = append(books, Book{ID: "2", Isbn: "7352552", Title: "Book Two", Author: &Author{Firstname: "Segun", Lastname: "Mustafa", Email: "segsalert@gmail.com"}})
	books = append(books, Book{ID: "3", Isbn: "2886868", Title: "Book Three", Author: &Author{Firstname: "Zara", Lastname: "Mustafa", Email: "zara@gmail.com"}})
	books = append(books, Book{ID: "4", Isbn: "6886868", Title: "Book Four", Author: &Author{Firstname: "Zayn", Lastname: "Mustafa", Email: "zayn@gmail.com"}})

	r.HandleFunc("/api/status", apiStatus).Methods("GET")
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{book}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{book}", updateBook).Methods("UPDATE")
	r.HandleFunc("/api/books/{book}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}
