package ddl

import (
//"log"
)

//implement Fetcher interface
type ddl_fetcher struct {
}

func (d *ddl_fetcher) HasUpdate() bool {
	//todo 更新逻辑
	return true
}

func (d *ddl_fetcher) Fetch() bool {
	return true
	// b, err := DownLoadRemoteFileByHTTP(remoteUrl)
	// if err != nil {

	// }
}
