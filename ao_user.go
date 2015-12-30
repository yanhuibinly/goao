package goao

type AoRsp struct {
	ErrMsg     string
	Result     uint32
	OutReserve string
}

/*UserInfoByUid start*/
type AoUserInfoByUidReq struct {
	Uid uint64
}

func NewAoUserInfoByUidReq() *AoUserInfoByUidReq {

	model := &AoUserInfoByUidReq{}

	return model
}

func (u *AoUserInfoByUidReq) Serialize(bs *ByteStream, req *Request) bool {
	bs.PushString(req.MachineKey)
	bs.PushString(req.Source)
	bs.PushUint32(req.SceneId)
	bs.PushUint64(u.Uid)
	bs.PushString(req.InReserve)

	return bs.IsGood()
}

type AoUserInfoByUidRsp struct {
	AoRsp
	User *AoUserInfoXXOO
}

func NewAoUserInfoByUidRsp() *AoUserInfoByUidRsp {
	model := &AoUserInfoByUidRsp{}

	return model
}

func (u *AoUserInfoByUidRsp) UnSerialize(bs *ByteStream) (bool, error) {

	var errResult error
	u.Result, errResult = bs.PopUint32()
	if errResult != nil {
		return false, errResult
	}
	userInfo := NewAoUserInfoXXOO()
	errUser := userInfo.UnSerialize(bs)

	if errUser != nil {
		return false, errUser
	}

	u.User = userInfo
	var errMsg error
	u.ErrMsg, errMsg = bs.PopString()
	if errMsg != nil {
		return false, errMsg
	}
	var errOutResever error
	u.OutReserve, errOutResever = bs.PopString()

	if errOutResever != nil {
		return false, errOutResever
	}

	return bs.IsGood(), nil
}

/*UserInfoByUid end*/
