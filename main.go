package main

import (
	"flag"
	"web_map/app"
	"web_map/store"

	"github.com/gin-gonic/gin"
)

func main() {
	filePath := ""
	flag.StringVar(&filePath, "config", "./", "cache file path")
	flag.Parse()
	store.StoreInstance = store.Init(filePath)

	r := gin.Default()
	r.LoadHTMLGlob("view/*")
	r.Static("/static/", "./static")
	r.GET("/", app.Index)
	r.Run("0.0.0.0:8888")
}
