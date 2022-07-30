package store

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"web_map/model"

	"gopkg.in/yaml.v3"
)

//FileStore .
type FileStore struct {
	Data map[string]model.WebInfo
	Path string
}

//Init .
func InitFileStore(path string) *FileStore {
	store := &FileStore{
		Path: path,
		Data: make(map[string]model.WebInfo),
	}
	info, err := os.Stat(path)

	if os.IsNotExist(err) {

	} else {
		if info.IsDir() {
			store.Path = filepath.Join(store.Path, "cache.info")
		} else {
			file, err := os.Open(store.Path)
			if err != nil {
				panic(err)
			}
			defer file.Close()
			yamlFile, err := ioutil.ReadFile(path)
			if err != nil {
				panic(fmt.Sprintf("failed to read config file: %s, with error: %+v", path, err.Error()))
			}
			data := []model.WebInfo{}
			err = yaml.Unmarshal(yamlFile, data)
			if err != nil {
				panic(err)
			}
			for _, d := range data {
				store.Data[d.ID] = d
			}
		}
	}
	return store
}

//Insert .
func (store *FileStore) Insert(key string, value model.WebInfo) error {
	store.Data[key] = value
	store.Dump2File()
	return nil
}

//Delete .
func (store *FileStore) Delete(key string) error {
	delete(store.Data, key)
	store.Dump2File()
	return nil
}

//Get .
func (store *FileStore) Get(key string) (model.WebInfo, error) {
	return model.WebInfo{}, nil
}

//GetAll .
func (store *FileStore) GetAll() ([]model.WebInfo, error) {
	data := []model.WebInfo{}
	for _, v := range store.Data {
		data = append(data, v)
	}
	return data, nil
}

func (store *FileStore) Dump2File() {
	dumpData := []model.WebInfo{}
	for _, v := range store.Data {
		dumpData = append(dumpData, v)
	}
	v, _ := yaml.Marshal(dumpData)
	file, err := os.OpenFile(store.Path, os.O_TRUNC, 066)
	if err != nil {
		fmt.Println(err)
		return
	}
	file.Write(v)
}
