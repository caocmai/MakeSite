package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

type Post struct {
	Content string
}

func createHTMLFile(name string) {
	fileContent, err := ioutil.ReadFile(name + ".txt")
	if err != nil {
		panic(err)
	}
	// fmt.Print(string(fileContent))

	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

	text := Post{Content: string(fileContent)}
	f, err := os.Create(name + ".html")
	err = t.Execute(f, text)

	if err != nil {
		panic(err)
	}
}

func createHTMLsFromDir() {
	directory := "."
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if strings.Contains(file.Name(), "txt") {
			var fileName string = file.Name()
			fileName = strings.TrimSuffix(fileName, ".txt")
			createHTMLFile(fileName)
		}
	}
}

func main() {
	// var fileName string
	// flag.StringVar(&fileName, "file", "defaultValue", "Sets flag to the file")
	// flag.Parse()
	// // fmt.Print(fileName)

	// createHTMLFile(fileName)
	createHTMLsFromDir()

}
