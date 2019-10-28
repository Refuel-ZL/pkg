package xingxing

import (
	"cdnservices/internal/pkg/wx"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"net/http"
)

const wxcheckdomainurl = "http://vip.xxweixin.com/weixin/wx_domain.php"

type XingX struct {
	User string `json:"user"`
	Key  string `json:"key"`
}

type WxCheckDomainResponse struct {
	Status  int8   `json:"status"`
	Domain  string `json:"domain"`
	Errmsg  string `json:"errmsg"`
	User    string `json:"user"`
	Endtime string `json:"endtime"`
}

func (x *XingX) Domain(domain string) (res wx.Wxdomianresponse, err error) {
	args := &fasthttp.Args{}
	args.Add("user", x.User)
	args.Add("domain", domain)
	args.Add("key", x.Key)
	code, resp, err := fasthttp.Post(nil, fmt.Sprintf("%v?%v", wxcheckdomainurl, string(args.QueryString())), args)
	if err != nil {
		return res, err
	}
	fmt.Println(string(resp))
	if code != http.StatusOK {
		return res, fmt.Errorf("%v", resp)
	}
	var _res WxCheckDomainResponse
	err = json.Unmarshal(resp, &_res)
	if err != nil {
		return res, err
	}
	res.Status = _res.Status
	res.Domain = _res.Domain
	return res, nil
}
