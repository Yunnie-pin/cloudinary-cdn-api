package data

type ResponseModel struct {
	Response   int         `json:"response"`
	Error      interface{} `json:"error"`
	AppID      string      `json:"appid"`
	Controller string      `json:"controller"`
	Action     string      `json:"action"`
	Result     interface{} `json:"result"`
}
