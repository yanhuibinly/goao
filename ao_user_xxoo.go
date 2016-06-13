package goao

import (
	"reflect"
	"errors"
)

type AoUserInfoXXOO struct {
	AoXXOO
	Uid                     uint64 //<uint64> uid, 用户id
	Uid_u                   uint8
	AccountType             uint8 //<uint8> 帐号类型，参见user_comm_define.h中的E_ACCOUNT_TYPE
	AccountType_u           uint8
	Mobile                  string //<std::string> 手机号
	Mobile_u                uint8
	Email                   string //<std::string> 电子邮箱
	Email_u                 uint8
	LoginAccount            string //<std::string> 个性化账号
	LoginAccount_u          uint8
	WechatOpenid            map[interface{}]interface{} //<std::map<std::string,std::string> > 微信OpenID：appid-openid
	WechatOpenid_u          uint8
	WechatUnionid           string //<std::string> 微信unionID
	WechatUnionid_u         uint8
	UserType                uint8 //<uint8> 用户类型，参见user_comm_define.h中的E_USER_TYPE
	UserType_u              uint8
	BitProperty             []uint8 //<std::bitset<64> > 用户属性位BitSet，具体意义参见b2b2c_define.h中的E_USER_PROPERTY
	BitProperty_u           uint8
	LoginPasswd             string //<std::string> 用户登录密码
	LoginPasswd_u           uint8
	TradePasswd             string //<std::string> 用户交易密码
	TradePasswd_u           uint8
	PasswordSecureLevel     uint8 //<uint8> 登录密码安全级别
	PasswordSecureLevel_u   uint8
	Skey                    string //<std::string> 用户sessionkey
	Skey_u                  uint8
	SkeyInvalidTime         uint32 //<uint32> 用户sessionkey失效时间
	SkeyInvalidTime_u       uint8
	DiffSrcRegTime          map[interface{}]interface{} //<std::map<uint32,uint32> > 不同来源用户注册时间, key为来源，value为注册时间
	DiffSrcRegTime_u        uint8
	Rating                  uint8 //<uint8> 用户评级
	Rating_u                uint8
	BabyInfo                []AoBabyInfoXXOO //<std::vector<b2b2c::user::po::CBabyInfoPo> > 宝宝列表
	BabyInfo_u              uint8
	RelationWithBaby        uint8 //<uint8> 与宝宝关系，参见user_comm_define.h中的E_RELATION_WITH_BABY
	RelationWithBaby_u      uint8
	Nickname                string //<std::string> 昵称
	Nickname_u              uint8
	Truename                string //<std::string> 真实姓名
	Truename_u              uint8
	Sex                     uint8 //<uint8> 性别：0-未知，1-女，2-男。参见E_USER_SEX
	Sex_u                   uint8
	Birthday                uint64 //<int64> 生日
	Birthday_u              uint8
	QQNumber                uint64 //<uint64> QQ号
	QQNumber_u              uint8
	Phone                   string //<std::string> 固定电话
	Phone_u                 uint8
	Fax                     string //<std::string> 传真
	Fax_u                   uint8
	Region                  string //<std::string> 所在地区id
	Region_u                uint8
	Address                 string //<std::string> 详细住址
	Address_u               uint8
	Community               string //<std::string> 小区
	Community_u             uint8
	Postcode                string //<std::string> 邮政编码
	Postcode_u              uint8
	IdentityType            uint8 //<uint8> 身份证件类型，参见user_comm_define.h中E_USER_IDTYPE
	IdentityType_u          uint8
	IdentityNum             string //<std::string> 身份证件号码
	IdentityNum_u           uint8
	Photo                   string //<std::string> 头像
	Photo_u                 uint8
	Signature               string //<std::string> 签名
	Signature_u             uint8
	Euid                    map[interface{}]interface{} //<std::map<uint32,std::string> > 外部导入用户euid, key为来源渠道（参见b2b2c_define.h中的E_USER_EUID_SOURCE）, value为外部uid
	Euid_u                  uint8
	Referrer                string //<std::string> 推荐人，招募人
	Referrer_u              uint8
	WechatAccount           string
	WechatAccount_u         uint8
	MemberLevel             uint8
	MemberLevel_u           uint8
	MembercardList          []AoMemberCardInfoXXOO
	MembercardList_u        uint8
	PregnantPassportList    []string
	PregnantPassportList_u  uint8
	Recruiter               string
	Recruiter_u             uint8
	Manager                 string
	Manager_u               uint8
	Creator                 string
	Creator_u               uint8
	CreatorDepartment       string
	CreatorDepartment_u     uint8
	RegisterSource          uint8
	RegisterSource_u        uint8
	MemberSource            uint8
	MemberSource_u          uint8
	IfSendErp               uint8
	IfSendErp_u             uint8
	SendErpTime             uint32
	SendErpTime_u           uint8
	DiffChannelActiveTime   map[interface{}]interface{} //<std::map<uint32,uint32> >
	DiffChannelActiveTime_u uint8
	PromoteActive           string
	PromoteActive_u         uint8
	BuyerProperty           AoBuyerPropertyXXOO
	BuyerProperty_u         uint8
	UserLableList           []string
	UserLableList_u         uint8
	UserLableRemarks        string
	UserLableRemarks_u      uint8
	MobileStatus            uint8    //<uint8_t> 手机号状态(版本>=20160413)
	MobileStatus_u          uint8    // uint8_t mobileStatus的flag位
	MemberBitProperty       []uint8  //<std::bitset<64> > 会员属性(版本>=20160413)
	MemberBitProperty_u     uint8    // uint8_t memberBitProperty的flag位
	PregnantPlan            uint8    //<uint8_t> 备孕计划(版本>=20160516)
	PregnantPlan_u          uint8    // uint8_t pregnantPlan的flag位
	UserPictureLableList    []string //<std::vector<std::string> > 用户图片标签列表(版本>=20160517)
	UserPictureLableList_u  uint8    // uint8_t userPictureLableList的flag位
}

