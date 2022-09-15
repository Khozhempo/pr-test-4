package delivery

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"pr-test-4/books"
	"pr-test-4/books/models"
	"time"
)

type Handler struct {
	useCase books.UseCase
}

func NewHandler(useCase books.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

type deleteInput struct {
	ID int `json:"book_Id"`
}

func (h *Handler) Delete(c *gin.Context) {
	inp := new(deleteInput)
	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := h.useCase.DeleteBook(c.Request.Context(), inp.ID); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

type BookResp struct {
	Id            int       `json:"id"`
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	PublisherYear time.Time `json:"publisher_year"`
}

func (h *Handler) GetBooks(c *gin.Context) {
	books, err := h.useCase.GetBooks(c.Request.Context())
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, toBooks(books))
}

func toBooks(books []*models.Book) []*BookResp {
	resp := make([]*BookResp, 0, len(books))

	for _, value := range books {
		resp = append(resp, &BookResp{
			Id:            value.Id,
			Title:         value.Title,
			Author:        value.Author,
			PublisherYear: value.PublisherYear})
	}

	return resp
}
