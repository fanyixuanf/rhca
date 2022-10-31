/*
@Time : 2022/10/31 20:53
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : templatemethod_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package Behavioral_Pattern

import (
	"fmt"
	"testing"
)

type Downloader interface {
	Download(uri string)
}

type implement interface {
	download()
	save()
}

type template struct {
	implement
	uri string
}

func NewTemplate(impl implement) *template {
	return &template{
		implement: impl,
	}
}

func (t *template) Download(uri string) {
	t.uri = uri
	fmt.Println("prepare downloading ...")
	t.implement.download()
	t.implement.save()
	fmt.Println("finish download save ...")
}

func (t *template) save() {
	fmt.Println("template default save")
}

type HTTPDownloader struct {
	*template
}

func NewHttpDownloader() Downloader {
	downloader := &HTTPDownloader{}
	template := NewTemplate(downloader)
	downloader.template = template
	return downloader
}

func (h *HTTPDownloader) download() {
	fmt.Printf("download %s via http \n", h.uri)
}

func (*HTTPDownloader) save() {
	fmt.Println("Http save ....")
}

type FTPDownloader struct {
	*template
}

func NewFTPDownloader() Downloader {
	download := &FTPDownloader{}
	template := NewTemplate(download)
	download.template = template
	return download
}

func (f *FTPDownloader) download() {
	fmt.Printf("download %s vis ftm \n", f.uri)
}

func TestTemplateMethod(t *testing.T) {
	var http Downloader = NewHttpDownloader()
	http.Download("http://127.0.0.1/download/filename.zip")

	var ftp Downloader = NewFTPDownloader()
	ftp.Download("http://127.0.0.1/download/filename.zip")
}
