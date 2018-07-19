package main

import (
    "net/http"
    "math/rand"
    "strconv"
    "github.com/gorilla/mux"
    "log"
    "encoding/json"
)

//Book Struct (Model)
type Book struct {
    ID string `json: "id"`
    Isbn string `json: "isbn"`
    Title string `json: "title"`
    Author *Author `json: "author"`

}

// Author struct

type Author struct {

   Firstname string `json: "firstname"`
   Lastname string `json: "lastname"`
}

//Init books var as a slice Book struct
var  books []Book

// Get all books
func getBooks(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(books)
}

//Get book of specific id
func getBook(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")
    // Get params
    params := mux.Vars(r)
    // Loop through books and find with id
    for _, item := range books {
	if item.ID == params["id"]{
	    json.NewEncoder(w).Encode(item)
	    return
	}
    }
    json.NewEncoder(w).Encode(&Book{})
}

//create new book
func createBook(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")
    var book Book
    _ = json.NewDecoder(r.Body).Decode(&book)
    book.ID = strconv.Itoa(rand.Intn(1000000))
    books = append(books, book)
    json.NewEncoder(w).Encode(book)
}

//
func updateBook(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index,item := range books{
        if item.ID == params["id"]{
            books = append(books[:index], books[index+1:]...)
            var book Book
            _ = json.NewDecoder(r.Body).Decode(&book)
            book.ID = params["id"]
            books = append(books, book)
            json.NewEncoder(w).Encode(book)
            return
        }
    }
    json.NewEncoder(w).Encode(books)
}

// delete book from books data of specific id
func deleteBook(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index,item := range books{
        if item.ID == params["id"]{
		books = append(books[:index], books[index+1:]...)
	    break
        }
    }
    json.NewEncoder(w).Encode(books)
}

func main() {
    // Init router
    r := mux.NewRouter()

    // Mock data
    books = append(books, Book{ID: "1", Isbn: "7890", Title: "The Brief History of Time",
    Author: &Author{Firstname: "Stephen", Lastname: "Hawking"}})
    books = append(books, Book{ID: "2", Isbn: "7891", Title: "Chasing the Monsoon",
    Author: &Author{Firstname: "Alexendar", Lastname: "Frater"}})
    books = append(books, Book{ID: "3", Isbn: "7892", Title: "Parallel Worlds",
    Author: &Author{Firstname: "Michiao", Lastname: "Kaku"}})

    // Route handlers /Endpoints
    r.HandleFunc("/api/books", getBooks).Methods("GET")
    r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
    r.HandleFunc("/api/books", createBook).Methods("POST")
    r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
    r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

    log.Fatal(http.ListenAndServe(":8002", r))
}
