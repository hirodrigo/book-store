package service

import (
	"bookstore/domain"
	port "bookstore/port/out"
	"fmt"
)

type FictionBookService struct {
	adapter port.CreateBookPortOut
}

func NewFictionBookService(adapter port.CreateBookPortOut) BookService {
	return &FictionBookService{adapter: adapter}
}

func (s *FictionBookService) CreateBook(b *domain.Book) (*domain.Book, error) {
	print(fmt.Sprintf("Processing Fiction Book: %v \n", b))
	return s.adapter.CreateBook(b)
}