func (ao *AoUserInfoXXOO) getClassLen() int {

	bsLen := NewByteStream()
	bsLen.SetRealWrite(false)
	ao.serialize_internal(bsLen)
	length := bsLen.GetWrittenLength()

	return length
}
func (ao *AoUserInfoXXOO) Serialize(bs *ByteStream) bool {
	bs.PushUint32(uint32(ao.getClassLen()))
	return ao.serialize_internal(bs)
}

func (a *AoUserInfoXXOO) serialize_internal(bs *ByteStream) bool {
	bs.PushUint32(a.Version) //<uint32> 版本号
	bs.PushUint64(a.Uid)     //<uint64> uid, 用户id
	bs.PushUint8(a.Uid_u)
	bs.PushUint8(a.AccountType) //<uint8> 帐号类型，参见user_comm_define.h中的E_ACCOUNT_TYPE
	bs.PushUint8(a.AccountType_u)
	bs.PushString(a.Mobile) //<std::string> 手机号
	bs.PushUint8(a.Mobile_u)
	bs.PushString(a.Email) //<std::string> 电子邮箱
	bs.PushUint8(a.Email_u)
	bs.PushString(a.LoginAccount) //<std::string> 个性化账号
	bs.PushUint8(a.LoginAccount_u)
	bs.PushMap(a.WechatOpenid) //<std::map<std::string,std::string> > 微信OpenID：appid-openid
	bs.PushUint8(a.WechatOpenid_u)
	bs.PushString(a.WechatUnionid) //<std::string> 微信unionID
	bs.PushUint8(a.WechatUnionid_u)
	bs.PushUint8(a.UserType) //<uint8> 用户类型，参见user_comm_define.h中的E_USER_TYPE
	bs.PushUint8(a.UserType_u)
	bs.PushBitSet(a.BitProperty) //<std::bitset<64> > 用户属性位BitSet，具体意义参见b2b2c_define.h中的E_USER_PROPERTY
	bs.PushUint8(a.BitProperty_u)
	bs.PushString(a.LoginPasswd) //<std::string> 用户登录密码
	bs.PushUint8(a.LoginPasswd_u)
	bs.PushString(a.TradePasswd) //<std::string> 用户交易密码
	bs.PushUint8(a.TradePasswd_u)
	bs.PushUint8(a.PasswordSecureLevel) //<uint8> 登录密码安全级别
	bs.PushUint8(a.PasswordSecureLevel_u)
	bs.PushString(a.Skey) //<std::string> 用户sessionkey
	bs.PushUint8(a.Skey_u)
	bs.PushUint32(a.SkeyInvalidTime) //<uint32> 用户sessionkey失效时间
	bs.PushUint8(a.SkeyInvalidTime_u)
	bs.PushMap(a.DiffSrcRegTime) //<std::map<uint32,uint32> > 不同来源用户注册时间, key为来源，value为注册时间
	bs.PushUint8(a.DiffSrcRegTime_u)
	bs.PushUint8(a.Rating) //<uint8> 用户评级
	bs.PushUint8(a.Rating_u)

	var babyInfo []IAoXXOO
	for _, v := range a.BabyInfo {
		babyInfo = append(babyInfo, IAoXXOO(&v))
	}
	bs.PushSet(babyInfo) //<std::vector<b2b2c::user::po::CBabyInfoPo> > 宝宝列表
	bs.PushUint8(a.BabyInfo_u)
	bs.PushUint8(a.RelationWithBaby) //<uint8> 与宝宝关系，参见user_comm_define.h中的E_RELATION_WITH_BABY
	bs.PushUint8(a.RelationWithBaby_u)
	bs.PushString(a.Nickname) //<std::string> 昵称
	bs.PushUint8(a.Nickname_u)
	bs.PushString(a.Truename) //<std::string> 真实姓名
	bs.PushUint8(a.Truename_u)
	bs.PushUint8(a.Sex) //<uint8> 性别：0-未知，1-女，2-男。参见E_USER_SEX
	bs.PushUint8(a.Sex_u)
	bs.PushUint64(a.Birthday) //<uint64> 生日
	bs.PushUint8(a.Birthday_u)
	bs.PushUint64(a.QQNumber) //<uint64> QQ号
	bs.PushUint8(a.QQNumber_u)
	bs.PushString(a.Phone) //<std::string> 固定电话
	bs.PushUint8(a.Phone_u)
	bs.PushString(a.Fax) //<std::string> 传真
	bs.PushUint8(a.Fax_u)
	bs.PushString(a.Region) //<std::string> 所在地区id
	bs.PushUint8(a.Region_u)
	bs.PushString(a.Address) //<std::string> 详细住址
	bs.PushUint8(a.Address_u)
	bs.PushString(a.Community) //<std::string> 小区
	bs.PushUint8(a.Community_u)
	bs.PushString(a.Postcode) //<std::string> 邮政编码
	bs.PushUint8(a.Postcode_u)
	bs.PushUint8(a.IdentityType) //<uint8> 身份证件类型，参见user_comm_define.h中E_USER_IDTYPE
	bs.PushUint8(a.IdentityType_u)
	bs.PushString(a.IdentityNum) //<std::string> 身份证件号码
	bs.PushUint8(a.IdentityNum_u)
	bs.PushString(a.Photo) //<std::string> 头像
	bs.PushUint8(a.Photo_u)
	bs.PushString(a.Signature) //<std::string> 签名
	bs.PushUint8(a.Signature_u)
	bs.PushMap(a.Euid) //<std::map<uint32,std::string> > 外部导入用户euid, key为来源渠道（参见b2b2c_define.h中的E_USER_EUID_SOURCE）, value为外部uid
	bs.PushUint8(a.Euid_u)
	bs.PushUint32(a.AddTime) //<uint32> 添加时间
	bs.PushUint8(a.AddTime_u)
	bs.PushString(a.Referrer) //<std::string> 推荐人
	bs.PushUint8(a.Referrer_u)
	bs.PushUint32(a.LastUpdateTime) //<uint32> 最后修改时间
	bs.PushUint8(a.LastUpdateTime_u)
	bs.PushString(a.Reserve) //<std::string> 预留字段
	bs.PushUint8(a.Reserve_u)
	if a.Version >= 20151110 {
		bs.PushString(a.WechatAccount) //<std::string> 微信账号
		bs.PushUint8(a.WechatAccount_u)
	}
	if a.Version >= 20151110 {
		bs.PushUint8(a.MemberLevel) //<uint8> 会员等级(1-银卡，2-金卡，3-铂金卡)
		bs.PushUint8(a.MemberLevel_u)
	}
	if a.Version >= 20151110 {
		var membercardList []IAoXXOO
		for _, m := range a.MembercardList {
			membercardList = append(membercardList, IAoXXOO(&m))
		}
		bs.PushObject(membercardList) //<std::vector<b2b2c::user::po::CMembercardInfoPo> > 会员卡列表(会员卡类型:外卡号,内卡号…)(卡类型：1-实体卡，2-联名卡，3-虚拟卡)(最多绑定10张)
		bs.PushUint8(a.MembercardList_u)
	}
	if a.Version >= 20151110 {
		bs.PushStringSet(a.PregnantPassportList) //<std::vector<std::string> > 好孕护照列表
		bs.PushUint8(a.PregnantPassportList_u)
	}
	if a.Version >= 20151110 {
		bs.PushString(a.Recruiter) //<std::string> 招募人(员工工号)
		bs.PushUint8(a.Recruiter_u)
	}
	if a.Version >= 20151110 {
		bs.PushString(a.Manager) //<std::string> 顾客关系经营人(原维护人)(员工工号)
		bs.PushUint8(a.Manager_u)
	}
	if a.Version >= 20151110 {
		bs.PushString(a.Creator) //<std::string> 建档人(员工工号)
		bs.PushUint8(a.Creator_u)
	}
	if a.Version >= 20151110 {
		bs.PushString(a.CreatorDepartment) //<std::string> 建档部门(门店编号)
		bs.PushUint8(a.CreatorDepartment_u)
	}
	if a.Version >= 20151110 {
		bs.PushUint8(a.RegisterSource) //<uint8> 注册来源
		bs.PushUint8(a.RegisterSource_u)
	}
	if a.Version >= 20151110 {
		bs.PushUint8(a.MemberSource) //<uint8> 会员来源
		bs.PushUint8(a.MemberSource_u)
	}
	if a.Version >= 20151110 {
		bs.PushUint8(a.IfSendErp) //<uint8> 是否同步给ERP
		bs.PushUint8(a.IfSendErp_u)
	}
	if a.Version >= 20151110 {
		bs.PushUint32(a.SendErpTime) //<uint32> 首次同步给ERP时间
		bs.PushUint8(a.SendErpTime_u)
	}
	if a.Version >= 20151231 {
		bs.PushMap(a.DiffChannelActiveTime) //<std::map<uint32,uint32> > 用户不同渠道激活时间, key为渠道，value为激活时间
		bs.PushUint8(a.DiffChannelActiveTime_u)
	}
	if a.Version >= 20160119 {
		bs.PushString(a.PromoteActive) //<std::string> 促活人
		bs.PushUint8(a.PromoteActive_u)
	}
	if a.Version >= 20160120 {
		a.BuyerProperty.Serialize(bs) //<b2b2c::user::po::CBuyerPropertyPo> 用户属性，参见E_USER_PROPERTY
		bs.PushUint8(a.BuyerProperty_u)
	}
	if a.Version >= 20160317 {
		bs.PushStringSet(a.UserLableList) //<std::set<std::string> > 用户标签列表
		bs.PushUint8(a.UserLableList_u)
	}
	if a.Version >= 20160317 {
		bs.PushString(a.UserLableRemarks) //<std::string> 用户标签备注
		bs.PushUint8(a.UserLableRemarks_u)
	}
	if a.Version >= 20160413 {
		bs.PushUint8(a.MobileStatus) //<uint8> 手机号状态
		bs.PushUint8(a.MobileStatus_u)
	}
	if a.Version >= 20160413 {
		bs.PushBitSet(a.MemberBitProperty) //<std::bitset<64> > 会员属性
		bs.PushUint8(a.MemberBitProperty_u)
	}
	if a.Version >= 20160516 {
		bs.PushUint8(a.PregnantPlan) //<uint8> 备孕计划
		bs.PushUint8(a.PregnantPlan_u)
	}
	if a.Version >= 20160517 {
		bs.PushStringSet(a.UserPictureLableList) //<std::vector<std::string> > 用户图片标签列表
		bs.PushUint8(a.UserPictureLableList_u)
	}
	return bs.IsGood()
}

