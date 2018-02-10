package ddl

import (
	"io/ioutil"
	//"log"
	"net/http"
	"os"
	"os/exec"
)

func DownLoadRemoteFileByHTTP(remoteUrl string) ([]byte, error) {
	resp, err := http.Get(remoteUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func WriteFileToTarget(file string, b []byte) error {
	err := ioutil.WriteFile(file, b, 0777)
	return err
}

func CheckFileIsExist(file string) bool {
	var exist = true
	if _, err := os.Stat(file); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func ExecCmd(c string) error {
	cmd := exec.Command("/bin/sh", "-c", c)
	err := cmd.Run()
	return err
}
