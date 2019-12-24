package main

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
)

const url = "https://api.stage.voiapp.io/v1/auth/verify/code"
const content = "application/json"

const concurrent = 750

func main() {
	fmt.Println("Hello scripter")
	var wg sync.WaitGroup

	for i := 0; i < concurrent; i++ {
		wg.Add(1)
		go request(i, &wg)
	}

	wg.Wait()

	fmt.Println("done")
}

func request(i int, wg *sync.WaitGroup) {
	payload := bytes.NewBufferString(`{
	"token": "0073172d-beef-4795-9ab4-615d38c9f05f",
	"code": "590575"
}`)
	resp, err := http.Post(url, content, payload)

	if err != nil {
		fmt.Println("%d error %s", i, err.Error())
	}

	if resp != nil {
		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(resp.Body)

		if err != nil {
			fmt.Println("%d error %s", i, err.Error())
		}

		respStr := buf.String()

		result := fmt.Sprintf("request: %d status: %d body: %s", i, resp.StatusCode, respStr)
		fmt.Println(result)
	}

	wg.Done()
}
