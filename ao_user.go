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

/*CheckLoginByUid start*/
type AoCheckLoginByUidReq struct {
}

func NewAoCheckLoginByUidReq() *AoCheckLoginByUidReq {

	model := &AoCheckLoginByUidReq{}
	return model
}

func (u *AoCheckLoginByUidReq) GetCmdId() uint32 {
	return 0x40061811
}
func (u *AoCheckLoginByUidReq) Serialize(bs *ByteStream, req *Request) bool {
	bs.PushString(req.MachineKey)
	bs.PushString(req.Source)
	bs.PushUint32(req.SceneId)
	bs.PushString(req.InReserve)

	return bs.IsGood()
}

type AoCheckLoginByUidRsp struct {
	AoRsp
}

func NewAoCheckLoginByUidRsp() *AoCheckLoginByUidRsp {
	model := &AoCheckLoginByUidRsp{}
	return model
}

func (u *AoCheckLoginByUidRsp) UnSerialize(bs *ByteStream) (bool, error) {

	var errResult error
	u.Result, errResult = bs.PopUint32()
	if errResult != nil {
		return false, errResult
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

/*CheckLoginByUid end*/

/*ModifyBasicUserInfoByUid start*/
type AoModifyBasicUserInfoByUidReq struct {
	Uid       uint64
	Signature string
	AuthCode  string
	UserInfo  *AoUserInfoXXOO
}

func NewAoModifyBasicUserInfoByUidReq() *AoModifyBasicUserInfoByUidReq {

	model := &AoModifyBasicUserInfoByUidReq{}
	model.AuthCode = "^YHN,ki8"
	return model
}

func (u *AoModifyBasicUserInfoByUidReq) GetCmdId() uint32 {
	return 0x40061804
}
func (u *AoModifyBasicUserInfoByUidReq) Serialize(bs *ByteStream, req *Request) bool {
	bs.PushString(req.MachineKey)
	bs.PushString(req.Source)
	bs.PushUint32(req.SceneId)
	bs.PushString(u.AuthCode)

	u.UserInfo.Serialize(bs)

	bs.PushString(req.InReserve)
	return bs.IsGood()
}

type AoModifyBasicUserInfoByUidRsp struct {
	AoRsp
}

func NewAoModifyBasicUserInfoByUidRsp() *AoModifyBasicUserInfoByUidRsp {
	model := &AoModifyBasicUserInfoByUidRsp{}
	return model
}

func (u *AoModifyBasicUserInfoByUidRsp) UnSerialize(bs *ByteStream) (bool, error) {

	var errResult error
	u.Result, errResult = bs.PopUint32()
	if errResult != nil {
		return false, errResult
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

/*ModifyBasicUserInfoByUid end*/

/*GetUserInfoByBindInfo start*/
type AoGetUserInfoByBindInfoReq struct {
	BindInfoType uint32 //<uint32_t> 绑定信息类型：1-手机号，2-邮箱，3-个性化账号，详见E_BINDINFO_TYPE，必填(版本>=0)
	BindInfo     string //<std::string> 绑定信息：手机号、邮箱、个性化账号, 必填(版本>=0)
}

func NewAoGetUserInfoByBindInfoReq() *AoGetUserInfoByBindInfoReq {

	model := &AoGetUserInfoByBindInfoReq{}
	return model
}

func (u *AoGetUserInfoByBindInfoReq) GetCmdId() uint32 {
	return 0x40061802
}

func (u *AoGetUserInfoByBindInfoReq) Serialize(bs *ByteStream, req *Request) bool {
	bs.PushString(req.MachineKey) //<std::string> 机器码，请取cookie里面的visitkey，无法获得visitkey的可以填随机字符串，必填
	bs.PushString(req.Source)     //<std::string> 调用来源, 请填调用方自己的源文件名称，必填
	bs.PushUint32(req.SceneId)    //<uint32_t> 场景id：1，只获取用户信息；2，获取用户信息和宝宝信息。必填
	bs.PushUint32(u.BindInfoType) //<uint32_t> 绑定信息类型：1-手机号，2-邮箱，3-个性化账号，详见E_BINDINFO_TYPE，必填
	bs.PushString(u.BindInfo)     //<std::string> 绑定信息：手机号、邮箱、个性化账号, 必填
	bs.PushString(req.InReserve)  //<std::string> 输入保留字

	return bs.IsGood()
}

type AoGetUserInfoByBindInfoRsp struct {
	AoRsp
	User *AoUserInfoXXOO
}

func NewAoGetUserInfoByBindInfoRsp() *AoGetUserInfoByBindInfoRsp {
	model := &AoGetUserInfoByBindInfoRsp{}
	return model
}

func (u *AoGetUserInfoByBindInfoRsp) UnSerialize(bs *ByteStream) (bool, error) {
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

/*GetUserInfoByBindInfo end*/
