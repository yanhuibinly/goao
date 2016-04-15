package goao

type AoScoreXXOO struct {
	AoXXOO

	ScorePoints uint32 //<uint32> 积分值(版本>=0)
	OrderId     string //<std::string> 单据号(版本>=0)
	RefundId    string //<std::string> 退款单据号(版本>=0)
	OperateType uint32 //<uint32> 操作类型(版本>=0)
	OperateTime uint64 //<uint64> 操作时间(版本>=0)
	Remarks     string //<std::string> 发放备注(版本>=0)

	ScorePoints_u uint8 //<uint8> (版本>=0)
	OrderId_u     uint8 //<uint8> (版本>=0)
	RefundId_u    uint8 //<uint8> (版本>=0)
	OperateType_u uint8 //<uint8> (版本>=0)
	OperateTime_u uint8 //<uint8> (版本>=0)
	Remarks_u     uint8 //<uint8> (版本>=0)
	Reserve_u     uint8 //<uint8> (版本>=0)
	Version_u     uint8
}

func NewAoScoreXXOO() *AoScoreXXOO {
	model := &AoScoreXXOO{}

	return model
}
func (ao *AoScoreXXOO) Serialize(bs *ByteStream) bool {

	bs.PushUint32(uint32(ao.getClassLen()))

	return ao.serialize_internal(bs)
}

func (ao *AoScoreXXOO) getClassLen() int {

	bsLen := NewByteStream()
	bsLen.SetRealWrite(false)
	ao.serialize_internal(bsLen)
	length := bsLen.GetWrittenLength()

	return length
}
func (ao *AoScoreXXOO) serialize_internal(bs *ByteStream) bool {
	bs.PushUint32(ao.Version)      //<uint32> 版本号
	bs.PushUint32(ao.ScorePoints)  //<uint32> 积分值
	bs.PushString(ao.OrderId)      //<std::string> 单据号
	bs.PushString(ao.RefundId)     //<std::string> 退款单据号
	bs.PushUint32(ao.OperateType)  //<uint32> 操作类型
	bs.PushUint64(ao.OperateTime)  //<uint64> 操作时间
	bs.PushString(ao.Remarks)      //<std::string> 发放备注
	bs.PushString(ao.Reserve)      //<std::string> 预留字段,无用
	bs.PushUint8(ao.Version_u)     //<uint8>
	bs.PushUint8(ao.ScorePoints_u) //<uint8>
	bs.PushUint8(ao.OrderId_u)     //<uint8>
	bs.PushUint8(ao.RefundId_u)    //<uint8>
	bs.PushUint8(ao.OperateType_u) //<uint8>
	bs.PushUint8(ao.OperateTime_u) //<uint8>
	bs.PushUint8(ao.Remarks_u)     //<uint8>
	bs.PushUint8(ao.Reserve_u)     //<uint8>

	return bs.IsGood()
}

type AoScoreAccountInterfaceOperatorXXOO struct {
	AoXXOO
	OperatorName         string //<std::string> 操作者名,业务方控制,可用于后台查询入账的操作人员(版本>=0)
	OperatorDepartment   string //<std::string> 操作者部门(版本>=0)
	OperatorName_u       uint8  //<uint8> (版本>=0)
	OperatorDepartment_u uint8  //<uint8> (版本>=0)
}

func NewAoScoreAccountInterfaceOperatorXXOO() *AoScoreAccountInterfaceOperatorXXOO {
	model := &AoScoreAccountInterfaceOperatorXXOO{}

	return model
}
func (ao *AoScoreAccountInterfaceOperatorXXOO) Serialize(bs *ByteStream) bool {

	bs.PushUint32(uint32(ao.getClassLen()))

	return ao.serialize_internal(bs)
}

func (ao *AoScoreAccountInterfaceOperatorXXOO) getClassLen() int {

	bsLen := NewByteStream()
	bsLen.SetRealWrite(false)
	ao.serialize_internal(bsLen)
	length := bsLen.GetWrittenLength()

	return length
}