func (a *AoUserInfoXXOO) UnSerialize(bs *ByteStream) error {
	classLen, err := bs.PopUint32()

	if err != nil {
		return err
	}
	startPop := bs.GetReadLength()

	a.Version, err = bs.PopUint32() //<uint32> 版本号

	if err != nil {
		return err
	}
	a.Uid, err = bs.PopUint64() //<uint64> uid, 用户id
	if err != nil {
		return errors.New("uid error " + err.Error())
	}

	a.Uid_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	a.AccountType, err = bs.PopUint8() //<uint8> 帐号类型，参见user_comm_define.h中的E_ACCOUNT_TYPE
	if err != nil {
		return err
	}
	a.AccountType_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	a.Mobile, err = bs.PopString() //<std::string> 手机号
	if err != nil {
		return err
	}
	a.Mobile_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	a.Email, err = bs.PopString() //<std::string> 电子邮箱
	if err != nil {
		return err
	}
	a.Email_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	a.LoginAccount, err = bs.PopString() //<std::string> 个性化账号
	if err != nil {
		return err
	}
	a.LoginAccount_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	a.WechatOpenid, err = bs.PopMap(reflect.TypeOf(""), reflect.TypeOf("")) //<std::map<std::string,std::string> > 微信OpenID：appid-openid

	if err != nil {
		return err
	}
	a.WechatOpenid_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	a.WechatUnionid, err = bs.PopString() //<std::string> 微信unionID
	if err != nil {
		return err
	}
	a.WechatUnionid_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	a.UserType, err = bs.PopUint8() //<uint8> 用户类型，参见user_comm_define.h中的E_USER_TYPE
	if err != nil {
		return err
	}
	a.UserType_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	a.BitProperty, err = bs.PopBitSet() //<std::bitset<64> > 用户属性位BitSet，具体意义参见b2b2c_define.h中的E_USER_PROPERTY
	if err != nil {
		return err
	}
	a.BitProperty_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	a.LoginPasswd, err = bs.PopString() //<std::string> 用户登录密码
	if err != nil {
		return err
	}
	a.LoginPasswd_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	a.TradePasswd, err = bs.PopString() //<std::string> 用户交易密码
	if err != nil {
		return err
	}
	a.TradePasswd_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	a.PasswordSecureLevel, err = bs.PopUint8() //<uint8> 登录密码安全级别
	if err != nil {
		return err
	}
	a.PasswordSecureLevel_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	a.Skey, err = bs.PopString() //<std::string> 用户sessionkey
	if err != nil {
		return err
	}
	a.Skey_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	a.SkeyInvalidTime, err = bs.PopUint32() //<uint32> 用户sessionkey失效时间
	if err != nil {
		return err
	}
	a.SkeyInvalidTime_u, err = bs.PopUint8()
	if err != nil {
		return err
	}

	a.DiffSrcRegTime, err = bs.PopMap(reflect.TypeOf(uint32(0)), reflect.TypeOf(uint32(0))) //<std::map<uint32,uint32> > 不同来源用户注册时间, key为来源，value为注册时间
	if err != nil {
		return err
	}
	a.DiffSrcRegTime_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	a.Rating, err = bs.PopUint8() //<uint8> 用户评级
	if err != nil {
		return err
	}
	a.Rating_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	var sizeBaby uint32
	sizeBaby, err = bs.PopUint32()
	if err != nil {
		return err
	}

	for i := uint32(0); i < sizeBaby; i++ {
		babyInfo := NewAoBabyInfoXXOO()
		err = babyInfo.UnSerialize(bs)
		if err != nil {
			return err
		}
		a.BabyInfo = append(a.BabyInfo, *babyInfo)
	}

	a.BabyInfo_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	
	a.RelationWithBaby, err = bs.PopUint8() //<uint8> 与宝宝关系，参见user_comm_define.h中的E_RELATION_WITH_BABY
	if err != nil {
		return err
	}
	a.RelationWithBaby_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	a.Nickname, err = bs.PopString() //<std::string> 昵称
	if err != nil {
		return err
	}
	a.Nickname_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	a.Truename, err = bs.PopString() //<std::string> 真实姓名
	if err != nil {
		return err
	}
	a.Truename_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	a.Sex, err = bs.PopUint8() //<uint8> 性别：0-未知，1-女，2-男。参见E_USER_SEX
	if err != nil {
		return err
	}
	a.Sex_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	a.Birthday, err = bs.PopUint64() //<int64> 生日
	if err != nil {
		return errors.New("Birthday error" + err.Error())
	}
	a.Birthday_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	a.QQNumber, err = bs.PopUint64() //<uint64> QQ号
	if err != nil {
		return errors.New("QQNumber error" + err.Error())
	}
	a.QQNumber_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	a.Phone, err = bs.PopString() //<std::string> 固定电话
	if err != nil {
		return err
	}
	a.Phone_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	a.Fax, err = bs.PopString() //<std::string> 传真
	if err != nil {
		return err
	}
	a.Fax_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	a.Region, err = bs.PopString() //<std::string> 所在地区id
	if err != nil {
		return err
	}
	a.Region_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	a.Address, err = bs.PopString() //<std::string> 详细住址
	if err != nil {
		return err
	}
	a.Address_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	a.Community, err = bs.PopString() //<std::string> 小区
	if err != nil {
		return err
	}
	a.Community_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	a.Postcode, err = bs.PopString() //<std::string> 邮政编码
	if err != nil {
		return err
	}
	a.Postcode_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	a.IdentityType, err = bs.PopUint8() //<uint8> 身份证件类型，参见user_comm_define.h中E_USER_IDTYPE
	if err != nil {
		return err
	}
	a.IdentityType_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	a.IdentityNum, err = bs.PopString() //<std::string> 身份证件号码
	if err != nil {
		return err
	}
	a.IdentityNum_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	a.Photo, err = bs.PopString() //<std::string> 头像
	if err != nil {
		return err
	}
	a.Photo_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	a.Signature, err = bs.PopString() //<std::string> 签名
	if err != nil {
		return err
	}
	a.Signature_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	a.Euid, err = bs.PopMap(reflect.TypeOf(uint32(0)), reflect.TypeOf("")) //<std::map<uint32,std::string> > 外部导入用户euid, key为来源渠道（参见b2b2c_define.h中的E_USER_EUID_SOURCE）, value为外部uid
	if err != nil {
		return err
	}
	a.Euid_u, err = bs.PopUint8()
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
	a.Referrer, err = bs.PopString() //<std::string> 推荐人，招募人
	if err != nil {
		return err
	}
	a.Referrer_u, err = bs.PopUint8()
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
	a.Reserve, err = bs.PopString() //<std::string> 预留字段
	if err != nil {
		return err
	}
	a.Reserve_u, err = bs.PopUint8()

	if err != nil {
		return err
	}
	a.Compat(bs, classLen, startPop)
	return nil
}

