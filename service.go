package main

import (
	"errors"
	"log"
	"strconv"
	"strings"
)

type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

type stringService struct{}

type BookService interface {
	Books() map[string]Book
	GetBook(int)  (Book, error)
	SetBook(Book) error
}
// bookService is a concrete implementation of BookService
type bookService struct{}

type Book struct {
	ID int
	Title string
	Author string
	Date string
	Publisher string
}
var bookmap map[string]Book  = make(map[string]Book)


func (stringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (stringService) Count(s string) int {
	return len(s)
}


func (bookService) Books() map[string]Book {
	if len(bookmap) == 0 {
		log.Println("Init Books")
		bookmap["1"]=Book{
			ID:1,
			Title:"A Spell for Chameleon",
			Author:"Piers Anthony",
			Date: "1977",
			Publisher:" Del Rey ",
		}
		bookmap["2"]=Book{
			ID:        2,
			Title:     "The Source of Magic",
			Author:    "Piers Anthony",
			Date:      "1979",
			Publisher: " Del Rey ",
		}
	}
	return bookmap
}

func (bookService) SetBook( book Book) error {
	id  := strconv.Itoa(book.ID)
	bookmap[id] = book
	return nil
}

func (bookService) GetBook(id int) (Book, error) {
	strid := strconv.Itoa(id)
	//Book = bookmap[id]
	book, _ := bookmap[strid]
	return book, nil
}


// ErrEmpty is returned when input string is empty
var ErrEmpty = errors.New("Empty string")

