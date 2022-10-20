package main

import (
	_ "github.com/whacosta/API-Facturacion-Ecuador/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	controller "github.com/whacosta/API-Facturacion-Ecuador/src/controllers"
)

// @title           API Facturación Ecuador
// @version         1.0
// @description     Esta API, permite enviar a validar comprobantes de manera gratuita.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    https://github.com/whacosta/API-Facturacion-Ecuador/issues

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func main() {
	r := gin.Default()
	c := controller.NewController()

	v1 := r.Group("/api/v1")
	{
		roots := v1.Group("/")
		{
			roots.GET("", c.HelloWorld)
		}
		accounts := v1.Group("/invoice")
		{
			accounts.POST("", c.GenerateInvoice)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
