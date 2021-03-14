package main

import (
    "fmt"
    "os"
    "path/filepath"
	"io/ioutil"
	"log"
	"io"
	"regexp"
)

func main() {
    var files []string

    root := "."
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}   
		return nil		
    })
    if err != nil {
        panic(err)
    }
    for _, file := range files {
        fmt.Println(file)
		removeText(file)
    }
}

//remove line if contains TODO
func removeText(fileName string){
	//Skip main.go file
	if(fileName!="main.go"){
		input, err := ioutil.ReadFile(fileName)
		if err != nil {
			log.Fatalln(err)
		}
		
		text := string(input)
		
		re := regexp.MustCompile(".*TODO.*\r?\n")
		lines := re.ReplaceAllString(text, "")
		
		err = WriteToFile(fileName, lines)
			if err != nil {
			log.Fatal(err)
		}
	}
}

func WriteToFile(filename string, data string) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    _, err = io.WriteString(file, data)
    if err != nil {
        return err
    }
    return file.Sync()
}