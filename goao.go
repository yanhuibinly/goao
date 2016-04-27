package goao

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"errors"
	"fmt"
	"github.com/fatih/pool"
	"github.com/tonyjt/tgo"
	"net"
	"strings"
	"sync"
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

var (
	poolsAo map[string]pool.Pool = make(map[string]pool.Pool)

	poolMux sync.Mutex
)

func (g *GoAo) dail(host string) (net.Conn, error) {
	c, err := net.Dial("tcp", host)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("open ao pool error: %s", err.Error()))
	}
	return c, err
}
func (g *GoAo) GetConn(host string) (net.Conn, error) {

	poolAo := poolsAo[host]

	if poolAo == nil {
		poolMux.Lock()

		defer poolMux.Unlock()

		poolAo := poolsAo[host]

		if poolAo == nil {
			var errPool error

			var newPool pool.Pool
			newPool, errPool = pool.NewChannelPool(1, 500, func() (net.Conn, error) {
				return g.dail(host)
			})

			if errPool != nil {
				return nil, errors.New(fmt.Sprintf("create ao pool channel error: %s", errPool.Error()))
			}
			poolsAo[host] = newPool
		}
	}

	poolAo = poolsAo[host]

	if poolAo != nil {

		conn, err := poolAo.Get()

		if err != nil {
			return nil, errors.New(fmt.Sprintf("open ao pool error: %s", err.Error()))
		}

		return conn, err
	}

	return nil, errors.New("ao pool is null")

}

func (g *GoAo) redailConn(c net.Conn, host string) (net.Conn, error) {

	//var one []byte

	//tgo.UtilLogError("ao connection reset try")

	//if _, err := c.Read(one); err != nil {
	//链接断开了
	tgo.UtilLogError("ao connection from pool redail")

	cp := c.(*pool.PoolConn)

	cp.Conn.Close()

	var err error
	cp.Conn, err = g.dail(host)

	if err != nil {
		tgo.UtilLogErrorf("ao connection pool redail failed:%s", err.Error())
		return cp, errors.New(fmt.Sprintf("redail failed,%s", err.Error()))
	}
	return cp, nil

	//}

	//return false, c, nil
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

	conn, errConn := g.GetConn(req.Host)

	if errConn != nil {
		return nil, errConn
	}

	defer conn.Close()

	conn.SetReadDeadline(time.Now().Add(30 * time.Second))

	defer conn.SetReadDeadline(time.Now().Add(72000 * time.Hour))

	dataWrite := g.serialize(iao, req)

	_, errWrite := conn.Write(dataWrite)

	if errWrite != nil {

		tgo.UtilLogErrorf("write data to ao failed:%s", errWrite)

		var errRedail error
		conn, errRedail = g.redailConn(conn, req.Host)

		if errRedail == nil {
			_, errWrite = conn.Write(dataWrite)

			if errWrite != nil {
				return nil, errWrite
			}
		} else {
			return nil, errRedail
		}
	}

	var byteStx = make([]byte, 2)

	_, errRead := conn.Read(byteStx)

	if errRead != nil {

		tgo.UtilLogErrorf("read byte stx failed:%s", errRead)

		if errRead.Error() == "connection reset by peer" {
			var errRedail error
			conn, errRedail = g.redailConn(conn, req.Host)

			if errRedail == nil {
				_, errRead = conn.Read(byteStx)

				if errRead != nil {
					return nil, errRead
				}
			} else {
				return nil, errRedail
			}
		} else {

			conn, _ = g.redailConn(conn, req.Host)

		}

		return nil, errRead
	}

	if binary.BigEndian.Uint16(byteStx) != 21930 {
		return nil, errors.New(fmt.Sprintf("stx is error :%d", byteStx))
	}

	var byteLength = make([]byte, 4)

	_, errRead = conn.Read(byteLength)

	if errRead != nil {
		tgo.UtilLogErrorf("read byte length failed:%s", errRead)

		return nil, errRead
	}

	dwLength := binary.BigEndian.Uint32(byteLength)

	var byteRes = make([]byte, dwLength)

	_, errRead = conn.Read(byteRes)

	if errRead != nil {

		tgo.UtilLogErrorf("read response failed:%s", errRead)

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
