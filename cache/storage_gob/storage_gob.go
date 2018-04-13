//Save, Load from to gob file
package storage_gob

import (
	"encoding/gob"
	"log"
	"os"
)

//Save object to file from path
func Save(path string, object interface{}) error {
	file, err := os.Create(path)
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(object)
	}
	file.Close()
	log.Println("storage_gob Save err=", err)
	return err
}

// Decode Gob file
func Load(path string, object interface{}) error {
	file, err := os.Open(path)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(object)
	}
	file.Close()
	log.Println("storage_gob Load err=", err)
	return err
}

//Delete file
func DeleteFile(path string) error {
	//err := os.Remove(path)
	return os.Remove(path)
}

//func check(e error) {
//	if e != nil {
//		_, file, line, _ := runtime.Caller(1)
//		log.Println(line, "\t", file, "\n", e)
//	}
//}
