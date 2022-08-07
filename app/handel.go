package app

import (
	"net/http"
	"strings"
	"web_map/model"
	"web_map/store"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

//Index .
func Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

//AddWeb .
func AddWeb(ctx *gin.Context) {
	item := model.WebInfo{}
	item.ID = uuid.NewV4().String()
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

	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "not find item"})
		return
	}
	store.StoreInstance.Delete(id)
	ctx.JSON(http.StatusOK, gin.H{"state": "ok"})
}

//GetWebList .
func GetWebList(ctx *gin.Context) {
	list, err := store.StoreInstance.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	relData := []model.WebInfo{}
	for _, item := range list {
		item.URL = item.IntnetURL
		if strings.HasPrefix(ctx.ClientIP(), "192.168.") {
			if item.IntranetURL != "" {
				item.URL = item.IntranetURL
			}
		}
		relData = append(relData, item)
	}
	ctx.JSON(http.StatusOK, gin.H{"error": nil, "data": relData})
}
