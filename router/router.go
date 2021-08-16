package router

import (
	"codeql-ct/config"
	"codeql-ct/docs"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
)

type ResponsePing struct {
	Message string `json:"message" example:"pong"`
}

// Ping godoc
// @Summary ping pong
// @Description return pong
// @ID ping-pong
// @Accept  json
// @Produce  json
// @Success 200 {object} ResponsePing
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, ResponsePing{Message: "pong"})
}

func InitRouter() *gin.Engine {
	router := gin.Default()
	// cors
	if config.AllowOrigins != nil {
		corsConfig := cors.DefaultConfig()
		//corsConfig.AllowAllOrigins = true
		corsConfig.AllowCredentials = true
		corsConfig.AllowOrigins = append(corsConfig.AllowOrigins, config.AllowOrigins...)
		router.Use(cors.New(corsConfig))
	}
	// ping pong
	{
		router.GET("/ping", Ping)
	}
	// swagger
	{
		// programmatically set swagger info
		docs.SwaggerInfo.Title = "Codeql-ct API"
		docs.SwaggerInfo.Description = "Codeql Continuous Test Platform"
		docs.SwaggerInfo.Version = "1.0"
		docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", config.SwaggerHost, config.LocalListenPort)
		docs.SwaggerInfo.BasePath = "/"
		docs.SwaggerInfo.Schemes = []string{"http"}
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	return router
}
