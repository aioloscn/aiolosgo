package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever1 struct {
	UserAgent string
	TimeOut time.Duration
}

/**
 实现者，只要实现里面的方法
 */
func (r *Retriever1) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	result, err := httputil.DumpResponse(resp, true)
	resp.Body.Close()

	if err != nil {
		panic(err)
	}
	return string(result)
}

