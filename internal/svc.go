package internal

import (
	"bytes"
	"fmt"
	"github.com/gammazero/workerpool"
	"net/http"
	"strings"
	"sync"
)

const (
	content = "application/json"
)

type Svc struct {
	url        string
	concurrent int
	payload    string
	method     string
}

func NewSvc(url string, num int, method string, payload string) Svc {
	return Svc{
		url:        url,
		concurrent: num,
		payload:    payload,
		method:     method,
	}
}

func (s Svc) Init() {
	var wg sync.WaitGroup
	wp := workerpool.New(20)

	for i := 0; i < s.concurrent; i++ {
		wg.Add(1)
		wp.Submit(func (){
			s.call(&wg, i, s.method, s.payload)
		})
	}

	wg.Wait()
	wp.Stop()
}

func (s Svc) call(wg *sync.WaitGroup, i int, method, payload string) {
	var resp *http.Response
	var err error

	switch strings.ToUpper(method) {
	case http.MethodPost:
		payloadBuff := bytes.NewBufferString(payload)
		resp, err = http.Post(s.url, content, payloadBuff)
	case http.MethodGet:
		resp, err = http.Get(s.url)
	case http.MethodHead:
		resp, err = http.Head(s.url)
	}

	if err != nil {
		fmt.Println("error ", i, err.Error())
	}

	if resp != nil {
		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(resp.Body)

		if err != nil {
			fmt.Println("error ", i, err.Error())
		}

		respStr := buf.String()
		result := fmt.Sprintf("call: %d status: %d body: %s", i+1, resp.StatusCode, respStr)
		fmt.Println(result)
	}

	wg.Done()
}
