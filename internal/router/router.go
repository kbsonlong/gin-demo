package router

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/kbsonlong/gin-demo/docs"
	"github.com/kbsonlong/gin-demo/internal/middleware"
	"github.com/kbsonlong/gin-demo/pkg/global"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	gin.SetMode(global.CONFIG.System.Mode)
	// 关闭控制台输出
	gin.DefaultWriter = ioutil.Discard
	r := gin.Default()
	r.Use(middleware.GinLogger())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})
	// docs.SwaggerInfo.BasePath = "/"
	// url := ginSwagger.URL("http://127.0.0.1:8300/swagger/doc.json")
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// apiv1 := r.Group("/v1")
	// {
	// 	apiv1.GET("ocr/result", v1.GeneralBasic)
	// }

	// auth := r.Group("/v1")
	// auth.Use(middleware.JWT())
	// {
	// 	auth.GET("/info", v1.GetUserInfo)
	// }
	return r
}
