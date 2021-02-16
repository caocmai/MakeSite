package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/russross/blackfriday"
)

type Post struct {
	Content string
}

// func createHTMLFileFromTxt(name string) {
// 	fileContent, err := ioutil.ReadFile(name + ".txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	// fmt.Print(string(fileContent))

// 	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

// 	text := Post{Content: string(fileContent)}
// 	f, err := os.Create(name + ".html")
// 	err = t.Execute(f, text)

// 	if err != nil {
// 		panic(err)
// 	}
// }

// func createHTMLFileFromMD(name string) {
// 	fileContent, err := ioutil.ReadFile(name + ".md")
// 	if err != nil {
// 		panic(err)
// 	}
// 	// fmt.Print(string(fileContent))

// 	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
// 	output := blackfriday.MarkdownBasic(fileContent)

// 	text := Post{Content: string(output)}
// 	f, err := os.Create(name + ".html")
// 	err = t.Execute(f, text)

// 	// err = t.ExecuteTemplate(w, "page", Page{Content: template.HTML(fileContent)})

// 	if err != nil {
// 		panic(err)
// 	}
// }

// func createHTMLsFromMDInDir() {
// 	directory := "."
// 	files, err := ioutil.ReadDir(directory)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for _, file := range files {
// 		if strings.Contains(file.Name(), "md") {
// 			var fileName string = file.Name()
// 			fileName = strings.TrimSuffix(fileName, ".md")
// 			createHTMLFileFromMD(fileName)
// 		}
// 	}
// }

// func createHTMLsFromTxtFileDir() {
// 	directory := "."
// 	files, err := ioutil.ReadDir(directory)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for _, file := range files {
// 		if strings.Contains(file.Name(), "txt") {
// 			var fileName string = file.Name()
// 			fileName = strings.TrimSuffix(fileName, ".txt")
// 			createHTMLFileFromTxt(fileName)
// 		}
// 	}
// }

func createHTMLFileFromType(name string, fileExtension string) {
	fileContent, err := ioutil.ReadFile(name + "." + fileExtension)
	if err != nil {
		panic(err)
	}
	// fmt.Print(string(fileContent))

	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

	if fileExtension == "md" {
		fileContent = blackfriday.MarkdownBasic(fileContent)
	}

	text := Post{Content: string(fileContent)}
	f, err := os.Create(name + ".html")
	err = t.Execute(f, text)

	if err != nil {
		panic(err)
	}
}

func createHTMLsFromTypeInDir(fileExtension string) {
	directory := "."
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if strings.Contains(file.Name(), fileExtension) {
			var fileName string = file.Name()
			fileName = strings.TrimSuffix(fileName, "."+fileExtension)
			createHTMLFileFromType(fileName, fileExtension)
		}
	}
}

func main() {
	// var fileName string
	// flag.StringVar(&fileName, "file", "defaultValue", "Sets flag to the file")
	// flag.Parse()
	// // fmt.Print(fileName)

	// createHTMLFile(fileName)
	// createHTMLsFromTxtFileDir()
	// createHTMLsFromMDInDir()

	// refactor code so function takes in a type and creates HTML based on the type
	// for md files in dir use as follows
	createHTMLsFromTypeInDir("md")
	// for txt files in dir use as follows
	createHTMLsFromTypeInDir("txt")

}
