package goao

import (
	"fmt"
)

type Pkg struct {
	dwLength   uint32
	dwSerialNo uint32
	wVersion   uint16 //2
	wCommand   uint16 //= 0
	dwUin      uint32 //= 0

	dwFlag            uint32
	dwResult          uint32 //= 0
	dwClientIP        uint32 //= 0
	wClientPort       uint16 //= 0
	dwAccessServerIP  uint32 //= 0
	wAccessServerPort uint16 //= 0
	dwAppInterfaceIP  uint32 //= 0
	wAppInterfacePort uint16 //= 0
	dwAppServerIP     uint32 //= 0
	wAppServerPort    uint16 //= 0

	dwOperatorId uint32 //= 0
	sPassport    string //= "1234567890"

	wSeconds   uint16 //= 0
	dwUSeconds uint32 //= 0

	wCookie2Length uint16 //= 35
	cLegacy        byte   //= 0
	iSockfd        uint32 //= 0
	dwSockChannel  uint32 //= 0
	wPeerPort      uint16 //= 0
	dwCommand      uint32 //= 0
	wSvrLevelIndex uint16 //= 0
	bCookie2       string //= "123456789012345678"

	iPkgHeadLength int //= 105
}

func NewPkg() *Pkg {
	p := Pkg{}
	p.dwLength = 0
	p.wCommand = 0
	p.dwUin = 0
	p.dwResult = 0
	p.dwClientIP = 0
	p.wClientPort = 0
	p.dwAccessServerIP = 0
	p.wAccessServerPort = 0
	p.dwAppInterfaceIP = 0
	p.wAppInterfacePort = 0
	p.dwAppServerIP = 0
	p.wAppServerPort = 0
	p.dwOperatorId = 0
	p.wSeconds = 0
	p.dwUSeconds = 0
	p.cLegacy = 0
	p.iSockfd = 0
	p.dwSockChannel = 0
	p.wPeerPort = 0
	p.dwCommand = 0
	p.wSvrLevelIndex = 0

	p.wVersion = 2
	p.dwSerialNo = 0
	p.wCookie2Length = 35
	p.sPassport = "1234567890"        //[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	p.bCookie2 = "123456789012345678" //[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8}
	p.iPkgHeadLength = 105
	return &p
}

func (p *Pkg) Serialize(bs *ByteStream) bool {
	bs.PushUint32(p.dwLength)             //4
	bs.PushUint32(p.dwSerialNo)           //8
	bs.PushUint16(p.wVersion)             //10
	bs.PushUint16(p.wCommand)             //12
	bs.PushUint32(p.dwUin)                //16
	bs.PushUint32(p.dwFlag)               //20
	bs.PushUint32(p.dwResult)             //24
	bs.PushUint32(p.dwClientIP)           //28
	bs.PushUint16(p.wClientPort)          //30
	bs.PushUint32(p.dwAccessServerIP)     //34
	bs.PushUint16(p.wAccessServerPort)    //36
	bs.PushUint32(p.dwAppInterfaceIP)     //40
	bs.PushUint16(p.wAppInterfacePort)    //42
	bs.PushUint32(p.dwAppServerIP)        //46
	bs.PushUint16(p.wAppServerPort)       //48
	bs.PushUint32(p.dwOperatorId)         //52
	bs.PushBytes([]byte(p.sPassport), 10) //62
	bs.PushUint16(p.wSeconds)             //64
	bs.PushUint32(p.dwUSeconds)           //68
	bs.PushUint16(p.wCookie2Length)       //70
	bs.PushByte(p.cLegacy)                //71
	bs.PushUint32(p.iSockfd)              //75
	bs.PushUint32(p.dwSockChannel)        //79
	bs.PushUint16(p.wPeerPort)            //81
	bs.PushUint32(p.dwCommand)            //85
	bs.PushUint16(p.wSvrLevelIndex)       //87
	bs.PushBytes([]byte(p.bCookie2), 18)  //105
	return bs.IsGood()
}
func (p *Pkg) UnSerialize(bs *ByteStream) error {
	var err error
	p.dwLength, err = bs.PopUint32() //4
	if err != nil {
		return err
	}

	p.dwSerialNo, err = bs.PopUint32() //8
	if err != nil {
		return err
	}
	p.wVersion, err = bs.PopUint16() //10
	if err != nil {
		return err
	}
	p.wCommand, err = bs.PopUint16() //12
	if err != nil {
		return err
	}
	p.dwUin, err = bs.PopUint32() //16

	if err != nil {
		return err
	}
	p.dwFlag, err = bs.PopUint32() //20
	if err != nil {
		return err
	}
	p.dwResult, err = bs.PopUint32() //24
	if err != nil {
		return err
	}
	p.dwClientIP, err = bs.PopUint32() //28
	if err != nil {
		return err
	}
	p.wClientPort, err = bs.PopUint16() //30
	if err != nil {
		return err
	}
	p.dwAccessServerIP, err = bs.PopUint32() //34
	if err != nil {
		return err
	}
	p.wAccessServerPort, err = bs.PopUint16() //36
	if err != nil {
		return err
	}
	p.dwAppInterfaceIP, err = bs.PopUint32() //40
	if err != nil {
		return err
	}
	p.wAppInterfacePort, err = bs.PopUint16() //42
	if err != nil {
		return err
	}
	p.dwAppServerIP, err = bs.PopUint32() //46
	if err != nil {
		return err
	}
	p.wAppServerPort, err = bs.PopUint16() //48
	if err != nil {
		return err
	}
	p.dwOperatorId, err = bs.PopUint32() //52
	if err != nil {
		return err
	}
	var bytePassport []byte

	bytePassport, err = bs.PopBytes(10) //62

	if err != nil {
		return err
	}

	p.sPassport = string(bytePassport)

	p.wSeconds, err = bs.PopUint16() //64
	if err != nil {
		return err
	}
	p.dwUSeconds, err = bs.PopUint32() //68
	if err != nil {
		return err
	}
	p.wCookie2Length, err = bs.PopUint16() //70
	if err != nil {
		return err
	}
	p.cLegacy, err = bs.PopByte() //71
	if err != nil {
		return err
	}
	p.iSockfd, err = bs.PopUint32() //75
	if err != nil {
		return err
	}
	p.dwSockChannel, err = bs.PopUint32() //79
	if err != nil {
		return err
	}
	p.wPeerPort, err = bs.PopUint16() //81
	if err != nil {
		return err
	}
	p.dwCommand, err = bs.PopUint32() //85
	if err != nil {
		return err
	}
	p.wSvrLevelIndex, err = bs.PopUint16() //87
	if err != nil {
		return err
	}
	var byteCookie []byte
	byteCookie, err = bs.PopBytes(18) //105
	if err != nil {
		return err
	}
	p.bCookie2 = string(byteCookie)

	return nil //bs.iBufLength
}
func (p *Pkg) SetDwCommand(cmd uint32) {
	p.dwCommand = cmd
	p.wCommand = uint16(cmd / 0x10000)
}

func (p *Pkg) SetDwLength(length uint32) {
	p.dwLength = length
}

func (p *Pkg) SetSPassword(sps string) {
	p.sPassport = sps[:10]
}

func (p *Pkg) SetDwOperatorId(operatorId int64) {
	p.dwOperatorId = uint32(operatorId)
}
