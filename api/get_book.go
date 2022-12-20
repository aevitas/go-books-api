package api

import (
	"net/http"

	"aevitas.dev/go-books/pkg"
	"github.com/gin-gonic/gin"
)

func (s *Server) HandleGetByISBN(ctx *gin.Context) {
	var book pkg.Book

	isbn := ctx.Param("isbn")

	ret := s.DB.First(&book, "isbn = ?", isbn)

	if ret.RowsAffected == 0 {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	if ret.Error != nil {
		ctx.AbortWithError(http.StatusBadRequest, ret.Error)
		return
	}

	ctx.JSON(http.StatusOK, book)
}
