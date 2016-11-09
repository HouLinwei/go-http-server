package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
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

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "hi, I love %s", r.URL.Path[1:])
// }

func indexHandler(res http.ResponseWriter, request *http.Request) {
	title := "index"
	p, error := loadPage(title)
	if error != nil {
		p = &Page{Title: title, Content: []byte("index page")}
	}
	tpl, _ := template.ParseFiles("template/index.html")
	// tpl, _ := template.ParseGlob("./template/index.html")
	tpl.Execute(res, p)
}

func main() {
	fmt.Println(" run ")
	// p1 := &Page{Title: "index", Content: []byte("hello,this is main page.")}
	// p1.save()
	// p2, _ := loadPage("index")
	// fmt.Println(string(p2.Content))

	// http server
	// http.HandleFunc("/", handler)
	http.HandleFunc("/index/", indexHandler)
	http.ListenAndServe(":8080", nil)
}
