// kcores.com index builder
// 

package main

import (
	"fmt"
	"ioutil"
)

const (
	INDEX_CONTENT_FOLDER = "./index-builder/index-content"

)

type Entry struct {
	Title  string
	Link   string
	Author string
	Date   string
	Cover  string
}

func main() {
	// load index-content folder files
	files, _ := ioutil.ReadDir(INDEX_CONTENT_FOLDER)
    for _, f := range files {
            fmt.Println(f.Name())
    }

	// load index-content
	var Entrys []Entry
	err := json.Unmarshal(jsonBlob , &Entrys) 
    if err != nil { 
        fatal("json Unmarshal failed.") 
    } 
	// generate content
	
	// write to local 
}