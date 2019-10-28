package fastsms

import (
	"fmt"
	"testing"
)

func TestSend(t *testing.T) {
	sms := FasrSms{
		Account: 73,
		Secret:  "IYFB2wHK0PBX/WHklOgc1g==",
	}
	if res, err := sms.Send([]string{"17773701949"}, "【测试】 你有收到信息吗"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%+v\n", res)
	}
}
