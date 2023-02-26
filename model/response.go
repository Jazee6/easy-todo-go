package model

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type GetTodoListResp struct {
	Id     int    `json:"id"`
	Todo   string `json:"todo"`
	IsDone bool   `json:"is_done"`
}
