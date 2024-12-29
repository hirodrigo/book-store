package domain

type BookType string

const (
	Fiction    BookType = "Fiction"
	NonFiction BookType = "Non-Fiction"
	Science    BookType = "Science"
	Biography  BookType = "Biography"
	Horror     BookType = "Horror"
)

func FromString(value string) BookType {
	return BookType(value)
}
