package goao
type AoOrderListDealPoXXOO struct {
	AoXXOO

	Version uint32	//<uint32_t> 协议版本号(版本>=0)
	Version_u uint8	//<uint8_t> (版本>=0)
	DealCode string 	//<std::string> 字符型带分表规则包裹单Code(版本>=0)
	DealCode_u uint8	//<uint8_t> (版本>=0)
	DealId uint64	//<uint64_t> 全局唯一包裹订单号(版本>=0)
	DealId_u uint8	//<uint8_t> (版本>=0)
	BdealId uint64	//<uint64_t> 交易单号(版本>=0)
	BdealId_u uint8	//<uint8_t> (版本>=0)
	BuyerId uint64	//<uint64_t> 用户信息_用户ID(版本>=0)
	BuyerId_u uint8	//<uint8_t> (版本>=0)
	BuyerNickName string	//<std::string> 用户信息_用户昵称(版本>=0)
	BuyerNickName_u uint8	//<uint8_t> (版本>=0)
	SellerId uint64	//<uint64_t> 卖家信息_卖家用户ID(版本>=0)
	SellerId_u uint8	//<uint8_t> (版本>=0)
	SellerShopId uint32	//<uint32_t> 卖家信息_卖家商户ID(版本>=0)
	SellerShopId_u uint8	//<uint8_t> (版本>=0)
	SellerName string	//<std::string> 卖家信息_卖家名称(版本>=0)
	SellerName_u uint8	//<uint8_t> (版本>=0)
	DealType uint32	//<uint32_t> 交易信息_订单类型：1-普通商品订单  2-游乐商品订单  3-跨境购商品订单 4-闪购商品订单 5-预售商品订单(版本>=0)
	DealType_u uint8	//<uint8_t> (版本>=0)
	DealSource uint32	//<uint32_t> 交易信息_订单渠道：1-线上精选APP订单  2-收银订单  3-POS订单  4-线上商城APP订单 5-淘宝订单 6-京东订单(版本>=0)
	DealSource_u uint8	//<uint8_t> (版本>=0)
	DealState uint32	//<uint32_t> 交易信息_订单状态：1-待支付  2-待出库  3-已出库  4-已打包 5-待评论 6-已完成 7-已取消(版本>=0)
	DealState_u uint8	//<uint8_t> (版本>=0)
	DealProduceState uint32	//<uint32_t> 交易信息_订单生产状态：1-待审核  2-已审核  3-已捡货  4-已打包 5-已发货（发货单以及物流单） 6-开始配送 7-已签收 10-暂停 11-异常(版本>=0)
	DealProduceState_u uint8	//<uint8_t> (版本>=0)
	SiteId uint32	//<uint32_t> 交易信息_区域ID(版本>=0)
	SiteId_u uint8	//<uint8_t> (版本>=0)
	DealPayType uint32	//<uint32_t> 交易信息_支付方式：1-在线支付（微信支付） 2-在线支付（支付宝支付） 3-在线支付（银联支付） 4-在线支付（ 信用支付） 10-货到付款 11-门店POS(版本>=0)
	DealPayType_u uint8	//<uint8_t> (版本>=0)
	DealBitProperty1 uint64	//<uint64_t> 交易信息_订单属性值bit1(版本>=0)
	DealBitProperty1_u uint8	//<uint8_t> (版本>=0)
	DealBitProperty2 uint64	//<uint64_t> 交易信息_订单属性值bit2(版本>=0)
	DealBitProperty2_u uint8	//<uint8_t> (版本>=0)
	DealCharProperty string	//<std::string> 交易信息_订单属性值char，长度128(版本>=0)
	DealCharProperty_u uint8	//<uint8_t> (版本>=0)
	DealItemSkuList string	//<std::string> 交易信息_商品skuid列表,长度512(版本>=0)
	DealItemSkuList_u uint8	//<uint8_t> (版本>=0)
	DealObtainScore uint32	//<uint32_t> 交易信息_获得积分值(版本>=0)
	DealObtainScore_u uint8	//<uint8_t> (版本>=0)
	DealObtainExperience uint32	//<uint32_t> 交易信息_获得经验值(版本>=0)
	DealObtainExperience_u uint8	//<uint8_t> (版本>=0)
	DealSaleMode uint32	//<uint32_t> 交易信息_销售模式(版本>=0)
	DealSaleMode_u uint8	//<uint8_t> (版本>=0)
	DealTotalFee uint64	//<uint64_t> 价格信息_订单总金额(版本>=0)
	DealTotalFee_u uint8	//<uint8_t> (版本>=0)
	DealDownFee uint32	//<uint32_t> 价格信息_预售定金金额(版本>=0)
	DealDownFee_u uint8	//<uint8_t> (版本>=0)
	DealShippingFee uint32	//<uint32_t> 价格信息_运费金额(版本>=0)
	DealShippingFee_u uint8	//<uint8_t> (版本>=0)
	DealOverseasTaxFree uint32	//<uint32_t> 价格信息_海外购税额(版本>=0)
	DealOverseasTaxFree_u uint8	//<uint8_t> (版本>=0)
	DealCodFee uint32	//<uint32_t> 价格信息_COD手续费(版本>=0)
	DealCodFee_u uint8	//<uint8_t> (版本>=0)
	DealInsuranceFee uint32	//<uint32_t> 价格信息_保险费(版本>=0)
	DealInsuranceFee_u uint8	//<uint8_t> (版本>=0)
	DealTotalPayment uint64	//<uint64_t> 支付信息_实付金额(版本>=0)
	DealTotalPayment_u uint8	//<uint8_t> (版本>=0)
	DealDownPayment uint32	//<uint32_t> 支付信息_预售定金支付(版本>=0)
	DealDownPayment_u uint8	//<uint8_t> (版本>=0)
	DealCouponPayment uint32	//<uint32_t> 支付信息_优惠券支付值(版本>=0)
	DealCouponPayment_u uint8	//<uint8_t> (版本>=0)
	DealCouponId string	//<std::string> 交易信息_优惠券代码，长度64(版本>=0)
	DealCouponId_u uint8	//<uint8_t> (版本>=0)
	DealScorePayment uint32	//<uint32_t> 支付信息_积分支付值(版本>=0)
	DealScorePayment_u uint8	//<uint8_t> (版本>=0)
	DealPrepaidPayment uint32	//<uint32_t> 交易信息_预付卡支付金额(版本>=0)
	DealPrepaidPayment_u uint8	//<uint8_t> (版本>=0)
	DealPrepaidId string	//<std::string> 交易信息_预付卡id，长度64(版本>=0)
	DealPrepaidId_u uint8	//<uint8_t> (版本>=0)
	DealDiscountPayment uint32	//<uint32_t> 交易信息_优惠折扣总金额(版本>=0)
	DealDiscountPayment_u uint8	//<uint8_t> (版本>=0)
	DealPromotionDesc string	//<std::string> 交易信息_促销信息描述(版本>=0)
	DealPromotionDesc_u uint8	//<uint8_t> (版本>=0)
	DealRefundTotalPayment uint64	//<uint64_t> 退款信息_退款总金额(版本>=0)
	DealRefundTotalPayment_u uint8	//<uint8_t> (版本>=0)
	DealRefundState uint32	//<uint32_t> 退款信息_订单退款状态(版本>=0)
	DealRefundState_u uint8	//<uint8_t> (版本>=0)
	DealOrderMd5 string	//<std::string> 安全信息_下单操作MD5，长度32(版本>=0)
	DealOrderMd5_u uint8	//<uint8_t> (版本>=0)
	DealOrderIP string	//<std::string> 安全信息_下单操作MD5，长度64(版本>=0)
	DealOrderIP_u uint8	//<uint8_t> (版本>=0)
	DealVisitkey string	//<std::string> 安全信息_下单visitkey，长度32(版本>=0)
	DealVisitkey_u uint8	//<uint8_t> (版本>=0)
	DealRefer string	//<std::string> 安全信息_下单refer，长度32(版本>=0)
	DealRefer_u uint8	//<uint8_t> (版本>=0)
	BuyerNote string	//<std::string> 客服信息_客户备注，长度128(版本>=0)
	BuyerNote_u uint8	//<uint8_t> (版本>=0)
	CsNote string	//<std::string> 客服信息_客服备注，长度128(版本>=0)
	CsNote_u uint8	//<uint8_t> (版本>=0)
	SellerNote string	//<std::string> 客服信息_卖家备注，长度128(版本>=0)
	SellerNote_u uint8	//<uint8_t> (版本>=0)
	DealGenTime uint32	//<uint32_t> 操作信息_订单生成时间）(版本>=0)
	DealGenTime_u uint8	//<uint8_t> (版本>=0)
	DealVerifyTime uint32	//<uint32_t> 操作信息_审核时间(版本>=0)
	DealVerifyTime_u uint8	//<uint8_t> (版本>=0)
	DealVerifyType uint32	//<uint32_t> 操作信息_审核类型，如单笔消费大于X元，单笔购买商品数量大于Y(版本>=0)
	DealVerifyType_u uint8	//<uint8_t> (版本>=0)
	DealVerifyDesc string	//<std::string> 操作信息_审核描述,长度128(版本>=0)
	DealVerifyDesc_u uint8	//<uint8_t> (版本>=0)
	DealCancelTime uint32	//<uint32_t> 操作信息_取消时间(版本>=0)
	DealCancelTime_u uint8	//<uint8_t> (版本>=0)
	DealCancelType uint32	//<uint32_t> 操作信息_取消类型，如单笔消费大于X元，单笔购买商品数量大于Y(版本>=0)
	DealCancelType_u uint8	//<uint8_t> (版本>=0)
	DealCancelDesc string	//<std::string> 操作信息_取消描述,长度128(版本>=0)
	DealCancelDesc_u uint8	//<uint8_t> (版本>=0)
	DealStopTime uint32	//<uint32_t> 操作信息_暂停时间(版本>=0)
	DealStopTime_u uint8	//<uint8_t> (版本>=0)
	DealStopType uint32	//<uint32_t> 操作信息_暂停类型，如单笔消费大于X元，单笔购买商品数量大于Y(版本>=0)
	DealStopType_u uint8	//<uint8_t> (版本>=0)
	DealStopDesc string	//<std::string> 操作信息_暂停描述,长度128(版本>=0)
	DealStopDesc_u uint8	//<uint8_t> (版本>=0)
	DealAnomalyTime uint32	//<uint32_t> 操作信息_异常时间(版本>=0)
	DealAnomalyTime_u uint8	//<uint8_t> (版本>=0)
	DealAnomalyType uint32	//<uint32_t> 操作信息_异常类型，如单笔消费大于X元，单笔购买商品数量大于Y(版本>=0)
	DealAnomalyType_u uint8	//<uint8_t> (版本>=0)
	DealAnomalyDesc string	//<std::string> 操作信息_异常描述,长度128(版本>=0)
	DealAnomalyDesc_u uint8	//<uint8_t> (版本>=0)
	DealPaySussTime uint32	//<uint32_t> 操作信息_订单支付成功时间(版本>=0)
	DealPaySussTime_u uint8	//<uint8_t> (版本>=0)
	DealConsignTime uint32	//<uint32_t> 操作信息_卖家发货时间(版本>=0)
	DealConsignTime_u uint8	//<uint8_t> (版本>=0)
	DealPackTime uint32	//<uint32_t> 操作信息_卖家发货时间(版本>=0)
	DealPackTime_u uint8	//<uint8_t> (版本>=0)
	DealDelayConfirmHours uint32	//<uint32_t> 操作信息_延长确认收货小时(版本>=0)
	DealDelayConfirmHours_u uint8	//<uint8_t> (版本>=0)
	DealRecvConfirmTime uint32	//<uint32_t> 操作时间_签收时间(版本>=0)
	DealRecvConfirmTime_u uint8	//<uint8_t> (版本>=0)
	DealEndTime uint32	//<uint32_t> 操作信息_结束时间(版本>=0)
	DealEndTime_u uint8	//<uint8_t> (版本>=0)
	DealEvalState uint32	//<uint32_t> 操作信息_评论评价状态：0 - 未评论 1 - 部分评论 2 - 全部评论(版本>=0)
	DealEvalState_u uint8	//<uint8_t> (版本>=0)
	DealDelFlag uint32	//<uint32_t> 操作信息_订单有效标记:1-有效  2-无效(版本>=0)
	DealDelFlag_u uint8	//<uint8_t> (版本>=0)
	DealVisibleState uint32	//<uint32_t> 操作信息_可见标识:0-全局可见 1-对客户不可见（订单删除功能）  2-对全局不可见（不推荐）(版本>=0)
	DealVisibleState_u uint8	//<uint8_t> (版本>=0)
	DealExtData string	//<std::string> 扩展信息_可扩展业务数据(版本>=0)
	DealExtData_u uint8	//<uint8_t> (版本>=0)
	LastUpdateTime uint32	//<uint32_t> 最后更新时间(版本>=0)
	LastUpdateTime_u uint8	//<uint8_t> (版本>=0)
	TradeInfoList uint8	//<haiziwang::oms::po::CTradePoList> 包裹订单中的商品单列表(版本>=0)
	TradeInfoList_u uint8	//<uint8_t> (版本>=0)
	DealRecvInfo uint8	//<haiziwang::oms::po::CRecvPo> 收货信息(版本>=0)
	DealRecvInfo_u uint8	//<uint8_t> (版本>=0)
	DealInvoiceInfo uint8	//<haiziwang::oms::po::CInvoicePo> 发票信息(版本>=0)
	DealInvoiceInfo_u uint8	//<uint8_t> (版本>=0)
	DealExpressInfo uint8	//<haiziwang::oms::po::CExpressPo> 物流信息(版本>=0)
	DealExpressInfo_u uint8	//<uint8_t> (版本>=0)
	ActionLogInfoList uint8	//<haiziwang::oms::po::CDealActionLogPoList> 流水日志表(版本>=0)
	ActionLogInfoList_u uint8	//<uint8_t> (版本>=0)
	ReserveStr string	//<std::string> 保留字段str(版本>=0)
	ReserveStr_u uint8	//<uint8_t> (版本>=0)
	BdealCode string	//<std::string> 交易单编码,带分库分表规则(版本>=20151022)
	BdealCode_u uint8	//<uint8_t> 交易单编码,带分库分表规则Flag(版本>=20151022)
	DealOutputTax uint32	//<uint32_t> 销项税率(版本>=20160324)
	DealOutputTax_u uint8	//<uint8_t> (版本>=20160324)
	DealInvoiceFee uint32	//<uint32_t> 开票金额(版本>=20160324)
	DealInvoiceFee_u uint8	//<uint8_t> (版本>=20160324)
	BuyerDeviceSource string	//<std::string> 设备来源(版本>=20160429)
	BuyerDeviceSource_u uint8	//<uint8_t> (版本>=20160429)
	PartnerType uint32	//<uint32_t> 商家类型(版本>=20160523)
	PartnerType_u uint8	//<uint8_t> (版本>=20160523)
	SyncCustomsFlag uint8	//<uint8_t> 是否同步海关 1:是,0:否(版本>=20160523)
	SyncCustomsFlag_u uint8	//<uint8_t> (版本>=20160523)
}

