package service

import (
	"bookstore/domain"
	port "bookstore/port/out"
)

type BookServiceFactory struct {
	implementations map[domain.BookType]BookService
}

func NewBookServiceFactory(adapter port.CreateBookPortOut) *BookServiceFactory {
	implementations := make(map[domain.BookType]BookService)

	implementations[domain.Fiction] = NewFictionBookService(adapter)
	implementations[domain.Horror] = NewHorrorBookService(adapter)

	return &BookServiceFactory{
		implementations: implementations,
	}
}

func (f *BookServiceFactory) GetService(bookType domain.BookType) BookService {
	return f.implementations[bookType]
}
