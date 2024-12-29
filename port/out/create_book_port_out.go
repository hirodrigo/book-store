package port

import (
	"bookstore/domain"
)

type CreateBookPortOut interface {
	CreateBook(b *domain.Book) (*domain.Book, error)
}
