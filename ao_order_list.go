package goao

/*UserReceiveCoupon start*/
type AoOrderListSysQueryReq struct {
	QueryType		uint32
	AoOrderListDealListQueryPoXXOO
	VerifyToken		string
	HistoryFlag		uint8
}

func NewAoOrderListSysQueryReq() *AoOrderListSysQueryReq {

	model := &AoOrderListSysQueryReq{}
	return model
}

func (u *AoOrderListSysQueryReq) Serialize(bs *ByteStream, req *Request) bool {

	bs.PushString(req.Source)
	bs.PushString(req.MachineKey)
	bs.PushString(u.VerifyToken)
	bs.PushUint8(u.HistoryFlag)
	bs.PushUint32(u.Version)
	bs.PushUint32(u.QueryType)

	aoUserRecvCouponParamXXOO := NewAoOrderListDealListQueryPoXXOO()
	aoUserRecvCouponParamXXOO.SetBuyerId(u.BuyerId)
	aoUserRecvCouponParamXXOO.SetDealCode(u.DealCode)
	aoUserRecvCouponParamXXOO.SetBDealCode(u.BdealCode)
	aoUserRecvCouponParamXXOO.SetStartIndex(u.StartIndex)
	aoUserRecvCouponParamXXOO.SetPageSize(u.PageSize)

	aoUserRecvCouponParamXXOO.Serialize(bs)

	bs.PushString(req.InReserve)

	return bs.IsGood()
}

func (u *AoOrderListSysQueryReq) GetCmdId() uint32 {
	return 0x30231803
}

type AoOrderListSysQueryRsp struct {
	AoRsp
	TotalNum uint32
	OrderListInfo []AoOrderListDealPoXXOO
}

func NewAoOrderListSysQueryRsp() *AoOrderListSysQueryRsp {
	model := &AoOrderListSysQueryRsp{}

	return model
}

func (u *AoOrderListSysQueryRsp) UnSerialize(bs *ByteStream) (bool, error) {
	var errResult error
	u.Result, errResult = bs.PopUint32()
	if errResult != nil {
		return false, errResult
	}

	var errTotalNum error
	u.TotalNum, errTotalNum = bs.PopUint32()
	if errTotalNum != nil {
		return false, errTotalNum
	}

	length, err := bs.PopUint32()
	if err != nil {
		return false, err
	}

	var i uint32
	for i =1; i<= length; i++ {
		info := NewAoOrderListDealPoXXOO()
		errOrderListInfo := info.UnSerialize(bs)
		if errOrderListInfo != nil {
			//return false, errOrderListInfo
		}
		u.OrderListInfo = append(u.OrderListInfo, *info)
	}

	u.ErrMsg, _ = bs.PopString()

	u.OutReserve, _ = bs.PopString()

	return bs.IsGood(), nil
}
