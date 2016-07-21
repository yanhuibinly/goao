package goao

/*ModifyTaskInfo start*/
type AoModifyTaskInfoReq struct {
    AuthCode string
    ChildTaskInfo  *AoChildTaskInfoPoXXOO
}

func NewAoModifyTaskInfoReq() *AoModifyTaskInfoReq {
    model := &AoModifyTaskInfoReq{}
    model.AuthCode = "^YHN,ki8"
    return model
}

func (u *AoModifyTaskInfoReq) GetCmdId() uint32 {
    return 0x40291802
}
func (u *AoModifyTaskInfoReq) Serialize(bs *ByteStream, req *Request) bool {
    bs.PushString(req.MachineKey)
    bs.PushString(req.Source)
    bs.PushUint32(req.SceneId)
    bs.PushString(u.AuthCode)
    u.ChildTaskInfo.Serialize(bs)
    bs.PushString(req.InReserve)

    return bs.IsGood()
}

type AoModifyTaskInfoRsp struct {
    AoRsp
}

func NewAoModifyTaskInfoRsp() *AoModifyTaskInfoRsp {
    model := &AoModifyTaskInfoRsp{}
    return model
}

func (u *AoModifyTaskInfoRsp) UnSerialize(bs *ByteStream) (bool, error) {

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
    var errOutResever error
    u.OutReserve, errOutResever = bs.PopString()
    if errOutResever != nil {
        return false, errOutResever
    }

    return bs.IsGood(), nil
}
/*ModifyTaskInfo end*/