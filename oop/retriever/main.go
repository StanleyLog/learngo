package main

import (
	"fmt"
	"learngo/oop/retriever/real"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(utl string, form map[string]string ) string
} 

func download(r Retriever) string{
	return r.Get("http://www.baidu.com")
}

func main() {
	//var r Retriever = mock.Retriever{Content:"http://www.stanleylog.com"}
	var r Retriever = &real.Retriever{"sldkfjsldf", 40}
	fmt.Println(download(r))

	fmt.Println()
}
