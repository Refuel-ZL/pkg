package chuanglan253

type Chuanglan struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type SendRequest struct {
	*Chuanglan
	Msg      string `json:"msg"`
	Phone    string `json:"phone"`
	Sendtime string `json:"sendtime,omitempty"`
	Report   string `json:"report,omitempty"`
	Extend   string `json:"extend,omitempty"`
	Uid      string `json:"uid,omitempty"`
}

type SendResponse struct {
	Code     string `json:"code"`
	Time     string `json:"time"`
	MsgId    string `json:"msgId"`
	ErrorMsg string `json:"errorMsg"`
}
