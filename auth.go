package main
import (
    "github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func AuthAli(objdomain,objkeyid,objkey,objbucket string) (bucket *oss.Bucket,err error) {
    client, _ := oss.New(objdomain, objkeyid, objkey)
    bucket, err = client.Bucket(objbucket)
    return bucket,err
}
