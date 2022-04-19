package main

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net"
	"net/http"
	"strings"
	"time"
)

func main() {
	// 创建连接池
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second, //连接超时
			KeepAlive: 30 * time.Second, //探活时间
		}).DialContext,
		MaxIdleConns:          100,              //最大空闲连接
		IdleConnTimeout:       90 * time.Second, //空闲超时时间
		TLSHandshakeTimeout:   10 * time.Second, //tls握手超时时间
		ExpectContinueTimeout: 1 * time.Second,  //100-continue状态码超时时间
	}
	// 创建客户端
	client := &http.Client{
		Timeout:   time.Second * 30, //请求超时时间
		Transport: transport,
	}
	// 请求数据
	resp, err := client.Get("http://127.0.0.1:1210/bye")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// 读取内容
	bds, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bds))
}

// StringVerifyRequest请求执行
func StringVerifyRequest(url string, arrs []string) {
	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}}
	rsp, err := client.Do(PostFormRequest(url, arrs))
	if err != nil {
		// return nil, err
		return
	}

	if rsp.StatusCode != http.StatusOK {
		log.Printf("Request failed with response code: %d", rsp.StatusCode)
	}
	// var res stringVerifyResponse
	// bodyBytes, _ := ioutil.ReadAll(rsp.Body)
	// json.Unmarshal(bodyBytes, &res)
	// return res.Data, nil
}

// PostFormRequest 构造post form表单请求
func PostFormRequest(url string, arrs []string) *http.Request {
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

func Handle(url string, method string, kind string, data []byte) ([]byte, error) {
	client := &http.Client{}

	switch method {
	case "post":
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
		if err != nil {
			return nil, err
		}
		req.SetBasicAuth("sendmail", "1f018f3c27b3330d0f")
		req.Header.Set("Content-Type", "application/json")
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		return ioutil.ReadAll(resp.Body)
	default:
		return []byte{}, errors.New("暂未支持该方法")
	}
}