func NewAoUserInfoXXOO() *AoUserInfoXXOO {
	obj := &AoUserInfoXXOO{}
	return obj
}

type AoBabyInfoXXOO struct {
	AoXXOO
	Bid           uint64 //<uint64> bid, 宝宝id(版本>=0)
	Bid_u         uint8  // uint8 bid的flag位
	Uid           uint64 //<uint64> uid, 用户id(版本>=0)
	Uid_u         uint8  // uint8 uid的flag位
	Nickname      string //<std::string> 昵称(版本>=0)
	Nickname_u    uint8  // uint8 nickname的flag位
	Truename      string //<std::string> 真实姓名(版本>=0)
	Truename_u    uint8  // uint8 truename的flag位
	Sex           uint8  //<uint8> 性别：0-未知，1-女，2-男。参见E_USER_SEX(版本>=0)
	Sex_u         uint8  // uint8 sex的flag位
	Birthday      uint64 //<int64> 生日/预产期(版本>=0)
	Birthday_u    uint8  // uint8 birthday的flag位
	BitProperty   []uint8 //<std::bitset<64> > 宝宝属性位BitSet，具体意义参见user_comm_define.h中的E_USER_PROPERTY(版本>=0)
	BitProperty_u uint8  // uint8 bitProperty的flag位
}

