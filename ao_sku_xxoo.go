package goao

import (
	"errors"
	"fmt"
)

/* AoDetailmainPoItemInfoXXOO START */
type AoDetailmainPoItemInfoXXOO struct {
	AoXXOO
	BasicInfo   AoDetailmainPoItemBasicInfoXXOO
	BasicInfo_u uint8
	NavInfo     AoDetailmainPoItemNavInfoXXOO
	NavInfo_u   uint8
	PmInfo      AoDetailmainPoItemPmInfoXXOO
	PmInfo_u    uint8
	StockInfo   AoDetailmainPoItemStockInfoXXOO
	StockInfo_u uint8
}

func NewAoDetailmainPoItemInfoXXOO() *AoDetailmainPoItemInfoXXOO {
	model := &AoDetailmainPoItemInfoXXOO{}

	return model
}

/*
func (ao *AoDetailmainPoItemInfoXXOO) Serialize(bs *ByteStream) bool {

	bs.PushUint32(uint32(ao.getClassLen()))

	return ao.serialize_internal(bs)
}

func (ao *AoDetailmainPoItemInfoXXOO) serialize_internal(bs *ByteStream) bool {
	bs.PushUint32(ao.Version)
	bs.PushUint8(ao.Version_u)

	ao.BasicInfo.Serialize(bs)
	bs.PushUint8(ao.BasicInfo_u)

	ao.NavInfo.Serialize(bs)
	bs.PushUint8(ao.NavInfo_u)

	ao.PmInfo.Serialize(bs)
	bs.PushUint8(ao.PmInfo_u)

	ao.StockInfo.Serialize(bs)
	bs.PushUint8(ao.StockInfo_u)

	return bs.IsGood()
}

func (ao *AoDetailmainPoItemInfoXXOO) getClassLen() int {

	bsLen := NewByteStream()
	bsLen.SetRealWrite(false)
	ao.serialize_internal(bsLen)
	length := bsLen.GetWrittenLength()

	return length
}*/

func (ao *AoDetailmainPoItemInfoXXOO) UnSerialize(bs *ByteStream) error {
	classLen, err := bs.PopUint32()

	if err != nil {
		return err
	}
	startPop := bs.GetReadLength()

	ao.Version, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.Version_u, err = bs.PopUint8()
	if err != nil {
		return err
	}

	ao.BasicInfo.UnSerialize(bs)
	if err != nil {
		return err
	}
	ao.BasicInfo_u, err = bs.PopUint8()
	if err != nil {
		return err
	}

	err = ao.NavInfo.UnSerialize(bs)
	if err != nil {
		return err
	}

	ao.NavInfo_u, err = bs.PopUint8()
	if err != nil {
		return err
	}

	err = ao.PmInfo.UnSerialize(bs)
	if err != nil {
		return err
	}
	ao.PmInfo_u, err = bs.PopUint8()
	if err != nil {
		return err
	}

	err = ao.StockInfo.UnSerialize(bs)
	if err != nil {
		return err
	}
	ao.StockInfo_u, err = bs.PopUint8()
	if err != nil {
		return err
	}

	ao.Compat(bs, classLen, startPop)
	return nil
}

/* AoDetailmainPoItemInfoXXOO END */

/* AoDetailmainPoItemBasicInfoXXOO START */

type AoDetailmainPoItemBasicInfoXXOO struct {
	AoXXOO
	UpdateTime               uint32
	UpdateTime_u             uint8
	SkuAddTime               uint32
	SkuAddTime_u             uint8
	SkuLastUpTime            uint32
	SkuLastUpTime_u          uint8
	SkuLastDownTime          uint32
	SkuLastDownTime_u        uint8
	DataState                uint32
	DataState_u              uint8
	SkuId                    uint64
	SkuId_u                  uint8
	SkuLocalCode             string
	SkuLocalCode_u           uint8
	CooperatorId             uint64
	CooperatorId_u           uint8
	CooperatorSubaccountId   uint64
	CooperatorSubaccountId_u uint8
	Name                     string
	Name_u                   uint8
	PromotionText            string
	PromotionText_u          uint8
	SkuPrice                 uint32
	SkuPrice_u               uint8
	AreaId                   uint32
	AreaId_u                 uint8
	AreaPrice                uint32
	AreaPrice_u              uint8
	SkuList                  []AoDetailmainPoItemSkuInfoXXOO
	SkuList_u                uint8
	BuyMin                   uint32
	BuyMin_u                 uint8
	BuyMax                   uint32
	BuyMax_u                 uint8
	BuyTimes                 uint32
	BuyTimes_u               uint8
	SkuProperty              []uint8 //<std::bitset<128> >
	SkuProperty_u            uint8
	SkuPicList               []AoDetailmainPoItemPicInfoXXOO
	SkuPicList_u             uint8
	PicCdnUrl                string
	PicCdnUrl_u              uint8
	DetailInfo               []AoDetailmainPoItemDetailInfoXXOO
	DetailInfo_u             uint8
	DetailCdnUrl             string
	DetailCdnUrl_u           uint8
	ExpertRecommend          string
	ExpertRecommend_u        uint8
	C1ids                    uint32
	C1ids_u                  uint8
	C2ids                    uint32
	C2ids_u                  uint8
	C3ids                    uint32
	C3ids_u                  uint8
	C1name                   string
	C1name_u                 uint8
	C2name                   string
	C2name_u                 uint8
	C3name                   string
	C3name_u                 uint8
	CategoryId               uint32
	CategoryId_u             uint8
	SkuSaleAttr              string
	SkuSaleAttr_u            uint8
	SkuSaleAttrText          string
	SkuSaleAttrText_u        uint8
	SkuKeyAttr               string
	SkuKeyAttr_u             uint8
	SkuKeyAttrText           string
	SkuKeyAttrText_u         uint8
	CategoryAttr             string
	CategoryAttr_u           uint8
	CategoryAttrText         string
	CategoryAttrText_u       uint8
	BrandNo                  uint64
	BrandNo_u                uint8
	BrandName                string
	BrandName_u              uint8
	ProductModel             string
	ProductModel_u           uint8
	SkuState                 uint32
	SkuState_u               uint8
	AreaState                uint32
	AreaState_u              uint8
	ItemState                uint32
	ItemState_u              uint8
	SaleMode                 uint32
	SaleMode_u               uint8
	SkuType                  uint32
	SkuType_u                uint8
	SkuWeight                uint32
	SkuWeight_u              uint8
	SkuVolume                uint32
	SkuVolume_u              uint8
	SkuSizeX                 uint32
	SkuSizeX_u               uint8
	SkuSizeY                 uint32
	SkuSizeY_u               uint8
	SkuSizeZ                 uint32
	SkuSizeZ_u               uint8
	BarCode                  []string
	BarCode_u                uint8
	SkuKeywords              string
	SkuKeywords_u            uint8
	ExchangePoint            uint32
	ExchangePoint_u          uint8
	MainSkuId                uint64
	MainSkuId_u              uint8
	TaxRate                  uint32
	TaxRate_u                uint8
	Sort                     uint32
	Sort_u                   uint8
	ReserveStr               string
	ReserveStr_u             uint8
	CooperName               string
	CooperName_u             uint8
	CooperShortName          string
	CooperShortName_u        uint8
	CooperInfo               string
	CooperInfo_u             uint8
	CooperFreight            string
	CooperFreight_u          uint8
}

func NewAoDetailmainPoItemBasicInfoXXOO() *AoDetailmainPoItemBasicInfoXXOO {
	model := &AoDetailmainPoItemBasicInfoXXOO{}

	return model
}

func (ao *AoDetailmainPoItemBasicInfoXXOO) Serialize(bs *ByteStream) bool {

	bs.PushUint32(uint32(ao.getClassLen()))

	return ao.serialize_internal(bs)
}

func (ao *AoDetailmainPoItemBasicInfoXXOO) serialize_internal(bs *ByteStream) bool {
	bs.PushUint32(ao.Version)
	bs.PushUint8(ao.Version_u)
	bs.PushUint32(ao.UpdateTime)
	bs.PushUint8(ao.UpdateTime_u)
	bs.PushUint32(ao.SkuAddTime)
	bs.PushUint8(ao.SkuAddTime_u)
	bs.PushUint32(ao.SkuLastUpTime)
	bs.PushUint8(ao.SkuLastUpTime_u)
	bs.PushUint32(ao.SkuLastDownTime)
	bs.PushUint8(ao.SkuLastDownTime_u)
	bs.PushUint32(ao.DataState)
	bs.PushUint8(ao.DataState_u)
	bs.PushUint64(ao.SkuId)
	bs.PushUint8(ao.SkuId_u)
	bs.PushString(ao.SkuLocalCode)
	bs.PushUint8(ao.SkuLocalCode_u)
	bs.PushUint64(ao.CooperatorId)
	bs.PushUint8(ao.CooperatorId_u)
	bs.PushUint64(ao.CooperatorSubaccountId)
	bs.PushUint8(ao.CooperatorSubaccountId_u)
	bs.PushString(ao.Name)
	bs.PushUint8(ao.Name_u)
	bs.PushString(ao.PromotionText)
	bs.PushUint8(ao.PromotionText_u)
	bs.PushUint32(ao.SkuPrice)
	bs.PushUint8(ao.SkuPrice_u)
	bs.PushUint32(ao.AreaId)
	bs.PushUint8(ao.AreaId_u)
	bs.PushUint32(ao.AreaPrice)
	bs.PushUint8(ao.AreaPrice_u)

	skuListLen := uint32(len(ao.SkuList))
	bs.PushUint32(skuListLen)
	for _, v := range ao.SkuList {
		v.Serialize(bs)
	}
	bs.PushUint8(ao.SkuList_u)

	bs.PushUint32(ao.BuyMin)
	bs.PushUint8(ao.BuyMin_u)
	bs.PushUint32(ao.BuyMax)
	bs.PushUint8(ao.BuyMax_u)
	bs.PushUint32(ao.BuyTimes)
	bs.PushUint8(ao.BuyTimes_u)
	bs.PushBitSet(ao.SkuProperty)
	bs.PushUint8(ao.SkuProperty_u)

	skuPicListLen := uint32(len(ao.SkuPicList))
	bs.PushUint32(skuPicListLen)
	for _, v := range ao.SkuPicList {
		v.Serialize(bs)
	}
	bs.PushUint8(ao.SkuPicList_u)

	bs.PushString(ao.PicCdnUrl)
	bs.PushUint8(ao.PicCdnUrl_u)

	detailInfoLen := uint32(len(ao.DetailInfo))
	bs.PushUint32(detailInfoLen)
	for _, v := range ao.DetailInfo {
		v.Serialize(bs)
	}
	bs.PushUint8(ao.DetailInfo_u)

	bs.PushString(ao.DetailCdnUrl)
	bs.PushUint8(ao.DetailCdnUrl_u)
	bs.PushString(ao.ExpertRecommend)
	bs.PushUint8(ao.ExpertRecommend_u)
	bs.PushUint32(ao.C1ids)
	bs.PushUint8(ao.C1ids_u)
	bs.PushUint32(ao.C2ids)
	bs.PushUint8(ao.C2ids_u)
	bs.PushUint32(ao.C3ids)
	bs.PushUint8(ao.C3ids_u)
	bs.PushString(ao.C1name)
	bs.PushUint8(ao.C1name_u)
	bs.PushString(ao.C2name)
	bs.PushUint8(ao.C2name_u)
	bs.PushString(ao.C3name)
	bs.PushUint8(ao.C3name_u)
	bs.PushUint32(ao.CategoryId)
	bs.PushUint8(ao.CategoryId_u)
	bs.PushString(ao.SkuSaleAttr)
	bs.PushUint8(ao.SkuSaleAttr_u)
	bs.PushString(ao.SkuSaleAttrText)
	bs.PushUint8(ao.SkuSaleAttrText_u)
	bs.PushString(ao.SkuKeyAttr)
	bs.PushUint8(ao.SkuKeyAttr_u)
	bs.PushString(ao.SkuKeyAttrText)
	bs.PushUint8(ao.SkuKeyAttrText_u)
	bs.PushString(ao.CategoryAttr)
	bs.PushUint8(ao.CategoryAttr_u)
	bs.PushString(ao.CategoryAttrText)
	bs.PushUint8(ao.CategoryAttrText_u)
	bs.PushUint64(ao.BrandNo)
	bs.PushUint8(ao.BrandNo_u)
	bs.PushString(ao.BrandName)
	bs.PushUint8(ao.BrandName_u)
	bs.PushString(ao.ProductModel)
	bs.PushUint8(ao.ProductModel_u)
	bs.PushUint32(ao.SkuState)
	bs.PushUint8(ao.SkuState_u)
	bs.PushUint32(ao.AreaState)
	bs.PushUint8(ao.AreaState_u)
	bs.PushUint32(ao.ItemState)
	bs.PushUint8(ao.ItemState_u)
	bs.PushUint32(ao.SaleMode)
	bs.PushUint8(ao.SaleMode_u)
	bs.PushUint32(ao.SkuType)
	bs.PushUint8(ao.SkuType_u)
	bs.PushUint32(ao.SkuWeight)
	bs.PushUint8(ao.SkuWeight_u)
	bs.PushUint32(ao.SkuVolume)
	bs.PushUint8(ao.SkuVolume_u)
	bs.PushUint32(ao.SkuSizeX)
	bs.PushUint8(ao.SkuSizeX_u)
	bs.PushUint32(ao.SkuSizeY)
	bs.PushUint8(ao.SkuSizeY_u)
	bs.PushUint32(ao.SkuSizeZ)
	bs.PushUint8(ao.SkuSizeZ_u)
	bs.PushStringSet(ao.BarCode)
	bs.PushUint8(ao.BarCode_u)
	bs.PushString(ao.SkuKeywords)
	bs.PushUint8(ao.SkuKeywords_u)
	bs.PushUint32(ao.ExchangePoint)
	bs.PushUint8(ao.ExchangePoint_u)
	bs.PushUint64(ao.MainSkuId)
	bs.PushUint8(ao.MainSkuId_u)
	bs.PushUint32(ao.TaxRate)
	bs.PushUint8(ao.TaxRate_u)
	bs.PushUint32(ao.Sort)
	bs.PushUint8(ao.Sort_u)
	bs.PushString(ao.ReserveStr)
	bs.PushUint8(ao.ReserveStr_u)

	if ao.Version >= uint32(1) {
		bs.PushString(ao.CooperName)
		bs.PushUint8(ao.CooperName_u)
		bs.PushString(ao.CooperShortName)
		bs.PushUint8(ao.CooperShortName_u)
		bs.PushString(ao.CooperInfo)
		bs.PushUint8(ao.CooperInfo_u)
		bs.PushString(ao.CooperFreight)
		bs.PushUint8(ao.CooperFreight_u)
	}

	return bs.IsGood()
}

