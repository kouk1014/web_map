package app

import (
	"net/http"
	"web_map/model"
	"web_map/store"

	"github.com/gin-gonic/gin"
)

//Index .
func Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "test.html", nil)
}

//AddWeb .
func AddWeb(ctx *gin.Context) {
	item := model.WebInfo{}
	err := ctx.ShouldBindJSON(&item)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err2 := store.StoreInstance.Insert(item.ID, item)
	if err2 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
	}
}

//DelWeb .
func DelWeb(ctx *gin.Context) {

}

//GetWebList .
func GetWebList(ctx *gin.Context) {

}
