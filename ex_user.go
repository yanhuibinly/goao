package goao

import (
	"errors"
	"fmt"
	"strconv"
)

type ModelUser struct {
	Uid              int64       `json:"uid"`
	AccountType      uint8       `json:"accountType"`
	Mobile           string      `json:"mobile"`
	Email            string      `json:"Email"`
	LoginAccount     string      `json:"loginAccount"`
	WechatOpenid     string      `json:"wechatOpenid"`
	WechatUnionid    string      `json:"wechatUnionid"`
	UserType         uint8       `json:"userType"`
	DiffSrcRegTime   string      `json:"diffSrcRegTime"`
	Rating           uint8       `json:"rating"`
	BabyInfo         interface{} `json:"babyInfo"`
	RelationWithBaby uint8       `json:"relationWithBaby"`
	Nickname         string      `json:"nickname"`
	Truename         string      `json:"truename"`
	Sex              uint8       `json:"sex"`
	Birthday         int64       `json:"birthday"`
	QQNumber         uint64      `json:"qQNumber"`
	Phone            string      `json:"phone"`
	Fax              string      `json:"fax"`
	Region           string      `json:"region"`
	Address          string      `json:"address"`
	Community        string      `json:"community"`
	Postcode         string      `json:"postcode"`
	IdentityType     uint8       `json:"identityType"`
	IdentityNum      string      `json:"identityNum"`
	Photo            string      `json:"photo"`
	Signature        string      `json:"signature"`
	UserTypeName     string      `json:"userTypeName"`
	Full_photo       string      `json:"full_photo"`
	Short_address    string      `json:"short_address"`
}

func ExUserGetUserInfos(uids []int64, host string, machineKey string, source string) (map[int64]ModelUser, error) {

	var req = NewRequest()

	req.Host = host

	req.MachineKey = machineKey

	req.SceneId = 1

	req.Source = source

	goClient := New()

	aoReq := NewAoBatchGetUserInfoByUidReq()

	//aoReq.Uids = uids_uint64
	var uids_uint64 []uint64

	for _, value := range uids {
		uids_uint64 = append(uids_uint64, uint64(value))
	}

	aoReq.Uids = append(aoReq.Uids, uids_uint64...)

	aoRes := NewAoBatchGetUserInfoByUidRsp()

	info, err := goClient.Call(req, aoReq, aoRes)

	if err == nil {
		info_struct := info.(*AoBatchGetUserInfoByUidRsp)

		info_map_struct := info_struct.Users

		returnData := make(map[int64]ModelUser)

		for uidOne, infoOne := range info_map_struct {
			returnData[int64(uidOne)] = exGetUserInfoFromAoData(infoOne)
		}
		return returnData, err
	}
	if aoRes.Result != 0 {
		err = errors.New(fmt.Sprintf("result code is %d", aoRes.Result))
	}
	return nil,err
}

func exUserGetNickName(uid uint64, nickName string) string {
	uid_string := strconv.FormatInt(int64(uid), 10)
	if nickName == "" {
		nickName = "用户" + uid_string
	}
	return nickName
}

func exUserGetSignature(signature string) string {
	if signature == "" {
		signature = "分享每一个有温度的瞬间"
	}
	return signature
}

func exUserGetPhoto(photo string) string {
	return photo
}

func exUserGetNameByUserType(userType uint8) string {
	userTypeName := [4]string{"会员", "育儿顾问", "官方运营", "育儿专家"}
	return userTypeName[userType]
}

func exUserGetFullPhotoByPhoto(photo string) string {
	var FullPhoto string
	if photo == "" {
		FullPhoto = "http://st.haiziwang.com/dynamic/m/img/personal-default-avatar.png"
	} else {
		FullPhoto = photo
	}
	return FullPhoto
}

