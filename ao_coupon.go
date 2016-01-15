package goao

import ()

/*UserReceiveCoupon start*/
type AoUserReceiveCouponReq struct {
	Uid              uint64
	ResourceBatchMd5 string
}

func NewAoUserReceiveCouponReq() *AoUserReceiveCouponReq {

	model := &AoUserReceiveCouponReq{}
	return model
}

func (u *AoUserReceiveCouponReq) Serialize(bs *ByteStream, req *Request) bool {
	bs.PushString(req.MachineKey)
	bs.PushString(req.Source)
	bs.PushUint32(req.SceneId)
	bs.PushUint32(0) //option
	aoUserRecvCouponParamXXOO := NewAoUserRecvCouponParamXXOO()
	aoUserRecvCouponParamXXOO.SetUserId(u.Uid)
	aoUserRecvCouponParamXXOO.SetResourceBatchMd5(u.ResourceBatchMd5)
	aoUserRecvCouponParamXXOO.Serialize(bs)

	bs.PushString(req.InReserve)

	return bs.IsGood()
}

func (u *AoUserReceiveCouponReq) GetCmdId() uint32 {
	return 0x30361801
}

type AoUserReceiveCouponRsp struct {
	AoRsp
}

func NewAoUserReceiveCouponRsp() *AoUserReceiveCouponRsp {
	model := &AoUserReceiveCouponRsp{}

	return model
}

func (u *AoUserReceiveCouponRsp) UnSerialize(bs *ByteStream) (bool, error) {
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
