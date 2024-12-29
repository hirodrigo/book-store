package service

import (
	"bookstore/domain"
	port "bookstore/port/out"
	"fmt"
)

type HorrorBookService struct {
	adapter port.CreateBookPortOut
}

func NewHorrorBookService(adapter port.CreateBookPortOut) BookService {
	return &HorrorBookService{adapter: adapter}
}

func (s *HorrorBookService) CreateBook(b *domain.Book) (*domain.Book, error) {
	print(fmt.Sprintf("Processing Horror Book: %v \n", b))
	return s.adapter.CreateBook(b)
}