func NewAoBabyInfoXXOO() *AoBabyInfoXXOO {
	obj := &AoBabyInfoXXOO{}

	return obj
}

func (ao *AoBabyInfoXXOO) Serialize(bs *ByteStream) bool {

	bs.PushUint32(uint32(ao.getClassLen()))

	return ao.serialize_internal(bs)
}

func (ao *AoBabyInfoXXOO) serialize_internal(bs *ByteStream) bool {
	bs.PushUint32(ao.Version)
	bs.PushUint64(ao.Bid)
	bs.PushUint8(ao.Bid_u)
	bs.PushUint64(ao.Uid)
	bs.PushUint8(ao.Uid_u)
	bs.PushString(ao.Nickname)
	bs.PushUint8(ao.Nickname_u)
	bs.PushString(ao.Truename)
	bs.PushUint8(ao.Truename_u)
	bs.PushUint8(ao.Sex)
	bs.PushUint8(ao.Sex_u)
	bs.PushUint64(ao.Birthday)
	bs.PushUint8(ao.Birthday_u)
	bs.PushBitSet(ao.BitProperty)
	bs.PushUint8(ao.BitProperty_u)
	bs.PushUint32(ao.AddTime)
	bs.PushUint8(ao.AddTime_u)
	bs.PushUint32(ao.LastUpdateTime)
	bs.PushUint8(ao.LastUpdateTime_u)
	bs.PushString(ao.Reserve)
	bs.PushUint8(ao.Reserve_u)

	return bs.IsGood()
}

func (ao *AoBabyInfoXXOO) getClassLen() int {

	bsLen := NewByteStream()
	bsLen.SetRealWrite(false)
	ao.serialize_internal(bsLen)
	length := bsLen.GetWrittenLength()

	return length
}

