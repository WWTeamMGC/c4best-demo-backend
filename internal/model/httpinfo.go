package model

// HttpInfo kafka 传进参数的Struct结构体
type HttpInfo struct {
	ClientIP string
	Method   string
	UrlPath  string
	Header   string
	body     string
}
