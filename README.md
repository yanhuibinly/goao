# goao
通过go访问appplatform的ao接口，需要传入ip

实例，通过用户id获取用户信息:
var uid uint64

	uid = uint64(1000015005)

	var req = NewRequest()

	req.CmdId = 0x40061801

	req.Host = "172.172.177.5:53101"

	req.MachineKey = "5680a103c606e"

	req.SceneId = 1

	req.Source = "/data/code/user_iapi/vendor/kidswant/ao-user/src/Lib/Info.php"

	req.Uid = uid

	req.InReserve = "utf8"

	aoUserReq := NewAoUserInfoByUidReq()

	goTest := New()

	aoUserReq.Uid = req.Uid

	aoUserRes := NewAoUserInfoByUidRsp()

	res, err := goTest.Call(req, aoUserReq, aoUserRes)

	if err != nil {
		t.Errorf("response is %v,err is %s", res, err.Error())
	}

	t.Errorf("response is %v", aoUserRes.User.Mobile)