func (a *AoBabyInfoXXOO) UnSerialize(bs *ByteStream) error {
	classLen, err := bs.PopUint32()
	if err != nil {
		return err
	}
	startPop := bs.GetReadLength()

	a.Version, err = bs.PopUint32() //<uint32> 版本号
	if err != nil {
		return err
	}
	a.Bid, err = bs.PopUint64()
	if err != nil {
		return errors.New("Bid error " + err.Error())
	}
	a.Bid_u, err = bs.PopUint8() // uint8 bid的flag位
	if err != nil {
		return err
	}
	a.Uid, err = bs.PopUint64() //<uint64> uid, 用户id(版本>=0)
	if err != nil {
		return errors.New("Uid error " + err.Error())
	}
	a.Uid_u, err = bs.PopUint8() // uint8 uid的flag位
	if err != nil {
		return err
	}
	a.Nickname, err = bs.PopString() //<std::string> 昵称(版本>=0)
	if err != nil {
		return err
	}
	a.Nickname_u, err = bs.PopUint8() // uint8 nickname的flag位
	if err != nil {
		return err
	}
	a.Truename, err = bs.PopString() //<std::string> 真实姓名(版本>=0)
	if err != nil {
		return err
	}
	a.Truename_u, err = bs.PopUint8() // uint8 truename的flag位
	if err != nil {
		return err
	}
	a.Sex, err = bs.PopUint8() //<uint8> 性别：0-未知，1-女，2-男。参见E_USER_SEX(版本>=0)
	if err != nil {
		return err
	}
	a.Sex_u, err = bs.PopUint8() // uint8 sex的flag位
	if err != nil {
		return err
	}
	a.Birthday, err = bs.PopUint64() //<int64> 生日/预产期(版本>=0)
	if err != nil {
		return err
	}
	a.Birthday_u, err = bs.PopUint8() // uint8 birthday的flag位
	if err != nil {
		return err
	}
	a.BitProperty, err = bs.PopBitSet() //<std::bitset<64> > 宝宝属性位BitSet，具体意义参见user_comm_define.h中的E_USER_PROPERTY(版本>=0)
	if err != nil {
		return errors.New("BitProperty error " + err.Error())
	}
	a.BitProperty_u, err = bs.PopUint8() // uint8 bitProperty的flag位
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
	a.Reserve, err = bs.PopString() //<std::string> 预留字段
	if err != nil {
		return err
	}
	a.Reserve_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	a.Compat(bs, classLen, startPop)
	return nil
}

type AoBuyerPropertyXXOO struct {
	AoXXOO
	IsPaidMember                    bool
	IsPaidMember_u                  uint8
	IsHonorableMember               bool
	IsHonorableMember_u             uint8
	IsPreparePregnant               bool
	IsPreparePregnant_u             uint8
	IsSHTMember                     bool
	IsSHTMember_u                   uint8
	IsSubscribePromotionalMessage   bool
	IsSubscribePromotionalMessage_u uint8
	IsSubscribePromotionalPhone     bool
	IsSubscribePromotionalPhone_u   uint8
	IsSubscribePromotionalWechat    bool
	IsSubscribePromotionalWechat_u  uint8
	IsAcceptVisitMessage            bool
	IsAcceptVisitMessage_u          uint8
	IsAcceptVisitPhone              bool
	IsAcceptVisitPhone_u            uint8
	IsSetOnlineActiveTime           bool
	IsSetOnlineActiveTime_u         uint8
	IsSetOfflineActiveTime          bool
	IsSetOfflineActiveTime_u        uint8
	IsBindMobile                    bool
	IsBindMobile_u                  uint8
	IsBindEmail                     bool
	IsBindEmail_u                   uint8
	IsBindLoginAccount              bool
	IsBindLoginAccount_u            uint8
	IsBindWechat                    bool
	IsBindWechat_u                  uint8
	IsEntityMembercard              bool
	IsEntityMembercard_u            uint8
	IsBankCobrandedMembercard       bool
	IsBankCobrandedMembercard_u     uint8
	IsFictitiousMembercard          bool
	IsFictitiousMembercard_u        uint8
	IsSourceTMApp                   bool
	IsSourceTMApp_u                 uint8
	IsSourceWechat                  bool
	IsSourceWechat_u                uint8
	IsSourceCRM                     bool
	IsSourceCRM_u                   uint8
	IsSourceERP                     bool
	IsSourceERP_u                   uint8
	IsHonorableCounselor            bool
	IsHonorableCounselor_u          uint8
}

func NewAoBuyerPropertyXXOOXXOO() *AoBuyerPropertyXXOO {
	obj := &AoBuyerPropertyXXOO{}
	obj.Version = uint32(20160120)
	return obj
}

func (ao *AoBuyerPropertyXXOO) Serialize(bs *ByteStream) bool {

	bs.PushUint32(uint32(ao.getClassLen()))

	return ao.serialize_internal(bs)
}

