package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title   string
	Content []byte
}

func (page *Page) save() error {
	filename := page.Title + ".txt"
	return ioutil.WriteFile(filename, page.Content, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	content, error := ioutil.ReadFile(filename)
	if error != nil {
		return nil, error
	}
	return &Page{Title: title, Content: content}, nil
}

func main() {
	p1 := &Page{Title: "index", Content: []byte("hello,this is main page.")}
	p1.save()
	p2, _ := loadPage("index")
	fmt.Println(string(p2.Content))
}
