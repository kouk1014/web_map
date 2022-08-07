package main

import (
	"flag"
	"web_map/app"
	"web_map/store"

	"github.com/gin-gonic/gin"
)

func main() {
	filePath := ""
	flag.StringVar(&filePath, "config", "./cache.info", "cache file path")
	flag.Parse()
	store.StoreInstance = store.InitFileStore(filePath)

	r := gin.Default()
	r.Delims("{[", "]}")
	r.LoadHTMLGlob("view/*")
	r.Static("/static/", "./static")
	r.GET("/", app.Index)
	r.POST("/web", app.AddWeb)
	r.DELETE("/web/:id", app.DelWeb)
	r.GET("/web/", app.GetWebList)
	r.Run("0.0.0.0:8888")
}
