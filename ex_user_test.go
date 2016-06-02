package goao

import (
	//"strconv"
	//"fmt"
	"testing"
)

func Test_ExGetUserInfos(t *testing.T) {
  var uids []int64
  uids = append(uids,1000159002)
  host:= "172.172.177.5:53101"
  source:= "goaotest"
  machineKey:= "goaotest"
  _,err:=ExUserGetUserInfos(uids,host,machineKey,source)
  if err!=nil{
    t.Errorf("error:%s", err.Error())
  }
}
