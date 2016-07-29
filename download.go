package main

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func DownloadAli(fid string, fd int, bucket *oss.Bucket, routines int) error {
	objtes, err := bucket.GetObjectMeta(fid)
	filepath := GetMetaData(objtes, "Filepath")
	filepath = strings.Replace(filepath, "/", " ", -1)
	filename := strings.Fields(filepath)[len(strings.Fields(filepath))-1]
	if err != nil {
		return err
	}
	body, err := bucket.GetObject(fid)
	if err != nil {
		return err
	}
	defer body.Close()
	fdp := os.NewFile(uintptr(fd), filename)
	defer fdp.Close()
	io.Copy(fdp, body)
	return err
}

func GetObjSize(objtes http.Header) (objsize int64) {
	for _, v := range objtes["Content-Length"] {
		objsize, _ = strconv.ParseInt(v, 10, 64)
	}
	return objsize
}

func GetMetaData(objtes http.Header, metaname string) (value string) {
	for _, value = range objtes["X-Oss-Meta-"+metaname] { //metaname="Filepath"
	}
	return value
}
