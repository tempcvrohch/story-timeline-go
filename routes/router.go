package routes

import (
	"fmt"

	"eksrow.com/story-timeline-go/global"
	"eksrow.com/story-timeline-go/services"
	"github.com/gin-gonic/gin"
)

const InterfaceAdress = "0.0.0.0"

func New(config *global.Config) *gin.Engine {
	r := gin.New()

	ctx := global.Context{Config: config, Router: r}
	bookService := services.BookService{Context: &ctx}

	r.POST("/book", bookService.CreateBook)

	r.Run(fmt.Sprintf("%s:%d", InterfaceAdress, config.Server.Port))

	return r
}
