package goao

type AoUserRecvCouponParamXXOO struct {
	AoXXOO

	/**
	 * 版本 >= 0
	 */

	Version_u uint8

	/**
	 * 批次号MD5值
	 *
	 * 版本 >= 0
	 */
	ResourceBatchMd5 string

	/**
	 * 批次号MD5值flag
	 *
	 * 版本 >= 0
	 */
	ResourceBatchMd5_u uint8
	/**
	 * 用户ID
	 *
	 * 版本 >= 0
	 */
	UserId uint64

	/**
	 * 版本 >= 0
	 */
	UserId_u uint8

	/**
	 * 手机号码
	 *
	 * 版本 >= 0
	 */
	UserMobile uint32

	/**
	 * 版本 >= 0
	 */
	UserMobile_u uint8

	/**
	 * 保留字段str
	 *
	 * 版本 >= 0
	 */
	ReserveStr string

	/**
	 * 版本 >= 0
	 */
	ReserveStr_u uint8
}

func NewAoUserRecvCouponParamXXOO() *AoUserRecvCouponParamXXOO {
	model := &AoUserRecvCouponParamXXOO{}

	return model
}

func (ao *AoUserRecvCouponParamXXOO) SetUserId(uid uint64) {
	ao.UserId = uid
	ao.UserId_u = 1
}

func (ao *AoUserRecvCouponParamXXOO) SetResourceBatchMd5(resourceBatchMd5 string) {
	ao.ResourceBatchMd5 = resourceBatchMd5
	ao.ResourceBatchMd5_u = 1
}

func (ao *AoUserRecvCouponParamXXOO) Serialize(bs *ByteStream) bool {

	bs.PushUint32(uint32(ao.getClassLen()))

	return ao.serialize_internal(bs)
}

func (ao *AoUserRecvCouponParamXXOO) serialize_internal(bs *ByteStream) bool {
	bs.PushUint32(ao.Version)
	bs.PushUint8(ao.Version_u)
	bs.PushString(ao.ResourceBatchMd5)
	bs.PushUint8(ao.ResourceBatchMd5_u)
	bs.PushUint64(ao.UserId)
	bs.PushUint8(ao.UserId_u)
	bs.PushUint32(ao.UserMobile)
	bs.PushUint8(ao.UserMobile_u)
	bs.PushString(ao.ReserveStr)
	bs.PushUint8(ao.ReserveStr_u)

	return bs.IsGood()
}

func (ao *AoUserRecvCouponParamXXOO) getClassLen() int {

	bsLen := NewByteStream()
	bsLen.SetRealWrite(false)
	ao.serialize_internal(bsLen)
	length := bsLen.GetWrittenLength()

	return length
}

func (ao *AoUserRecvCouponParamXXOO) UnSerialize(bs *ByteStream) error {
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
	ao.ResourceBatchMd5, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.ResourceBatchMd5_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.UserId, err = bs.PopUint64()
	if err != nil {
		return err
	}
	ao.UserId_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.UserMobile, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.UserMobile_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.ReserveStr, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.ReserveStr_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.Compat(bs, classLen, startPop)
	return nil
}
