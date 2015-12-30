package goao

type Request struct {
	CmdId      uint32                 //命令字
	Host       string                 //接口地址
	Source     string                 //来源
	MachineKey string                 //机器key
	Uid        uint64                 //用户id
	SceneId    uint32                 //场景id
	InReserve  string                 //编码（如utf8)
	Paras      map[string]interface{} //参数
}

func NewRequest() *Request {
	return &Request{}
}

func (r *Request) Serialize(bs *ByteStream) bool {
	bs.PushString(r.MachineKey)
	bs.PushString(r.Source)
	bs.PushUint32(r.SceneId)
	bs.PushUint64(r.Uid)
	bs.PushString(r.InReserve)

	return bs.IsGood()
}

/*
func (r *Request) UnSerialize(bs *ByteStream) bool {
	r.MachineKey = bs.PushString()
	r.Source = bs.PushString()
	r.SceneId = bs.PushUint32()
	r.Uid = bs.PushUint64()
	r.InReserve = bs.PushString()

	return bs.IsGood()
}*/
