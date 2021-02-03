package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
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

	// Adding flag?
	filePtr := flag.String(name, "noName", "Sets flag to the file")
	flag.Parse()
	fmt.Print(*filePtr)

	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

	text := Post{Content: string(fileContent)}
	f, err := os.Create(name + ".html")
	err = t.Execute(f, text)

	if err != nil {
		panic(err)
	}
}

func dir() {
	directory := "."
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func main() {
	createHTMLFile("latest-post")
	// dir()

}
