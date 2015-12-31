package goao

type IAoReq interface {
	Serialize(bs *ByteStream, req *Request) bool

	GetCmdId() uint32
}

type IAoRsp interface {
	UnSerialize(bs *ByteStream) (bool, error)
}