func (ao *AoScoreAccountInterfaceOperatorXXOO) serialize_internal(bs *ByteStream) bool {
	bs.PushUint32(ao.Version)             //<uint32> 版本号
	bs.PushString(ao.OperatorName)        //<std::string> 操作者名,业务方控制,可用于后台查询入账的操作人员
	bs.PushString(ao.OperatorDepartment)  //<std::string> 操作者部门
	bs.PushString(ao.Reserve)             //<std::string> reserve  预留字段,无用
	bs.PushUint8(ao.Version_u)            //<uint8>
	bs.PushUint8(ao.OperatorName_u)       //<uint8>
	bs.PushUint8(ao.OperatorDepartment_u) //<uint8>
	bs.PushUint8(ao.Reserve_u)            //<uint8>
	return bs.IsGood()
}

/*
type AoScorePointXXOO struct {
	AoXXOO
	Uid                        uint64                           //<uint64> 用户id(版本>=0)
	ConsumeTime                uint32                           //<uint32> 消费时间(版本>=0)
	DealCodeAndPointsParamPo   map[string]AoScorePointParamXXOO //<std::map<std::string,b2b2c::score::po::CPointsParamPo> > 扣减单号及对应积分值, 支持批量, 必填(版本>=0)
	Consumecode                string                           //<std::string> 消费单号, 最长不能超过64位(版本>=0)
	Uid_u                      uint8                            //<uint8> (版本>=0)
	ConsumeTime_u              uint8                            //<uint8> (版本>=0)
	DealCodeAndPointsParamPo_u uint8                            //<uint8> (版本>=0)
	Consumecode_u              uint8                            //<uint8> (版本>=0)
}

func NewAoScorePointXXOO() *AoScorePointXXOO {
	model := &AoScorePointXXOO{}

	return model
}

func (ao *AoScorePointXXOO) Serialize(bs *ByteStream) bool {

	bs.PushUint32(uint32(ao.getClassLen()))

	return ao.serialize_internal(bs)
}

func (ao *AoScorePointXXOO) getClassLen() int {

	bsLen := NewByteStream()
	bsLen.SetRealWrite(false)
	ao.serialize_internal(bsLen)
	length := bsLen.GetWrittenLength()

	return length
}
func (ao *AoScorePointXXOO) serialize_internal(bs *ByteStream) bool {
	bs.PushUint32(ao.Version)     //<uint32> 版本号
	bs.PushUint64(ao.Uid)         //<uint64> 用户id
	bs.PushUint32(ao.ConsumeTime) //<uint32> 消费时间
	bs.PushMapXXOO(ao.DealCodeAndPointsParamPo)
	bs.PushString(ao.Consumecode)               //<std::string> 消费单号, 最长不能超过64位
	bs.PushString(ao.Reserve)                   //<std::string> 其他保留
	bs.PushUint8(ao.Version_u)                  //<uint8>
	bs.PushUint8(ao.Uid_u)                      //<uint8>
	bs.PushUint8(ao.ConsumeTime_u)              //<uint8>
	bs.PushUint8(ao.DealCodeAndPointsParamPo_u) //<uint8>
	bs.PushUint8(ao.Consumecode_u)              //<uint8>
	bs.PushUint8(ao.Reserve_u)                  //<uint8>

	return bs.IsGood()
}

type AoScorePointParamXXOO struct {
	AoXXOO
	Type              uint32 //<uint32> 积分类型(版本>=0)
	PromotionPoints   uint32 //<uint32> 促销积分总值(版本>=0)
	CashPoints        uint32 //<uint32> 现金积分总值(版本>=0)
	Remarks           string //<std::string> 备注(版本>=0)
	Type_u            uint8  //<uint8> (版本>=0)
	PromotionPoints_u uint8  //<uint8> (版本>=0)
	CashPoints_u      uint8  //<uint8> (版本>=0)
	Remarks_u         uint8  //<uint8> (版本>=0)
}

func NewAoScorePointParamXXOO() *AoScorePointParamXXOO {
	model := &AoScorePointParamXXOO{}

	return model
}

func (ao *AoScorePointParamXXOO) Serialize(bs *ByteStream) bool {

	bs.PushUint32(uint32(ao.getClassLen()))

	return ao.serialize_internal(bs)
}

func (ao *AoScorePointParamXXOO) Serialize(bs *ByteStream) bool {

	bs.PushUint32(uint32(ao.getClassLen()))

	return ao.serialize_internal(bs)
}

func (ao *AoScorePointParamXXOO) serialize_internal(bs *ByteStream) bool {

	bs.PushUint32(ao.Version)          //<uint32> 版本号
	bs.PushUint32(ao.Type)             //<uint32> 积分类型
	bs.PushUint32(ao.PromotionPoints)  //<uint32> 促销积分总值
	bs.PushUint32(ao.CashPoints)       //<uint32> 现金积分总值
	bs.PushString(ao.Remarks)          //<std::string> 备注
	bs.PushString(ao.Reserve)          //<std::string> 其他保留
	bs.PushUint8(ao.Version_u)         //<uint8>
	bs.PushUint8(ao.Type_u)            //<uint8>
	bs.PushUint8(ao.PromotionPoints_u) //<uint8>
	bs.PushUint8(ao.CashPoints_u)      //<uint8>
	bs.PushUint8(ao.Remarks_u)         //<uint8>
	bs.PushUint8(ao.Reserve_u)         //<uint8>

	return bs.IsGood()
}
*/
type AoScoreOperateXXOO struct {
	AoXXOO
	Uid              uint64 //<uint64> 用户uid(版本>=0)
	OperateType      uint64 //<uint64> 操作类型(版本>=0)
	Score            uint64 //<uint64> 积分值(版本>=0)
	OrderId          string //<std::string> 单据号(版本>=0)
	ReturnId         string //<std::string> 退货单号(版本>=0)
	OperateTime      uint32 //<uint32> 操作时间(版本>=0)
	ExpirationTime   uint64 //<uint64> 有效时间(版本>=0)
	Uid_u            uint8  //<uint8> (版本>=0)
	OperateType_u    uint8  //<uint8> (版本>=0)
	Score_u          uint8  //<uint8> (版本>=0)
	OrderId_u        uint8  //<uint8> (版本>=0)
	ReturnId_u       uint8  //<uint8> (版本>=0)
	OperateTime_u    uint8  //<uint8> (版本>=0)
	ExpirationTime_u uint8  //<uint8> (版本>=0)
	Remarks          string //<std::string> 备注(版本>=20160222)
	Remarks_u        uint8  //<uint8> (版本>=20160222)
}

