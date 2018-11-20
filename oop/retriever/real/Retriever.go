package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
}

func (r *Retriever) Get(url string) string {
	if resp, err := http.Get(url); nil != err {
		panic(err)
	} else {
		response, err := httputil.DumpResponse(resp, true)
		defer resp.Body.Close()

		if nil != err {
			panic(err)
		}

		return string(response)
	}
}
