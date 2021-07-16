package server

import (
	"crypto/tls"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// 获取client对象，需要设置为跳过不安全的https
var client *http.Client = &http.Client{
	Transport: &http.Transport{
		MaxIdleConnsPerHost: 10,
		TLSClientConfig: &tls.Config{
			MaxVersion:         tls.VersionTLS11,
			InsecureSkipVerify: true,
		},
	},
}

// 定义客户端测试代码
func BcjClient(items []string) (res []bool, err error) {
	data := strings.NewReader(strings.Join(items, ","))
	req, err := http.NewRequest("GET", "https://localhost:8080/exists", data)
	req.Close = false
	if err != nil {
		return nil, errors.New("请求初始化失败")
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("客户端发起请求失败")
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("响应数据解析失败")
	}
	for _, item := range strings.Split(string(content), ",") {
		i, err := strconv.ParseBool(item)
		if err != nil {
			return nil, errors.New("响应数据格式不正确")
		}
		res = append(res, i)
	}
	return res, nil
}

func ClearCache() error {
	req, err := http.NewRequest("GET", "https://localhost:8080/clear", nil)
	if err != nil {
		return errors.New("请求初始化失败")
	}
	resp, err := client.Do(req)
	if err != nil {
		return errors.New("客户端发起请求失败")
	}
	defer resp.Body.Close()
	return nil
}
