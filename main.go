package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	action := flag.String("action", "push", "push or pull")
	flag.String("filepath", "filepath", "File path")
	fid := flag.String("fid", "0", "Lustre fid")
	fd := flag.Int("fd", 0, "File descriptor")
	flag.String("loglevel", "WARNING", "log level")
	lustre_root := flag.String("lustre_root", "/mnt", "Lustre client mount point")
	routines := flag.Int("routines", 3, "rotines number, default is 3")

	objdomain := flag.String("objdomain", "domain", "object storage oss domain")
	objkeyid := flag.String("objkeyid", "id", "object storage key id")
	objkey := flag.String("objkey", "key", "object storage key")
	objbucket := flag.String("objbucket", "test", "object storage bucket name")

	flag.Parse()

	bucket, err := AuthAli(*objdomain, *objkeyid, *objkey, *objbucket)
	if err != nil {
		fmt.Println("--AuthAli failed:", err)
		os.Exit(2)
	}
	filepath, err := CmdOutput("lfs", "fid2path", *lustre_root, *fid)
	filepath = strings.Replace(filepath, "\n", "", -1)
	if err != nil {
		fmt.Println("--lfs fid2path err:", err)
		os.Exit(2)
	}

	if strings.Contains(*action, "push") {
		fmt.Println("--begin upload")
		err := UploadAli(*fid, filepath, *fd, bucket, *routines)
		if err != nil {
			fmt.Println("--UploadAli err:", err)
			os.Exit(2)
		}
	} else if strings.Contains(*action, "pull") {
		fmt.Println("--begin download")
		err := DownloadAli(*fid, *fd, bucket, *routines)
		if err != nil {
			fmt.Println("--DownloadAli err:", err)
			os.Exit(2)
		}
	}
}