func (ao *AoDetailmainPoItemBasicInfoXXOO) getClassLen() int {

	bsLen := NewByteStream()
	bsLen.SetRealWrite(false)
	ao.serialize_internal(bsLen)
	length := bsLen.GetWrittenLength()

	return length
}

func (ao *AoDetailmainPoItemBasicInfoXXOO) UnSerialize(bs *ByteStream) error {
	classLen, err := bs.PopUint32()

	if err != nil {
		return err
	}
	startPop := bs.GetReadLength()

	ao.Version, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.Version_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.UpdateTime, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.UpdateTime_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuAddTime, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.SkuAddTime_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuLastUpTime, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.SkuLastUpTime_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuLastDownTime, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.SkuLastDownTime_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.DataState, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.DataState_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuId, err = bs.PopUint64()
	if err != nil {
		return err
	}
	ao.SkuId_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuLocalCode, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.SkuLocalCode_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.CooperatorId, err = bs.PopUint64()
	if err != nil {
		return err
	}
	ao.CooperatorId_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.CooperatorSubaccountId, err = bs.PopUint64()
	if err != nil {
		return err
	}
	ao.CooperatorSubaccountId_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.Name, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.Name_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.PromotionText, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.PromotionText_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuPrice, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.SkuPrice_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.AreaId, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.AreaId_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.AreaPrice, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.AreaPrice_u, err = bs.PopUint8()
	if err != nil {
		return err
	}

	var skuListLen uint32
	skuListLen, err = bs.PopUint32()
	if err != nil {
		return err
	}
	for i := uint32(0); i < skuListLen; i++ {
		info := NewAoDetailmainPoItemSkuInfoXXOO()
		err = info.UnSerialize(bs)
		if err != nil {
			return err
		}
		ao.SkuList = append(ao.SkuList, *info)
	}
	ao.SkuList_u, err = bs.PopUint8()
	if err != nil {
		return err
	}

	ao.BuyMin, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.BuyMin_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.BuyMax, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.BuyMax_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.BuyTimes, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.BuyTimes_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuProperty, err = bs.PopBitSet()
	if err != nil {
		return err
	}
	ao.SkuProperty_u, err = bs.PopUint8()
	if err != nil {
		return err
	}

	var skuPicListLen uint32
	skuPicListLen, err = bs.PopUint32()
	if err != nil {
		return err
	}
	for i := uint32(0); i < skuPicListLen; i++ {
		info := NewAoDetailmainPoItemPicInfoXXOO()
		err = info.UnSerialize(bs)
		if err != nil {
			return err
		}
		ao.SkuPicList = append(ao.SkuPicList, *info)
	}
	ao.SkuPicList_u, err = bs.PopUint8()
	if err != nil {
		return err
	}

	ao.PicCdnUrl, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.PicCdnUrl_u, err = bs.PopUint8()
	if err != nil {
		return err
	}

	var detailInfoLen uint32
	detailInfoLen, err = bs.PopUint32()
	if err != nil {
		return err
	}
	for i := uint32(0); i < detailInfoLen; i++ {
		info := NewAoDetailmainPoItemDetailInfoXXOO()
		err = info.UnSerialize(bs)
		if err != nil {
			return err
		}
		ao.DetailInfo = append(ao.DetailInfo, *info)
	}
	ao.DetailInfo_u, err = bs.PopUint8()
	if err != nil {
		return err
	}

	ao.DetailCdnUrl, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.DetailCdnUrl_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.ExpertRecommend, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.ExpertRecommend_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.C1ids, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.C1ids_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.C2ids, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.C2ids_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.C3ids, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.C3ids_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.C1name, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.C1name_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.C2name, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.C2name_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.C3name, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.C3name_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.CategoryId, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.CategoryId_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuSaleAttr, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.SkuSaleAttr_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuSaleAttrText, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.SkuSaleAttrText_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuKeyAttr, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.SkuKeyAttr_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuKeyAttrText, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.SkuKeyAttrText_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.CategoryAttr, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.CategoryAttr_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.CategoryAttrText, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.CategoryAttrText_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.BrandNo, err = bs.PopUint64()
	if err != nil {
		return err
	}
	ao.BrandNo_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.BrandName, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.BrandName_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.ProductModel, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.ProductModel_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuState, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.SkuState_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.AreaState, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.AreaState_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.ItemState, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.ItemState_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SaleMode, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.SaleMode_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuType, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.SkuType_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuWeight, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.SkuWeight_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuVolume, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.SkuVolume_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuSizeX, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.SkuSizeX_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuSizeY, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.SkuSizeY_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuSizeZ, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.SkuSizeZ_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.BarCode, err = bs.PopStringSet()
	if err != nil {
		return err
	}
	ao.BarCode_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuKeywords, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.SkuKeywords_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.ExchangePoint, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.ExchangePoint_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.MainSkuId, err = bs.PopUint64()
	if err != nil {
		return err
	}
	ao.MainSkuId_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.TaxRate, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.TaxRate_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.Sort, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.Sort_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.ReserveStr, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.ReserveStr_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	if ao.Version >= uint32(1) {
		ao.CooperName, err = bs.PopString()
		if err != nil {
			return err
		}
		ao.CooperName_u, err = bs.PopUint8()
		if err != nil {
			return err
		}
		ao.CooperShortName, err = bs.PopString()
		if err != nil {
			return err
		}
		ao.CooperShortName_u, err = bs.PopUint8()
		if err != nil {
			return err
		}
		ao.CooperInfo, err = bs.PopString()
		if err != nil {
			return err
		}
		ao.CooperInfo_u, err = bs.PopUint8()
		if err != nil {
			return err
		}
		ao.CooperFreight, err = bs.PopString()
		if err != nil {
			return err
		}
		ao.CooperFreight_u, err = bs.PopUint8()
		if err != nil {
			return err
		}
	}

	ao.Compat(bs, classLen, startPop)
	return nil
}

/* AoDetailmainPoItemBasicInfoXXOO END */

/* AoDetailmainPoItemBasicInfoXXOO START */

type AoDetailmainPoItemSkuInfoXXOO struct {
	AoXXOO
	SkuId              uint64
	SkuId_u            uint8
	SkuLocalCode       string
	SkuLocalCode_u     uint8
	CategoryId         uint32
	CategoryId_u       uint8
	SkuSaleAttr        string
	SkuSaleAttr_u      uint8
	SkuSaleAttrText    string
	SkuSaleAttrText_u  uint8
	SkuKeyAttr         string
	SkuKeyAttr_u       uint8
	SkuKeyAttrText     string
	SkuKeyAttrText_u   uint8
	CategoryAttr       string
	CategoryAttr_u     uint8
	CategoryAttrText   string
	CategoryAttrText_u uint8
	MinPic             string
	MinPic_u           uint8
	SkuItemProperty    []uint8 //<std::bitset<128> >
	SkuItemProperty_u  uint8
	Name               string
	Name_u             uint8
	ProductModel       string
	ProductModel_u     uint8
	SkuPrice           uint32
	SkuPrice_u         uint8
	TaxRate            uint32
	TaxRate_u          uint8
	BarCode            string
	BarCode_u          uint8
	SaleMode           uint32
	SaleMode_u         uint8
	SkuState           uint32
	SkuState_u         uint8
	MainSkuid          uint64
	MainSkuid_u        uint8
	Sort               uint32
	Sort_u             uint8
	CooperatorId       uint32
	CooperatorId_u     uint8
	ReserveStr         string
	ReserveStr_u       uint8
}

func NewAoDetailmainPoItemSkuInfoXXOO() *AoDetailmainPoItemSkuInfoXXOO {
	model := &AoDetailmainPoItemSkuInfoXXOO{}

	return model
}

func (ao *AoDetailmainPoItemSkuInfoXXOO) Serialize(bs *ByteStream) bool {

	bs.PushUint32(uint32(ao.getClassLen()))

	return ao.serialize_internal(bs)
}

