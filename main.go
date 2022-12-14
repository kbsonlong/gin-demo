package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/kbsonlong/gin-demo/internal/router"
	"github.com/kbsonlong/gin-demo/pkg/core"
	"github.com/kbsonlong/gin-demo/pkg/global"
)

func main() {
	global.CFG = core.Viper()
	global.LOG = core.Zap()
	//gin.SetMode("debug")
	r := router.NewRouter()
	// r.Use(middleware.GinLogger())
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", global.CONFIG.System.HttpPort),
		Handler:        r,
		ReadTimeout:    10 * time.Millisecond,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("s.ListenAndServe err: %v", err)
	}
}
