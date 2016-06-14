package goao

import (
	//"strconv"
	//"fmt"
	"testing"
)

type LogTest struct{
	T *testing.T
}


func (l *LogTest)Error(format string, a ...interface{}){
		l.T.Logf(format, a...)
}
func (l *LogTest)Info(format string, a ...interface{}){
		l.T.Logf(format, a...)
}

func testInit(t *testing.T){
	log = &LogTest{T:t}
}
func Test_ExGetUserInfos(t *testing.T) {
	testInit(t)

	for i:=0;i<50;i++{
		var uids []int64

		uids = append(uids,1000001809)
		uids = append(uids,1000001513)
		uids = append(uids,1000001690)
		//uids = append(uids,1000020601)
		//uids = append(uids,1000002024)

	//1000005355 1000001690 1000020601
	//uids = append(uids,1000000631)
		host := "172.172.178.24:53101"
		source := "gotest"
		machineKey := "gotest"
		_, err := ExUserGetUserInfos(uids, host, machineKey, source)
		if err != nil {
			t.Errorf("error:%s,index:%d", err.Error(),i)
			return
		}
	}
}

func Test_ExUserModify(t *testing.T) {
	testInit(t)
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
	testInit(t)
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


func Test_ExUserCheckLogin(t *testing.T){
	testInit(t)
	host := "172.172.178.24:53101"
	source := "gotest"
	machineKey := "gotest"
	var uid int64
	uid = 1000001513
	skey := "H5D5E9D246"
	for i:=0;i<1;i++{
	result,err:= ExUserCheckLogin(uid, skey, host, machineKey, source)

	if err!=nil{
			t.Errorf("error:%s,index:%d", err.Error(),i)
	}else if !result{
			t.Error("login false")
		}
	}
}