func (ao *AoDetailmainPoItemSkuInfoXXOO) serialize_internal(bs *ByteStream) bool {
	bs.PushUint32(ao.Version)
	bs.PushUint8(ao.Version_u)
	bs.PushUint64(ao.SkuId)
	bs.PushUint8(ao.SkuId_u)
	bs.PushString(ao.SkuLocalCode)
	bs.PushUint8(ao.SkuLocalCode_u)
	bs.PushUint32(ao.CategoryId)
	bs.PushUint8(ao.CategoryId_u)
	bs.PushString(ao.SkuSaleAttr)
	bs.PushUint8(ao.SkuSaleAttr_u)
	bs.PushString(ao.SkuSaleAttrText)
	bs.PushUint8(ao.SkuSaleAttrText_u)
	bs.PushString(ao.SkuKeyAttr)
	bs.PushUint8(ao.SkuKeyAttr_u)
	bs.PushString(ao.SkuKeyAttrText)
	bs.PushUint8(ao.SkuKeyAttrText_u)
	bs.PushString(ao.CategoryAttr)
	bs.PushUint8(ao.CategoryAttr_u)
	bs.PushString(ao.CategoryAttrText)
	bs.PushUint8(ao.CategoryAttrText_u)
	bs.PushString(ao.MinPic)
	bs.PushUint8(ao.MinPic_u)
	bs.PushBitSet(ao.SkuItemProperty)
	bs.PushUint8(ao.SkuItemProperty_u)
	bs.PushString(ao.Name)
	bs.PushUint8(ao.Name_u)
	bs.PushString(ao.ProductModel)
	bs.PushUint8(ao.ProductModel_u)
	bs.PushUint32(ao.SkuPrice)
	bs.PushUint8(ao.SkuPrice_u)
	bs.PushUint32(ao.TaxRate)
	bs.PushUint8(ao.TaxRate_u)
	bs.PushString(ao.BarCode)
	bs.PushUint8(ao.BarCode_u)
	bs.PushUint32(ao.SaleMode)
	bs.PushUint8(ao.SaleMode_u)
	bs.PushUint32(ao.SkuState)
	bs.PushUint8(ao.SkuState_u)
	bs.PushUint64(ao.MainSkuid)
	bs.PushUint8(ao.MainSkuid_u)
	bs.PushUint32(ao.Sort)
	bs.PushUint8(ao.Sort_u)
	bs.PushUint32(ao.CooperatorId)
	bs.PushUint8(ao.CooperatorId_u)
	bs.PushString(ao.ReserveStr)
	bs.PushUint8(ao.ReserveStr_u)

	return bs.IsGood()
}

func (ao *AoDetailmainPoItemSkuInfoXXOO) getClassLen() int {

	bsLen := NewByteStream()
	bsLen.SetRealWrite(false)
	ao.serialize_internal(bsLen)
	length := bsLen.GetWrittenLength()

	return length
}

func (ao *AoDetailmainPoItemSkuInfoXXOO) UnSerialize(bs *ByteStream) error {
	classLen, err := bs.PopUint32()

	if err != nil {
		return err
	}
	startPop := bs.GetReadLength()

	ao.Version, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.Version_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuId, err = bs.PopUint64()
	if err != nil {
		return err
	}
	ao.SkuId_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuLocalCode, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.SkuLocalCode_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.CategoryId, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.CategoryId_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuSaleAttr, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.SkuSaleAttr_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuSaleAttrText, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.SkuSaleAttrText_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuKeyAttr, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.SkuKeyAttr_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuKeyAttrText, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.SkuKeyAttrText_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.CategoryAttr, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.CategoryAttr_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.CategoryAttrText, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.CategoryAttrText_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.MinPic, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.MinPic_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuItemProperty, err = bs.PopBitSet()
	if err != nil {
		return err
	}
	ao.SkuItemProperty_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.Name, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.Name_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.ProductModel, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.ProductModel_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuPrice, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.SkuPrice_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.TaxRate, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.TaxRate_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.BarCode, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.BarCode_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SaleMode, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.SaleMode_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuState, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.SkuState_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.MainSkuid, err = bs.PopUint64()
	if err != nil {
		return err
	}
	ao.MainSkuid_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.Sort, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.Sort_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.CooperatorId, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.CooperatorId_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.ReserveStr, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.ReserveStr_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.Compat(bs, classLen, startPop)
	return nil
}

/* AoDetailmainPoItemBasicInfoXXOO END */

/* AoDetailmainPoItemPicInfoXXOO START */

type AoDetailmainPoItemPicInfoXXOO struct {
	AoXXOO
	SkuId         uint64
	SkuId_u       uint8
	Channel       uint32
	Channel_u     uint8
	PicUrl        string
	PicUrl_u      uint8
	PicSize       string
	PicSize_u     uint8
	PicIndex      uint8
	PicIndex_u    uint8
	PicType       uint8
	PicType_u     uint8
	PicProperty   []uint8
	PicProperty_u uint8
	PicDesc       string
	PicDesc_u     uint8
	ReserveStr    string
	ReserveStr_u  uint8
}

func NewAoDetailmainPoItemPicInfoXXOO() *AoDetailmainPoItemPicInfoXXOO {
	model := &AoDetailmainPoItemPicInfoXXOO{}

	return model
}

func (ao *AoDetailmainPoItemPicInfoXXOO) Serialize(bs *ByteStream) bool {

	bs.PushUint32(uint32(ao.getClassLen()))

	return ao.serialize_internal(bs)
}

func (ao *AoDetailmainPoItemPicInfoXXOO) serialize_internal(bs *ByteStream) bool {
	bs.PushUint32(ao.Version)
	bs.PushUint8(ao.Version_u)
	bs.PushUint64(ao.SkuId)
	bs.PushUint8(ao.SkuId_u)
	bs.PushUint32(ao.Channel)
	bs.PushUint8(ao.Channel_u)
	bs.PushString(ao.PicUrl)
	bs.PushUint8(ao.PicUrl_u)
	bs.PushString(ao.PicSize)
	bs.PushUint8(ao.PicSize_u)
	bs.PushUint8(ao.PicIndex)
	bs.PushUint8(ao.PicIndex_u)
	bs.PushUint8(ao.PicType)
	bs.PushUint8(ao.PicType_u)
	bs.PushBitSet(ao.PicProperty)
	bs.PushUint8(ao.PicProperty_u)
	bs.PushString(ao.PicDesc)
	bs.PushUint8(ao.PicDesc_u)
	bs.PushString(ao.ReserveStr)
	bs.PushUint8(ao.ReserveStr_u)

	return bs.IsGood()
}

func (ao *AoDetailmainPoItemPicInfoXXOO) getClassLen() int {

	bsLen := NewByteStream()
	bsLen.SetRealWrite(false)
	ao.serialize_internal(bsLen)
	length := bsLen.GetWrittenLength()

	return length
}

func (ao *AoDetailmainPoItemPicInfoXXOO) UnSerialize(bs *ByteStream) error {
	classLen, err := bs.PopUint32()

	if err != nil {
		return err
	}
	startPop := bs.GetReadLength()

	ao.Version, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.Version_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuId, err = bs.PopUint64()
	if err != nil {
		return err
	}
	ao.SkuId_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.Channel, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.Channel_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.PicUrl, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.PicUrl_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.PicSize, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.PicSize_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.PicIndex, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.PicIndex_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.PicType, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.PicType_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.PicProperty, err = bs.PopBitSet()
	if err != nil {
		return err
	}
	ao.PicProperty_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.PicDesc, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.PicDesc_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.ReserveStr, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.ReserveStr_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.Compat(bs, classLen, startPop)
	return nil
}

/* AoDetailmainPoItemPicInfoXXOO END */

/* AoDetailmainPoItemDetailInfoXXOO START */

type AoDetailmainPoItemDetailInfoXXOO struct {
	AoXXOO
	SkuId        uint64
	SkuId_u      uint8
	DetailType   uint8
	DetailType_u uint8
	Channel      uint32
	Channel_u    uint8
	DetailUrl    string
	DetailUrl_u  uint8
	DetailInfo   string
	DetailInfo_u uint8
	ReserveStr   string
	ReserveStr_u uint8
}

func NewAoDetailmainPoItemDetailInfoXXOO() *AoDetailmainPoItemDetailInfoXXOO {
	model := &AoDetailmainPoItemDetailInfoXXOO{}

	return model
}

func (ao *AoDetailmainPoItemDetailInfoXXOO) Serialize(bs *ByteStream) bool {

	bs.PushUint32(uint32(ao.getClassLen()))

	return ao.serialize_internal(bs)
}

func (ao *AoDetailmainPoItemDetailInfoXXOO) serialize_internal(bs *ByteStream) bool {
	bs.PushUint32(ao.Version)
	bs.PushUint8(ao.Version_u)
	bs.PushUint64(ao.SkuId)
	bs.PushUint8(ao.SkuId_u)
	bs.PushUint8(ao.DetailType)
	bs.PushUint8(ao.DetailType_u)
	bs.PushUint32(ao.Channel)
	bs.PushUint8(ao.Channel_u)
	bs.PushString(ao.DetailUrl)
	bs.PushUint8(ao.DetailUrl_u)
	bs.PushString(ao.DetailInfo)
	bs.PushUint8(ao.DetailInfo_u)
	bs.PushString(ao.ReserveStr)
	bs.PushUint8(ao.ReserveStr_u)

	return bs.IsGood()
}

func (ao *AoDetailmainPoItemDetailInfoXXOO) getClassLen() int {

	bsLen := NewByteStream()
	bsLen.SetRealWrite(false)
	ao.serialize_internal(bsLen)
	length := bsLen.GetWrittenLength()

	return length
}

func (ao *AoDetailmainPoItemDetailInfoXXOO) UnSerialize(bs *ByteStream) error {
	classLen, err := bs.PopUint32()

	if err != nil {
		return err
	}
	startPop := bs.GetReadLength()

	ao.Version, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.Version_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuId, err = bs.PopUint64()
	if err != nil {
		return err
	}
	ao.SkuId_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.DetailType, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.DetailType_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.Channel, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.Channel_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.DetailUrl, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.DetailUrl_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.DetailInfo, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.DetailInfo_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.ReserveStr, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.ReserveStr_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.Compat(bs, classLen, startPop)
	return nil
}

/* AoDetailmainPoItemDetailInfoXXOO END */

/* AoDetailmainPoItemNavInfoXXOO START */

type AoDetailmainPoItemNavInfoXXOO struct {
	AoXXOO
	UpdateTime    uint32
	UpdateTime_u  uint8
	DataState     uint32
	DataState_u   uint8
	NavList       []AoDetailmainPoItemNavNodeXXOO
	NavList_u     uint8
	DetailList    []AoDetailmainPoItemDetailNodeXXOO
	DetailList_u  uint8
	NavProperty   []uint8
	NavProperty_u uint8
	NewNav        string
	NewNav_u      uint8
	ReserveStr    string
	ReserveStr_u  uint8
}

func NewAoDetailmainPoItemNavInfoXXOO() *AoDetailmainPoItemNavInfoXXOO {
	model := &AoDetailmainPoItemNavInfoXXOO{}

	return model
}

func (ao *AoDetailmainPoItemNavInfoXXOO) Serialize(bs *ByteStream) bool {

	bs.PushUint32(uint32(ao.getClassLen()))

	return ao.serialize_internal(bs)
}

