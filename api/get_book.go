package api

import (
	"errors"
	"net/http"

	"aevitas.dev/go-books/pkg"
	"github.com/gin-gonic/gin"
)

func (s *Server) HandleGetByISBN(ctx *gin.Context) {
	var book pkg.Book

	isbn := ctx.Param("isbn")

	if len(isbn) == 0 {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("invalid isbn"))
	}

	ret := s.DB.First(&book, "isbn = ?", isbn)

	if ret.Error != nil {
		ctx.AbortWithError(http.StatusBadRequest, ret.Error)
	}

	if ret.RowsAffected == 0 {
		ctx.AbortWithError(http.StatusNotFound, nil)
	}

	ctx.JSON(http.StatusOK, book)
}
