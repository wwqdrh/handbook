package gof

import "fmt"

type Downloader interface {
	Download()
}

type ConcreteDownload struct {
	Url string
}

// Download 下载
func (download ConcreteDownload) Download() {
	fmt.Println(fmt.Sprintf("%s 在下载中", download.Url))
}

type DownloadProxy struct {
	Url        string
	Downloader Downloader
}

// Download 下载
func (proxy DownloadProxy) Download() {
	fmt.Println(fmt.Sprintf("准备开始下载%s", proxy.Url))
	proxy.Downloader.Download()
	fmt.Println(fmt.Sprintf("下载%s完成", proxy.Url))
}
