package filelisting

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func HandlerFileList(request *http.Request, response http.ResponseWriter) (err error) {

	path := request.URL.Path[len("/list/"):]
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
