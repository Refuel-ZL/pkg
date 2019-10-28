package monkey

import (
	"fmt"
	"testing"
)

func TestXingX_Domain(t *testing.T) {
	aa := MonkeyX{
		Key: "asdasdasd",
	}
	res, err := aa.Domain("baidu.com")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%+v", res)
}