func (ao *AoBuyerPropertyXXOO) serialize_internal(bs *ByteStream) bool {
	bs.PushUint32(ao.Version)
	bs.PushBool(ao.IsPaidMember)
	bs.PushUint8(ao.IsPaidMember_u)
	bs.PushBool(ao.IsHonorableMember)
	bs.PushUint8(ao.IsHonorableMember_u)
	bs.PushBool(ao.IsPreparePregnant)
	bs.PushUint8(ao.IsPreparePregnant_u)
	bs.PushBool(ao.IsSHTMember)
	bs.PushUint8(ao.IsSHTMember_u)
	bs.PushBool(ao.IsSubscribePromotionalMessage)
	bs.PushUint8(ao.IsSubscribePromotionalMessage_u)
	bs.PushBool(ao.IsSubscribePromotionalPhone)
	bs.PushUint8(ao.IsSubscribePromotionalPhone_u)
	bs.PushBool(ao.IsSubscribePromotionalWechat)
	bs.PushUint8(ao.IsSubscribePromotionalWechat_u)
	bs.PushBool(ao.IsAcceptVisitMessage)
	bs.PushUint8(ao.IsAcceptVisitMessage_u)
	bs.PushBool(ao.IsAcceptVisitPhone)
	bs.PushUint8(ao.IsAcceptVisitPhone_u)
	bs.PushString(ao.Reserve)
	bs.PushUint8(ao.Reserve_u)
	if ao.Version >= uint32(20151231) {
		bs.PushBool(ao.IsSetOnlineActiveTime)
		bs.PushUint8(ao.IsSetOnlineActiveTime_u)
		bs.PushBool(ao.IsSetOfflineActiveTime)
		bs.PushUint8(ao.IsSetOfflineActiveTime_u)
	}
	if ao.Version >= uint32(20160120) {
		bs.PushBool(ao.IsBindMobile)
		bs.PushUint8(ao.IsBindMobile_u)
		bs.PushBool(ao.IsBindEmail)
		bs.PushUint8(ao.IsBindEmail_u)
		bs.PushBool(ao.IsBindLoginAccount)
		bs.PushUint8(ao.IsBindLoginAccount_u)
		bs.PushBool(ao.IsBindWechat)
		bs.PushUint8(ao.IsBindWechat_u)
		bs.PushBool(ao.IsEntityMembercard)
		bs.PushUint8(ao.IsEntityMembercard_u)
		bs.PushBool(ao.IsBankCobrandedMembercard)
		bs.PushUint8(ao.IsBankCobrandedMembercard_u)
		bs.PushBool(ao.IsFictitiousMembercard)
		bs.PushUint8(ao.IsFictitiousMembercard_u)
		bs.PushBool(ao.IsSourceTMApp)
		bs.PushUint8(ao.IsSourceTMApp_u)
		bs.PushBool(ao.IsSourceWechat)
		bs.PushUint8(ao.IsSourceWechat_u)
		bs.PushBool(ao.IsSourceCRM)
		bs.PushUint8(ao.IsSourceCRM_u)
		bs.PushBool(ao.IsSourceERP)
		bs.PushUint8(ao.IsSourceERP_u)
		bs.PushBool(ao.IsHonorableCounselor)
		bs.PushUint8(ao.IsHonorableCounselor_u)
	}

	return bs.IsGood()
}

func (ao *AoBuyerPropertyXXOO) getClassLen() int {

	bsLen := NewByteStream()
	bsLen.SetRealWrite(false)
	ao.serialize_internal(bsLen)
	length := bsLen.GetWrittenLength()

	return length
}

func (a *AoBuyerPropertyXXOO) UnSerialize(bs *ByteStream) error {
	classLen, err := bs.PopUint32()
	if err != nil {
		return nil
	}
	startPop := bs.GetReadLength()

	a.Version, err = bs.PopUint32() //<uint32> 版本号
	if err != nil {
		return nil
	}
	a.IsPaidMember, err = bs.PopBool()
	if err != nil {
		return nil
	}
	a.IsPaidMember_u, err = bs.PopUint8()
	if err != nil {
		return nil
	}
	a.IsHonorableMember, err = bs.PopBool()
	if err != nil {
		return nil
	}
	a.IsHonorableMember_u, err = bs.PopUint8()
	if err != nil {
		return nil
	}
	a.IsPreparePregnant, err = bs.PopBool()
	if err != nil {
		return nil
	}
	a.IsPreparePregnant_u, err = bs.PopUint8()
	if err != nil {
		return nil
	}
	a.IsSHTMember, err = bs.PopBool()
	if err != nil {
		return nil
	}
	a.IsSHTMember_u, err = bs.PopUint8()
	if err != nil {
		return nil
	}
	a.IsSubscribePromotionalMessage, err = bs.PopBool()
	if err != nil {
		return nil
	}
	a.IsSubscribePromotionalMessage_u, err = bs.PopUint8()
	if err != nil {
		return nil
	}
	a.IsSubscribePromotionalPhone, err = bs.PopBool()
	if err != nil {
		return nil
	}
	a.IsSubscribePromotionalPhone_u, err = bs.PopUint8()
	if err != nil {
		return nil
	}
	a.IsSubscribePromotionalWechat, err = bs.PopBool()
	if err != nil {
		return nil
	}
	a.IsSubscribePromotionalWechat_u, err = bs.PopUint8()
	if err != nil {
		return nil
	}
	a.IsAcceptVisitMessage, err = bs.PopBool()
	if err != nil {
		return nil
	}
	a.IsAcceptVisitMessage_u, err = bs.PopUint8()
	if err != nil {
		return nil
	}
	a.IsAcceptVisitPhone, err = bs.PopBool()
	if err != nil {
		return nil
	}
	a.IsAcceptVisitPhone_u, err = bs.PopUint8()
	if err != nil {
		return nil
	}
	a.Reserve, err = bs.PopString() //<std::string> 预留字段
	if err != nil {
		return nil
	}
	a.Reserve_u, err = bs.PopUint8()
	if err != nil {
		return nil
	}
	if a.Version >= uint32(20151231) {
		a.IsSetOnlineActiveTime, err = bs.PopBool()
		if err != nil {
			return nil
		}
		a.IsSetOnlineActiveTime_u, err = bs.PopUint8()
		if err != nil {
			return nil
		}
		a.IsSetOfflineActiveTime, err = bs.PopBool()
		if err != nil {
			return nil
		}
		a.IsSetOfflineActiveTime_u, err = bs.PopUint8()
		if err != nil {
			return nil
		}
	}
	if a.Version >= uint32(20160120) {
		a.IsBindMobile, err = bs.PopBool()
		if err != nil {
			return nil
		}
		a.IsBindMobile_u, err = bs.PopUint8()
		if err != nil {
			return nil
		}
		a.IsBindEmail, err = bs.PopBool()
		if err != nil {
			return nil
		}
		a.IsBindEmail_u, err = bs.PopUint8()
		if err != nil {
			return nil
		}
		a.IsBindLoginAccount, err = bs.PopBool()
		if err != nil {
			return nil
		}
		a.IsBindLoginAccount_u, err = bs.PopUint8()
		if err != nil {
			return nil
		}
		a.IsBindWechat, err = bs.PopBool()
		if err != nil {
			return nil
		}
		a.IsBindWechat_u, err = bs.PopUint8()
		if err != nil {
			return nil
		}
		a.IsEntityMembercard, err = bs.PopBool()
		if err != nil {
			return nil
		}
		a.IsEntityMembercard_u, err = bs.PopUint8()
		if err != nil {
			return nil
		}
		a.IsBankCobrandedMembercard, err = bs.PopBool()
		if err != nil {
			return nil
		}
		a.IsBankCobrandedMembercard_u, err = bs.PopUint8()
		if err != nil {
			return nil
		}
		a.IsFictitiousMembercard, err = bs.PopBool()
		if err != nil {
			return nil
		}
		a.IsFictitiousMembercard_u, err = bs.PopUint8()
		if err != nil {
			return nil
		}
		a.IsSourceTMApp, err = bs.PopBool()
		if err != nil {
			return nil
		}
		a.IsSourceTMApp_u, err = bs.PopUint8()
		if err != nil {
			return nil
		}
		a.IsSourceWechat, err = bs.PopBool()
		if err != nil {
			return nil
		}
		a.IsSourceWechat_u, err = bs.PopUint8()
		if err != nil {
			return nil
		}
		a.IsSourceCRM, err = bs.PopBool()
		if err != nil {
			return nil
		}
		a.IsSourceCRM_u, err = bs.PopUint8()
		if err != nil {
			return nil
		}
		a.IsSourceERP, err = bs.PopBool()
		if err != nil {
			return nil
		}
		a.IsSourceERP_u, err = bs.PopUint8()
		if err != nil {
			return nil
		}
		a.IsHonorableCounselor, err = bs.PopBool()
		if err != nil {
			return nil
		}
		a.IsHonorableCounselor_u, err = bs.PopUint8()
		if err != nil {
			return nil
		}
	}
	a.Compat(bs, classLen, startPop)
	return nil
}

