package books

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := router.Group("/books")
	routes.POST("/", h.AddBook)
	routes.GET("/", h.GetBook)
	routes.GET("/", h.GetBooks)
	routes.PUT("/", h.UpdateBook)
	routes.DELETE("/", h.DeleteBook)
}