func (ao *AoDetailmainPoItemNavInfoXXOO) serialize_internal(bs *ByteStream) bool {
	bs.PushUint32(ao.Version)
	bs.PushUint8(ao.Version_u)
	bs.PushUint32(ao.UpdateTime)
	bs.PushUint8(ao.UpdateTime_u)
	bs.PushUint32(ao.DataState)
	bs.PushUint8(ao.DataState_u)

	navListLen := uint32(len(ao.NavList))
	bs.PushUint32(navListLen)
	for _, v := range ao.NavList {
		v.Serialize(bs)
	}
	bs.PushUint8(ao.NavList_u)

	detailListLen := uint32(len(ao.DetailList))
	bs.PushUint32(detailListLen)
	for _, v := range ao.DetailList {
		v.Serialize(bs)
	}
	bs.PushUint8(ao.DetailList_u)

	bs.PushBitSet(ao.NavProperty)
	bs.PushUint8(ao.NavProperty_u)
	bs.PushString(ao.NewNav)
	bs.PushUint8(ao.NewNav_u)
	bs.PushString(ao.ReserveStr)
	bs.PushUint8(ao.ReserveStr_u)

	return bs.IsGood()
}

func (ao *AoDetailmainPoItemNavInfoXXOO) getClassLen() int {

	bsLen := NewByteStream()
	bsLen.SetRealWrite(false)
	ao.serialize_internal(bsLen)
	length := bsLen.GetWrittenLength()

	return length
}

func (ao *AoDetailmainPoItemNavInfoXXOO) UnSerialize(bs *ByteStream) error {
	classLen, err := bs.PopUint32()

	if err != nil {
		return err
	}
	startPop := bs.GetReadLength()

	ao.Version, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.Version_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.UpdateTime, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.UpdateTime_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.DataState, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.DataState_u, err = bs.PopUint8()
	if err != nil {
		return err
	}

	var navListLen uint32
	navListLen, err = bs.PopUint32()
	if err != nil {
		return err
	}
	for i := uint32(0); i < navListLen; i++ {
		info := NewAoDetailmainPoItemNavNodeXXOO()
		err = info.UnSerialize(bs)
		if err != nil {
			return err
		}
		ao.NavList = append(ao.NavList, *info)
	}
	ao.NavList_u, err = bs.PopUint8()
	if err != nil {
		return err
	}

	var detailListLen uint32
	detailListLen, err = bs.PopUint32()
	if err != nil {
		return err
	}
	for i := uint32(0); i < detailListLen; i++ {
		info := NewAoDetailmainPoItemDetailNodeXXOO()
		err = info.UnSerialize(bs)
		if err != nil {
			return err
		}
		ao.DetailList = append(ao.DetailList, *info)
	}
	ao.DetailList_u, err = bs.PopUint8()
	if err != nil {
		return err
	}

	ao.NavProperty, err = bs.PopBitSet()
	if err != nil {
		return err
	}
	ao.NavProperty_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.NewNav, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.NewNav_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.ReserveStr, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.ReserveStr_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.Compat(bs, classLen, startPop)
	return nil
}

/* AoDetailmainPoItemNavInfoXXOO END */

/* AoDetailmainPoItemPmInfoXXOO START */

type AoDetailmainPoItemPmInfoXXOO struct {
	AoXXOO
	SkuId                uint64
	SkuId_u              uint8
	ItemUnitPrice        uint32
	ItemUnitPrice_u      uint8
	ItemPmPrice          uint32
	ItemPmPrice_u        uint8
	ItemPriceName        string
	ItemPriceName_u      uint8
	ItemPriceEndTime     uint32
	ItemPriceEndTime_u   uint8
	ItemPriceLimitDesc   string
	ItemPriceLimitDesc_u uint8
	PromotionInfoList    []AoDetailmainPoPromotionInfoXXOO
	PromotionInfoList_u  uint8
	ReserveStr           string
	ReserveStr_u         uint8
	ItemPriceStartTime   uint32
	ItemPriceStartTime_u uint8
	PlusItemList         []AoDetailmainPoGiftItemInfoXXOO
	PlusItemList_u       uint8
	ItemOpInfo           string
	ItemOpInfo_u         uint8
}

func NewAoDetailmainPoItemPmInfoXXOO() *AoDetailmainPoItemPmInfoXXOO {
	model := &AoDetailmainPoItemPmInfoXXOO{}

	return model
}

func (ao *AoDetailmainPoItemPmInfoXXOO) Serialize(bs *ByteStream) bool {

	bs.PushUint32(uint32(ao.getClassLen()))

	return ao.serialize_internal(bs)
}

func (ao *AoDetailmainPoItemPmInfoXXOO) serialize_internal(bs *ByteStream) bool {
	bs.PushUint32(ao.Version)
	bs.PushUint8(ao.Version_u)
	bs.PushUint64(ao.SkuId)
	bs.PushUint8(ao.SkuId_u)
	bs.PushUint32(ao.ItemUnitPrice)
	bs.PushUint8(ao.ItemUnitPrice_u)
	bs.PushUint32(ao.ItemPmPrice)
	bs.PushUint8(ao.ItemPmPrice_u)
	bs.PushString(ao.ItemPriceName)
	bs.PushUint8(ao.ItemPriceName_u)
	bs.PushUint32(ao.ItemPriceEndTime)
	bs.PushUint8(ao.ItemPriceEndTime_u)
	bs.PushString(ao.ItemPriceLimitDesc)
	bs.PushUint8(ao.ItemPriceLimitDesc_u)
	bs.PushUint32(ao.ItemUnitPrice)

	promotionInfoListLen := uint32(len(ao.PromotionInfoList))
	bs.PushUint32(promotionInfoListLen)
	for _, v := range ao.PromotionInfoList {
		v.Serialize(bs)
	}
	bs.PushUint8(ao.PromotionInfoList_u)

	bs.PushString(ao.ReserveStr)
	bs.PushUint8(ao.ReserveStr_u)

	if ao.Version >= uint32(1) {
		bs.PushUint32(ao.ItemPriceStartTime)
		bs.PushUint8(ao.ItemPriceStartTime_u)

		plusItemListLen := uint32(len(ao.PlusItemList))
		bs.PushUint32(plusItemListLen)
		for _, v := range ao.PlusItemList {
			v.Serialize(bs)
		}
		bs.PushUint8(ao.PlusItemList_u)
	}

	if ao.Version >= uint32(2) {
		bs.PushString(ao.ItemOpInfo)
		bs.PushUint8(ao.ItemOpInfo_u)
	}

	return bs.IsGood()
}

func (ao *AoDetailmainPoItemPmInfoXXOO) getClassLen() int {

	bsLen := NewByteStream()
	bsLen.SetRealWrite(false)
	ao.serialize_internal(bsLen)
	length := bsLen.GetWrittenLength()

	return length
}

func (ao *AoDetailmainPoItemPmInfoXXOO) UnSerialize(bs *ByteStream) error {
	classLen, err := bs.PopUint32()

	if err != nil {
		return err
	}
	startPop := bs.GetReadLength()

	ao.Version, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.Version_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuId, err = bs.PopUint64()
	if err != nil {
		return err
	}
	ao.SkuId_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.ItemUnitPrice, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.ItemUnitPrice_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.ItemPmPrice, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.ItemPmPrice_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.ItemPriceName, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.ItemPriceName_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.ItemPriceEndTime, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.ItemPriceEndTime_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.ItemPriceLimitDesc, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.ItemPriceLimitDesc_u, err = bs.PopUint8()
	if err != nil {
		return err
	}

	var promotionInfoListLen uint32
	promotionInfoListLen, err = bs.PopUint32()
	if err != nil {
		return err
	}
	for i := uint32(0); i < promotionInfoListLen; i++ {
		info := NewAoDetailmainPoPromotionInfoXXOO()
		err = info.UnSerialize(bs)
		if err != nil {
			return err
		}
		ao.PromotionInfoList = append(ao.PromotionInfoList, *info)
	}
	ao.PromotionInfoList_u, err = bs.PopUint8()
	if err != nil {
		return err
	}

	ao.ReserveStr, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.ReserveStr_u, err = bs.PopUint8()
	if err != nil {
		return err
	}

	if ao.Version >= uint32(1) {
		ao.ItemPriceStartTime, err = bs.PopUint32()
		if err != nil {
			return err
		}
		ao.ItemPriceStartTime_u, err = bs.PopUint8()
		if err != nil {
			return err
		}

		var plusItemListLen uint32
		plusItemListLen, err = bs.PopUint32()
		if err != nil {
			return err
		}
		for i := uint32(0); i < plusItemListLen; i++ {
			info := NewAoDetailmainPoGiftItemInfoXXOO()
			err = info.UnSerialize(bs)
			if err != nil {
				return err
			}
			ao.PlusItemList = append(ao.PlusItemList, *info)
		}
		ao.PlusItemList_u, err = bs.PopUint8()
		if err != nil {
			return err
		}
	}

	if ao.Version >= uint32(2) {
		ao.ItemOpInfo, err = bs.PopString()
		if err != nil {
			return err
		}
		ao.ItemOpInfo_u, err = bs.PopUint8()
		if err != nil {
			return err
		}
	}

	ao.Compat(bs, classLen, startPop)
	return nil
}

/* AoDetailmainPoItemPmInfoXXOO END */

/* AoDetailmainPoPromotionInfoXXOO START */

type AoDetailmainPoPromotionInfoXXOO struct {
	AoXXOO
	RuleId               uint64
	RuleId_u             uint8
	RuleType             uint32
	RuleType_u           uint8
	RuleTypeDesc         string
	RuleTypeDesc_u       uint8
	RuleName             string
	RuleName_u           uint8
	RuleDesc             string
	RuleDesc_u           uint8
	OpInfo               string
	OpInfo_u             uint8
	OpUrl                string
	OpUrl_u              uint8
	PromotionLimitDesc   string
	PromotionLimitDesc_u uint8
	PromotionEndTime     uint32
	PromotionEndTime_u   uint8
	PmSkuId              []uint64
	PmSkuId_u            uint8
	CouponBatchId        []string
	CouponBatchId_u      uint8
	PromotionStartTime   uint32
	PromotionStartTime_u uint8
	GiftList             []AoDetailmainPoGiftItemInfoXXOO
	GiftList_u           uint8
	ReserveStr           string
	ReserveStr_u         uint8
}

func NewAoDetailmainPoPromotionInfoXXOO() *AoDetailmainPoPromotionInfoXXOO {
	model := &AoDetailmainPoPromotionInfoXXOO{}
	model.Version = uint32(1)

	return model
}

func (ao *AoDetailmainPoPromotionInfoXXOO) Serialize(bs *ByteStream) bool {

	bs.PushUint32(uint32(ao.getClassLen()))

	return ao.serialize_internal(bs)
}