func NewAoOrderListDealPoXXOO() *AoOrderListDealPoXXOO {
	model := &AoOrderListDealPoXXOO{}

	return model
}

func (ao *AoOrderListDealPoXXOO) getClassLen() int {

	bsLen := NewByteStream()
	bsLen.SetRealWrite(false)
	ao.serialize_internal(bsLen)
	length := bsLen.GetWrittenLength()

	return length
}

func (ao *AoOrderListDealPoXXOO) Serialize(bs *ByteStream) bool {
	bs.PushUint32(uint32(ao.getClassLen()))

	return ao.serialize_internal(bs)
}

func (ao *AoOrderListDealPoXXOO) serialize_internal(bs *ByteStream) bool {

	bs.PushUint32(ao.Version)	//<uint32_t> 协议版本号
	bs.PushUint8(ao.Version_u)	//<uint8_t>
	bs.PushString(ao.DealCode)	//<std::string> 字符型带分表规则包裹单Code
	bs.PushUint8(ao.DealCode_u)	//<uint8_t>
	bs.PushUint64(ao.DealId)	//<uint64_t> 全局唯一包裹订单号
	bs.PushUint8(ao.DealId_u)	//<uint8_t>
	bs.PushUint64(ao.BdealId)	//<uint64_t> 交易单号
	bs.PushUint8(ao.BdealId_u)	//<uint8_t>
	bs.PushUint64(ao.BuyerId)	//<uint64_t> 用户信息_用户ID
	bs.PushUint8(ao.BuyerId_u)	//<uint8_t>
	bs.PushString(ao.BuyerNickName)	//<std::string> 用户信息_用户昵称
	bs.PushUint8(ao.BuyerNickName_u)	//<uint8_t>
	bs.PushUint64(ao.SellerId)	//<uint64_t> 卖家信息_卖家用户ID
	bs.PushUint8(ao.SellerId_u)	//<uint8_t>
	bs.PushUint32(ao.SellerShopId)	//<uint32_t> 卖家信息_卖家商户ID
	bs.PushUint8(ao.SellerShopId_u)	//<uint8_t>
	bs.PushString(ao.SellerName)	//<std::string> 卖家信息_卖家名称
	bs.PushUint8(ao.SellerName_u)	//<uint8_t>
	bs.PushUint32(ao.DealType)	//<uint32_t> 交易信息_订单类型：1-普通商品订单  2-游乐商品订单  3-跨境购商品订单 4-闪购商品订单 5-预售商品订单
	bs.PushUint8(ao.DealType_u)	//<uint8_t>
	bs.PushUint32(ao.DealSource)	//<uint32_t> 交易信息_订单渠道：1-线上精选APP订单  2-收银订单  3-POS订单  4-线上商城APP订单 5-淘宝订单 6-京东订单
	bs.PushUint8(ao.DealSource_u)	//<uint8_t>
	bs.PushUint32(ao.DealState)	//<uint32_t> 交易信息_订单状态：1-待支付  2-待出库  3-已出库  4-已打包 5-待评论 6-已完成 7-已取消
	bs.PushUint8(ao.DealState_u)	//<uint8_t>
	bs.PushUint32(ao.DealProduceState)	//<uint32_t> 交易信息_订单生产状态：1-待审核  2-已审核  3-已捡货  4-已打包 5-已发货（发货单以及物流单） 6-开始配送 7-已签收 10-暂停 11-异常
	bs.PushUint8(ao.DealProduceState_u)	//<uint8_t>
	bs.PushUint32(ao.SiteId)	//<uint32_t> 交易信息_区域ID
	bs.PushUint8(ao.SiteId_u)	//<uint8_t>
	bs.PushUint32(ao.DealPayType)	//<uint32_t> 交易信息_支付方式：1-在线支付（微信支付） 2-在线支付（支付宝支付） 3-在线支付（银联支付） 4-在线支付（ 信用支付） 10-货到付款 11-门店POS
	bs.PushUint8(ao.DealPayType_u)	//<uint8_t>
	bs.PushUint64(ao.DealBitProperty1)	//<uint64_t> 交易信息_订单属性值bit1
	bs.PushUint8(ao.DealBitProperty1_u)	//<uint8_t>
	bs.PushUint64(ao.DealBitProperty2)	//<uint64_t> 交易信息_订单属性值bit2
	bs.PushUint8(ao.DealBitProperty2_u)	//<uint8_t>
	bs.PushString(ao.DealCharProperty)	//<std::string> 交易信息_订单属性值char，长度128
	bs.PushUint8(ao.DealCharProperty_u)	//<uint8_t>
	bs.PushString(ao.DealItemSkuList)	//<std::string> 交易信息_商品skuid列表,长度512
	bs.PushUint8(ao.DealItemSkuList_u)	//<uint8_t>
	bs.PushUint32(ao.DealObtainScore)	//<uint32_t> 交易信息_获得积分值
	bs.PushUint8(ao.DealObtainScore_u)	//<uint8_t>
	bs.PushUint32(ao.DealObtainExperience)	//<uint32_t> 交易信息_获得经验值
	bs.PushUint8(ao.DealObtainExperience_u)	//<uint8_t>
	bs.PushUint32(ao.DealSaleMode)	//<uint32_t> 交易信息_销售模式
	bs.PushUint8(ao.DealSaleMode_u)	//<uint8_t>
	bs.PushUint64(ao.DealTotalFee)	//<uint64_t> 价格信息_订单总金额
	bs.PushUint8(ao.DealTotalFee_u)	//<uint8_t>
	bs.PushUint32(ao.DealDownFee)	//<uint32_t> 价格信息_预售定金金额
	bs.PushUint8(ao.DealDownFee_u)	//<uint8_t>
	bs.PushUint32(ao.DealShippingFee)	//<uint32_t> 价格信息_运费金额
	bs.PushUint8(ao.DealShippingFee_u)	//<uint8_t>
	bs.PushUint32(ao.DealOverseasTaxFree)	//<uint32_t> 价格信息_海外购税额
	bs.PushUint8(ao.DealOverseasTaxFree_u)	//<uint8_t>
	bs.PushUint32(ao.DealCodFee)	//<uint32_t> 价格信息_COD手续费
	bs.PushUint8(ao.DealCodFee_u)	//<uint8_t>
	bs.PushUint32(ao.DealInsuranceFee)	//<uint32_t> 价格信息_保险费
	bs.PushUint8(ao.DealInsuranceFee_u)	//<uint8_t>
	bs.PushUint64(ao.DealTotalPayment)	//<uint64_t> 支付信息_实付金额
	bs.PushUint8(ao.DealTotalPayment_u)	//<uint8_t>
	bs.PushUint32(ao.DealDownPayment)	//<uint32_t> 支付信息_预售定金支付
	bs.PushUint8(ao.DealDownPayment_u)	//<uint8_t>
	bs.PushUint32(ao.DealCouponPayment)	//<uint32_t> 支付信息_优惠券支付值
	bs.PushUint8(ao.DealCouponPayment_u)	//<uint8_t>
	bs.PushString(ao.DealCouponId)	//<std::string> 交易信息_优惠券代码，长度64
	bs.PushUint8(ao.DealCouponId_u)	//<uint8_t>
	bs.PushUint32(ao.DealScorePayment)	//<uint32_t> 支付信息_积分支付值
	bs.PushUint8(ao.DealScorePayment_u)	//<uint8_t>
	bs.PushUint32(ao.DealPrepaidPayment)	//<uint32_t> 交易信息_预付卡支付金额
	bs.PushUint8(ao.DealPrepaidPayment_u)	//<uint8_t>
	bs.PushString(ao.DealPrepaidId)	//<std::string> 交易信息_预付卡id，长度64
	bs.PushUint8(ao.DealPrepaidId_u)	//<uint8_t>
	bs.PushUint32(ao.DealDiscountPayment)	//<uint32_t> 交易信息_优惠折扣总金额
	bs.PushUint8(ao.DealDiscountPayment_u)	//<uint8_t>
	bs.PushString(ao.DealPromotionDesc)	//<std::string> 交易信息_促销信息描述
	bs.PushUint8(ao.DealPromotionDesc_u)	//<uint8_t>
	bs.PushUint64(ao.DealRefundTotalPayment)	//<uint64_t> 退款信息_退款总金额
	bs.PushUint8(ao.DealRefundTotalPayment_u)	//<uint8_t>
	bs.PushUint32(ao.DealRefundState)	//<uint32_t> 退款信息_订单退款状态
	bs.PushUint8(ao.DealRefundState_u)	//<uint8_t>
	bs.PushString(ao.DealOrderMd5)	//<std::string> 安全信息_下单操作MD5，长度32
	bs.PushUint8(ao.DealOrderMd5_u)	//<uint8_t>
	bs.PushString(ao.DealOrderIP)	//<std::string> 安全信息_下单操作MD5，长度64
	bs.PushUint8(ao.DealOrderIP_u)	//<uint8_t>
	bs.PushString(ao.DealVisitkey)	//<std::string> 安全信息_下单visitkey，长度32
	bs.PushUint8(ao.DealVisitkey_u)	//<uint8_t>
	bs.PushString(ao.DealRefer)	//<std::string> 安全信息_下单refer，长度32
	bs.PushUint8(ao.DealRefer_u)	//<uint8_t>
	bs.PushString(ao.BuyerNote)	//<std::string> 客服信息_客户备注，长度128
	bs.PushUint8(ao.BuyerNote_u)	//<uint8_t>
	bs.PushString(ao.CsNote)	//<std::string> 客服信息_客服备注，长度128
	bs.PushUint8(ao.CsNote_u)	//<uint8_t>
	bs.PushString(ao.SellerNote)	//<std::string> 客服信息_卖家备注，长度128
	bs.PushUint8(ao.SellerNote_u)	//<uint8_t>
	bs.PushUint32(ao.DealGenTime)	//<uint32_t> 操作信息_订单生成时间）
	bs.PushUint8(ao.DealGenTime_u)	//<uint8_t>
	bs.PushUint32(ao.DealVerifyTime)	//<uint32_t> 操作信息_审核时间
	bs.PushUint8(ao.DealVerifyTime_u)	//<uint8_t>
	bs.PushUint32(ao.DealVerifyType)	//<uint32_t> 操作信息_审核类型，如单笔消费大于X元，单笔购买商品数量大于Y
	bs.PushUint8(ao.DealVerifyType_u)	//<uint8_t>
	bs.PushString(ao.DealVerifyDesc)	//<std::string> 操作信息_审核描述,长度128
	bs.PushUint8(ao.DealVerifyDesc_u)	//<uint8_t>
	bs.PushUint32(ao.DealCancelTime)	//<uint32_t> 操作信息_取消时间
	bs.PushUint8(ao.DealCancelTime_u)	//<uint8_t>
	bs.PushUint32(ao.DealCancelType)	//<uint32_t> 操作信息_取消类型，如单笔消费大于X元，单笔购买商品数量大于Y
	bs.PushUint8(ao.DealCancelType_u)	//<uint8_t>
	bs.PushString(ao.DealCancelDesc)	//<std::string> 操作信息_取消描述,长度128
	bs.PushUint8(ao.DealCancelDesc_u)	//<uint8_t>
	bs.PushUint32(ao.DealStopTime)	//<uint32_t> 操作信息_暂停时间
	bs.PushUint8(ao.DealStopTime_u)	//<uint8_t>
	bs.PushUint32(ao.DealStopType)	//<uint32_t> 操作信息_暂停类型，如单笔消费大于X元，单笔购买商品数量大于Y
	bs.PushUint8(ao.DealStopType_u)	//<uint8_t>
	bs.PushString(ao.DealStopDesc)	//<std::string> 操作信息_暂停描述,长度128
	bs.PushUint8(ao.DealStopDesc_u)	//<uint8_t>
	bs.PushUint32(ao.DealAnomalyTime)	//<uint32_t> 操作信息_异常时间
	bs.PushUint8(ao.DealAnomalyTime_u)	//<uint8_t>
	bs.PushUint32(ao.DealAnomalyType)	//<uint32_t> 操作信息_异常类型，如单笔消费大于X元，单笔购买商品数量大于Y
	bs.PushUint8(ao.DealAnomalyType_u)	//<uint8_t>
	bs.PushString(ao.DealAnomalyDesc)	//<std::string> 操作信息_异常描述,长度128
	bs.PushUint8(ao.DealAnomalyDesc_u)	//<uint8_t>
	bs.PushUint32(ao.DealPaySussTime)	//<uint32_t> 操作信息_订单支付成功时间
	bs.PushUint8(ao.DealPaySussTime_u)	//<uint8_t>
	bs.PushUint32(ao.DealConsignTime)	//<uint32_t> 操作信息_卖家发货时间
	bs.PushUint8(ao.DealConsignTime_u)	//<uint8_t>
	bs.PushUint32(ao.DealPackTime)	//<uint32_t> 操作信息_卖家发货时间
	bs.PushUint8(ao.DealPackTime_u)	//<uint8_t>
	bs.PushUint32(ao.DealDelayConfirmHours)	//<uint32_t> 操作信息_延长确认收货小时
	bs.PushUint8(ao.DealDelayConfirmHours_u)	//<uint8_t>
	bs.PushUint32(ao.DealRecvConfirmTime)	//<uint32_t> 操作时间_签收时间
	bs.PushUint8(ao.DealRecvConfirmTime_u)	//<uint8_t>
	bs.PushUint32(ao.DealEndTime)	//<uint32_t> 操作信息_结束时间
	bs.PushUint8(ao.DealEndTime_u)	//<uint8_t>
	bs.PushUint32(ao.DealEvalState)	//<uint32_t> 操作信息_评论评价状态：0 - 未评论 1 - 部分评论 2 - 全部评论
	bs.PushUint8(ao.DealEvalState_u)	//<uint8_t>
	bs.PushUint32(ao.DealDelFlag)	//<uint32_t> 操作信息_订单有效标记:1-有效  2-无效
	bs.PushUint8(ao.DealDelFlag_u)	//<uint8_t>
	bs.PushUint32(ao.DealVisibleState)	//<uint32_t> 操作信息_可见标识:0-全局可见 1-对客户不可见（订单删除功能）  2-对全局不可见（不推荐）
	bs.PushUint8(ao.DealVisibleState_u)	//<uint8_t>
	bs.PushString(ao.DealExtData)	//<std::string> 扩展信息_可扩展业务数据
	bs.PushUint8(ao.DealExtData_u)	//<uint8_t>
	bs.PushUint32(ao.LastUpdateTime)	//<uint32_t> 最后更新时间
	bs.PushUint8(ao.LastUpdateTime_u)	//<uint8_t>
	//$bs->pushObject($this->tradeInfoList,'\haiziwang\oms\po\TradePoList')	//<haiziwang::oms::po::CTradePoList> 包裹订单中的商品单列表
	bs.PushUint8(ao.TradeInfoList_u)	//<uint8_t>
	//$bs->pushObject($this->dealRecvInfo,'\haiziwang\oms\po\RecvPo')	//<haiziwang::oms::po::CRecvPo> 收货信息
	bs.PushUint8(ao.DealRecvInfo_u)	//<uint8_t>
	//$bs->pushObject($this->dealInvoiceInfo,'\haiziwang\oms\po\InvoicePo')	//<haiziwang::oms::po::CInvoicePo> 发票信息
	bs.PushUint8(ao.DealInvoiceInfo_u)	//<uint8_t>
	//$bs->pushObject($this->dealExpressInfo,'\haiziwang\oms\po\ExpressPo')	//<haiziwang::oms::po::CExpressPo> 物流信息
	bs.PushUint8(ao.DealExpressInfo_u)	//<uint8_t>
	//$bs->pushObject($this->actionLogInfoList,'\haiziwang\oms\po\DealActionLogPoList')	//<haiziwang::oms::po::CDealActionLogPoList> 流水日志表
	bs.PushUint8(ao.ActionLogInfoList_u)	//<uint8_t>
	bs.PushString(ao.ReserveStr)	//<std::string> 保留字段str
	bs.PushUint8(ao.ReserveStr_u)	//<uint8_t>
	if(ao.Version >= 20151022){
	bs.PushString(ao.BdealCode)	//<std::string> 交易单编码,带分库分表规则
	}
	if(ao.Version >= 20151022){
	bs.PushUint8(ao.BdealCode_u)	//<uint8_t> 交易单编码,带分库分表规则Flag
	}
	if(ao.Version >= 20160324){
	bs.PushUint32(ao.DealOutputTax)	//<uint32_t> 销项税率
	}
	if(ao.Version >= 20160324){
	bs.PushUint8(ao.DealOutputTax_u)	//<uint8_t>
	}
	if(ao.Version >= 20160324){
	bs.PushUint32(ao.DealInvoiceFee)	//<uint32_t> 开票金额
	}
	if(ao.Version >= 20160324){
	bs.PushUint8(ao.DealInvoiceFee_u)	//<uint8_t>
	}
	if(ao.Version >= 20160429){
	bs.PushString(ao.BuyerDeviceSource)	//<std::string> 设备来源
	}
	if(ao.Version >= 20160429){
	bs.PushUint8(ao.BuyerDeviceSource_u)	//<uint8_t>
	}
	if(ao.Version >= 20160523){
	bs.PushUint32(ao.PartnerType)	//<uint32_t> 商家类型
	}
	if(ao.Version >= 20160523){
	bs.PushUint8(ao.PartnerType_u)	//<uint8_t>
	}
	if(ao.Version >= 20160523){
	bs.PushUint8(ao.SyncCustomsFlag)	//<uint8_t> 是否同步海关 1:是,0:否
	}
	if(ao.Version >= 20160523){
	bs.PushUint8(ao.SyncCustomsFlag_u)	//<uint8_t>
	}

	return bs.IsGood()
}

