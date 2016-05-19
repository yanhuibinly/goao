package goao

/*BatchFetchInfoList start*/
type AoSkuBatchGetInfoReq struct {
	Skuid    []uint64
	Channel  uint32
	Address  string
	UserIp   string
	UserData map[interface{}]interface{}
}

func NewAoSkuBatchGetInfoReq() *AoSkuBatchGetInfoReq {

	model := &AoSkuBatchGetInfoReq{}
	return model
}

func (ao *AoSkuBatchGetInfoReq) Serialize(bs *ByteStream, req *Request) bool {
	bs.PushString(req.MachineKey)
	bs.PushString(req.Source)
	bs.PushUint32(req.SceneId)
	bs.PushUint32(req.Option) //option
	bs.PushUint64Set(ao.Skuid)
	bs.PushUint32(ao.Channel)
	bs.PushString(ao.Address)
	bs.PushString(ao.UserIp)
	bs.PushMap(ao.UserData)
	return bs.IsGood()
}

func (u *AoSkuBatchGetInfoReq) GetCmdId() uint32 {
	return 0x20091802
}

type AoSkuBatchGetInfoRsp struct {
	AoRsp
	ItemInfoList []AoDetailmainPoItemInfoXXOO
}

func NewAoSkuBatchGetInfoRsp() *AoSkuBatchGetInfoRsp {
	model := &AoSkuBatchGetInfoRsp{}

	return model
}

func (u *AoSkuBatchGetInfoRsp) UnSerialize(bs *ByteStream) (bool, error) {

	var errResult error
	u.Result, errResult = bs.PopUint32()
	if errResult != nil {
		return false, errResult
	}
	var itemInfoListLen uint32
	itemInfoListLen, err := bs.PopUint32()
	if err != nil {
		return false, err
	}
	for i := uint32(0); i < itemInfoListLen; i++ {
		info := NewAoDetailmainPoItemInfoXXOO()
		err = info.UnSerialize(bs)
		if err != nil {
			return false, err
		}
		u.ItemInfoList = append(u.ItemInfoList, *info)
	}

	var errMsg error
	u.ErrMsg, errMsg = bs.PopString()
	if errMsg != nil {
		return false, errMsg
	}

	return bs.IsGood(), nil
}

func (u *AoSkuBatchGetInfoRsp) getCmdId() uint32 {
	return 0x20098802
}