func exGetUserInfoFromAoData(infoOne AoUserInfoXXOO) ModelUser {
	u := new(ModelUser)
	u.Uid = int64(infoOne.Uid)
	u.AccountType = infoOne.AccountType
	u.Mobile = infoOne.Mobile
	u.Email = infoOne.Email
	u.LoginAccount = infoOne.LoginAccount
	//u.WechatOpenid = info_map_struct.WechatOpenid
	u.WechatUnionid = infoOne.WechatUnionid
	u.UserType = infoOne.UserType
	//u.DiffSrcRegTime = info_map_struct.DiffSrcRegTime
	u.Rating = infoOne.Rating
	u.BabyInfo = infoOne.BabyInfo
	u.RelationWithBaby = infoOne.RelationWithBaby
	u.Nickname = exUserGetNickName(infoOne.Uid, infoOne.Nickname)
	u.Truename = infoOne.Truename
	u.Sex = infoOne.Sex
	u.Birthday = int64(infoOne.Birthday)
	u.QQNumber = infoOne.QQNumber
	u.Phone = infoOne.Phone
	u.Fax = infoOne.Fax
	u.Region = infoOne.Region
	u.Address = infoOne.Address
	u.Community = infoOne.Community
	u.Postcode = infoOne.Postcode
	u.IdentityType = infoOne.IdentityType
	u.IdentityNum = infoOne.IdentityNum
	u.Photo = exUserGetPhoto(infoOne.Photo)
	u.Signature = exUserGetSignature(infoOne.Signature)

	u.UserTypeName = exUserGetNameByUserType(u.UserType)
	u.Full_photo = exUserGetFullPhotoByPhoto(u.Photo)
	u.Short_address = ""

	return *u
}

func ExUserModifyBasicByUid(uid int64, signature string, signatureFlag bool,
	photo string, photoFlag bool,
	nickName string, nickNameFlag bool,
	userType uint8, userTypeFlag bool,
	host string, machineKey string, source string) (int, error) {

	var req = NewRequest()

	req.Host = host

	req.MachineKey = machineKey

	req.SceneId = 1

	req.Source = source

	goClient := New()

	aoReq := NewAoModifyBasicUserInfoByUidReq()

	userInfo := NewAoUserInfoXXOO()

	if signatureFlag {
		userInfo.Signature_u = 1
		userInfo.Signature = signature
	}
	if photoFlag {
		userInfo.Photo_u = 1
		userInfo.Photo = photo
	}
	if nickNameFlag {
		userInfo.Nickname_u = 1
		userInfo.Nickname = nickName
	}
	if userTypeFlag {
		userInfo.UserType_u = 1
		userInfo.UserType = userType
	}
	userInfo.Uid_u = 1
	userInfo.Uid = uint64(uid)
	aoReq.UserInfo = userInfo
	aoReq.Uid = uint64(uid)

	aoRes := NewAoModifyBasicUserInfoByUidRsp()

	_, err := goClient.Call(req, aoReq, aoRes)

	if err != nil {
		return 0, err
	}

	return int(aoRes.Result), nil
}

func ExUserGetUserInfoByPhone(phone string, host string, machineKey string, source string) (*ModelUser, error) {

	var req = NewRequest()

	req.Host = host

	req.MachineKey = machineKey

	req.SceneId = 1

	req.Source = source

	goClient := New()

	aoReq := NewAoGetUserInfoByBindInfoReq()
	aoReq.BindInfoType = 1
	aoReq.BindInfo = phone
	aoRes := NewAoGetUserInfoByBindInfoRsp()

	info, err := goClient.Call(req, aoReq, aoRes)

	if err != nil {
		return nil, err
	} else if aoRes.Result != 0 {
		err = errors.New(fmt.Sprintf("result code is %d", aoRes.Result))
		return nil, err
	}

	info_struct := info.(*AoGetUserInfoByBindInfoRsp)

	if info_struct.User != nil {
		user := exGetUserInfoFromAoData(*info_struct.User)
		return &user, nil
	} else {
		return nil, nil
	}
}
