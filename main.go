package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"

	"net/http"

	"github.com/gorilla/mux"
)

// setting book model usin gthe book struct

type Book struct {
	Id     string  `json:"id"`
	Isbn   string  `json: "isbn"`
	Title  string  `json:" title"`
	Author *Author `json: "author"`
}

// author model using the author struct

type Author struct {
	Firstname string ` json: "first_name"`
	Lastname  string ` json: "last_name"`
}

// init the book slice
var books []Book

// Get the ALL books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)

}

// get a  Single book

func GetBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Tpye", "application/json")
	params := mux.Vars(r)
	for _, iteam := range books {
		if iteam.Id == params["id"] {
			json.NewEncoder(w).Encode(iteam)
			return
		}

	}
	json.NewEncoder(w).Encode(&Book{})

}

// creating a new book
func CreatBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	w.Header().Set("Content-Type", "applicatio/json")
	_ = json.NewDecoder(r.Body).Decode(books)
	book.Id = strconv.Itoa(rand.Intn(1000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)

}

// delete a book sing the unique id of the book

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, iteam := range books {
		if iteam.Id == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)

}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, iteam := range books {
		if iteam.Id == params["id"] {
			books = append(books[:index], books[:index+1]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.Id = strconv.Itoa(rand.Intn(10000))
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return

		}
	}

}

func main() {
	fmt.Println("hello welcome to te coding foethe golang programming")
	books = append(books, Book{Id: "1", Isbn: "1234", Title: "book one", Author: &Author{Firstname: "coding", Lastname: "master"}})
	books = append(books, Book{Id: "2", Isbn: "21345", Title: "book two", Author: &Author{Firstname: "javacodig", Lastname: "developer"}})
	router := mux.NewRouter()
	router.HandleFunc("/api/books", GetBooks).Methods("GET")
	router.HandleFunc("api/books/{id}", GetBook).Methods("GET")
	router.HandleFunc("/api/books/", CreatBook).Methods("POST")
	router.HandleFunc("/api/book/{id}", deleteBook).Methods("POST")
	router.HandleFunc("/api/book/{id}", updateBook).Methods("POST")

	http.ListenAndServe(":9000", router)
}
