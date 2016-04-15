package goao

import ()

type AoScoreManualAdjustReq struct {
	Uid         uint64
	OperateType uint64
	Score       uint64
	OrderId     string
	OperateTime uint32
}

func NewAoScoreManualAdjustReq() *AoScoreManualAdjustReq {

	model := &AoScoreManualAdjustReq{}
	return model
}

func (u *AoScoreManualAdjustReq) Serialize(bs *ByteStream, req *Request) bool {
	bs.PushString(req.MachineKey)
	bs.PushString(req.Source)
	bs.PushUint32(req.SceneId)

	accountInterfaceOperatorXXOO := NewAoScoreAccountInterfaceOperatorXXOO()
	accountInterfaceOperatorXXOO.Serialize(bs)

	scoreOperateXXOO := NewAoScoreOperateXXOO()
	scoreOperateXXOO.Uid = u.Uid
	scoreOperateXXOO.Uid_u = 1
	scoreOperateXXOO.OperateType = u.OperateType
	scoreOperateXXOO.OperateType_u = 1
	scoreOperateXXOO.Score = u.Score
	scoreOperateXXOO.Score_u = 1
	scoreOperateXXOO.OrderId = u.OrderId
	scoreOperateXXOO.OrderId_u = 1
	scoreOperateXXOO.OperateTime = u.OperateTime
	scoreOperateXXOO.OperateTime_u = 1
	scoreOperateXXOO.Serialize(bs)

	bs.PushString(req.InReserve)
	return bs.IsGood()
}

func (u *AoScoreManualAdjustReq) GetCmdId() uint32 {
	return 0x40111803
}

type AoScoreManualAdjustResp struct {
	AoRsp
}

func NewAoScoreManualAdjustResp() *AoScoreManualAdjustResp {
	model := &AoScoreManualAdjustResp{}

	return model
}
func (u *AoScoreManualAdjustResp) UnSerialize(bs *ByteStream) (bool, error) {
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

/*function Unserialize($bs){
	$this->_arr_value['result'] = $bs->popUint32_t();
	$this->_arr_value['errmsg'] = $bs->popString();	//<std::string> 错误提示信息
	$this->_arr_value['outReserve'] = $bs->popString();	//<std::string> 输出保留字

}*/
