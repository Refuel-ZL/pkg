package monkey

import (
	"cdnservices/internal/pkg/wx"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"net/http"
)

const DomainUri = "http://api.monkeyapi.com"

type MonkeyX struct {
	Key string `json:"appkey"`
}

type WxCheckDomainResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

func (x *MonkeyX) Domain(domain string) (res wx.Wxdomianresponse, err error) {
	args := &fasthttp.Args{}
	args.Add("url", domain)
	args.Add("key", x.Key)
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req) // 用完需要释放资源
	url := fmt.Sprintf("%v?%v", DomainUri, string(args.QueryString()))
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")

	req.SetRequestURI(url)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp) // 用完需要释放资源

	if err := fasthttp.Do(req, resp); err != nil {
		fmt.Println("请求失败:", err.Error())
		return res, err
	}
	code := resp.StatusCode()

	if code != http.StatusOK {
		return res, fmt.Errorf("%v", resp)
	}
	res.Domain = domain
	var _res WxCheckDomainResponse
	b := resp.Body()
	fmt.Printf("%v\n", string(b))
	err = json.Unmarshal(b, &_res)
	if err != nil {
		return res, err
	}
	if _res.Code == 200 {
		switch _res.Data { //0为正常，1为"非官方网址，请确认是否继续访问"，2为域名已封杀，3提示如需浏览，请长按网址复制后使用浏览器打开
		case "0":
			res.Status = 0
		case "1":
			fallthrough
		case "3":
			res.Status = 1
		case "2":
			res.Status = 2
		default:
			res.Status = 3
		}
	} else {
		switch _res.Code {
		case 513:
			res.Status = -1
		default:
			res.Status = 3
		}
	}

	return res, nil
}
