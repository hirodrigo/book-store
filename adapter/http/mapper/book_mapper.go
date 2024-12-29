package mapper

import "bookstore/domain"

func ToDomainBook(bookProperties map[string]interface{}) *domain.Book {
	return &domain.Book{
		Id:          bookProperties["book_id"].(string),
		ReleaseYear: bookProperties["release_year"].(string),
		BType:       domain.FromString(bookProperties["book_type"].(string)),
		Info:        bookProperties,
	}
}
