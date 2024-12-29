package route

import (
	controller "bookstore/adapter/http/book"

	"github.com/gin-gonic/gin"
)

type BookRoute struct {
	gin            *gin.Engine
	bookController *controller.BookController
}

func NewBookRoute(httpServer *gin.Engine, bookController *controller.BookController) *BookRoute {
	return &BookRoute{httpServer, bookController}
}

func (r *BookRoute) Setup() {
	api := r.gin.Group("/bookstore/v1")
	{
		api.POST("/books", r.bookController.CreateBook)
	}
}
