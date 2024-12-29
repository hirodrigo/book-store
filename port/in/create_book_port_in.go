package port

import (
	"bookstore/domain"
)

type CreateBookPortIn interface {
	CreateBook(b *domain.Book) (*domain.Book, error)
}
