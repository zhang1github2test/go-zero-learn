// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2

package types

type UserReq struct {
	Name    string `json:"name,optional"`
	Id      string `json:"id,optional"`
	Age     int    `json:"age,optional"`
	TraceId string `header:"X-Trace-Id"`
}

type UserReqResp struct {
	UserReq
	Status string `json:"status"`
}
