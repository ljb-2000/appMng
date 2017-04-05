package commons

import (
	"io"
	"net/http"
	"time"
	"log"
	"io/ioutil"
)

func MyTestHttpRequest(method, endpoint string, reqBody io.Reader, user, pwd string) ([]byte, error) {
	hc := &http.Client{
		Timeout: time.Second * 10,
	}
	var body []byte
	request, _ := http.NewRequest(method, endpoint, reqBody)

	if len(user) > 0 && len(pwd) > 0 {
		request.SetBasicAuth(user, pwd)
	}
	request.Header.Set("Content-Type", "application/json")
	//request.Header.Set("Cookie", cookie)

	log.Println("------------send request begin------------------")
	httpRes, err := hc.Do(request)
	if err != nil {
		log.Println(err.Error())
	} else {
		defer httpRes.Body.Close()
		log.Println(httpRes.StatusCode)
		log.Println(httpRes.Status)
		body, _ = ioutil.ReadAll(httpRes.Body)
	}
	log.Println("------------recv response rancher ------------------")
	return body, err
}


