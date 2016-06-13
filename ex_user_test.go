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
	source := "gotest"
	machineKey := "gotest"
	info, err := ExUserGetUserInfos(uids, host, machineKey, source)
	if err != nil {
		t.Errorf("error:%s", err.Error())
	} else {
		t.Errorf("info:%v", info)
	}
}

func Test_ExUserModify(t *testing.T) {
	var uid int64
	uid = 1000193003
	host := "172.172.177.5:53101"
	source := "gotest"
	machineKey := "gotest"
	signature := "newSignature"
	nickName := "newNickName"
	photo := "http://img4.18183.duoku.com/uploads/allimg/160420/32-160420160319-50.jpg"
	var userType uint8
	userType = 3

	code, err := ExUserModifyBasicByUid(uid, signature, true, photo, true, nickName, true, userType, true, host, machineKey, source)

	if err != nil {
		t.Errorf("error:%s", err.Error())
	} else if code != 0 {
		t.Errorf("code:%d", code)
	}
}

func Test_ExUserGetByPhone(t *testing.T) {
	host := "172.172.177.5:53101"
	source := "gotest"
	machineKey := "gotest"
	phone := "13852286536"

	user, err := ExUserGetUserInfoByPhone(phone, host, machineKey, source)

	if err != nil {
		t.Errorf("error:%s", err.Error())
	} else {
		t.Errorf("info:%v", user)
	}
}
