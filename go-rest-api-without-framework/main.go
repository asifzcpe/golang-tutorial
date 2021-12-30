package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	book1 := Book{"1010", "Algorith and Data Structure", "45678", 56}
	book2 := Book{"2010", "Programming with Golang", "12345", 656}
	books = append(books, book1, book2)
	http.HandleFunc("/api/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Test api"}`))
	})

	http.HandleFunc("/api/books", getBooks)
	http.HandleFunc("/api/create-book", createBook)
	http.HandleFunc("/api/get-book/{id}", getBook)
	err := http.ListenAndServe(":7090", nil)

	if err != nil {
		log.Fatal(err)
	}

}
