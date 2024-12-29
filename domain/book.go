package domain

type Book struct {
	Id          string
	ReleaseYear string
	BType       BookType
	Info        map[string]interface{}
}
