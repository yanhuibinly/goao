package goao

import (
	//"strconv"
	//"fmt"
	"testing"
)

func Test_ExGetUserInfos(t *testing.T) {
	var uids []int64
	uids = append(uids, 1000193003)
	host := "172.172.177.5:53101"
	source := "cmt"
	machineKey := "cmt"
	data, err := ExUserGetUserInfos(uids, host, machineKey, source)
	if err != nil {
		t.Errorf("error:%s", err.Error())
	} else {
		t.Errorf("data is :%v", data)
	}
}
