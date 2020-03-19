package internal

import (
	"bytes"
	"fmt"
	"github.com/gammazero/workerpool"
	"net/http"
	"strings"
	"sync"
	"time"
)

const (
	timeout = time.Second * 3
)

type Svc struct {
	url        string
	concurrent int
	payload    string
	method     string
	headers []string
}

func NewSvc(num int, url, method, payload string, headers []string) Svc {
	return Svc{
		url:        url,
		concurrent: num,
		payload:    payload,
		method:     method,
		headers: headers,
	}
}

func (s Svc) Init() {
	var wg sync.WaitGroup
	wp := workerpool.New(30)

	for i := 0; i < s.concurrent; i++ {
		wg.Add(1)
		wp.Submit(func (){
			s.call(&wg, i)
		})
	}

	wg.Wait()
	wp.Stop()
}

func (s Svc) call(wg *sync.WaitGroup, i int) {
	method := s.setHttpMethod()
	payloadBuff := s.setPayload()
	req, err := http.NewRequest(method, s.url, payloadBuff)
	if err != nil {
		fmt.Println("error voi", i, err.Error())
		return
	}

	s.setHeaders(req)

	if resp := httpCall(err, req, i); resp != nil {
		respStr := parseRespToStr(err, resp, i)
		result := fmt.Sprintf("call: %d status: %d body: %s", i+1, resp.StatusCode, respStr)
		fmt.Println(result)
	}

	wg.Done()
}

func parseRespToStr(err error, resp *http.Response, i int) string {
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)

	if err != nil {
		fmt.Println("error ", i, err.Error())
	}

	respStr := buf.String()
	return respStr
}

func httpCall(err error, req *http.Request, i int) *http.Response {
	client := &http.Client{
		Timeout: timeout,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error ", i, err.Error())
	}
	return resp
}

func (s Svc) setHeaders(req *http.Request) {
	for _, header := range s.headers {
		keyValue := strings.Split(header, ":")
		req.Header.Add(keyValue[0], keyValue[1])
	}
}

func (s Svc) setPayload() *bytes.Buffer {
	if s.payload != "" {
		return bytes.NewBufferString(s.payload)
	}
	return &bytes.Buffer{}
}

func (s Svc) setHttpMethod() string {
	// set the http method safe checks
	switch strings.ToUpper(s.method) {
	case http.MethodPost:
		return http.MethodPost
	case http.MethodHead:
		return http.MethodHead
	default:
		return http.MethodGet
	}
}
