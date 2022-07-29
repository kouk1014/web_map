package store

import "web_map/model"

//FileStore .
type FileStore struct {
	Data map[string]string
	Path string
}

//Init .
func Init(path string) *FileStore {
	store := &FileStore{
		Path: path,
		Data: make(map[string]string),
	}

	return store
}

//Insert .
func (file *FileStore) Insert(key string, value model.WebInfo) error {
	return nil
}

//Delete .
func (file *FileStore) Delete(key string) error {
	return nil
}

//Get .
func (file *FileStore) Get(key string) (model.WebInfo, error) {
	return model.WebInfo{}, nil
}

//GetAll .
func (file *FileStore) GetAll() ([]model.WebInfo, error) {
	return nil, nil
}
