package controller

import (
	"bookstore/adapter/http/mapper"
	port "bookstore/port/in"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	port port.CreateBookPortIn
}

func NewBookController(port port.CreateBookPortIn) *BookController {
	return &BookController{
		port: port,
	}
}

func (b *BookController) CreateBook(context *gin.Context) {
	var bookProperties map[string]interface{}

	if err := context.ShouldBindJSON(&bookProperties); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	b.port.CreateBook(mapper.ToDomainBook(bookProperties))

	context.JSON(http.StatusCreated, bookProperties)
}
