/*
 * @Author: kbsonlong kbsonlong@gmail.com
 * @Date: 2023-05-09 14:27:20
 * @LastEditors: kbsonlong kbsonlong@gmail.com
 * @LastEditTime: 2023-05-09 14:39:57
 * @FilePath: /gin-demo/internal/router/router.go
 */
package router

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/kbsonlong/gin-demo/docs"
	"github.com/kbsonlong/gin-demo/internal/middleware"
	"github.com/kbsonlong/gin-demo/pkg/global"
	"github.com/penglongli/gin-metrics/ginmetrics"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	gin.SetMode(global.CONFIG.System.Mode)
	// 关闭控制台输出
	gin.DefaultWriter = ioutil.Discard
	r := gin.Default()
	// get global Monitor object
	m := ginmetrics.GetMonitor()

	// +optional set metric path, default /debug/metrics
	m.SetMetricPath("/metrics")
	// +optional set slow time, default 5s
	m.SetSlowTime(10)
	// +optional set request duration, default {0.1, 0.3, 1.2, 5, 10}
	// used to p95, p99
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})

	// set middleware for gin
	m.Use(r)
	r.Use(middleware.GinLogger())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})
	// // # e *gin.Engine
	// r.GET("metrics", PromHandler(promhttp.Handler()))

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

func PromHandler(handler http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}
