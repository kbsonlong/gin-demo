/*
 * @Author: kbsonlong kbsonlong@gmail.com
 * @Date: 2023-05-09 14:27:20
 * @LastEditors: kbsonlong kbsonlong@gmail.com
 * @LastEditTime: 2023-05-09 16:47:10
 * @FilePath: /mysql-operator/Users/zengshenglong/Code/GoWorkSpace/go-demo/gin-demo/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	_ "embed"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/kbsonlong/gin-demo/internal/router"
	"github.com/kbsonlong/gin-demo/pkg/core"
	"github.com/kbsonlong/gin-demo/pkg/global"
)

// 获取环境变量信息
func GetEnvDefault(key, defVal string) string {
	val, ex := os.LookupEnv(key)
	if !ex {
		return defVal
	}
	return val
}

func main() {
	config_file := GetEnvDefault("CONFIG_FILE", "conf/config.yaml")
	global.CFG = core.Viper(config_file)
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
