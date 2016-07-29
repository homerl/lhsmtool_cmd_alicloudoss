package main
import (
    "bytes"
    "os/exec"
    "os"
    "fmt"
    "strings"
)

const largefile = 5000000000 //5GB

func GetFileSize(filepath string) (num int64,err error) {
        file, _ := os.Open( filepath )
	fi, err := file.Stat()
	fmt.Printf("--The file size %d Bytes\n", fi.Size())
	defer file.Close()
	num = fi.Size()
	return num, err
}

func LargeThan (num int64) bool {
	if num > largefile {
		return true
	}
	return false
}

func FilExist (filepath string) bool {
	if _, err := os.Stat(filepath); err == nil {
		return true
	}
	return false
}

func CmdOutput (arg ...string) (string,error) {
	cmd := exec.Command(arg[0], arg[1:]...)
	cmd.Stdin = strings.NewReader("")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("--syscmd output err:", err)
		return out.String(),err
	}
	fmt.Println("--syscmd output:", out.String())
	return out.String(),err
}
