package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	//Get file directory from user input
	fmt.Println("Enter the directory you want to clean (default-../TestFolder): ")
	var root string
	fmt.Scanln(&root)

	var files []string

	//Set default directory
	if root == "" {
		root = "../TestFolder"
	}

	//Walk walks the file tree rooted at root, calling fn for each file or directory in the tree, including root
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	//For each of the files walked through, print the file name, and call removeText func
	for _, file := range files {
		fmt.Println(file)
		removeText(file)
	}
}

//Func to remove lines if contains keyword 'TODO'
func removeText(fileName string) {
	//Skip main.go file
	if fileName != "main.go" {
		//Read file bytes from filename param, a success call return err==null, not err==EOF
		input, err := ioutil.ReadFile(fileName)
		if err != nil {
			log.Fatalln(err)
		}

		//Convert content to string
		text := string(input)

		//Replace keyword 'TODO' by regex
		re := regexp.MustCompile(".*TODO.*\r?\n")
		lines := re.ReplaceAllString(text, "")

		//Write string into a file
		err = WriteToFile(fileName, lines)
		if err != nil {
			log.Fatal(err)
		}
	}
}

//Func to write string into a file function, with filename param
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