func NewAoScoreOperateXXOO() *AoScoreOperateXXOO {
	return &AoScoreOperateXXOO{}
}
func (ao *AoScoreOperateXXOO) Serialize(bs *ByteStream) bool {
	bs.PushUint32(uint32(ao.getClassLen()))

	return ao.serialize_internal(bs)
}

func (ao *AoScoreOperateXXOO) serialize_internal(bs *ByteStream) bool {
	bs.PushUint32(ao.Version)         //<uint32> 版本号
	bs.PushUint64(ao.Uid)             //<uint64> 用户uid
	bs.PushUint64(ao.OperateType)     //<uint64> 操作类型
	bs.PushUint64(ao.Score)           //<uint64> 积分值
	bs.PushString(ao.OrderId)         //<std::string> 单据号
	bs.PushString(ao.ReturnId)        //<std::string> 退货单号
	bs.PushUint32(ao.OperateTime)     //<uint32> 操作时间
	bs.PushUint64(ao.ExpirationTime)  //<uint64> 有效时间
	bs.PushString(ao.Reserve)         //<std::string> 其他保留
	bs.PushUint8(ao.Version_u)        //<uint8>
	bs.PushUint8(ao.Uid_u)            //<uint8>
	bs.PushUint8(ao.OperateType_u)    //<uint8>
	bs.PushUint8(ao.Score_u)          //<uint8>
	bs.PushUint8(ao.OrderId_u)        //<uint8>
	bs.PushUint8(ao.ReturnId_u)       //<uint8>
	bs.PushUint8(ao.OperateTime_u)    //<uint8>
	bs.PushUint8(ao.ExpirationTime_u) //<uint8>
	bs.PushUint8(ao.Reserve_u)        //<uint8>
	if ao.Version >= 20160222 {
		bs.PushString(ao.Remarks) //<std::string> 备注
	}
	if ao.Version >= 20160222 {
		bs.PushUint8(ao.Remarks_u) //<uint8>
	}
	return bs.IsGood()
}

func (ao *AoScoreOperateXXOO) getClassLen() int {

	bsLen := NewByteStream()
	bsLen.SetRealWrite(false)
	ao.serialize_internal(bsLen)
	length := bsLen.GetWrittenLength()

	return length
}
