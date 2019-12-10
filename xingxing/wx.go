package xingxing

import (
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

type Wxdomianresponse struct {
	Status int8   `json:"status"` //0为正常，1为被转码，2为被封，3查询失败，-1接口到期
	Domain string `json:"domain"`
}

func (x *XingX) Domain(domain string) (res Wxdomianresponse, err error) {
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
	var _res struct {
		Status  int8   `json:"status"`
		Domain  string `json:"domain"`
		Errmsg  string `json:"errmsg"`
		User    string `json:"user"`
		Endtime string `json:"endtime"`
	}
	err = json.Unmarshal(resp, &_res)
	if err != nil {
		return res, err
	}
	res.Status = _res.Status
	res.Domain = _res.Domain
	return res, nil
}
