package store

import (
	"testing"
	"web_map/model"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	store := InitFileStore("./test.data")
	item := model.WebInfo{
		ID:          "1",
		Name:        "路由器",
		SortNumber:  2,
		IntnetURL:   "192.168.3.2",
		IntranetURL: "",
	}
	err := store.Insert(item.ID, item)
	assert.Equal(t, err, nil)
}
