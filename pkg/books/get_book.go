package books

import (
	"net/http"

	"github.com/benavio/Gorm_crud/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetBook(ctx *gin.Context) {
	var book models.Book
	id := ctx.Param("id")

	if result := h.DB.First(&book, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	ctx.JSON(http.StatusOK, &book)
}