func (ao *AoOrderListDealPoXXOO) UnSerialize(bs *ByteStream) error {
	classLen, err := bs.PopUint32()

	if err != nil {
		return err
	}
	startPop := bs.GetReadLength()

	ao.Version, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.Version_u, err = bs.PopUint8()
	if err != nil {
		return err
	}

	ao.DealCode , err = bs.PopString()	//<std::string> 字符型带分表规则包裹单Code
	ao.DealCode_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealId , err = bs.PopUint64()	//<uint64_t> 全局唯一包裹订单号
	ao.DealId_u , err = bs.PopUint8()	//<uint8_t>
	ao.BdealId , err = bs.PopUint64()	//<uint64_t> 交易单号
	ao.BdealId_u , err = bs.PopUint8()	//<uint8_t>
	ao.BuyerId , err = bs.PopUint64()	//<uint64_t> 用户信息_用户ID
	ao.BuyerId_u , err = bs.PopUint8()	//<uint8_t>
	ao.BuyerNickName , err = bs.PopString()	//<std::string> 用户信息_用户昵称
	ao.BuyerNickName_u , err = bs.PopUint8()	//<uint8_t>
	ao.SellerId , err = bs.PopUint64()	//<uint64_t> 卖家信息_卖家用户ID
	ao.SellerId_u , err = bs.PopUint8()	//<uint8_t>
	ao.SellerShopId , err = bs.PopUint32()	//<uint32_t> 卖家信息_卖家商户ID
	ao.SellerShopId_u , err = bs.PopUint8()	//<uint8_t>
	ao.SellerName , err = bs.PopString()	//<std::string> 卖家信息_卖家名称
	ao.SellerName_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealType , err = bs.PopUint32()	//<uint32_t> 交易信息_订单类型：1-普通商品订单  2-游乐商品订单  3-跨境购商品订单 4-闪购商品订单 5-预售商品订单
	ao.DealType_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealSource , err = bs.PopUint32()	//<uint32_t> 交易信息_订单渠道：1-线上精选APP订单  2-收银订单  3-POS订单  4-线上商城APP订单 5-淘宝订单 6-京东订单
	ao.DealSource_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealState , err = bs.PopUint32()	//<uint32_t> 交易信息_订单状态：1-待支付  2-待出库  3-已出库  4-已打包 5-待评论 6-已完成 7-已取消
	ao.DealState_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealProduceState , err = bs.PopUint32()	//<uint32_t> 交易信息_订单生产状态：1-待审核  2-已审核  3-已捡货  4-已打包 5-已发货（发货单以及物流单） 6-开始配送 7-已签收 10-暂停 11-异常
	ao.DealProduceState_u , err = bs.PopUint8()	//<uint8_t>
	ao.SiteId , err = bs.PopUint32()	//<uint32_t> 交易信息_区域ID
	ao.SiteId_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealPayType , err = bs.PopUint32()	//<uint32_t> 交易信息_支付方式：1-在线支付（微信支付） 2-在线支付（支付宝支付） 3-在线支付（银联支付） 4-在线支付（ 信用支付） 10-货到付款 11-门店POS
	ao.DealPayType_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealBitProperty1 , err = bs.PopUint64()	//<uint64_t> 交易信息_订单属性值bit1
	ao.DealBitProperty1_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealBitProperty2 , err = bs.PopUint64()	//<uint64_t> 交易信息_订单属性值bit2
	ao.DealBitProperty2_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealCharProperty , err = bs.PopString()	//<std::string> 交易信息_订单属性值char，长度128
	ao.DealCharProperty_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealItemSkuList , err = bs.PopString()	//<std::string> 交易信息_商品skuid列表,长度512
	ao.DealItemSkuList_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealObtainScore , err = bs.PopUint32()	//<uint32_t> 交易信息_获得积分值
	ao.DealObtainScore_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealObtainExperience , err = bs.PopUint32()	//<uint32_t> 交易信息_获得经验值
	ao.DealObtainExperience_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealSaleMode , err = bs.PopUint32()	//<uint32_t> 交易信息_销售模式
	ao.DealSaleMode_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealTotalFee , err = bs.PopUint64()	//<uint64_t> 价格信息_订单总金额
	ao.DealTotalFee_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealDownFee , err = bs.PopUint32()	//<uint32_t> 价格信息_预售定金金额
	ao.DealDownFee_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealShippingFee , err = bs.PopUint32()	//<uint32_t> 价格信息_运费金额
	ao.DealShippingFee_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealOverseasTaxFree , err = bs.PopUint32()	//<uint32_t> 价格信息_海外购税额
	ao.DealOverseasTaxFree_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealCodFee , err = bs.PopUint32()	//<uint32_t> 价格信息_COD手续费
	ao.DealCodFee_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealInsuranceFee , err = bs.PopUint32()	//<uint32_t> 价格信息_保险费
	ao.DealInsuranceFee_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealTotalPayment , err = bs.PopUint64()	//<uint64_t> 支付信息_实付金额
	ao.DealTotalPayment_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealDownPayment , err = bs.PopUint32()	//<uint32_t> 支付信息_预售定金支付
	ao.DealDownPayment_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealCouponPayment , err = bs.PopUint32()	//<uint32_t> 支付信息_优惠券支付值
	ao.DealCouponPayment_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealCouponId , err = bs.PopString()	//<std::string> 交易信息_优惠券代码，长度64
	ao.DealCouponId_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealScorePayment , err = bs.PopUint32()	//<uint32_t> 支付信息_积分支付值
	ao.DealScorePayment_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealPrepaidPayment , err = bs.PopUint32()	//<uint32_t> 交易信息_预付卡支付金额
	ao.DealPrepaidPayment_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealPrepaidId , err = bs.PopString()	//<std::string> 交易信息_预付卡id，长度64
	ao.DealPrepaidId_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealDiscountPayment , err = bs.PopUint32()	//<uint32_t> 交易信息_优惠折扣总金额
	ao.DealDiscountPayment_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealPromotionDesc , err = bs.PopString()	//<std::string> 交易信息_促销信息描述
	ao.DealPromotionDesc_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealRefundTotalPayment , err = bs.PopUint64()	//<uint64_t> 退款信息_退款总金额
	ao.DealRefundTotalPayment_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealRefundState , err = bs.PopUint32()	//<uint32_t> 退款信息_订单退款状态
	ao.DealRefundState_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealOrderMd5 , err = bs.PopString()	//<std::string> 安全信息_下单操作MD5，长度32
	ao.DealOrderMd5_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealOrderIP , err = bs.PopString()	//<std::string> 安全信息_下单操作MD5，长度64
	ao.DealOrderIP_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealVisitkey , err = bs.PopString()	//<std::string> 安全信息_下单visitkey，长度32
	ao.DealVisitkey_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealRefer , err = bs.PopString()	//<std::string> 安全信息_下单refer，长度32
	ao.DealRefer_u , err = bs.PopUint8()	//<uint8_t>
	ao.BuyerNote , err = bs.PopString()	//<std::string> 客服信息_客户备注，长度128
	ao.BuyerNote_u , err = bs.PopUint8()	//<uint8_t>
	ao.CsNote , err = bs.PopString()	//<std::string> 客服信息_客服备注，长度128
	ao.CsNote_u , err = bs.PopUint8()	//<uint8_t>
	ao.SellerNote , err = bs.PopString()	//<std::string> 客服信息_卖家备注，长度128
	ao.SellerNote_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealGenTime , err = bs.PopUint32()	//<uint32_t> 操作信息_订单生成时间）
	ao.DealGenTime_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealVerifyTime , err = bs.PopUint32()	//<uint32_t> 操作信息_审核时间
	ao.DealVerifyTime_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealVerifyType , err = bs.PopUint32()	//<uint32_t> 操作信息_审核类型，如单笔消费大于X元，单笔购买商品数量大于Y
	ao.DealVerifyType_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealVerifyDesc , err = bs.PopString()	//<std::string> 操作信息_审核描述,长度128
	ao.DealVerifyDesc_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealCancelTime , err = bs.PopUint32()	//<uint32_t> 操作信息_取消时间
	ao.DealCancelTime_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealCancelType , err = bs.PopUint32()	//<uint32_t> 操作信息_取消类型，如单笔消费大于X元，单笔购买商品数量大于Y
	ao.DealCancelType_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealCancelDesc , err = bs.PopString()	//<std::string> 操作信息_取消描述,长度128
	ao.DealCancelDesc_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealStopTime , err = bs.PopUint32()	//<uint32_t> 操作信息_暂停时间
	ao.DealStopTime_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealStopType , err = bs.PopUint32()	//<uint32_t> 操作信息_暂停类型，如单笔消费大于X元，单笔购买商品数量大于Y
	ao.DealStopType_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealStopDesc , err = bs.PopString()	//<std::string> 操作信息_暂停描述,长度128
	ao.DealStopDesc_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealAnomalyTime , err = bs.PopUint32()	//<uint32_t> 操作信息_异常时间
	ao.DealAnomalyTime_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealAnomalyType , err = bs.PopUint32()	//<uint32_t> 操作信息_异常类型，如单笔消费大于X元，单笔购买商品数量大于Y
	ao.DealAnomalyType_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealAnomalyDesc , err = bs.PopString()	//<std::string> 操作信息_异常描述,长度128
	ao.DealAnomalyDesc_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealPaySussTime , err = bs.PopUint32()	//<uint32_t> 操作信息_订单支付成功时间
	ao.DealPaySussTime_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealConsignTime , err = bs.PopUint32()	//<uint32_t> 操作信息_卖家发货时间
	ao.DealConsignTime_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealPackTime , err = bs.PopUint32()	//<uint32_t> 操作信息_卖家发货时间
	ao.DealPackTime_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealDelayConfirmHours , err = bs.PopUint32()	//<uint32_t> 操作信息_延长确认收货小时
	ao.DealDelayConfirmHours_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealRecvConfirmTime , err = bs.PopUint32()	//<uint32_t> 操作时间_签收时间
	ao.DealRecvConfirmTime_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealEndTime , err = bs.PopUint32()	//<uint32_t> 操作信息_结束时间
	ao.DealEndTime_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealEvalState , err = bs.PopUint32()	//<uint32_t> 操作信息_评论评价状态：0 - 未评论 1 - 部分评论 2 - 全部评论
	ao.DealEvalState_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealDelFlag , err = bs.PopUint32()	//<uint32_t> 操作信息_订单有效标记:1-有效  2-无效
	ao.DealDelFlag_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealVisibleState , err = bs.PopUint32()	//<uint32_t> 操作信息_可见标识:0-全局可见 1-对客户不可见（订单删除功能）  2-对全局不可见（不推荐）
	ao.DealVisibleState_u , err = bs.PopUint8()	//<uint8_t>
	ao.DealExtData , err = bs.PopString()	//<std::string> 扩展信息_可扩展业务数据
	ao.DealExtData_u , err = bs.PopUint8()	//<uint8_t>
	ao.LastUpdateTime , err = bs.PopUint32()	//<uint32_t> 最后更新时间
	ao.LastUpdateTime_u , err = bs.PopUint8()	//<uint8_t>
	//ao.TradeInfoList , err = bs.PopObject('\haiziwang\oms\po\TradePoList')	//<haiziwang::oms::po::CTradePoList> 包裹订单中的商品单列表
	ao.TradeInfoList_u , err = bs.PopUint8()	//<uint8_t>
	//ao.DealRecvInfo , err = bs.PopObject('\haiziwang\oms\po\RecvPo')	//<haiziwang::oms::po::CRecvPo> 收货信息
	ao.DealRecvInfo_u , err = bs.PopUint8()	//<uint8_t>
	//ao.DealInvoiceInfo , err = bs.PopObject('\haiziwang\oms\po\InvoicePo')	//<haiziwang::oms::po::CInvoicePo> 发票信息
	ao.DealInvoiceInfo_u , err = bs.PopUint8()	//<uint8_t>
	//ao.DealExpressInfo , err = bs.PopObject('\haiziwang\oms\po\ExpressPo')	//<haiziwang::oms::po::CExpressPo> 物流信息
	ao.DealExpressInfo_u , err = bs.PopUint8()	//<uint8_t>
	//ao.ActionLogInfoList , err = bs.PopObject('\haiziwang\oms\po\DealActionLogPoList')	//<haiziwang::oms::po::CDealActionLogPoList> 流水日志表
	ao.ActionLogInfoList_u , err = bs.PopUint8()	//<uint8_t>
	ao.ReserveStr , err = bs.PopString()	//<std::string> 保留字段str
	ao.ReserveStr_u , err = bs.PopUint8()	//<uint8_t>
	if(ao.Version >= 20151022){
	ao.BdealCode, err = bs.PopString()	//<std::string> 交易单编码,带分库分表规则
	}
	if(ao.Version >= 20151022){
	ao.BdealCode_u, err = bs.PopUint8()	//<uint8_t> 交易单编码,带分库分表规则Flag
	}
	if(ao.Version >= 20160324){
	ao.DealOutputTax, err = bs.PopUint32()	//<uint32_t> 销项税率
	}
	if(ao.Version >= 20160324){
	ao.DealOutputTax_u, err = bs.PopUint8()	//<uint8_t>
	}
	if(ao.Version >= 20160324){
	ao.DealInvoiceFee, err = bs.PopUint32()	//<uint32_t> 开票金额
	}
	if(ao.Version >= 20160324){
	ao.DealInvoiceFee_u, err = bs.PopUint8()	//<uint8_t>
	}
	if(ao.Version >= 20160429){
	ao.BuyerDeviceSource, err = bs.PopString()	//<std::string> 设备来源
	}
	if(ao.Version >= 20160429){
	ao.BuyerDeviceSource_u, err = bs.PopUint8()	//<uint8_t>
	}
	if(ao.Version >= 20160523){
	ao.PartnerType, err = bs.PopUint32()	//<uint32_t> 商家类型
	}
	if(ao.Version >= 20160523){
	ao.PartnerType_u, err = bs.PopUint8()	//<uint8_t>
	}
	if(ao.Version >= 20160523){
	ao.SyncCustomsFlag, err = bs.PopUint8()	//<uint8_t> 是否同步海关 1:是,0:否
	}
	if(ao.Version >= 20160523){
	ao.SyncCustomsFlag_u, err = bs.PopUint8()	//<uint8_t>
	}

	ao.Compat(bs, classLen, startPop)
	return nil
}


