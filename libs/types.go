package libs

import (
	"encoding/json"
	"fmt"
)

//Header api header
type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}


//Query api url query
type Query struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (q *Query)String()string {
	return fmt.Sprintf("%s=%s", q.Key, q.Value)
}
//Message running message
type Message struct {
	Service string 
	Running bool
}

//ResponseMessage HTTP响应信息
type ResponseMessage struct {
	Code uint `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

//ToBytes 返回byte数组
func (r *ResponseMessage)ToBytes()[]byte {
	data, _ := json.MarshalIndent(r, "", "\t")
	return data
}

