package store

import "web_map/model"

//StoreInstance .
var StoreInstance Store

//Store .
type Store interface {
	Insert(key string, value model.WebInfo) error
	Delete(key string) error
	Get(key string) (*model.WebInfo, error)
	GetAll() ([]model.WebInfo, error)
}