func (ao *AoDetailmainPoPromotionInfoXXOO) serialize_internal(bs *ByteStream) bool {
	bs.PushUint32(ao.Version)
	bs.PushUint8(ao.Version_u)
	bs.PushUint64(ao.RuleId)
	bs.PushUint8(ao.RuleId_u)
	bs.PushUint32(ao.RuleType)
	bs.PushUint8(ao.RuleType_u)
	bs.PushString(ao.RuleTypeDesc)
	bs.PushUint8(ao.RuleTypeDesc_u)
	bs.PushString(ao.RuleName)
	bs.PushUint8(ao.RuleName_u)
	bs.PushString(ao.RuleDesc)
	bs.PushUint8(ao.RuleDesc_u)
	bs.PushString(ao.OpInfo)
	bs.PushUint8(ao.OpInfo_u)
	bs.PushString(ao.OpUrl)
	bs.PushUint8(ao.OpUrl_u)
	bs.PushString(ao.PromotionLimitDesc)
	bs.PushUint8(ao.PromotionLimitDesc_u)
	bs.PushUint32(ao.PromotionEndTime)
	bs.PushUint8(ao.PromotionEndTime_u)
	bs.PushUint64Set(ao.PmSkuId)
	bs.PushUint8(ao.PmSkuId_u)
	bs.PushStringSet(ao.CouponBatchId)
	bs.PushUint8(ao.CouponBatchId_u)
	bs.PushString(ao.ReserveStr)
	bs.PushUint8(ao.ReserveStr_u)

	if ao.Version >= uint32(1) {
		bs.PushUint32(ao.PromotionStartTime)
		bs.PushUint8(ao.PromotionStartTime_u)

		giftListLen := uint32(len(ao.GiftList))
		bs.PushUint32(giftListLen)
		for _, v := range ao.GiftList {
			v.Serialize(bs)
		}
		bs.PushUint8(ao.GiftList_u)
	}

	return bs.IsGood()
}

func (ao *AoDetailmainPoPromotionInfoXXOO) getClassLen() int {

	bsLen := NewByteStream()
	bsLen.SetRealWrite(false)
	ao.serialize_internal(bsLen)
	length := bsLen.GetWrittenLength()

	return length
}

func (ao *AoDetailmainPoPromotionInfoXXOO) UnSerialize(bs *ByteStream) error {
	classLen, err := bs.PopUint32()

	if err != nil {
		return err
	}
	startPop := bs.GetReadLength()

	ao.Version, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.Version_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.RuleId, err = bs.PopUint64()
	if err != nil {
		return err
	}
	ao.RuleId_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.RuleType, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.RuleType_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.RuleTypeDesc, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.RuleTypeDesc_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.RuleName, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.RuleName_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.RuleDesc, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.RuleDesc_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.OpInfo, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.OpInfo_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.OpUrl, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.OpUrl_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.PromotionLimitDesc, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.PromotionLimitDesc_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.PromotionEndTime, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.PromotionEndTime_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.PmSkuId, err = bs.PopUint64Set()
	if err != nil {
		return err
	}
	ao.PmSkuId_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.CouponBatchId, err = bs.PopStringSet()
	if err != nil {
		return err
	}
	ao.CouponBatchId_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.ReserveStr, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.ReserveStr_u, err = bs.PopUint8()
	if err != nil {
		return err
	}

	if ao.Version >= uint32(1) {
		ao.PromotionStartTime, err = bs.PopUint32()
		if err != nil {
			return err
		}
		ao.PromotionStartTime_u, err = bs.PopUint8()
		if err != nil {
			return err
		}

		var giftListLen uint32
		giftListLen, err = bs.PopUint32()
		if err != nil {
			return err
		}
		for i := uint32(0); i < giftListLen; i++ {
			info := NewAoDetailmainPoGiftItemInfoXXOO()
			err = info.UnSerialize(bs)
			if err != nil {
				return err
			}
			ao.GiftList = append(ao.GiftList, *info)
		}
		ao.GiftList_u, err = bs.PopUint8()
		if err != nil {
			return err
		}
	}

	ao.Compat(bs, classLen, startPop)
	return nil
}

/* AoDetailmainPoPromotionInfoXXOO END */

/* AoDetailmainPoGiftItemInfoXXOO START */

type AoDetailmainPoGiftItemInfoXXOO struct {
	AoXXOO
	SkuId           uint64
	SkuId_u         uint8
	ItemName        string
	ItemName_u      uint8
	ItemImgUrl      string
	ItemImgUrl_u    uint8
	ItemNum         uint32
	ItemNum_u       uint8
	ItemUnitPrice   uint64
	ItemUnitPrice_u uint8
	ItemPmPrice     uint64
	ItemPmPrice_u   uint8
	ItemStock       uint32
	ItemStock_u     uint8
	ReserveStr      string
	ReserveStr_u    uint8
}

func NewAoDetailmainPoGiftItemInfoXXOO() *AoDetailmainPoGiftItemInfoXXOO {
	model := &AoDetailmainPoGiftItemInfoXXOO{}

	return model
}

func (ao *AoDetailmainPoGiftItemInfoXXOO) Serialize(bs *ByteStream) bool {

	bs.PushUint32(uint32(ao.getClassLen()))

	return ao.serialize_internal(bs)
}

func (ao *AoDetailmainPoGiftItemInfoXXOO) serialize_internal(bs *ByteStream) bool {
	bs.PushUint32(ao.Version)
	bs.PushUint8(ao.Version_u)
	bs.PushUint64(ao.SkuId)
	bs.PushUint8(ao.SkuId_u)
	bs.PushString(ao.ItemName)
	bs.PushUint8(ao.ItemName_u)
	bs.PushString(ao.ItemImgUrl)
	bs.PushUint8(ao.ItemImgUrl_u)
	bs.PushUint32(ao.ItemNum)
	bs.PushUint8(ao.ItemNum_u)
	bs.PushUint64(ao.ItemUnitPrice)
	bs.PushUint8(ao.ItemUnitPrice_u)
	bs.PushUint64(ao.ItemPmPrice)
	bs.PushUint8(ao.ItemPmPrice_u)
	bs.PushUint32(ao.ItemStock)
	bs.PushUint8(ao.ItemStock_u)
	bs.PushString(ao.ReserveStr)
	bs.PushUint8(ao.ReserveStr_u)

	return bs.IsGood()
}

func (ao *AoDetailmainPoGiftItemInfoXXOO) getClassLen() int {

	bsLen := NewByteStream()
	bsLen.SetRealWrite(false)
	ao.serialize_internal(bsLen)
	length := bsLen.GetWrittenLength()

	return length
}

func (ao *AoDetailmainPoGiftItemInfoXXOO) UnSerialize(bs *ByteStream) error {
	classLen, err := bs.PopUint32()

	if err != nil {
		return err
	}
	startPop := bs.GetReadLength()

	ao.Version, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.Version_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuId, err = bs.PopUint64()
	if err != nil {
		return err
	}
	ao.SkuId_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.ItemName, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.ItemName_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.ItemImgUrl, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.ItemImgUrl_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.ItemNum, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.ItemNum_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.ItemUnitPrice, err = bs.PopUint64()
	if err != nil {
		return err
	}
	ao.ItemUnitPrice_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.ItemPmPrice, err = bs.PopUint64()
	if err != nil {
		return err
	}
	ao.ItemPmPrice_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.ItemStock, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.ItemStock_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.ReserveStr, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.ReserveStr_u, err = bs.PopUint8()
	if err != nil {
		return err
	}

	ao.Compat(bs, classLen, startPop)
	return nil
}

/* AoDetailmainPoGiftItemInfoXXOO END */

/* AoDetailmainPoItemNavNodeXXOO START */

type AoDetailmainPoItemNavNodeXXOO struct {
	AoXXOO
	NavId        uint32
	NavId_u      uint8
	Name         string
	Name_u       uint8
	ReserveStr   string
	ReserveStr_u uint8
}

func NewAoDetailmainPoItemNavNodeXXOO() *AoDetailmainPoItemNavNodeXXOO {
	model := &AoDetailmainPoItemNavNodeXXOO{}
	model.Version = uint32(1)

	return model
}

func (ao *AoDetailmainPoItemNavNodeXXOO) Serialize(bs *ByteStream) bool {

	bs.PushUint32(uint32(ao.getClassLen()))

	return ao.serialize_internal(bs)
}

func (ao *AoDetailmainPoItemNavNodeXXOO) serialize_internal(bs *ByteStream) bool {
	bs.PushUint32(ao.Version)
	bs.PushUint8(ao.Version_u)
	bs.PushUint32(ao.NavId)
	bs.PushUint8(ao.NavId_u)
	bs.PushString(ao.Name)
	bs.PushUint8(ao.Name_u)
	bs.PushString(ao.ReserveStr)
	bs.PushUint8(ao.ReserveStr_u)

	return bs.IsGood()
}

func (ao *AoDetailmainPoItemNavNodeXXOO) getClassLen() int {

	bsLen := NewByteStream()
	bsLen.SetRealWrite(false)
	ao.serialize_internal(bsLen)
	length := bsLen.GetWrittenLength()

	return length
}

func (ao *AoDetailmainPoItemNavNodeXXOO) UnSerialize(bs *ByteStream) error {
	classLen, err := bs.PopUint32()

	if err != nil {
		return err
	}
	startPop := bs.GetReadLength()

	ao.Version, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.Version_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.NavId, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.NavId_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.Name, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.Name_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.ReserveStr, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.ReserveStr_u, err = bs.PopUint8()
	if err != nil {
		return err
	}

	ao.Compat(bs, classLen, startPop)
	return nil
}

/* AoDetailmainPoItemNavNodeXXOO END */

/* AoDetailmainPoItemDetailNodeXXOO START */

type AoDetailmainPoItemDetailNodeXXOO struct {
	AoXXOO
	AttrId       uint32
	AttrId_u     uint8
	AttrName     string
	AttrName_u   uint8
	StrOption    string
	StrOption_u  uint8
	GroupId      uint32
	GroupId_u    uint8
	GroupName    string
	GroupName_u  uint8
	Order        uint32
	Order_u      uint8
	GroupOrder   uint32
	GroupOrder_u uint8
	ReserveStr   string
	ReserveStr_u uint8
}

func NewAoDetailmainPoItemDetailNodeXXOO() *AoDetailmainPoItemDetailNodeXXOO {
	model := &AoDetailmainPoItemDetailNodeXXOO{}
	model.Version = uint32(1)

	return model
}

func (ao *AoDetailmainPoItemDetailNodeXXOO) Serialize(bs *ByteStream) bool {

	bs.PushUint32(uint32(ao.getClassLen()))

	return ao.serialize_internal(bs)
}

func (ao *AoDetailmainPoItemDetailNodeXXOO) serialize_internal(bs *ByteStream) bool {
	bs.PushUint32(ao.Version)
	bs.PushUint8(ao.Version_u)
	bs.PushUint32(ao.AttrId)
	bs.PushUint8(ao.AttrId_u)
	bs.PushString(ao.AttrName)
	bs.PushUint8(ao.AttrName_u)
	bs.PushString(ao.StrOption)
	bs.PushUint8(ao.StrOption_u)
	bs.PushUint32(ao.GroupId)
	bs.PushUint8(ao.GroupId_u)
	bs.PushString(ao.GroupName)
	bs.PushUint8(ao.GroupName_u)
	bs.PushUint32(ao.Order)
	bs.PushUint8(ao.Order_u)
	bs.PushUint32(ao.GroupOrder)
	bs.PushUint8(ao.GroupOrder_u)
	bs.PushString(ao.ReserveStr)
	bs.PushUint8(ao.ReserveStr_u)

	return bs.IsGood()
}

func (ao *AoDetailmainPoItemDetailNodeXXOO) getClassLen() int {

	bsLen := NewByteStream()
	bsLen.SetRealWrite(false)
	ao.serialize_internal(bsLen)
	length := bsLen.GetWrittenLength()

	return length
}

