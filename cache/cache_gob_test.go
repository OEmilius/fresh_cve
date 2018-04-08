package cache

//	c := cve.Cve{
//		Id:        "test ID",
//		Published: "2018 01 01",
//		Summary:   "test summary",
//	}

import (
	"fmt"
	//	storage "fresh_cve/cache/storage_gob"
	"fresh_cve/cve"
	"os"
)

func ExampleNewCache() {
	cache := NewCache()
	fmt.Println(cache.cveMap)
	//Output:map[]
}

func ExampleCache_AddList() {
	cache := NewCache()
	c1 := cve.Cve{Id: "1", Summary: "first"}
	c2 := cve.Cve{Id: "2", Summary: "second"}
	l := []cve.Cve{c1, c2}
	cache.AddList(l)
	fmt.Println(len(cache.cveMap))
	//Output:2
}

func ExampleCache_GetAllCve() {
	cache := NewCache()
	c1 := cve.Cve{Id: "1", Summary: "first"}
	c2 := cve.Cve{Id: "2", Summary: "second"}
	l := []cve.Cve{c1, c2}
	cache.AddList(l)
	fmt.Println(len(cache.GetAllCve()))
	//Output:2
}

func Example_saveToGobFile() {
	//var fileName string = "cache.gob"
	cache := NewCache()
	cache.FileGobName = "storage.gob"
	c1 := cve.Cve{Id: "1", Summary: "first"}
	c2 := cve.Cve{Id: "2", Summary: "second"}
	l := []cve.Cve{c1, c2}
	cache.AddList(l)
	//	err := storage.Save(fileName, cache)
	err := cache.Save()
	fmt.Println("saving err=", err)
	if _, err := os.Stat(cache.FileGobName); err == nil {
		fmt.Println("file Exisists")
	}

	//Output:saving err= <nil>
	//file Exisists
}

func Example_loadFromGobFile() {
	cache := NewCache()
	cache.FileGobName = "storage.gob"
	err := cache.Load()
	fmt.Println(err)
	fmt.Println(len(cache.cveMap))
	cache.DeleteFile()
	//Output:<nil>
	//2
}

//func Example_save_load_del() {
//	var fileName string = "cache.gob"
//	cache := NewCache()
//	c1 := cve.Cve{Id: "1", Summary: "first"}
//	c2 := cve.Cve{Id: "2", Summary: "second"}
//	l := []cve.Cve{c1, c2}
//	cache.AddList(l)
//	err := storage.Save(fileName, cache)
//	fmt.Println(err)
//	err = storage.Load(fileName, cache)
//	fmt.Println(err)
//	fmt.Println("objects in=", len(cache.GetAllCve()))
//	err = storage.DeleteFile(fileName)
//	fmt.Println(err)
//	//Output:<nil>
//	//<nil>
//	//objects in= 2
//	//<nil>
//}
