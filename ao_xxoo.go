package goao

type AoXXOO struct {
	Version          uint32 //<uint32_t> 版本号
	AddTime          uint32 //<uint32_t> 添加时间
	AddTime_u        uint8
	LastUpdateTime   uint32 //<uint32_t> 最后修改时间
	LastUpdateTime_u uint8
	Reserve          string //<std::string> 预留字段
	Reserve_u        uint8
}

func (a *AoXXOO) Compat(bs *ByteStream, classLen uint32, startPop int) {
	/**********************为了支持多个版本的客户端************************/
	needPopLen := int(classLen) - (bs.GetReadLength() - startPop)

	for i := 0; i < needPopLen; i++ {
		bs.PopUint8()
	}
	/**********************为了支持多个版本的客户端************************/
}
