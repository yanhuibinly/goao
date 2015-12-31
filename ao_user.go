package goao

type AoRsp struct {
	ErrMsg     string
	Result     uint32
	OutReserve string
}

/*UserInfoByUid start*/
type AoGetUserInfoByUidReq struct {
	Uid uint64
}

func NewAoGetUserInfoByUidReq() *AoGetUserInfoByUidReq {

	model := &AoGetUserInfoByUidReq{}
	return model
}

func (u *AoGetUserInfoByUidReq) Serialize(bs *ByteStream, req *Request) bool {
	bs.PushString(req.MachineKey)
	bs.PushString(req.Source)
	bs.PushUint32(req.SceneId)
	bs.PushUint64(u.Uid)
	bs.PushString(req.InReserve)

	return bs.IsGood()
}

func (u *AoGetUserInfoByUidReq) GetCmdId() uint32 {
	return 0x40061801
}

type AoGetUserInfoByUidRsp struct {
	AoRsp
	User *AoUserInfoXXOO
}

func NewAoGetUserInfoByUidRsp() *AoGetUserInfoByUidRsp {
	model := &AoGetUserInfoByUidRsp{}

	return model
}

func (u *AoGetUserInfoByUidRsp) UnSerialize(bs *ByteStream) (bool, error) {

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

/*GetUserInfoByUid end*/

/*BatchGetUserInfoByUid start*/
type AoBatchGetUserInfoByUidReq struct {
	Uids []uint64
}

func NewAoBatchGetUserInfoByUidReq() *AoBatchGetUserInfoByUidReq {

	model := &AoBatchGetUserInfoByUidReq{}
	return model
}

func (u *AoBatchGetUserInfoByUidReq) GetCmdId() uint32 {
	return 0x40061814
}
func (u *AoBatchGetUserInfoByUidReq) Serialize(bs *ByteStream, req *Request) bool {
	bs.PushString(req.MachineKey)
	bs.PushString(req.Source)
	bs.PushUint32(req.SceneId)
	bs.PushUint64Set(u.Uids)
	bs.PushString(req.InReserve)

	return bs.IsGood()
}

type AoBatchGetUserInfoByUidRsp struct {
	AoRsp
	Users map[uint64]AoUserInfoXXOO
}

func NewAoBatchGetUserInfoByUidRsp() *AoBatchGetUserInfoByUidRsp {
	model := &AoBatchGetUserInfoByUidRsp{}
	model.Users = make(map[uint64]AoUserInfoXXOO)
	return model
}

func (u *AoBatchGetUserInfoByUidRsp) UnSerialize(bs *ByteStream) (bool, error) {

	var errResult error
	u.Result, errResult = bs.PopUint32()
	if errResult != nil {
		return false, errResult
	}
	length, err := bs.PopUint32()
	if err != nil {
		return false, err
	}

	for i := uint32(0); i < length; i++ {

		uid, err := bs.PopUint64()

		if err != nil {
			return false, err
		}

		userInfo := NewAoUserInfoXXOO()

		errUser := userInfo.UnSerialize(bs)

		if errUser != nil {
			return false, errUser
		}

		u.Users[uid] = *userInfo
	}

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

/*BatchGetUserInfoByUid end*/
