package goao

import (
	"reflect"
)

type AoUserInfoXXOO struct {
	AoXXOO
	Uid                   uint64 //<uint64_t> uid, 用户id
	Uid_u                 uint8
	AccountType           uint8 //<uint8_t> 帐号类型，参见user_comm_define.h中的E_ACCOUNT_TYPE
	AccountType_u         uint8
	Mobile                string //<std::string> 手机号
	Mobile_u              uint8
	Email                 string //<std::string> 电子邮箱
	Email_u               uint8
	LoginAccount          string //<std::string> 个性化账号
	LoginAccount_u        uint8
	WechatOpenid          map[interface{}]interface{} //<std::map<std::string,std::string> > 微信OpenID：appid-openid
	WechatOpenid_u        uint8
	WechatUnionid         string //<std::string> 微信unionID
	WechatUnionid_u       uint8
	UserType              uint8 //<uint8_t> 用户类型，参见user_comm_define.h中的E_USER_TYPE
	UserType_u            uint8
	BitProperty           []uint8 //<std::bitset<64> > 用户属性位BitSet，具体意义参见b2b2c_define.h中的E_USER_PROPERTY
	BitProperty_u         uint8
	LoginPasswd           string //<std::string> 用户登录密码
	LoginPasswd_u         uint8
	TradePasswd           string //<std::string> 用户交易密码
	TradePasswd_u         uint8
	PasswordSecureLevel   uint8 //<uint8_t> 登录密码安全级别
	PasswordSecureLevel_u uint8
	Skey                  string //<std::string> 用户sessionkey
	Skey_u                uint8
	SkeyInvalidTime       uint32 //<uint32_t> 用户sessionkey失效时间
	SkeyInvalidTime_u     uint8
	DiffSrcRegTime        map[interface{}]interface{} //<std::map<uint32_t,uint32_t> > 不同来源用户注册时间, key为来源，value为注册时间
	DiffSrcRegTime_u      uint8
	Rating                uint8 //<uint8_t> 用户评级
	Rating_u              uint8
	BabyInfo              []AoBabyInfoXXOO //<std::vector<b2b2c::user::po::CBabyInfoPo> > 宝宝列表
	BabyInfo_u            uint8
	RelationWithBaby      uint8 //<uint8_t> 与宝宝关系，参见user_comm_define.h中的E_RELATION_WITH_BABY
	RelationWithBaby_u    uint8
	Nickname              string //<std::string> 昵称
	Nickname_u            uint8
	Truename              string //<std::string> 真实姓名
	Truename_u            uint8
	Sex                   uint8 //<uint8_t> 性别：0-未知，1-女，2-男。参见E_USER_SEX
	Sex_u                 uint8
	Birthday              int64 //<int64_t> 生日
	Birthday_u            uint8
	QQNumber              uint64 //<uint64_t> QQ号
	QQNumber_u            uint8
	Phone                 string //<std::string> 固定电话
	Phone_u               uint8
	Fax                   string //<std::string> 传真
	Fax_u                 uint8
	Region                string //<std::string> 所在地区id
	Region_u              uint8
	Address               string //<std::string> 详细住址
	Address_u             uint8
	Community             string //<std::string> 小区
	Community_u           uint8
	Postcode              string //<std::string> 邮政编码
	Postcode_u            uint8
	IdentityType          uint8 //<uint8_t> 身份证件类型，参见user_comm_define.h中E_USER_IDTYPE
	IdentityType_u        uint8
	IdentityNum           string //<std::string> 身份证件号码
	IdentityNum_u         uint8
	Photo                 string //<std::string> 头像
	Photo_u               uint8
	Signature             string //<std::string> 签名
	Signature_u           uint8
	Euid                  map[interface{}]interface{} //<std::map<uint32_t,std::string> > 外部导入用户euid, key为来源渠道（参见b2b2c_define.h中的E_USER_EUID_SOURCE）, value为外部uid
	Euid_u                uint8

	Referrer   string //<std::string> 推荐人，招募人
	Referrer_u uint8
}