func (ao *AoDetailmainPoItemDetailNodeXXOO) UnSerialize(bs *ByteStream) error {
	classLen, err := bs.PopUint32()

	if err != nil {
		return err
	}
	startPop := bs.GetReadLength()

	ao.Version, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.Version_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.AttrId, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.AttrId_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.AttrName, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.AttrName_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.StrOption, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.StrOption_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.GroupId, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.GroupId_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.GroupName, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.GroupName_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.Order, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.Order_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.GroupOrder, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.GroupOrder_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.ReserveStr, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.ReserveStr_u, err = bs.PopUint8()
	if err != nil {
		return err
	}

	ao.Compat(bs, classLen, startPop)
	return nil
}

/* AoDetailmainPoItemDetailNodeXXOO END */

/* AoDetailmainPoItemStockInfoXXOO START */

type AoDetailmainPoItemStockInfoXXOO struct {
	AoXXOO
	UpdateTime      uint32
	UpdateTime_u    uint8
	DataState       uint32
	DataState_u     uint8
	StockNum        int32
	StockNum_u      uint8
	AreaId          uint32
	AreaId_u        uint8
	SupplyStockId   uint32
	SupplyStockId_u uint8
	DcId            string
	DcId_u          uint8
	ReserveStr      string
	ReserveStr_u    uint8
	StockSysno      uint64
	StockSysno_u    uint8
}

func NewAoDetailmainPoItemStockInfoXXOO() *AoDetailmainPoItemStockInfoXXOO {
	model := &AoDetailmainPoItemStockInfoXXOO{}
	model.Version = uint32(1)

	return model
}

/*
func (ao *AoDetailmainPoItemStockInfoXXOO) Serialize(bs *ByteStream) bool {

	bs.PushUint32(uint32(ao.getClassLen()))

	return ao.serialize_internal(bs)
}

func (ao *AoDetailmainPoItemStockInfoXXOO) serialize_internal(bs *ByteStream) bool {
	bs.PushUint32(ao.Version)
	bs.PushUint8(ao.Version_u)
	bs.PushUint32(ao.UpdateTime)
	bs.PushUint8(ao.UpdateTime_u)
	bs.PushUint32(ao.DataState)
	bs.PushUint8(ao.DataState_u)
	bs.PushInt32(ao.StockNum)
	bs.PushUint8(ao.StockNum_u)
	bs.PushUint32(ao.AreaId)
	bs.PushUint8(ao.AreaId_u)
	bs.PushUint32(ao.SupplyStockId)
	bs.PushUint8(ao.SupplyStockId_u)
	bs.PushString(ao.DcId)
	bs.PushUint8(ao.DcId_u)
	bs.PushString(ao.ReserveStr)
	bs.PushUint8(ao.ReserveStr_u)
	if ao.Version >= uint32(1) {
		bs.PushUint64(ao.StockSysno)
		bs.PushUint8(ao.StockSysno_u)
	}

	return bs.IsGood()
}

func (ao *AoDetailmainPoItemStockInfoXXOO) getClassLen() int {

	bsLen := NewByteStream()
	bsLen.SetRealWrite(false)
	ao.serialize_internal(bsLen)
	length := bsLen.GetWrittenLength()

	return length
}
*/
func (ao *AoDetailmainPoItemStockInfoXXOO) UnSerialize(bs *ByteStream) error {
	classLen, err := bs.PopUint32()

	if err != nil {
		return err
	}
	startPop := bs.GetReadLength()

	ao.Version, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.Version_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.UpdateTime, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.UpdateTime_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.DataState, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.DataState_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.StockNum, err = bs.PopInt32()
	if err != nil {
		return err
	}
	ao.StockNum_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.AreaId, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.AreaId_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SupplyStockId, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.SupplyStockId_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.DcId, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.DcId_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.ReserveStr, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.ReserveStr_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	if ao.Version >= uint32(1) {
		ao.StockSysno, err = bs.PopUint64()
		if err != nil {
			return err
		}
		ao.StockSysno_u, err = bs.PopUint8()
		if err != nil {
			return err
		}
	}

	ao.Compat(bs, classLen, startPop)
	return nil
}

/* AoDetailmainPoItemStockInfoXXOO END */



/*	AoSkuInfoById   start	*/
type AoSkuInfoByIdFilterXXOO struct {
	AoXXOO
	Skuid   uint64
	Skuid_u uint8
}

func NewAoSkuInfoByIdFilterXXOO() *AoSkuInfoByIdFilterXXOO {
	model := &AoSkuInfoByIdFilterXXOO{}

	return model
}

func (ao *AoSkuInfoByIdFilterXXOO) SetSkuid(skuid uint64) {
	ao.Skuid = skuid
	ao.Skuid_u = 1
}

func (ao *AoSkuInfoByIdFilterXXOO) Serialize(bs *ByteStream) bool {

	bs.PushUint32(uint32(ao.getClassLen()))

	return ao.serialize_internal(bs)
}

func (ao *AoSkuInfoByIdFilterXXOO) serialize_internal(bs *ByteStream) bool {
	bs.PushUint32(ao.Version)
	bs.PushUint8(ao.Version_u)
	bs.PushUint64(ao.Skuid)
	bs.PushUint8(ao.Skuid_u)

	return bs.IsGood()
}

func (ao *AoSkuInfoByIdFilterXXOO) getClassLen() int {

	bsLen := NewByteStream()
	bsLen.SetRealWrite(false)
	ao.serialize_internal(bsLen)
	length := bsLen.GetWrittenLength()

	return length
}

/**			SkuAllInfo				**/
type AoSkuInfoByIdSkuAllInfoXXOO struct {
	AoXXOO
	Skuid             uint64
	Skuid_u           uint8
	OrdinaryInfoPo    AoSkuInfoByIdOrdinaryInfoXXOO
	OrdinaryInfoPo_u  uint8
	SkuDetailInfoPo   []AoSkuInfoByIdSkuDetailInfoXXOO
	SkuDetailInfoPo_u uint8
	AreaInfoPo        map[uint32]AoSkuInfoByIdAreaInfoXXOO
	AreaInfoPo_u      uint8
	SkuPicPo          []AoSkuInfoByIdSkuPicXXOO
	SkuPicPo_u        uint8
	SkuCombPo         []AoSkuInfoByIdSkuCombXXOO
	SkuCombPo_u       uint8
}

func NewAoSkuInfoByIdSkuAllInfoXXOO() *AoSkuInfoByIdSkuAllInfoXXOO {
	model := &AoSkuInfoByIdSkuAllInfoXXOO{}
	return model
}

func (ao *AoSkuInfoByIdSkuAllInfoXXOO) UnSerialize(bs *ByteStream) error {
	classLen, err := bs.PopUint32()
	if err != nil {
		return err
	}
	startPop := bs.GetReadLength()

	ao.Version, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.Version_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	//fmt.Printf("version:%d", ao.Version)
	ao.Skuid, err = bs.PopUint64()
	if err != nil {
		return err
	}
	ao.Skuid_u, err = bs.PopUint8()
	if err != nil {
		return err
	}

	ordinaryInfoPo := NewAoSkuInfoByIdOrdinaryInfoXXOO()
	err = ordinaryInfoPo.UnSerialize(bs)
	if err != nil {
		return err
	}
	ao.OrdinaryInfoPo = *ordinaryInfoPo
	ao.OrdinaryInfoPo_u, err = bs.PopUint8()
	if err != nil {
		return err
	}

	lengthSkuDetail, err := bs.PopUint32()
	if err != nil {
		return err
	}
	for i := uint32(0); i < lengthSkuDetail; i++ {

		skuDetailInfoPo := NewAoSkuInfoByIdSkuDetailInfoXXOO()
		err = skuDetailInfoPo.UnSerialize(bs)
		if err != nil {
			return err
		}
		ao.SkuDetailInfoPo = append(ao.SkuDetailInfoPo, *skuDetailInfoPo)
	}
	ao.SkuDetailInfoPo_u, err = bs.PopUint8()
	if err != nil {
		return err
	}

	lengthAreaInfoPo, err := bs.PopUint32()
	if err != nil {
		return err
	}
	var areaId uint32
	//fmt.Printf("len:%d", lengthAreaInfoPo)
	ao.AreaInfoPo = make(map[uint32]AoSkuInfoByIdAreaInfoXXOO, lengthAreaInfoPo)
	//var areaInfoPo *AoSkuInfoByIdAreaInfoXXOO
	for i := uint32(0); i < lengthAreaInfoPo; i++ {
		//fmt.Printf("i:%d", i)
		areaId, err = bs.PopUint32()
		if err != nil {
			return errors.New(fmt.Sprintf("areaInfo error: %s,length:%d,index:%d", err.Error(), lengthAreaInfoPo, i))
		}
		areaInfoPo := NewAoSkuInfoByIdAreaInfoXXOO()
		err = areaInfoPo.UnSerialize(bs)
		if err != nil {
			return err
		}
		//fmt.Printf("areaId:%d,areaInfo:%+v", areaId, *areaInfoPo)
		ao.AreaInfoPo[areaId] = *areaInfoPo
	}
	//fmt.Printf("infoPros:%+v", ao.AreaInfoPo)
	ao.AreaInfoPo_u, err = bs.PopUint8()
	if err != nil {
		return err
	}

	lengthSkuPic, err := bs.PopUint32()
	if err != nil {
		return err
	}
	for i := uint32(0); i < lengthSkuPic; i++ {

		SkuPicPo := NewAoSkuInfoByIdSkuPicXXOO()
		err = SkuPicPo.UnSerialize(bs)
		if err != nil {
			return err
		}
		ao.SkuPicPo = append(ao.SkuPicPo, *SkuPicPo)
	}
	ao.SkuPicPo_u, err = bs.PopUint8()
	if err != nil {
		return err
	}

	if ao.Version >= 1 {
		lengthskuCombPo, err := bs.PopUint32()
		if err != nil {
			return err
		}
		for i := uint32(0); i < lengthskuCombPo; i++ {

			skuCombPo := NewAoSkuInfoByIdSkuCombXXOO()
			err = skuCombPo.UnSerialize(bs)
			if err != nil {
				return err
			}
			ao.SkuCombPo = append(ao.SkuCombPo, *skuCombPo)
		}
		ao.SkuCombPo_u, err = bs.PopUint8()
		if err != nil {
			return err
		}

	}

	ao.Compat(bs, classLen, startPop)
	return nil
}

