// kcores.com index builder
// 

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"encoding/json"
	"encoding/hex"
	"crypto/md5"
	"os"
	"image"
	_ "image/jpeg"
    _ "image/png"
    _ "image/gif"
)

const (
	CONTENT_BUILDER_VERSION = "0.01"
)

const (
	REPO_DIR = "../../"
	INDEX_CONTENT_FOLDER = "../database/"
	TOPIC_PAGE_FILE_NAME = "../../topics.html"
	READING_PAGE_FILE_NAME = "../../reading.html"
	LIST_PAGE_ADDR_PREFIX = "../../"
	LIST_PAGE_ADDR_SUFFIX = ".html"
)

type Topic struct {
	ToipcName       string  `json:"toipc"`
	TopicIcon       string  `json:"topicIcon"`
	TopicDesc       string  `json:"topicDesc"`
	EntryList       []Entry `json:"entryList"`
	EntryCount      int     
	ListPageId      string
	LatestEntryDate string
}

type Entry struct {
	Title  string `json:title`
	Desc   string `json:desc`
	Link   string `json:link`
	Author string `json:author`
	Date   string `json:date`
	Cover  string `json:cover`
}

type Topics []Topic

type AllEntrys []Entry

func main() {
    fmt.Printf("\nKCORES Content Builder %s\n", CONTENT_BUILDER_VERSION)
    fmt.Printf("---------------------------\n\n")

	// load database folder files
	var Topics Topics
	var AllEntrys AllEntrys
	files, _ := ioutil.ReadDir(INDEX_CONTENT_FOLDER)
    for _, f := range files {
            fmt.Printf("processing file: %s\n", f.Name())
            bytes, err := ioutil.ReadFile(INDEX_CONTENT_FOLDER+f.Name())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("string read: %d bytes\n", len(bytes))
			// process file content
			var Topic Topic
			err = json.Unmarshal(bytes , &Topic) 
    		if err != nil { 
    		    log.Fatal("json Unmarshal failed.\n", err) 
    		} 
    		// fill EntryCount & ListPageId
    		Topic.EntryCount = len(Topic.EntryList)
    		tmd5 := md5.Sum([]byte(Topic.ToipcName))
    		Topic.ListPageId = hex.EncodeToString(tmd5[:])
    		// load Topics & AllEntrys
    		Topics = append(Topics, Topic)
    		for _, v := range Topic.EntryList {
    			AllEntrys = append(AllEntrys, v)
    		}
    		// log content
    		DumpTopic(Topic)

    }

	// generate toipcs page
    GenerateTopicsPage(Topics)
	// generate reading page
	GenerateReadingPage(AllEntrys)
	// generate list page
	for _, Topic := range Topics {
		GenerateListPage(Topic)
	}

	fmt.Printf("\n\nFINISH, System exit.")
}

func getImageDimension(imagePath string) (int, int) {
    file, err := os.Open(imagePath)
    defer file.Close();
    if err != nil {
        fmt.Fprintf(os.Stderr, "%v\n", err)
    }

    img, _, err := image.DecodeConfig(file)
    if err != nil {
        fmt.Fprintf(os.Stderr, "%s: %v\n", imagePath, err)
    }
    return img.Width, img.Height
}


func DumpTopic(Topic Topic) {
	fmt.Printf("\nDump Content:\n")
    fmt.Printf("Topic.ToipcName :	%s\n", Topic.ToipcName)
	fmt.Printf("Topic.TopicIcon :	%s\n", Topic.TopicIcon)
	fmt.Printf("Topic.TopicDesc :	%s\n", Topic.TopicDesc)
	fmt.Printf("Topic.EntryCount :	%s\n", Topic.EntryCount)
	fmt.Printf("Topic.ListPageId :	%s\n", Topic.ListPageId)
	fmt.Printf("Topic.EntryList :\n")
	for k, v := range Topic.EntryList {
		fmt.Printf("    Topic.EntryList.Title: %d. %s\n", k+1, v.Title)
	}
	fmt.Printf("\n\n")
}

