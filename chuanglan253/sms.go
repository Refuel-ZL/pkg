package chuanglan253

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

func (C *Chuanglan) Send(phone string, msg string) (err error) {
	var params SendRequest
	params = SendRequest{
		Chuanglan: C,
		Msg:       msg,
		Phone:     phone,
		Report:    "true",
	}
	resp, err := send(params)
	if err != nil {
		return err
	}
	if resp.Code != "0" {
		return errors.New(resp.ErrorMsg)
	}
	return nil
}

func send(params SendRequest) (srep SendResponse, err error) {
	bytesData, err := json.Marshal(params)
	if err != nil {
		return srep, err
	}
	url := "http://smssh1.253.com/msg/send/json" //短信发送URL
	request, err := http.NewRequest("POST", url, bytes.NewReader(bytesData))
	if err != nil {
		return srep, err
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{
		Timeout: time.Second * 15,
	}
	resp, err := client.Do(request)
	if err != nil {
		return srep, err
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return srep, err
	}
	err = json.Unmarshal(respBytes, &srep)
	return srep, err
}
