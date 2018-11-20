package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type httpHandler func(http.ResponseWriter, *http.Request)

type appHandler func(http.ResponseWriter, *http.Request) error

type userError interface {
	error
	Messager() string
}

type customerError string

func (err customerError) Error() string {
	return err.Messager()
}

func (err customerError) Messager() string  {
	return string(err)
}

func errWrapper(handler appHandler) httpHandler {

	return func(writer http.ResponseWriter, request *http.Request) {

		log.Println("新进会话")

		defer func() {
			r := recover()

			if nil != r {
				if err, ok := r.(error); ok {
					log.Println("发生错误：", err)
				} else {
					log.Println("未知异常：", r)
				}

				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}

		}()

		err := handler(writer, request)
		if err != nil {
			log.Printf("Error Handling request, %s", err.Error())
			code := http.StatusOK


			if err, ok := err.(userError) ; ok {
				http.Error(writer, err.Messager(), http.StatusBadRequest)
				return
			}

			switch {

			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}

			http.Error(writer, http.StatusText(code), code)
			//fmt.Println(code)
		}
	}

}

const prefix  = "/list/"

func HandlerFileList(response http.ResponseWriter, request *http.Request) (err error) {

	log.Println("访问路径：", request.URL.Path)

	if strings.Index(request.URL.Path, prefix) != 0 {
		//return errors.New("the path is must start with " + prefix)
		return customerError("the path is must start with " + prefix)
	}

	path := request.URL.Path[len(prefix):]
	fmt.Println(path)

	file, err := os.Open(path)
	if err != nil {
		//http.Error(response, err.Error(), http.StatusInternalServerError)
		//return
		return err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		//panic(err)
		return err
	}

	response.Write(bytes)

	return nil
}

func FuncHandler(writer http.ResponseWriter, request *http.Request) {

	fmt.Println("进入handler函数")
	b, err := ioutil.ReadAll(strings.NewReader("页面首页"))
	fmt.Println(b)

	if err == nil {
		writer.Write(b)
	}

}


func main() {


	http.HandleFunc("/", errWrapper(HandlerFileList))

	//http.HandleFunc("/list/", FuncHandler)
	//http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	fmt.Println("进入handler函数")
	//})

	fmt.Println("程序启动")
	err := http.ListenAndServe(":8088", nil)

	if err != nil {
		panic(err)
	}
}

