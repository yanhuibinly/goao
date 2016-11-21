package goao


type AoGetPathByNavIdNavEntryXXOO struct {
	AoXXOO
	Version			uint8
	NavId			uint32
	MapId			uint32
	PNavId			uint32
	Name			string
	Type			uint32
	Catalog			uint32
	Note			string
	Order			uint32
	PropertyStr		string
	SearchCond		string
	HasAttr			uint32
	CustomStr1		string
	CustomStr2		string
	CustomUint1		uint32
	CustomUint2		uint32
	IsPreDelete		uint32
	IsCooperatorFirst	uint32
	IsLowPriceFirst		uint32
	IsHighPriceFirst	uint32
	PropertyMask		[]uint8
	ExtraData		map[uint32]string
}

func NewAoGetPathByNavIdNavEntryXXOO() *AoGetPathByNavIdNavEntryXXOO {
	model := &AoGetPathByNavIdNavEntryXXOO{}

	return model
}

func (ao *AoGetPathByNavIdNavEntryXXOO) UnSerialize(bs *ByteStream) error {
	classLen, err := bs.PopUint32()
	if err != nil {
		return err
	}
	startPop := bs.GetReadLength()

	ao.Version, err = bs.PopUint8()
	if err != nil {
		return err
	}

	ao.NavId, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.MapId, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.PNavId, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.Name, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.Type, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.Catalog, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.Note, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.Order, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.PropertyStr, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.SearchCond, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.HasAttr, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.CustomStr1, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.CustomStr2, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.CustomUint1, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.CustomUint2, err = bs.PopUint32()
	if err != nil {
		return err
	}

	if ao.Version >= 140 {
		ao.IsPreDelete, err = bs.PopUint32()
		if err != nil {
			return err
		}
		ao.IsCooperatorFirst, err = bs.PopUint32()
		if err != nil {
			return err
		}
		ao.IsLowPriceFirst, err = bs.PopUint32()
		if err != nil {
			return err
		}
		ao.IsHighPriceFirst, err = bs.PopUint32()
		if err != nil {
			return err
		}
	}
	if ao.Version >= 141 {
		ao.PropertyMask, err = bs.PopBitSet()
		if err != nil {
			return err
		}
	}
	if ao.Version >= 142 {
		//ao.ExtraData, err = bs.PopMap(reflect.TypeOf(uint32(0)), reflect.TypeOf(""))
		ao.ExtraData, err = bs.PopMapUInt32String()
		if err != nil {
			return err
		}
	}

	ao.Compat(bs, classLen, startPop)
	return nil
}
