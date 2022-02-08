package strverify

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"strings"
)

var (
	client *http.Client
)

func HttpClientExample(ctx context.Context) {
	go ServerStart(ctx, true)
	client = &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}}
}

// StringVerifyRequest请求执行
func StringVerifyRequest(url string, arrs []string) ([]bool, error) {
	rsp, err := client.Do(stringVerifyConstruct(url, arrs))
	if err != nil {
		return nil, err
	}

	if rsp.StatusCode != http.StatusOK {
		log.Printf("Request failed with response code: %d", rsp.StatusCode)
	}
	var res stringVerifyResponse
	bodyBytes, _ := ioutil.ReadAll(rsp.Body)
	json.Unmarshal(bodyBytes, &res)
	return res.Data, nil
}

func stringVerifyConstruct(url string, arrs []string) *http.Request {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	// 填充数据
	for _, item := range arrs {
		fw, err := writer.CreateFormField("arrs")
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(fw, strings.NewReader(item))
		if err != nil {
			log.Fatal(err)
		}
	}
	// Close multipart writer.
	writer.Close()
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body.Bytes()))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req
}
