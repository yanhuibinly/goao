package goao

import (
	//"strconv"
	//"fmt"
	"testing"
)

func Test_CallGetUserInfo(t *testing.T) {

	var uid uint64

	uid = uint64(1000015005)

	var req = NewRequest()

	req.Host = "172.172.177.5:53101"

	req.MachineKey = "5680a103c606e"

	req.SceneId = 1

	req.Source = "go ao test"

	req.Uid = uid

	aoUserReq := NewAoGetUserInfoByUidReq()

	goTest := New()

	aoUserReq.Uid = req.Uid

	aoUserRes := NewAoGetUserInfoByUidRsp()

	res, err := goTest.Call(req, aoUserReq, aoUserRes)

	if err != nil {
		t.Errorf("response is %v,err is %s", res, err.Error())
	} else if aoUserRes.Result != 0 {
		t.Errorf("response is %v,err is %s", res)

	}
}

func Test_CallBatchGetUserInfo(t *testing.T) {

	var uids []uint64

	uids = []uint64{1000015005, 1000015006}

	var req = NewRequest()

	req.Host = "172.172.177.5:53101"

	req.MachineKey = "5680a103c606e"

	req.SceneId = 1

	req.Source = "go ao test"

	aoUserReq := NewAoBatchGetUserInfoByUidReq()

	goTest := New()

	aoUserReq.Uids = uids

	aoUserRes := NewAoBatchGetUserInfoByUidRsp()

	res, err := goTest.Call(req, aoUserReq, aoUserRes)

	if err != nil {
		t.Errorf("response is %v,err is %s", res, err.Error())
	}
	if aoUserRes.Result != 0 {
		t.Errorf("response is %v", aoUserRes)
	}
}

func Test_CallCheckLoginByUid(t *testing.T) {

	var uid int64

	uid = 1000072894

	skey := "H5BEEE7D991"

	var req = NewRequest()

	req.Host = "172.172.177.5:53101"

	req.MachineKey = "5680a103c606e"

	req.SceneId = 1

	req.Source = "go ao test"

	goTest := New()

	goTest.SetDwOperatorId(uid)

	goTest.SetSPassword(skey)

	aoReq := NewAoCheckLoginByUidReq()

	aoRes := NewAoCheckLoginByUidRsp()

	res, err := goTest.Call(req, aoReq, aoRes)

	if err != nil {
		t.Errorf("response is %v,err is %s", res, err.Error())
	}
	t.Errorf("response is %v", aoRes.Result)
}
