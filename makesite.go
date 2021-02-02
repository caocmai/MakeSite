package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
)

type Text struct {
	Content string
}

func main() {
	FileContents, err := ioutil.ReadFile("first-post.txt")
	if err != nil {
		panic(err)
	}

	fmt.Print(string(FileContents))

	t := template.Must(template.New("template.tmpl").ParseFiles("new.html"))

	fmt.Print(t)
	text := Text{Content: string(FileContents)}
	err = t.Execute(os.Stdout, text)
	// err = t.ExecuteTemplate(wr, "fileName", FileContents)

	// t := template.Must(template.New("template.tmpl").Parse("new.html"))
	// err = t.Execute(os.Stdout, FileContents)

	if err != nil {
		panic(err)
	}

	// bytesToWrite := []byte("new.html")
	// err = ioutil.WriteFile("new.html", bytesToWrite, 0644)
	// if err != nil {
	// 	panic(err)
	// }

}