type AoOrderListDealListQueryPoXXOO struct {
	AoXXOO

	/**
	 * 买家id
	 *
	 * 版本>=0
	 */
	BuyerId uint64

	/**
	 * 版本>=0
	 */
	BuyerId_u uint8

	/**
	 *  卖家id
	 *
	 *  版本>=0
	 */
	SellerId uint64

	SellerId_u uint8

	SellerShopId uint64 // 商铺id(版本>=0)
	SellerShopId_u 	uint8	//<uint8_t> (版本>=0)
	DealSource 	uint32	//<uint32_t> 订单来源(版本>=0)
	DealSource_u 	uint8	//<uint8_t> (版本>=0)
	DealType 	uint32	//<uint32_t> 订单类型(版本>=0)
	DealType_u 	uint8	//<uint8_t> (版本>=0)
	DealPayType 	uint32	//<uint32_t> 订单支付类型(版本>=0)
	DealPayType_u 	uint8	//<uint8_t> (版本>=0)
	DealState 	uint32	//<uint32_t> 订单状态(版本>=0)
	DealState_u 	uint8	//<uint8_t> (版本>=0)
	DealRefundState 	uint32	//<uint32_t> 退款状态(版本>=0)
	DealRefundState_u 	uint8	//<uint8_t> (版本>=0)
	DealEvalState 	uint32	//<uint32_t> 评价状态(版本>=0)
	DealEvalState_u 	uint8	//<uint8_t> (版本>=0)
	DealVisibleState 	uint32	//<uint32_t> 可见状态(版本>=0)
	DealVisibleState_u 	uint8	//<uint8_t> (版本>=0)
	DealBitProperty1 	uint32	//<uint32_t> 订单通用属性1(版本>=0)
	DealBitProperty1_u 	uint8	//<uint8_t> (版本>=0)
	DealBitProperty2 	uint32	//<uint32_t> 订单通用属性2(版本>=0)
	DealBitProperty2_u 	uint8	//<uint8_t> (版本>=0)
	ItemId 	string	//<std::string> 商品id(版本>=0)
	ItemId_u 	uint8	//<uint8_t> (版本>=0)
	ItemSkuId 	string	//<std::string> 商品sku id(版本>=0)
	ItemSkuId_u 	uint8	//<uint8_t> (版本>=0)
	RecvName 	string	//<std::string> 收货人名称(版本>=0)
	RecvName_u 	uint8	//<uint8_t> (版本>=0)
	RecvMobile 	uint64	//<uint64_t> 收货人名称(版本>=0)
	RecvMobile_u 	uint8	//<uint8_t> (版本>=0)
	DealIdList 	[]uint64	//<std::vector<uint64_t> > 订单id列表(版本>=0)
	DealIdList_t 	uint8	//<uint8_t> (版本>=0)
	TimeType 	uint32	//<uint32_t> 时间类型(版本>=0)
	TimeType_u 	uint8	//<uint8_t> (版本>=0)
	StartTime 	uint32	//<uint32_t> 开始时间，以天为单位的时间戳，即某天的0点0分0秒(版本>=0)
	StartTime_u 	uint8	//<uint8_t> (版本>=0)
	EndTime 	uint32	//<uint32_t> 结束时间，以天为单位的时间戳，即某天的0点0分0秒(版本>=0)
	EndTime_u 	uint8	//<uint8_t> (版本>=0)
	SortType 	uint32	//<uint32_t> 排序类型(版本>=0)
	SortType_u 	uint8	//<uint8_t> (版本>=0)
	StartIndex 	uint32	//<uint32_t> 开始序号(版本>=0)
	StartIndex_u 	uint8	//<uint8_t> (版本>=0)
	PageSize 	uint32	//<uint32_t> 页码(版本>=0)
	PageSize_u 	uint8	//<uint8_t> (版本>=0)
	Reserve 	string	//<std::string> 保留字段(版本>=0)
	Reserve_u 	uint8	//<uint8_t> (版本>=0)
	BdealCode 	string	//<std::string> 交易单编码,带分库分表规则，暂不用(版本>=0)
	BdealCode_u 	uint8	//<uint8_t> (版本>=0)
	DealCode 	string	//<std::string> 订单编码,带分库分表规则，暂不用(版本>=0)
	DealCode_u 	uint8	//<uint8_t> (版本>=0)
	DealSourceOpt 	uint32	//<uint32_t> 订单来源选项(200-全部订单，101-线下订单,102-线上订单,其他待扩展)(版本>=20151130)
	DealSourceOpt_u 	uint8	//<uint8_t> 订单来源选项Flag(版本>=20151130)

}

