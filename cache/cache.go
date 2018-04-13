//Memory cashe implemented using simply map
package cache

import (
	"encoding/json"
	storage "fresh_cve/cache/storage_gob"
	"log"

	"github.com/OEmilius/fresh_cve/cve"
)

//TODO: RACE подумать ОПАСНОСТЬ добавить mutexы
//TODO: избежать гонки использовать sync.map
type Cache struct {
	FileGobName string
	cveMap      map[string]cve.Cve
}

func NewCache() *Cache {
	return &Cache{cveMap: make(map[string]cve.Cve)}
}

func (cache *Cache) AddList(list []cve.Cve) {
	for _, cve := range list {
		cache.cveMap[cve.Id] = cve
	}
}

func (cache *Cache) GetAllCve() (list []cve.Cve) {
	for _, cve := range cache.cveMap {
		list = append(list, cve)
	}
	return list
}

func (cache *Cache) GetAllcveJson() string {
	b, err := json.Marshal(cache.GetAllCve())
	if err != nil {
		log.Println("cache error in GetAllcveJson", err)
	}
	return string(b)
}

func (cache *Cache) Save() error {
	log.Println("saving cache")
	return storage.Save(cache.FileGobName, cache.cveMap)
}

func (cache *Cache) Load() error {
	log.Println("loading cache")
	var workmap map[string]cve.Cve
	if err := storage.Load(cache.FileGobName, &workmap); err == nil {
		cache.cveMap = workmap
		return nil
	} else {
		return err
	}

}

func (cache *Cache) DeleteFile() error {
	return storage.DeleteFile(cache.FileGobName)
}
