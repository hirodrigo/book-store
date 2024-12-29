package service

import (
	"bookstore/domain"
	port "bookstore/port/in"
)

type BookService interface {
	CreateBook(b *domain.Book) (*domain.Book, error)
}

type bookService struct {
	factory BookServiceFactory
}

func NewBookService(factory *BookServiceFactory) port.CreateBookPortIn {
	return &bookService{
		factory: *factory,
	}
}

func (s *bookService) CreateBook(b *domain.Book) (*domain.Book, error) {
	service := s.factory.GetService(b.BType)
	return service.CreateBook(b)
}