func NewAoOrderListDealListQueryPoXXOO() *AoOrderListDealListQueryPoXXOO {
	model := &AoOrderListDealListQueryPoXXOO{}

	return model
}

func (ao *AoOrderListDealListQueryPoXXOO) SetBuyerId(buyerId uint64) {
	ao.BuyerId = buyerId
	ao.BuyerId_u = 1
}

func (ao *AoOrderListDealListQueryPoXXOO) SetDealCode(dealCode string)  {
	ao.DealCode = dealCode
	ao.DealCode_u = 1
}

func (ao *AoOrderListDealListQueryPoXXOO) SetBDealCode(bDealCode string)  {
	ao.BdealCode = bDealCode
	ao.BdealCode_u = 1
}

func (ao *AoOrderListDealListQueryPoXXOO) SetStartIndex(startIndex uint32)  {
	ao.StartIndex = startIndex
	ao.StartIndex_u = 1
}

func (ao *AoOrderListDealListQueryPoXXOO) SetPageSize(pageSize uint32)  {
	ao.PageSize = pageSize
	ao.PageSize_u = 1
}

func (ao *AoOrderListDealListQueryPoXXOO) Serialize(bs *ByteStream) bool {

	bs.PushUint32(uint32(ao.getClassLen()))

	return ao.serialize_internal(bs)
}

