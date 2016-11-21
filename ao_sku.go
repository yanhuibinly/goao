package goao

import (
	"errors"
	"fmt"
)

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
/*BatchFetchInfoList end*/


/*AoSkuInfoById start*/
type AoSkuInfoByIdReq struct {
	Option             	 uint64
	Param			 uint32
	Skuid			 uint64
}

func NewAoSkuInfoByIdReq() *AoSkuInfoByIdReq {
	model := &AoSkuInfoByIdReq{Option:0, Param:13}
	return model
}

func (u *AoSkuInfoByIdReq) Serialize(bs *ByteStream, req *Request) bool {
	bs.PushString(req.MachineKey)
	bs.PushString(req.Source)
	bs.PushUint32(req.SceneId)
	bs.PushUint64(u.Option)

	filter := NewAoSkuInfoByIdFilterXXOO()
	filter.SetSkuid(u.Skuid)
	var filters []IAoXXOO
	filters = append(filters, filter)
	bs.PushSet(filters)

	bs.PushUint32(u.Param)
	bs.PushString(req.InReserve)

	return bs.IsGood()
}

func (u *AoSkuInfoByIdReq) GetCmdId() uint32 {
	return 0x20051802
}

type AoSkuInfoByIdRsp struct {
	AoRsp
	SkuAllInfo		map[uint64]AoSkuInfoByIdSkuAllInfoXXOO
}

func NewAoSkuInfoByIdRsp() *AoSkuInfoByIdRsp {
	model := &AoSkuInfoByIdRsp{}
	return model
}

func (u *AoSkuInfoByIdRsp) UnSerialize(bs *ByteStream) (bool, error) {
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

	length, err := bs.PopUint32()
	if err != nil {
		return false, err
	}
	var skuId uint64
	u.SkuAllInfo = make(map[uint64]AoSkuInfoByIdSkuAllInfoXXOO, length)
	for i := uint32(0); i < length; i++ {
		skuId, err = bs.PopUint64()
		if err != nil {
			return false, errors.New(fmt.Sprintf("skuAllInfo error: %s,length:%d,index:%d",err.Error(),length,i))
		}
		skuInfo := NewAoSkuInfoByIdSkuAllInfoXXOO()
		err = skuInfo.UnSerialize(bs)
		if err != nil {
			return false, err
		}
		u.SkuAllInfo[skuId] = *skuInfo
	}

	var errOutResever error
	u.OutReserve, errOutResever = bs.PopString()
	if errOutResever != nil {
		return false, errOutResever
	}

	return bs.IsGood(), nil
}
/*AoSkuInfoById start*/
