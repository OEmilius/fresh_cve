package storage_gob

import (
	"fmt"
	"os"
)

var fileName = "cache.gob"

func ExampleSave() {
	var savemap = make(map[string]int)
	savemap["1"] = 1
	savemap["2"] = 2

	err := Save(fileName, savemap)
	if _, err := os.Stat(fileName); err == nil {
		fmt.Println("file Exisists")
	}
	_ = err
	//Output:file Exisists
}

func ExampleLoad() {
	var workmap map[string]int
	err := Load(fileName, &workmap)
	fmt.Println(len(workmap))
	fmt.Println(err)
	//Output: 2
	//<nil>
}

func ExampleDeleFile() {
	fmt.Println(DeleteFile(fileName))
	//Output:<nil>

}