func (ao *AoOrderListDealListQueryPoXXOO) serialize_internal(bs *ByteStream) bool {
	bs.PushUint32(ao.Version)
	bs.PushUint8(ao.Version_u)

	bs.PushUint64(ao.BuyerId);	//<uint64_t> 买家id
	bs.PushUint8(ao.BuyerId_u);	//<uint8_t>
	bs.PushUint64(ao.SellerId);	//<uint64_t> 卖家id
	bs.PushUint8(ao.SellerId_u);	//<uint8_t>
	bs.PushUint64(ao.SellerShopId);	//<uint64_t> 商铺id
	bs.PushUint8(ao.SellerShopId_u);	//<uint8_t>
	bs.PushUint32(ao.DealSource);	//<uint32_t> 订单来源
	bs.PushUint8(ao.DealSource_u);	//<uint8_t>
	bs.PushUint32(ao.DealType);	//<uint32_t> 订单类型
	bs.PushUint8(ao.DealType_u);	//<uint8_t>
	bs.PushUint32(ao.DealPayType);	//<uint32_t> 订单支付类型
	bs.PushUint8(ao.DealPayType_u);	//<uint8_t>
	bs.PushUint32(ao.DealState);	//<uint32_t> 订单状态
	bs.PushUint8(ao.DealState_u);	//<uint8_t>
	bs.PushUint32(ao.DealRefundState);	//<uint32_t> 退款状态
	bs.PushUint8(ao.DealRefundState_u);	//<uint8_t>
	bs.PushUint32(ao.DealEvalState);	//<uint32_t> 评价状态
	bs.PushUint8(ao.DealEvalState_u);	//<uint8_t>
	bs.PushUint32(ao.DealVisibleState);	//<uint32_t> 可见状态
	bs.PushUint8(ao.DealVisibleState_u);	//<uint8_t>
	bs.PushUint32(ao.DealBitProperty1);	//<uint32_t> 订单通用属性1
	bs.PushUint8(ao.DealBitProperty1_u);	//<uint8_t>
	bs.PushUint32(ao.DealBitProperty2);	//<uint32_t> 订单通用属性2
	bs.PushUint8(ao.DealBitProperty2_u);	//<uint8_t>
	bs.PushString(ao.ItemId);	//<std::string> 商品id
	bs.PushUint8(ao.ItemId_u);	//<uint8_t>
	bs.PushString(ao.ItemSkuId);	//<std::string> 商品sku id
	bs.PushUint8(ao.ItemSkuId_u);	//<uint8_t>
	bs.PushString(ao.RecvName);	//<std::string> 收货人名称
	bs.PushUint8(ao.RecvName_u);	//<uint8_t>
	bs.PushUint64(ao.RecvMobile);	//<uint64_t> 收货人名称
	bs.PushUint8(ao.RecvMobile_u);	//<uint8_t>

	var dealIdList []uint64
	for _, v := range ao.DealIdList {
		dealIdList = append(dealIdList, v)
	}
	bs.PushUint64Set(dealIdList) // <std::vector<uint64_t> > 订单id列表
	//$bs->pushObject($this->dealIdList,'stl_vector');	//<std::vector<uint64_t> > 订单id列表

	bs.PushUint8(ao.DealIdList_t);	//<uint8_t>
	bs.PushUint32(ao.TimeType);	//<uint32_t> 时间类型
	bs.PushUint8(ao.TimeType_u);	//<uint8_t>
	bs.PushUint32(ao.StartTime);	//<uint32_t> 开始时间，以天为单位的时间戳，即某天的0点0分0秒
	bs.PushUint8(ao.StartTime_u);	//<uint8_t>
	bs.PushUint32(ao.EndTime);	//<uint32_t> 结束时间，以天为单位的时间戳，即某天的0点0分0秒
	bs.PushUint8(ao.EndTime_u);	//<uint8_t>
	bs.PushUint32(ao.SortType);	//<uint32_t> 排序类型
	bs.PushUint8(ao.SortType_u);	//<uint8_t>
	bs.PushUint32(ao.StartIndex);	//<uint32_t> 开始序号
	bs.PushUint8(ao.StartIndex_u);	//<uint8_t>
	bs.PushUint32(ao.PageSize);	//<uint32_t> 页码
	bs.PushUint8(ao.PageSize_u);	//<uint8_t>
	bs.PushString(ao.Reserve);	//<std::string> 保留字段
	bs.PushUint8(ao.Reserve_u);	//<uint8_t>
	bs.PushString(ao.BdealCode);	//<std::string> 交易单编码,带分库分表规则，暂不用
	bs.PushUint8(ao.BdealCode_u);	//<uint8_t>
	bs.PushString(ao.DealCode);	//<std::string> 订单编码,带分库分表规则，暂不用
	bs.PushUint8(ao.DealCode_u);	//<uint8_t>

	if(ao.Version >= 20151130){
		bs.PushUint32(ao.DealSourceOpt) //<uint32_t> 订单来源选项(200-全部订单，101-线下订单,102-线上订单,其他待扩展)
	}
	if(ao.Version >= 20151130){
		bs.PushUint8(ao.DealSourceOpt_u);	//<uint8_t> 订单来源选项Flag
	}
	return bs.IsGood()
}

