package goao

type AoChildTaskInfoPoXXOO struct {
    AoXXOO
    TaskId uint64
    TaskId_u uint8
    FatherTaskId uint64
    FatherTaskId_u uint8
    TaskBatchId uint32
    TaskBatchId_u uint8
    Uid uint64
    Uid_u uint8
    TaskName string
    TaskName_u uint8
    TaskDesc string
    TaskDesc_u uint8
    TaskPic string
    TaskPic_u uint8
    GuideJumpUrl string
    GuideJumpUrl_u uint8
    TaskStatus map[uint32]uint32
    TaskStatus_u uint8
}

func NewAoChildTaskInfoPoXXOO() *AoChildTaskInfoPoXXOO {
    obj := &AoChildTaskInfoPoXXOO{}
    return obj
}

func (ao *AoChildTaskInfoPoXXOO) Serialize(bs *ByteStream) bool {
    bs.PushUint32(uint32(ao.getClassLen()))
    return ao.serialize_internal(bs)
}

func (ao *AoChildTaskInfoPoXXOO) serialize_internal(bs *ByteStream) bool {
    bs.PushUint32(ao.Version)
    bs.PushUint64(ao.TaskId)
    bs.PushUint8(ao.TaskId_u)
    bs.PushUint64(ao.FatherTaskId)
    bs.PushUint8(ao.FatherTaskId_u)
    bs.PushUint32(ao.TaskBatchId)
    bs.PushUint8(ao.TaskBatchId_u)
    bs.PushUint64(ao.Uid)
    bs.PushUint8(ao.Uid_u)
    bs.PushString(ao.TaskName)
    bs.PushUint8(ao.TaskName_u)
    bs.PushString(ao.TaskDesc)
    bs.PushUint8(ao.TaskDesc_u)
    bs.PushString(ao.TaskPic)
    bs.PushUint8(ao.TaskPic_u)
    bs.PushString(ao.GuideJumpUrl)
    bs.PushUint8(ao.GuideJumpUrl_u)
    bs.PushMapUInt32UInt32(ao.TaskStatus)
    bs.PushUint8(ao.TaskStatus_u)
    bs.PushString(ao.Reserve)
    bs.PushUint8(ao.Reserve_u)
    bs.PushUint32(ao.AddTime)
    bs.PushUint8(ao.AddTime_u)
    bs.PushUint32(ao.LastUpdateTime)
    bs.PushUint8(ao.LastUpdateTime_u)

    return bs.IsGood()
}

func (ao *AoChildTaskInfoPoXXOO) getClassLen() int {
    bsLen := NewByteStream()
    bsLen.SetRealWrite(false)
    ao.serialize_internal(bsLen)
    length := bsLen.GetWrittenLength()

    return length
}

func (a *AoChildTaskInfoPoXXOO) UnSerialize(bs *ByteStream) error {
    classLen, err := bs.PopUint32()
    if err != nil {
        return err
    }
    startPop := bs.GetReadLength()

    a.Version, err = bs.PopUint32() //<uint32> 版本号
    if err != nil {
        return err
    }
    a.TaskId, err = bs.PopUint64()
    if err != nil {
        return err
    }
    a.TaskId_u, err = bs.PopUint8()
    if err != nil {
        return err
    }
    a.FatherTaskId, err = bs.PopUint64()
    if err != nil {
        return err
    }
    a.FatherTaskId_u, err = bs.PopUint8()
    if err != nil {
        return err
    }
    a.TaskBatchId, err = bs.PopUint32()
    if err != nil {
        return err
    }
    a.TaskBatchId_u, err = bs.PopUint8()
    if err != nil {
        return err
    }
    a.Uid, err = bs.PopUint64()
    if err != nil {
        return err
    }
    a.Uid_u, err = bs.PopUint8()
    if err != nil {
        return err
    }
    a.TaskName, err = bs.PopString()
    if err != nil {
        return err
    }
    a.TaskName_u, err = bs.PopUint8()
    if err != nil {
        return err
    }
    a.TaskDesc, err = bs.PopString()
    if err != nil {
        return err
    }
    a.TaskDesc_u, err = bs.PopUint8()
    if err != nil {
        return err
    }
    a.TaskPic, err = bs.PopString()
    if err != nil {
        return err
    }
    a.TaskPic_u, err = bs.PopUint8()
    if err != nil {
        return err
    }
    a.GuideJumpUrl, err = bs.PopString()
    if err != nil {
        return err
    }
    a.GuideJumpUrl_u, err = bs.PopUint8()
    if err != nil {
        return err
    }
    a.TaskStatus, err = bs.PopMapUInt32UInt32()
    if err != nil {
        return err
    }
    a.TaskStatus_u, err = bs.PopUint8()
    if err != nil {
        return err
    }
    a.Reserve, err = bs.PopString() //<std::string> 预留字段
    if err != nil {
        return err
    }
    a.Reserve_u, err = bs.PopUint8()
    if err != nil {
        return err
    }
    a.AddTime, err = bs.PopUint32() //<uint32> 添加时间
    if err != nil {
        return err
    }
    a.AddTime_u, err = bs.PopUint8()
    if err != nil {
        return err
    }
    a.LastUpdateTime, err = bs.PopUint32() //<uint32> 最后修改时间
    if err != nil {
        return err
    }
    a.LastUpdateTime_u, err = bs.PopUint8()
    if err != nil {
        return err
    }

    a.Compat(bs, classLen, startPop)
    return nil
}