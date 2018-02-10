package main

import (
	"github.com/zieckey/ddloader/ddl"
	"os"
)

func main() {
	// 这个main只是一个示例代码，报警和日志都是默认的方式（记录本地日志）
	// 各个业务方可以根据ddl库，定制自己的下载器，报警和日志都可以自定义

	m := new(ddl.Main)
	err := m.Init()
	if err != nil {
		ddl.Fatalf("Init failed. %v", err.Error())
		os.Exit(-1)
	}

	err = m.Run()
	if err != nil {
		ddl.Fatalf("Run failed. %v", err.Error())
		os.Exit(-2)
	}

	err = m.Uninit()
	if err != nil {
		ddl.Fatalf("Uninit failed. %v", err.Error())
		os.Exit(-3)
	}

	os.Exit(0)
}
