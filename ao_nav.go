
package goao


type AoGetPathByNavIdReq struct {
	MapId             	 uint32
	NavId			 uint32
}

func NewAoGetPathByNavIdReq() *AoGetPathByNavIdReq {
	model := &AoGetPathByNavIdReq{MapId:1}
	return model
}

func (u *AoGetPathByNavIdReq) Serialize(bs *ByteStream, req *Request) bool {
	bs.PushString(req.MachineKey)
	bs.PushString(req.Source)
	bs.PushUint32(req.SceneId)
	bs.PushUint32(u.MapId)
	bs.PushUint32(u.NavId)

	bs.PushString(req.InReserve)

	return bs.IsGood()
}

func (u *AoGetPathByNavIdReq) GetCmdId() uint32 {
	return 0xA001180f
}

type AoGetPathByNavIdRsp struct {
	AoRsp
	FullPath		[]AoGetPathByNavIdNavEntryXXOO
	ChildNav		[]AoGetPathByNavIdNavEntryXXOO
}

func NewAoGetPathByNavIdRsp() *AoGetPathByNavIdRsp {
	model := &AoGetPathByNavIdRsp{}

	return model
}

func (u *AoGetPathByNavIdRsp) UnSerialize(bs *ByteStream) (bool, error) {
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

	lengthNavEntry, err := bs.PopUint32()
	if err != nil {
		return false, err
	}
	for i := uint32(0); i < lengthNavEntry; i++ {

		navEntry := NewAoGetPathByNavIdNavEntryXXOO()
		err = navEntry.UnSerialize(bs)
		if err != nil {
			return false, err
		}
		u.FullPath = append(u.FullPath, *navEntry)
	}

	lengthNavEntry, err = bs.PopUint32()
	if err != nil {
		return false, err
	}
	for i := uint32(0); i < lengthNavEntry; i++ {

		navEntry := NewAoGetPathByNavIdNavEntryXXOO()
		err = navEntry.UnSerialize(bs)
		if err != nil {
			return false, err
		}
		u.ChildNav = append(u.ChildNav, *navEntry)
	}

	var errOutResever error
	u.OutReserve, errOutResever = bs.PopString()
	if errOutResever != nil {
		return false, errOutResever
	}

	return bs.IsGood(), nil
}
