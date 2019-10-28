package fastsms

type SendRequest struct {
	Account int    `json:"account"`
	Mobile  string `json:"mobile"`
	Content string `json:"content"`
	Secret  string `json:"secret"`
}

type SendResponse struct {
	Code int `json:"code"`
	Data struct {
		Details []struct {
			Mobile string `json:"mobile"`
			MsgId  uint64 `json:"msg_id"`
		} `json:"details"`
	} `json:"data"`
	Message string `json:"message"`
}
