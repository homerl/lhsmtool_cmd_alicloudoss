package main
import (
    "github.com/aliyun/aliyun-oss-go-sdk/oss"
    "fmt"
)

func UploadFile(fid,filepath string,fd int, bucket *oss.Bucket,routines int) error {
	filesize,err := GetFileSize(filepath)
	if err != nil {
		return err
	}
	if LargeThan(filesize) {
		fmt.Println("--multiple uploadfile fid fd filepath:",fid,filepath)
		return bucket.UploadFile(fid, filepath, 524280*1024, oss.Routines(routines))
	} else {
		fmt.Println("--uploadfile fid fd filepath:",fid,fd,filepath)
		return bucket.PutObjectFromFile(fid, filepath)
	}
}

func AddMetaData(fid,metaname,metavalue string,bucket *oss.Bucket) error {
	return bucket.SetObjectMeta(fid, oss.Meta(metaname, metavalue))
}

func UploadAli(fid,filepath string,fd int, bucket *oss.Bucket,routines int) (err error) {
	err=UploadFile(fid,filepath,fd,bucket,routines)
	if err != nil {
		return err
	}
	err=AddMetaData(fid,"filepath",filepath,bucket)
	return err
}
