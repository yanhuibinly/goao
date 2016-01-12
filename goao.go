package goao

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"errors"
	"fmt"
	"net"
	"strings"
	"time"
)

type GoAo struct {
	ph      *Pkg
	stx     uint16
	byteStx byte
	etx     byte
}

func New() *GoAo {
	goao := &GoAo{}
	goao.ph = NewPkg()
	goao.stx = 0x55AA
	goao.byteStx = 2
	goao.etx = 3
	return goao
}

func (g *GoAo) Call(req *Request, iao IAoReq, iaoRsp IAoRsp) (interface{}, error) {
	if req == nil || iao == nil {
		return nil, errors.New("request or iao is null")
	}

	if strings.Trim(req.Host, " ") == "" {
		return nil, errors.New("address is null")
	}
	if iao.GetCmdId() <= 0 {
		return nil, errors.New("cmdId is invalid")
	}

	conn, errConn := net.Dial("tcp", req.Host)

	if errConn != nil {
		return nil, errConn
	}

	defer conn.Close()

	conn.SetDeadline(time.Now().Add(3000 * time.Millisecond))

	dataWrite := g.serialize(iao, req)

	_, errWrite := conn.Write(dataWrite)

	if errWrite != nil {
		return nil, errWrite
	}

	var byteStx = make([]byte, 2)

	_, errRead := conn.Read(byteStx)

	if errRead != nil {
		return nil, errRead
	}

	if binary.BigEndian.Uint16(byteStx) != 21930 {
		return nil, errors.New(fmt.Sprintf("stx is error :%d", byteStx))
	}

	var byteLength = make([]byte, 4)

	_, errRead = conn.Read(byteLength)

	if errRead != nil {
		return nil, errRead
	}

	dwLength := binary.BigEndian.Uint32(byteLength)

	var byteRes = make([]byte, dwLength)

	_, errRead = conn.Read(byteRes)

	if errRead != nil {
		return nil, errRead
	}

	return g.unSerialize(iaoRsp, byteRes, len(byteRes))
}

func (g *GoAo) serialize(iao IAoReq, req *Request) []byte {

	// 计算Request包长度
	bsdummy := NewByteStream()

	bsdummy.SetRealWrite(false)

	iao.Serialize(bsdummy, req)

	dwBodyLen := bsdummy.GetWrittenLength()

	// 计算整个包长度
	dwPkgLength := uint32(1 + g.ph.iPkgHeadLength + dwBodyLen + 1)

	// 构建包头
	g.ph.SetDwCommand(iao.GetCmdId())
	g.ph.SetDwLength(dwPkgLength)
	// 构建序列化buffer
	bs := NewByteStream()

	// 序列化
	bs.PushUint16(g.stx)
	bs.PushUint32(dwPkgLength)
	bs.PushByte(g.byteStx)
	g.ph.Serialize(bs)

	iao.Serialize(bs, req)
	bs.PushByte(g.etx)
	return bs.GetWriteBuffer()
}
func (g *GoAo) unSerialize(iao IAoRsp, byteRes []byte, length int) (interface{}, error) {

	bs := NewByteStream()
	bs.Reset(byteRes, length)

	stx, err := bs.PopByte()

	if err != nil {
		return nil, err
	}

	if stx != 2 {
		return nil, errors.New("stx is not 2")
	}

	g.ph.UnSerialize(bs)

	_, err = iao.UnSerialize(bs)

	if err != nil {
		return nil, err
	}
	return iao, nil
}

func getBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (g *GoAo) SetSPassword(sps string) {
	g.ph.SetSPassword(sps)
}

func (g *GoAo) SetDwOperatorId(operatorId int64) {
	g.ph.SetDwOperatorId(operatorId)
}
