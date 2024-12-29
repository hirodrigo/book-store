package database

import (
	"bookstore/domain"
	port "bookstore/port/out"
	"fmt"
)

type DatabaseAdapter interface {
}

type databaseAdapter struct {
}

func NewDatabaseAdapter() port.CreateBookPortOut {
	return &databaseAdapter{}
}

func (d *databaseAdapter) CreateBook(b *domain.Book) (*domain.Book, error) {
	print(fmt.Sprintf("Inserting book into database: %v\n", b))

	return b, nil
}