func (a *AoUserInfoXXOO) UnSerialize(bs *ByteStream) error {
	classLen, err := bs.PopUint32()

	if err != nil {
		return err
	}
	startPop := bs.GetReadLength()

	a.Version, err = bs.PopUint32() //<uint32_t> 版本号

	if err != nil {
		return err
	}
	a.Uid, err = bs.PopUint64() //<uint64_t> uid, 用户id
	if err != nil {
		return err
	}

	a.Uid_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	a.AccountType, err = bs.PopUint8() //<uint8_t> 帐号类型，参见user_comm_define.h中的E_ACCOUNT_TYPE
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
	a.UserType, err = bs.PopUint8() //<uint8_t> 用户类型，参见user_comm_define.h中的E_USER_TYPE
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
	a.PasswordSecureLevel, err = bs.PopUint8() //<uint8_t> 登录密码安全级别
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
	a.SkeyInvalidTime, err = bs.PopUint32() //<uint32_t> 用户sessionkey失效时间
	if err != nil {
		return err
	}
	a.SkeyInvalidTime_u, err = bs.PopUint8()
	if err != nil {
		return err
	}

	a.DiffSrcRegTime, err = bs.PopMap(reflect.TypeOf(uint32(0)), reflect.TypeOf(uint32(0))) //<std::map<uint32_t,uint32_t> > 不同来源用户注册时间, key为来源，value为注册时间
	if err != nil {
		return err
	}
	a.DiffSrcRegTime_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	a.Rating, err = bs.PopUint8() //<uint8_t> 用户评级
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
	a.RelationWithBaby, err = bs.PopUint8() //<uint8_t> 与宝宝关系，参见user_comm_define.h中的E_RELATION_WITH_BABY
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
	a.Sex, err = bs.PopUint8() //<uint8_t> 性别：0-未知，1-女，2-男。参见E_USER_SEX
	if err != nil {
		return err
	}
	a.Sex_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	a.Birthday, err = bs.PopInt64() //<int64_t> 生日
	if err != nil {
		return err
	}
	a.Birthday_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	a.QQNumber, err = bs.PopUint64() //<uint64_t> QQ号
	if err != nil {
		return err
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
	a.IdentityType, err = bs.PopUint8() //<uint8_t> 身份证件类型，参见user_comm_define.h中E_USER_IDTYPE
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
	a.Euid, err = bs.PopMap(reflect.TypeOf(uint32(0)), reflect.TypeOf("")) //<std::map<uint32_t,std::string> > 外部导入用户euid, key为来源渠道（参见b2b2c_define.h中的E_USER_EUID_SOURCE）, value为外部uid
	if err != nil {
		return err
	}
	a.Euid_u, err = bs.PopUint8()
	if err != nil {
		return err
	}
	a.AddTime, err = bs.PopUint32() //<uint32_t> 添加时间
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
	a.LastUpdateTime, err = bs.PopUint32() //<uint32_t> 最后修改时间
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
	Bid           uint64 //<uint64_t> bid, 宝宝id(版本>=0)
	Bid_u         uint8  // uint8_t bid的flag位
	Uid           uint64 //<uint64_t> uid, 用户id(版本>=0)
	Uid_u         uint8  // uint8_t uid的flag位
	Nickname      string //<std::string> 昵称(版本>=0)
	Nickname_u    uint8  // uint8_t nickname的flag位
	Truename      string //<std::string> 真实姓名(版本>=0)
	Truename_u    uint8  // uint8_t truename的flag位
	Sex           uint8  //<uint8_t> 性别：0-未知，1-女，2-男。参见E_USER_SEX(版本>=0)
	Sex_u         uint8  // uint8_t sex的flag位
	Birthday      int64  //<int64_t> 生日/预产期(版本>=0)
	Birthday_u    uint8  // uint8_t birthday的flag位
	BitProperty   uint64 //<std::bitset<64> > 宝宝属性位BitSet，具体意义参见user_comm_define.h中的E_USER_PROPERTY(版本>=0)
	BitProperty_u uint8  // uint8_t bitProperty的flag位
}

func NewAoBabyInfoXXOO() *AoBabyInfoXXOO {
	obj := &AoBabyInfoXXOO{}

	return obj
}

func (a *AoBabyInfoXXOO) UnSerialize(bs *ByteStream) error {
	classLen, err := bs.PopUint32()
	if err != nil {
		return nil
	}
	startPop := bs.GetReadLength()

	a.Version, err = bs.PopUint32() //<uint32_t> 版本号
	if err != nil {
		return nil
	}
	a.Bid, err = bs.PopUint64()
	if err != nil {
		return nil
	}
	a.Bid_u, err = bs.PopUint8() // uint8_t bid的flag位
	if err != nil {
		return nil
	}
	a.Uid, err = bs.PopUint64() //<uint64_t> uid, 用户id(版本>=0)
	if err != nil {
		return nil
	}
	a.Uid_u, err = bs.PopUint8() // uint8_t uid的flag位
	if err != nil {
		return nil
	}
	a.Nickname, err = bs.PopString() //<std::string> 昵称(版本>=0)
	if err != nil {
		return nil
	}
	a.Nickname_u, err = bs.PopUint8() // uint8_t nickname的flag位
	if err != nil {
		return nil
	}
	a.Truename, err = bs.PopString() //<std::string> 真实姓名(版本>=0)
	if err != nil {
		return nil
	}
	a.Truename_u, err = bs.PopUint8() // uint8_t truename的flag位
	if err != nil {
		return nil
	}
	a.Sex, err = bs.PopUint8() //<uint8_t> 性别：0-未知，1-女，2-男。参见E_USER_SEX(版本>=0)
	if err != nil {
		return nil
	}
	a.Sex_u, err = bs.PopUint8() // uint8_t sex的flag位
	if err != nil {
		return nil
	}
	a.Birthday, err = bs.PopInt64() //<int64_t> 生日/预产期(版本>=0)
	if err != nil {
		return nil
	}
	a.Birthday_u, err = bs.PopUint8() // uint8_t birthday的flag位
	if err != nil {
		return nil
	}
	a.BitProperty, err = bs.PopUint64() //<std::bitset<64> > 宝宝属性位BitSet，具体意义参见user_comm_define.h中的E_USER_PROPERTY(版本>=0)
	if err != nil {
		return nil
	}
	a.BitProperty_u, err = bs.PopUint8() // uint8_t bitProperty的flag位
	if err != nil {
		return nil
	}
	a.AddTime, err = bs.PopUint32() //<uint32_t> 添加时间
	if err != nil {
		return nil
	}
	a.AddTime_u, err = bs.PopUint8()
	if err != nil {
		return nil
	}
	a.LastUpdateTime, err = bs.PopUint32() //<uint32_t> 最后修改时间
	if err != nil {
		return nil
	}
	a.LastUpdateTime_u, err = bs.PopUint8()
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
	a.Compat(bs, classLen, startPop)
	return nil
}
