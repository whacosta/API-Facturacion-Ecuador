package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HelloWorld godoc
// @Summary      Hello World Controller
// @Router       / [get]
func (c *Controller) HelloWorld(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Hello World!")
}
