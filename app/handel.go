package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Index .
func Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "test.html", nil)
}

//AddWeb .
func AddWeb(ctx *gin.Context) {

}

//DelWeb .
func DelWeb(ctx *gin.Context) {

}

//GetWebList .
func GetWebList(ctx *gin.Context) {

}