type AoSkuInfoByIdOrdinaryInfoXXOO struct {
	AoXXOO
	Skuid                    uint64
	Skuid_u                  uint8
	CooperatorSubaccountId   uint64
	CooperatorSubaccountId_u uint8
	CooperatorId             uint64
	CooperatorId_u           uint8
	CategoryId               uint32
	CategoryId_u             uint8
	SkuType                  uint32
	SkuType_u                uint8
	SkuOperationModel        uint32
	SkuOperationModel_u      uint8
	BarCode                  string
	BarCode_u                uint8
	SkuLocalCode             string
	SkuLocalCode_u           uint8
	SkuDimensionalBarCode    string
	SkuDimensionalBarCode_u  uint8
	SkuTitle                 string
	SkuTitle_u               uint8
	SkuSubTitle              string
	SkuSubTitle_u            uint8
	SkuPromotDesc            string
	SkuPromotDesc_u          uint8
	SkuSaleAttr              string
	SkuSaleAttr_u            uint8
	SkuSaleAttrDesc          string
	SkuSaleAttrDesc_u        uint8
	SkuKeyAttr               string
	SkuKeyAttr_u             uint8
	SkuKeyAttrDesc           string
	SkuKeyAttrDesc_u         uint8
	SkuReferPrice            uint32
	SkuReferPrice_u          uint8
	SkuProperty              []uint8
	SkuProperty_u            uint8
	SkuState                 uint8
	SkuState_u               uint8
	SkuWeight                uint32
	SkuWeight_u              uint8
	SkuNetWeight             uint32
	SkuNetWeight_u           uint8
	SkuVolume                uint64
	SkuVolume_u              uint8
	SkuSizeX                 uint16
	SkuSizeX_u               uint8
	SkuSizeY                 uint16
	SkuSizeY_u               uint8
	SkuSizeZ                 uint16
	SkuSizeZ_u               uint8
	SkuCategoryAttr          string
	SkuCategoryAttr_u        uint8
	SkuKeywords              string
	SkuKeywords_u            uint8
	SkuVatrate               uint32
	SkuVatrate_u             uint8
	SkuSnapVersion           uint32
	SkuSnapVersion_u         uint8
	SkuBuyLimit              uint32
	SkuBuyLimit_u            uint8
	SkuBrand                 uint32
	SkuBrand_u               uint8
	SkuExchangePoint         uint32
	SkuExchangePoint_u       uint8
	MainSkuid                uint64
	MainSkuid_u              uint8
	SkuLastUpTime            uint32
	SkuLastUpTime_u          uint8
	SkuLastDownTime          uint32
	SkuLastDownTime_u        uint8
	SkuAddTime               uint32
	SkuAddTime_u             uint8
	SkuLastUpdateTime        uint32
	SkuLastUpdateTime_u      uint8
	LastUpdateTime           uint32
	LastUpdateTime_u         uint8
	Sort                     uint32
	Sort_u                   uint8
	ExpertRecommend          string
	ExpertRecommend_u        uint8
	TaxInfo                  AoSkuInfoByIdTaxXXOO
	TaxInfo_u                uint8
	PopSkuCode               string
	PopSkuCode_u             uint8
	SkuEffectiveTime         uint32
	SkuEffectiveTime_u       uint8
	SkuExpireTime            uint32
	SkuExpireTime_u          uint8
	SkuStockExpireTime       uint32
	SkuStockExpireTime_u     uint8
	SkuContacterNum          uint32
	SkuContacterNum_u        uint8
	SkuTravellerNum          uint32
	SkuTravellerNum_u        uint8
	SkuSaleArea              string
	SkuSaleArea_u            uint8
	SkuUnitCost              uint32
	SkuUnitCost_u            uint8
}

func NewAoSkuInfoByIdOrdinaryInfoXXOO() *AoSkuInfoByIdOrdinaryInfoXXOO {
	return &AoSkuInfoByIdOrdinaryInfoXXOO{}
}

func (ao *AoSkuInfoByIdOrdinaryInfoXXOO) UnSerialize(bs *ByteStream) error {

	classLen, err := bs.PopUint32()
	if err != nil {
		return err
	}
	startPop := bs.GetReadLength()

	ao.Version, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.Version_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.Skuid, err = bs.PopUint64()
	if err != nil {
		return err
	}
	ao.Skuid_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.CooperatorSubaccountId, err = bs.PopUint64()
	if err != nil {
		return err
	}
	ao.CooperatorSubaccountId_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.CooperatorId, err = bs.PopUint64()
	if err != nil {
		return err
	}
	ao.CooperatorId_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.CategoryId, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.CategoryId_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuType, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.SkuType_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuOperationModel, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.SkuOperationModel_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.BarCode, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.BarCode_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuLocalCode, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.SkuLocalCode_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuDimensionalBarCode, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.SkuDimensionalBarCode_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuTitle, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.SkuTitle_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuSubTitle, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.SkuSubTitle_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuPromotDesc, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.SkuPromotDesc_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuSaleAttr, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.SkuSaleAttr_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuSaleAttrDesc, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.SkuSaleAttrDesc_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuKeyAttr, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.SkuKeyAttr_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuKeyAttrDesc, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.SkuKeyAttrDesc_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuReferPrice, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.SkuReferPrice_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuProperty, err = bs.PopBitSet()
	if err != nil {
		return err
	}
	ao.SkuProperty_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuState, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuState_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuWeight, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.SkuWeight_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuNetWeight, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.SkuNetWeight_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuVolume, err = bs.PopUint64()
	if err != nil {
		return err
	}
	ao.SkuVolume_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuSizeX, err = bs.PopUint16()
	if err != nil {
		return err
	}
	ao.SkuSizeX_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuSizeY, err = bs.PopUint16()
	if err != nil {
		return err
	}
	ao.SkuSizeY_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuSizeZ, err = bs.PopUint16()
	if err != nil {
		return err
	}
	ao.SkuSizeZ_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuCategoryAttr, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.SkuCategoryAttr_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuKeywords, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.SkuKeywords_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuVatrate, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.SkuVatrate_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuSnapVersion, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.SkuSnapVersion_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuBuyLimit, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.SkuBuyLimit_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuBrand, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.SkuBrand_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuExchangePoint, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.SkuExchangePoint_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.MainSkuid, err = bs.PopUint64()
	if err != nil {
		return err
	}
	ao.MainSkuid_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuLastUpTime, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.SkuLastUpTime_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuLastDownTime, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.SkuLastDownTime_u, err = bs.PopUint8()
	if err != nil {
		return err
	}

	ao.SkuAddTime, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.SkuAddTime_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuLastUpdateTime, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.SkuLastUpdateTime_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.LastUpdateTime, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.LastUpdateTime_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.Sort, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.Sort_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.ExpertRecommend, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.ExpertRecommend_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.Reserve, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.Reserve_u, err = bs.PopUint8()
	if err != nil {
		return err
	}

	if ao.Version >= 1 {
		taxPro := NewAoSkuInfoByIdTaxXXOO()
		err = taxPro.UnSerialize(bs)
		if err != nil {
			return err
		}
		ao.TaxInfo = *taxPro
		ao.TaxInfo_u, err = bs.PopUint8()
		if err != nil {
			return err
		}
	}
	if ao.Version >= 2 {
		ao.PopSkuCode, err = bs.PopString()
		if err != nil {
			return err
		}
		ao.PopSkuCode_u, err = bs.PopUint8()
		if err != nil {
			return err
		}
	}
	if ao.Version >= 3 {
		ao.SkuEffectiveTime, err = bs.PopUint32()
		if err != nil {
			return err
		}
		ao.SkuEffectiveTime_u, err = bs.PopUint8()
		if err != nil {
			return err
		}
		ao.SkuExpireTime, err = bs.PopUint32()
		if err != nil {
			return err
		}
		ao.SkuExpireTime_u, err = bs.PopUint8()
		if err != nil {
			return err
		}
	}
	if ao.Version >= 4 {
		ao.SkuStockExpireTime, err = bs.PopUint32()
		if err != nil {
			return err
		}
		ao.SkuStockExpireTime_u, err = bs.PopUint8()
		if err != nil {

		}
	}
	if ao.Version >= 5 {
		ao.SkuContacterNum, err = bs.PopUint32()
		if err != nil {
			return err
		}
		ao.SkuContacterNum_u, err = bs.PopUint8()
		if err != nil {

		}
		ao.SkuTravellerNum, err = bs.PopUint32()
		if err != nil {
			return err
		}
		ao.SkuTravellerNum_u, err = bs.PopUint8()
		if err != nil {

		}
	}
	if ao.Version >= 6 {
		ao.SkuSaleArea, err = bs.PopString()
		if err != nil {
			return err
		}
		ao.SkuSaleArea_u, err = bs.PopUint8()
		if err != nil {

		}
	}
	if ao.Version >= 7 {
		ao.SkuUnitCost, err = bs.PopUint32()
		if err != nil {
			return err
		}
		ao.SkuUnitCost_u, err = bs.PopUint8()
		if err != nil {

		}
	}

	ao.Compat(bs, classLen, startPop)
	return nil
}

type AoSkuInfoByIdTaxXXOO struct {
	AoXXOO
	TaxRate             uint32
	TaxRate_u           uint8
	CustomNo            uint32
	CustomNo_u          uint8
	CustomRecordNo      string
	CustomRecordNo_u    uint8
	CustomRecordName    string
	CustomRecordName_u  uint8
	DistributionWay     uint32
	DistributionWay_u   uint8
	BrandNation         string
	BrandNation_u       uint8
	BrandFlag           string
	BrandFlag_u         uint8
	DistributionPlace   string
	DistributionPlace_u uint8
	SkucustomRecordNo   string
	SkucustomRecordNo_u uint8
}

func NewAoSkuInfoByIdTaxXXOO() *AoSkuInfoByIdTaxXXOO {
	return &AoSkuInfoByIdTaxXXOO{}
}

func (ao *AoSkuInfoByIdTaxXXOO) UnSerialize(bs *ByteStream) error {
	classLen, err := bs.PopUint32()
	if err != nil {
		return err
	}
	startPop := bs.GetReadLength()

	ao.Version, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.Version_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.TaxRate, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.TaxRate_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.Reserve, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.Reserve_u, err = bs.PopUint8()
	if err != nil {
		return err
	}

	if ao.Version >= 1 {
		ao.CustomNo, err = bs.PopUint32()
		if err != nil {
			return err
		}
		ao.CustomNo_u, err = bs.PopUint8()
		if err != nil {
			return err
		}
		ao.CustomRecordNo, err = bs.PopString()
		if err != nil {
			return err
		}
		ao.CustomRecordNo_u, err = bs.PopUint8()
		if err != nil {
			return err
		}
		ao.CustomRecordName, err = bs.PopString()
		if err != nil {
			return err
		}
		ao.CustomRecordName_u, err = bs.PopUint8()
		if err != nil {
			return err
		}
		ao.DistributionWay, err = bs.PopUint32()
		if err != nil {
			return err
		}
		ao.DistributionWay_u, err = bs.PopUint8()
		if err != nil {
			return err
		}
	}

	if ao.Version >= 2 {
		ao.BrandNation, err = bs.PopString()
		if err != nil {
			return err
		}
		ao.BrandNation_u, err = bs.PopUint8()
		if err != nil {
			return err
		}
		ao.BrandFlag, err = bs.PopString()
		if err != nil {
			return err
		}
		ao.BrandFlag_u, err = bs.PopUint8()
		if err != nil {
			return err
		}
		ao.DistributionPlace, err = bs.PopString()
		if err != nil {
			return err
		}
		ao.DistributionPlace_u, err = bs.PopUint8()
		if err != nil {
			return err
		}
	}

	if ao.Version >= 3 {
		ao.SkucustomRecordNo, err = bs.PopString()
		if err != nil {
			return err
		}
		ao.SkucustomRecordNo_u, err = bs.PopUint8()
		if err != nil {
			return err
		}
	}

	ao.Compat(bs, classLen, startPop)
	return nil
}

type AoSkuInfoByIdSkuDetailInfoXXOO struct {
	AoXXOO
	Skuid                     uint64
	Skuid_u                   uint8
	SkuDetailType             uint8
	SkuDetailType_u           uint8
	Channel                   uint32
	Channel_u                 uint8
	CooperatorSubaccountId    uint64
	CooperatorSubaccountId_u  uint8
	Property                  []uint8
	Property_u                uint8
	DetailUrl                 string
	DetailUrl_u               uint8
	SkuDetail                 string
	SkuDetail_u               uint8
	SkuDetailAddTime          uint32
	SkuDetailAddTime_u        uint8
	SkuDetailLastUpdateTime   uint32
	SkuDetailLastUpdateTime_u uint8
}

