package common

// 分页列表请求
type PageListReq struct {
	Page     int    `json:"page" query:"page" default:"1"`
	PageSize int    `json:"page_size" query:"page_size" default:"10"`
	Keyword  string `json:"keyword" query:"keyword"`
}

// 列表响应
type ListResp struct {
	List any `json:"list"`
}

// 分页列表响应
type PageListResp struct {
	List  any   `json:"list"`
	Total int64 `json:"total"`
}

// 规范化结果输出
type BaseResponseMetadata struct {
	RequestId string `json:"RequestId"`
	Action    string `json:"Action"`
	Version   string `json:"Version"`
	Service   string `json:"Service"`
	Region    string `json:"Region"`
}

type ErrorResponseMetadata struct {
	BaseResponseMetadata
	Error Error `json:"Error"`
}

type ResponseMetadata struct {
	ResponseMetadata *BaseResponseMetadata `json:"ResponseMetadata"`
	Result           interface{}           `json:"Result"`
}

type Error struct {
	Code    int    `json:"Code"`
	Message string `json:"Message"`
}
