package delivery

import (
	"github.com/gin-gonic/gin"
	"pr-test-4/books"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc books.UseCase) {
	h := NewHandler(uc)
	books := router.Group("/books")
	{
		books.GET("", h.GetBooks)
		books.DELETE("", h.Delete)
	}

}
