package main

import (
	"io/ioutil"
	"os"
	"text/template"
)

type Story struct {
	Content string
}

func main() {
	fileContents, err := ioutil.ReadFile("first-post.txt")
	if err != nil {
		panic(err)
	}
	// fmt.Print(string(FileContents))

	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

	// t, err := template.New("todos").Parse("You have a task named \"{{ .Content}}\"")

	text := Story{Content: string(fileContents)}

	f, err := os.Create("first-post.html")

	err = t.Execute(f, text)

	if err != nil {
		panic(err)
	}

}
