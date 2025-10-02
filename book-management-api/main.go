package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"sync"
)

type Book struct {
	ID              string `json:"id"`
	Title           string `json:"title"`
	Author          string `json:"author"`
	PublicationYear int    `json:"publication_year"`
}

var (
	books = make(map[string]Book)
	mutex = sync.Mutex{}

)

func main() {


	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getBooks(w, r)
		case http.MethodPost:
			createBook(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	
	http.HandleFunc("/books/", func(w http.ResponseWriter, r *http.Request) {	
		switch r.Method {
		case http.MethodGet:
			getBook(w, r)
		case http.MethodPut:
			udpateBook(w, r)
		case http.MethodDelete:
			deleteBook(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil)
}



func getBooks(w http.ResponseWriter, r *http.Request)  {
	mutex.Lock() // Lock the mutex to ensure thread-safe access to the books map
	defer mutex.Unlock() // Ensure mutex is unlocked after function completes
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// get book by id 
func getBook(w http.ResponseWriter, r *http.Request)  {
	mutex.Lock()
	defer mutex.Unlock()
	idStr := strings.TrimPrefix(r.URL.Path, "/books/")

	book, exists := books[idStr]
	if !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)

}

func createBook(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	w.Header().Set("Content-Type", "application/json")
	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	
	books[book.ID] = book
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}


func udpateBook(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	w.Header().Set("Content-Type", "application/json")
	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	book, exists := books[idStr]
	if !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	books[idStr] = book
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	w.Header().Set("Content-Type", "application/json")
	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	_, exists := books[idStr]
	if !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	delete(books, idStr)
	w.WriteHeader(http.StatusNoContent)
}

