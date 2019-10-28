package fastsms

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const sendurl = "http://fastsms.happycheer.com/send"

type FasrSms struct {
	Account int    `json:"account"`
	Secret  string `json:"secret"`
}

func (f *FasrSms) Send(phones []string, content string) (res SendResponse, err error) {
	sendRequest := SendRequest{
		Account: f.Account,
		Secret:  f.Secret,
		Mobile:  strings.Join(phones, ";"),
		Content: content,
	}
	form := url.Values{}
	form.Set("account", strconv.Itoa(sendRequest.Account))
	form.Set("mobile", sendRequest.Mobile)
	form.Set("content", sendRequest.Content)
	form.Set("secret", sendRequest.Secret)

	headers := make(map[string]string)
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	body := bytes.NewBufferString(form.Encode())
	resp, err := request("POST", sendurl, headers, body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(resp, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func request(method string, url string, headers map[string]string, body io.Reader) (res []byte, err error) {
	method = strings.ToUpper(method)
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		request.Header.Set(k, v)
	}
	client := http.Client{
		Timeout: time.Second * 15,
	}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBytes, nil
}