func (ao *AoOrderListDealListQueryPoXXOO) getClassLen() int {

	bsLen := NewByteStream()
	bsLen.SetRealWrite(false)
	ao.serialize_internal(bsLen)
	length := bsLen.GetWrittenLength()

	return length
}

func (ao *AoOrderListDealListQueryPoXXOO) UnSerialize(bs *ByteStream) error {
	classLen, err := bs.PopUint32()

	if err != nil {
		return err
	}
	startPop := bs.GetReadLength()

	ao.Version, err = bs.PopUint32();	//<uint32_t> 版本号
	if err != nil {
		return err
	}

	//$this->version = ao.Version, err;
	ao.Version_u, err = bs.PopUint8();	//<uint8_t>
	if err != nil {
		return err
	}

	ao.BuyerId, err = bs.PopUint64();	//<uint64_t> 买家id
	if err != nil {
		return err
	}

	ao.BuyerId_u, err = bs.PopUint8();	//<uint8_t>
	if err != nil {
		return err
	}

	ao.SellerId, err = bs.PopUint64();	//<uint64_t> 卖家id
	if err != nil {
		return err
	}

	ao.SellerId_u, err = bs.PopUint8();	//<uint8_t>
	if err != nil {
		return err
	}

	ao.SellerShopId, err = bs.PopUint64();	//<uint64_t> 商铺id
	if err != nil {
		return err
	}

	ao.SellerShopId_u, err = bs.PopUint8();	//<uint8_t>
	if err != nil {
		return err
	}

	ao.DealSource, err = bs.PopUint32();	//<uint32_t> 订单来源
	if err != nil {
		return err
	}

	ao.DealSource_u, err = bs.PopUint8();	//<uint8_t>
	if err != nil {
		return err
	}

	ao.DealType, err = bs.PopUint32();	//<uint32_t> 订单类型
	if err != nil {
		return err
	}

	ao.DealType_u, err = bs.PopUint8();	//<uint8_t>
	if err != nil {
		return err
	}

	ao.DealPayType, err = bs.PopUint32();	//<uint32_t> 订单支付类型
	if err != nil {
		return err
	}

	ao.DealPayType_u, err = bs.PopUint8();	//<uint8_t>
	if err != nil {
		return err
	}

	ao.DealState, err = bs.PopUint32();	//<uint32_t> 订单状态
	if err != nil {
		return err
	}

	ao.DealState_u, err = bs.PopUint8();	//<uint8_t>
	if err != nil {
		return err
	}

	ao.DealRefundState, err = bs.PopUint32();	//<uint32_t> 退款状态
	if err != nil {
		return err
	}

	ao.DealRefundState_u, err = bs.PopUint8();	//<uint8_t>
	if err != nil {
		return err
	}

	ao.DealEvalState, err = bs.PopUint32();	//<uint32_t> 评价状态
	if err != nil {
		return err
	}

	ao.DealEvalState_u, err = bs.PopUint8();	//<uint8_t>
	if err != nil {
		return err
	}

	ao.DealVisibleState, err = bs.PopUint32();	//<uint32_t> 可见状态
	if err != nil {
		return err
	}

	ao.DealVisibleState_u, err = bs.PopUint8();	//<uint8_t>
	if err != nil {
		return err
	}

	ao.DealBitProperty1, err = bs.PopUint32();	//<uint32_t> 订单通用属性1
	if err != nil {
		return err
	}

	ao.DealBitProperty1_u, err = bs.PopUint8();	//<uint8_t>
	if err != nil {
		return err
	}

	ao.DealBitProperty2, err = bs.PopUint32();	//<uint32_t> 订单通用属性2
	if err != nil {
		return err
	}

	ao.DealBitProperty2_u, err = bs.PopUint8();	//<uint8_t>
	if err != nil {
		return err
	}

	ao.ItemId, err = bs.PopString();	//<std::string> 商品id
	if err != nil {
		return err
	}

	ao.ItemId_u, err = bs.PopUint8();	//<uint8_t>
	if err != nil {
		return err
	}

	ao.ItemSkuId, err = bs.PopString();	//<std::string> 商品sku id
	if err != nil {
		return err
	}

	ao.ItemSkuId_u, err = bs.PopUint8();	//<uint8_t>
	if err != nil {
		return err
	}

	ao.RecvName, err = bs.PopString();	//<std::string> 收货人名称
	if err != nil {
		return err
	}

	ao.RecvName_u, err = bs.PopUint8();	//<uint8_t>
	if err != nil {
		return err
	}

	ao.RecvMobile, err = bs.PopUint64();	//<uint64_t> 收货人名称
	if err != nil {
		return err
	}

	ao.RecvMobile_u, err = bs.PopUint8();	//<uint8_t>
	if err != nil {
		return err
	}

	//ao.DealIdList, err = bs.PopObject('stl_vector<uint64_t>');	//<std::vector<uint64_t> > 订单id列表
	ao.DealIdList_t, err = bs.PopUint8();	//<uint8_t>
	if err != nil {
		return err
	}

	ao.TimeType, err = bs.PopUint32();	//<uint32_t> 时间类型
	if err != nil {
		return err
	}

	ao.TimeType_u, err = bs.PopUint8();	//<uint8_t>
	if err != nil {
		return err
	}

	ao.StartTime, err = bs.PopUint32();	//<uint32_t> 开始时间，以天为单位的时间戳，即某天的0点0分0秒
	ao.StartTime_u, err = bs.PopUint8();	//<uint8_t>
	if err != nil {
		return err
	}

	ao.EndTime, err = bs.PopUint32();	//<uint32_t> 结束时间，以天为单位的时间戳，即某天的0点0分0秒
	if err != nil {
		return err
	}

	ao.EndTime_u, err = bs.PopUint8();	//<uint8_t>
	if err != nil {
		return err
	}

	ao.SortType, err = bs.PopUint32();	//<uint32_t> 排序类型
	if err != nil {
		return err
	}

	ao.SortType_u, err = bs.PopUint8();	//<uint8_t>
	if err != nil {
		return err
	}

	ao.StartIndex, err = bs.PopUint32();	//<uint32_t> 开始序号
	if err != nil {
		return err
	}

	ao.StartIndex_u, err = bs.PopUint8();	//<uint8_t>
	if err != nil {
		return err
	}

	ao.PageSize, err = bs.PopUint32();	//<uint32_t> 页码
	if err != nil {
		return err
	}

	ao.PageSize_u, err = bs.PopUint8();	//<uint8_t>
	if err != nil {
		return err
	}

	ao.Reserve, err = bs.PopString();	//<std::string> 保留字段
	if err != nil {
		return err
	}

	ao.Reserve_u, err = bs.PopUint8();	//<uint8_t>
	if err != nil {
		return err
	}

	ao.BdealCode, err = bs.PopString();	//<std::string> 交易单编码,带分库分表规则，暂不用
	if err != nil {
		return err
	}

	ao.BdealCode_u, err = bs.PopUint8();	//<uint8_t>
	ao.DealCode, err = bs.PopString();	//<std::string> 订单编码,带分库分表规则，暂不用
	if err != nil {
		return err
	}

	ao.DealCode_u, err = bs.PopUint8();	//<uint8_t>
	if err != nil {
		return err
	}

	if(ao.Version >= 20151130){
		ao.DealSourceOpt, err = bs.PopUint32() //<uint32_t> 订单来源选项(200-全部订单，101-线下订单,102-线上订单,其他待扩展)
		if err != nil {
			return err
		}
	}
	if(ao.Version >= 20151130){
		ao.DealSourceOpt_u, err = bs.PopUint8() //<uint8_t> 订单来源选项Flag
		if err != nil {
			return err
		}
	}
	ao.Compat(bs, classLen, startPop)
	return nil
}