func GenerateTopicsPage(Topics Topics) {
    fmt.Printf("GenerateTopicsPage Start.\n")
	var topicsPage string
	// fill template
	topicsPage += TOPICS_P1
	for _, toipc := range Topics {
		topicsPage += fmt.Sprintf(TOPICS_P2, 
			toipc.ListPageId + LIST_PAGE_ADDR_SUFFIX, 
			toipc.TopicIcon, 
			toipc.ToipcName, 
			toipc.TopicDesc, 
			toipc.EntryCount)
	}
	topicsPage += TOPICS_P3
	// write file
	if _, err := os.Stat(TOPIC_PAGE_FILE_NAME); os.IsNotExist(err) {
		os.Create(TOPIC_PAGE_FILE_NAME)
	}
	f, err := os.OpenFile(TOPIC_PAGE_FILE_NAME, os.O_WRONLY|os.O_TRUNC, 0600)
	defer f.Close()
	_, err = f.WriteString(topicsPage) 
	if err != nil {
    	log.Fatal("GenerateTopicsPage failed.\n", err) 
    	return
	}
    fmt.Printf("%s generated.\n", TOPIC_PAGE_FILE_NAME)
    fmt.Printf("GenerateTopicsPage DONE.\n")
    return
}

func GenerateReadingPage(AllEntrys AllEntrys) {
	fmt.Printf("GenerateReadingPage Start.\n")
	var readingPage string
	var section1 string
	var section2 string
	var boxCounter int
	boxCounter = 1
	// fill template
	readingPage += READING_P1
	for _, entry := range AllEntrys {
		_, imgHeight := getImageDimension(REPO_DIR + entry.Cover)
		section1 += fmt.Sprintf(READING_P2, 
			boxCounter,
			imgHeight,
			entry.Cover)
		section2 += fmt.Sprintf(READING_P4,
			boxCounter,
			entry.Link,
			entry.Title,
			entry.Author,
			entry.Date)
		boxCounter ++
	}
	readingPage += section1
	readingPage += READING_P3
	readingPage += section2
	readingPage += READING_P5
	// write file
	if _, err := os.Stat(READING_PAGE_FILE_NAME); os.IsNotExist(err) {
		os.Create(READING_PAGE_FILE_NAME)
	}
	f, err := os.OpenFile(READING_PAGE_FILE_NAME, os.O_WRONLY|os.O_TRUNC, 0600)
	defer f.Close()
	_, err = f.WriteString(readingPage) 
	if err != nil {
    	log.Fatal("GenerateReadingPage failed.\n", err) 
    	return
	}
    fmt.Printf("%s generated.\n", READING_PAGE_FILE_NAME)
    fmt.Printf("GenerateReadingPage DONE.\n")
    return
}


func GenerateListPage(Topic Topic) {
	fmt.Printf("GenerateListPage Start.\n")
	var listPage string
	var section1 string
	var section2 string
	var boxCounter int
	boxCounter = 1
	// fill template
	targetFile := LIST_PAGE_ADDR_PREFIX + Topic.ListPageId + LIST_PAGE_ADDR_SUFFIX
	listPage += LIST_P1
	for _, entry := range Topic.EntryList {
		_, imgHeight := getImageDimension(REPO_DIR + entry.Cover)
		section1 += fmt.Sprintf(READING_P2, 
			boxCounter,
			imgHeight,
			entry.Cover)
		section2 += fmt.Sprintf(READING_P4,
			boxCounter,
			entry.Link,
			entry.Title,
			entry.Author,
			entry.Date)
		boxCounter ++
	}
	listPage += section1
	listPage += fmt.Sprintf(LIST_P3, 
		Topic.TopicIcon,
		Topic.ToipcName,
		Topic.TopicDesc)
	listPage += section2
	listPage += LIST_P5
	// write file
	if _, err := os.Stat(targetFile); os.IsNotExist(err) {
		os.Create(targetFile)
	}
	f, err := os.OpenFile(targetFile, os.O_WRONLY|os.O_TRUNC, 0600)
	defer f.Close()
	_, err = f.WriteString(listPage) 
	if err != nil {
    	log.Fatal("GenerateListPage failed.\n", err) 
    	return
	}
    fmt.Printf("%s generated.\n", targetFile)
    fmt.Printf("GenerateListPage DONE.\n")
    return
}