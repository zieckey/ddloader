package ddl

import (
	"log"
	"testing"
)

func Test_DownLoadRemoteFileByHTTP(t *testing.T) {
	r, err := DownLoadRemoteFileByHTTP("https://www.toutiao.com/robots.txt")
	if err != nil {
		panic(err)
	}
	log.Println(string(r))
}

func Test_WriteFileToTarget(t *testing.T) {
	file := "/tmp/toutiao.robots.txt"
	r, err := DownLoadRemoteFileByHTTP("https://www.toutiao.com/robots.txt")
	if err != nil {
		panic(err)
	}
	err = WriteFileToTarget(file, r)
	if err != nil {
		panic(err)
	}
}

func Test_ExecCmd(t *testing.T) {
	cmd := "rm -rf /tmp/toutiao.robots.txt"
	err := ExecCmd(cmd)
	if err != nil {
		panic(err)
	}
}