type AoMemberCardInfoXXOO struct {
	AoXXOO
	Type               uint8
	Type_u             uint8
	MemberCardNumPub   string
	MemberCardNumPub_u uint8
	MemberCardNumSec   string
	MemberCardNumSec_u uint8
	CardStatus         uint8
	CardStatus_u       uint8
}

func NewAoBuyerPropertyPoXXOO() *AoMemberCardInfoXXOO {
	obj := &AoMemberCardInfoXXOO{}
	obj.Version = uint32(20160120)
	return obj
}

func (ao *AoMemberCardInfoXXOO) Serialize(bs *ByteStream) bool {

	bs.PushUint32(uint32(ao.getClassLen()))

	return ao.serialize_internal(bs)
}

func (ao *AoMemberCardInfoXXOO) serialize_internal(bs *ByteStream) bool {
	bs.PushUint32(ao.Version)
	bs.PushUint8(ao.Type)
	bs.PushUint8(ao.Type_u)
	bs.PushString(ao.MemberCardNumPub)
	bs.PushUint8(ao.MemberCardNumPub_u)
	bs.PushString(ao.MemberCardNumSec)
	bs.PushUint8(ao.MemberCardNumSec_u)
	bs.PushString(ao.Reserve)
	bs.PushUint8(ao.Reserve_u)
	bs.PushUint8(ao.CardStatus)
	bs.PushUint8(ao.CardStatus_u)

	return bs.IsGood()
}

func (ao *AoMemberCardInfoXXOO) getClassLen() int {

	bsLen := NewByteStream()
	bsLen.SetRealWrite(false)
	ao.serialize_internal(bsLen)
	length := bsLen.GetWrittenLength()

	return length
}

func (a *AoMemberCardInfoXXOO) UnSerialize(bs *ByteStream) error {
	classLen, err := bs.PopUint32()
	if err != nil {
		return nil
	}
	startPop := bs.GetReadLength()

	a.Version, err = bs.PopUint32() //<uint32> 版本号
	if err != nil {
		return nil
	}
	a.Type, err = bs.PopUint8()
	if err != nil {
		return nil
	}
	a.Type_u, err = bs.PopUint8()
	if err != nil {
		return nil
	}
	a.MemberCardNumPub, err = bs.PopString()
	if err != nil {
		return nil
	}
	a.MemberCardNumPub_u, err = bs.PopUint8()
	if err != nil {
		return nil
	}
	a.MemberCardNumSec, err = bs.PopString()
	if err != nil {
		return nil
	}
	a.MemberCardNumSec_u, err = bs.PopUint8()
	if err != nil {
		return nil
	}
	a.Reserve, err = bs.PopString() //<std::string> 预留字段
	if err != nil {
		return nil
	}
	a.Reserve_u, err = bs.PopUint8()
	if err != nil {
		return nil
	}
	if a.Version >= uint32(20160120) {
		a.CardStatus, err = bs.PopUint8()
		if err != nil {
			return nil
		}
		a.CardStatus_u, err = bs.PopUint8()
		if err != nil {
			return nil
		}
	}
	a.Compat(bs, classLen, startPop)
	return nil
}
