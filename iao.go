package goao

type IAoReq interface {
	Serialize(bs *ByteStream, req *Request) bool
}

type IAoRsp interface {
	UnSerialize(bs *ByteStream) (bool, error)
}