func NewAoSkuInfoByIdSkuDetailInfoXXOO() *AoSkuInfoByIdSkuDetailInfoXXOO {
	return &AoSkuInfoByIdSkuDetailInfoXXOO{}
}

func (ao *AoSkuInfoByIdSkuDetailInfoXXOO) UnSerialize(bs *ByteStream) error {
	classLen, err := bs.PopUint32()
	if err != nil {
		return err
	}
	startPop := bs.GetReadLength()

	ao.Version, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.Version_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.Skuid, err = bs.PopUint64()
	if err != nil {
		return err
	}
	ao.Skuid_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuDetailType, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuDetailType_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.Channel, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.Channel_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.CooperatorSubaccountId, err = bs.PopUint64()
	if err != nil {
		return err
	}
	ao.CooperatorSubaccountId_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.Property, err = bs.PopBitSet()
	if err != nil {
		return err
	}
	ao.Property_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.DetailUrl, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.DetailUrl_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuDetail, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.SkuDetail_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuDetailAddTime, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.SkuDetailAddTime_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.SkuDetailLastUpdateTime, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.SkuDetailLastUpdateTime_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.Reserve, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.Reserve_u, err = bs.PopUint8()
	if err != nil {
		return err
	}

	ao.Compat(bs, classLen, startPop)
	return nil
}

type AoSkuInfoByIdAreaInfoXXOO struct {
	AoXXOO
	AreaId                        uint32
	AreaId_u                      uint8
	CoSubAccountId                uint64
	CoSubAccountId_u              uint8
	CoSkuAreaCode                 string
	CoSkuAreaCode_u               uint8
	CoBarCode                     string
	CoBarCode_u                   uint8
	PromotionType                 uint32
	PromotionType_u               uint8
	AreaPromotDesc                string
	AreaPromotDesc_u              uint8
	AreaPrice                     uint32
	AreaPrice_u                   uint8
	AreaPrePrice                  uint32
	AreaPrePrice_u                uint8
	AreaCostPrice                 uint32
	AreaCostPrice_u               uint8
	AreaBussiessPrice             uint32
	AreaBussiessPrice_u           uint8
	AreaProperty                  []uint8
	AreaProperty_u                uint8
	AreaState                     uint8
	AreaState_u                   uint8
	AreaMinBuyCount               uint32
	AreaMinBuyCount_u             uint8
	AreaMaxBuyCount               uint32
	AreaMaxBuyCount_u             uint8
	AreaAddTime                   uint32
	AreaAddTime_u                 uint8
	AreaLastUpdateTime            uint32
	AreaLastUpdateTime_u          uint8
	AreaBuyTimes                  uint32
	AreaBuyTimes_u                uint8
	AreaFirstPublishTime          uint32
	AreaFirstPublishTime_u        uint8
	AreaPromotDescEffectiveTime   uint32
	AreaPromotDescEffectiveTime_u uint8
	AreaPromotDescExpireTime      uint32
	AreaPromotDescExpireTime_u    uint8
	AreaShareCommission           uint32
	AreaShareCommission_u         uint8
	EntityId		     	uint32
	EntityId_u			uint8
	StoreId				string
	StoreId_u			uint8
	ChannelId			uint32
	ChannelId_u			uint8
	Detail				string
	Detail_u			uint8
	WorkStateCode			string
	WorkStateCode_u			uint8
	CirculationModeCode		string
	CirculationModeCode_u		uint8
	AreaBuyer			string
	AreaBuyer_u			uint8

}

func NewAoSkuInfoByIdAreaInfoXXOO() *AoSkuInfoByIdAreaInfoXXOO {
	return &AoSkuInfoByIdAreaInfoXXOO{}
}

func (ao *AoSkuInfoByIdAreaInfoXXOO) UnSerialize(bs *ByteStream) error {
	classLen, err := bs.PopUint32()
	if err != nil {
		return err
	}
	startPop := bs.GetReadLength()

	ao.Version, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.Version_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.AreaId, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.AreaId_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.CoSubAccountId, err = bs.PopUint64()
	if err != nil {
		return err
	}
	ao.CoSubAccountId_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.CoSkuAreaCode, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.CoSkuAreaCode_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.CoBarCode, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.CoBarCode_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.PromotionType, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.PromotionType_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.AreaPromotDesc, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.AreaPromotDesc_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.AreaPrice, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.AreaPrice_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.AreaPrePrice, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.AreaPrePrice_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.AreaCostPrice, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.AreaCostPrice_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.AreaBussiessPrice, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.AreaBussiessPrice_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.AreaProperty, err = bs.PopBitSet()
	if err != nil {
		return err
	}
	ao.AreaProperty_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.AreaState, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.AreaState_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.AreaMinBuyCount, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.AreaMinBuyCount_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.AreaMaxBuyCount, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.AreaMaxBuyCount_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.AreaAddTime, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.AreaAddTime_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.AreaLastUpdateTime, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.AreaLastUpdateTime_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.Reserve, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.Reserve_u, err = bs.PopUint8()
	if err != nil {
		return err
	}

	if ao.Version >= 1 {
		ao.AreaBuyTimes, err = bs.PopUint32()
		if err != nil {
			return err
		}
		ao.AreaBuyTimes_u, err = bs.PopUint8()
		if err != nil {
			return err
		}
	}

	if ao.Version >= 2 {
		ao.AreaFirstPublishTime, err = bs.PopUint32()
		if err != nil {
			return err
		}
		ao.AreaFirstPublishTime_u, err = bs.PopUint8()
		if err != nil {
			return err
		}
		ao.AreaPromotDescEffectiveTime, err = bs.PopUint32()
		if err != nil {
			return err
		}
		ao.AreaPromotDescEffectiveTime_u, err = bs.PopUint8()
		if err != nil {
			return err
		}
		ao.AreaPromotDescExpireTime, err = bs.PopUint32()
		if err != nil {
			return err
		}
		ao.AreaPromotDescExpireTime_u, err = bs.PopUint8()
		if err != nil {
			return err
		}
	}

	if ao.Version >= 3 {
		ao.AreaShareCommission, err = bs.PopUint32()
		if err != nil {
			return err
		}
		ao.AreaShareCommission_u, err = bs.PopUint8()
		if err != nil {
			return err
		}
	}

	if ao.Version >= 4 {
		ao.EntityId, err = bs.PopUint32()
		if err != nil {
			return err
		}
		ao.EntityId_u, err = bs.PopUint8()
		if err != nil {
			return err
		}
		ao.StoreId, err = bs.PopString()
		if err != nil {
			return err
		}
		ao.StoreId_u, err = bs.PopUint8()
		if err != nil {
			return err
		}
		ao.ChannelId, err = bs.PopUint32()
		if err != nil {
			return err
		}
		ao.ChannelId_u, err = bs.PopUint8()
		if err != nil {
			return err
		}
		ao.Detail, err = bs.PopString()
		if err != nil {
			return err
		}
		ao.Detail_u, err = bs.PopUint8()
		if err != nil {
			return err
		}
		ao.WorkStateCode, err = bs.PopString()
		if err != nil {
			return err
		}
		ao.WorkStateCode_u, err = bs.PopUint8()
		if err != nil {
			return err
		}
		ao.CirculationModeCode, err = bs.PopString()
		if err != nil {
			return err
		}
		ao.CirculationModeCode_u, err = bs.PopUint8()
		if err != nil {
			return err
		}
		ao.AreaBuyer, err = bs.PopString()
		if err != nil {
			return err
		}
		ao.AreaBuyer_u, err = bs.PopUint8()
		if err != nil {
			return err
		}

	}

	ao.Compat(bs, classLen, startPop)
	return nil
}

type AoSkuInfoByIdSkuPicXXOO struct {
	AoXXOO
	Channel             uint32
	Channel_u           uint8
	CoSubAccountId      uint64
	CoSubAccountId_u    uint8
	PicUrl              string
	PicUrl_u            uint8
	PicFileId           string
	PicFileId_u         uint8
	PicSize             string
	PicSize_u           uint8
	PicIndex            uint8
	PicIndex_u          uint8
	PicType             uint8
	PicType_u           uint8
	PicProperty         []uint8
	PicProperty_u       uint8
	PicDesc             string
	PicDesc_u           uint8
	PicUploadTime       uint32
	PicUploadTime_u     uint8
	PicLastUpdateTime   uint32
	PicLastUpdateTime_u uint8
}

func NewAoSkuInfoByIdSkuPicXXOO() *AoSkuInfoByIdSkuPicXXOO {
	return &AoSkuInfoByIdSkuPicXXOO{}
}

func (ao *AoSkuInfoByIdSkuPicXXOO) UnSerialize(bs *ByteStream) error {
	classLen, err := bs.PopUint32()
	if err != nil {
		return err
	}
	startPop := bs.GetReadLength()

	ao.Version, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.Version_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.Channel, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.Channel_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.CoSubAccountId, err = bs.PopUint64()
	if err != nil {
		return err
	}
	ao.CoSubAccountId_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.PicUrl, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.PicUrl_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.PicFileId, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.PicFileId_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.PicSize, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.PicSize_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.PicIndex, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.PicIndex_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.PicType, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.PicType_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.PicProperty, err = bs.PopBitSet()
	if err != nil {
		return err
	}
	ao.PicProperty_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.PicDesc, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.PicDesc_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.PicUploadTime, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.PicUploadTime_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.PicLastUpdateTime, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.PicLastUpdateTime_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.Reserve, err = bs.PopString()
	if err != nil {
		return err
	}
	ao.Reserve_u, err = bs.PopUint8()
	if err != nil {
		return err
	}

	ao.Compat(bs, classLen, startPop)
	return nil
}

type AoSkuInfoByIdSkuCombXXOO struct {
	AoXXOO
	CombSkuId         uint64
	CombSkuId_u       uint8
	CombParentSkuId   uint64
	CombParentSkuId_u uint8
	ShareRatio        uint64
	ShareRatio_u      uint8
	CombSkuNum        uint64
	CombSkuNum_u      uint8
}

func NewAoSkuInfoByIdSkuCombXXOO() *AoSkuInfoByIdSkuCombXXOO {
	return &AoSkuInfoByIdSkuCombXXOO{}
}

func (ao *AoSkuInfoByIdSkuCombXXOO) UnSerialize(bs *ByteStream) error {
	classLen, err := bs.PopUint32()
	if err != nil {
		return err
	}
	startPop := bs.GetReadLength()

	ao.Version, err = bs.PopUint32()
	if err != nil {
		return err
	}
	ao.Version_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.CombSkuId, err = bs.PopUint64()
	if err != nil {
		return err
	}
	ao.CombSkuId_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.CombParentSkuId, err = bs.PopUint64()
	if err != nil {
		return err
	}
	ao.CombParentSkuId_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.ShareRatio, err = bs.PopUint64()
	if err != nil {
		return err
	}
	ao.ShareRatio_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	ao.CombSkuNum, err = bs.PopUint64()
	if err != nil {
		return err
	}
	ao.CombSkuNum_u, err = bs.PopUint8()
	if err != nil {
		return err
	}

	ao.Compat(bs, classLen, startPop)
	return nil
}

/*	AoSkuInfoById 	end	*/
