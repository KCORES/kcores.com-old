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
    "sort"
    "image"
    _ "image/jpeg"
    _ "image/png"
    _ "image/gif"
)

const (
    CONTENT_BUILDER_VERSION = "0.0.2"
)

const (
    REPO_DIR 					= "../../"
    INDEX_CONTENT_FOLDER 		= "../../database/list/"
    TOPIC_PAGE_FILE_NAME 		= "../../generated/topics.html"
    READING_PAGE_FILE_NAME      = "../../generated/reading.html"
    LIST_PAGE_ADDR_PREFIX 		= "../../generated/"
    LIST_PAGE_ADDR_SUFFIX 		= ".html"
    OPENSOURCE_PAGE_FILE_NAME   = "../../generated/opensource.html"
    OPENSOURCE_CONFIG 	  		= "../../database/opensource/list-page/opensource.json"
)

type Topic struct {
    ToipcName       string  `json:"toipc"`
    TopicIcon       string  `json:"topicIcon"`
    TopicDesc       string  `json:"topicDesc"`
    LongTopicDesc   string  `json:"longTopicDesc"`
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
            // sort Entry list
            sort.Slice(Topic.EntryList, func(i, j int) bool {
              return Topic.EntryList[i].Date > Topic.EntryList[j].Date
            })
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

    // generate opensource page
    GenerateOpensourcePage()

    // ok, done
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
        fmt.Fprintf(os.Stderr, "\033[31m %s: %v\033[0m\n", imagePath, err)
    }
    return img.Width, img.Height
}


func DumpTopic(Topic Topic) {
    fmt.Printf("\nDump Content:\n")
    fmt.Printf("Topic.ToipcName :	%s\n", Topic.ToipcName)
    fmt.Printf("Topic.TopicIcon :	%s\n", Topic.TopicIcon)
    fmt.Printf("Topic.TopicDesc :	%s\n", Topic.TopicDesc)
    fmt.Printf("Topic.LongTopicDesc :	%s\n", Topic.LongTopicDesc)
    fmt.Printf("Topic.EntryCount :	%d\n", Topic.EntryCount)
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
    // sort by topic name
    sort.Slice(Topics, func(i, j int) bool {
      return Topics[i].ToipcName > Topics[j].ToipcName
    })
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
    var boxCounter int
    boxCounter = 1
    // sory all entry
    sort.Slice(AllEntrys, func(i, j int) bool {
      return AllEntrys[i].Date > AllEntrys[j].Date
    })
    // fill template
    readingPage += READING_P1
    for _, entry := range AllEntrys {
        _, imgHeight := getImageDimension(REPO_DIR + entry.Cover)
        fmt.Printf("Now image height: %dpx.\n", imgHeight)
        section1 += fmt.Sprintf(READING_P2, 
            boxCounter,
            entry.Cover,
            entry.Link,
            entry.Title,
            entry.Author,
            entry.Date)
        boxCounter ++
    }
    readingPage += section1
    readingPage += READING_P3
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
    var boxCounter int
    boxCounter = 1
    // check topic desc
    topicDesc := Topic.TopicDesc
    if len(Topic.LongTopicDesc) != 0 {
        topicDesc = Topic.LongTopicDesc
    }
    // fill template
    targetFile := LIST_PAGE_ADDR_PREFIX + Topic.ListPageId + LIST_PAGE_ADDR_SUFFIX
    listPage += LIST_P1
    for _, entry := range Topic.EntryList {
        _, imgHeight := getImageDimension(REPO_DIR + entry.Cover)
        fmt.Printf("Now image height: %dpx.\n", imgHeight)
        section1 += fmt.Sprintf(LIST_P3, 
            boxCounter,
            entry.Cover,
            entry.Link,
            entry.Title,
            entry.Author,
            entry.Date)
        boxCounter ++
    }
    listPage += fmt.Sprintf(LIST_P2, 
        Topic.TopicIcon,
        Topic.ToipcName,
        topicDesc)
    listPage += section1
    listPage += LIST_P4
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


func GenerateOpensourcePage() {
    // generate opensource page
    fmt.Printf("generate opensource page: %s\n", OPENSOURCE_CONFIG)
    bytes, err := ioutil.ReadFile(OPENSOURCE_CONFIG)
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
    // sort Entry list
    sort.Slice(Topic.EntryList, func(i, j int) bool {
      return Topic.EntryList[i].Date > Topic.EntryList[j].Date
    })
    
    // log content
    DumpTopic(Topic)

    fmt.Printf("GenerateOpensourcePage Start.\n")
    var page string
    var section1 string
    var boxCounter int
    boxCounter = 1
    // check topic desc
    topicDesc := Topic.TopicDesc
    if len(Topic.LongTopicDesc) != 0 {
        topicDesc = Topic.LongTopicDesc
    }
    // fill template
    targetFile := OPENSOURCE_PAGE_FILE_NAME
    page += OPENSOURCE_P1
    for _, entry := range Topic.EntryList {
        _, imgHeight := getImageDimension(REPO_DIR + entry.Cover)
        fmt.Printf("Now image height: %dpx.\n", imgHeight)
        section1 += fmt.Sprintf(OPENSOURCE_P3, 
            boxCounter,
            entry.Cover,
            entry.Link,
            entry.Title,
            entry.Author,
            entry.Date)
        boxCounter ++
    }
    page += fmt.Sprintf(OPENSOURCE_P2, 
        Topic.TopicIcon,
        Topic.ToipcName,
        topicDesc)
    page += section1
    page += OPENSOURCE_P4
    // write file
    if _, err := os.Stat(targetFile); os.IsNotExist(err) {
        os.Create(targetFile)
    }
    f, err := os.OpenFile(targetFile, os.O_WRONLY|os.O_TRUNC, 0600)
    defer f.Close()
    _, err = f.WriteString(page) 
    if err != nil {
        log.Fatal("GenerateOpensourcePage failed.\n", err) 
        return
    }
    fmt.Printf("%s generated.\n", targetFile)
    fmt.Printf("GenerateOpensourcePage DONE.\n")
    return
}