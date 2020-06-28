package libs

import (
	"fmt"
	"strings"
)

//Method api method
type Method string

//all http methods
const (
	GET Method = "GET"
	POST Method = "POST"
	PUT Method = "PUT"
	DELETE Method = "DELETE"
)

//NewMethod 使用string创建一个Method
func NewMethod(m string)Method{
	switch strings.ToUpper(m) {
		case "GET":
			return GET
		case "PUT":
			return PUT
		case "POST":
			return POST
		case "DELETE":
			return DELETE
		default:
			return GET
	}
}

//APIArgs Api conf
type APIArgs struct {
	Responese string   `json:"responese"`
	Request   string   `json:"request"`
	Headers   map[string]string  `json:"headers"`
	Queries   map[string]string  `json:"queries"`
}

//API Api conf
type API struct {
	Methods map[string]*APIArgs	`json:"methods"`
	Auth string	`json:"auth"`
} 

//NewAPI create a new api
func NewAPI() *API {
	return &API{
		Methods: make(map[string]*APIArgs),
		Auth: "",
	}
}

//AddMethod 为API添加一个method
func (api *API)AddMethod(method string, args *APIArgs)error{
	for m := range api.Methods {
		if m == method {
			return fmt.Errorf("[%s] method is exists", method)
		}
	}
	api.Methods[strings.ToUpper(method)] = args
	return nil
}

//GetMethod 获取一个对应Method的配置
func (api *API)GetMethod(method string)(*APIArgs, error){
	if args, ok := api.Methods[method]; ok {
		return args, nil
	}
	return nil, fmt.Errorf("404 not found")
}

//AddHeader 添加一个header
func (a *APIArgs)AddHeader(key string, value string) error {
	for k := range a.Headers {
		if key == k {
			return fmt.Errorf("header [%s] exists", k)
		}
	}
	a.Headers[key] = value
	return nil
}

//AddQuery 添加一个Query
func (a *APIArgs)AddQuery(key string, value string)error{
	for k := range a.Queries {
		if k == key {
			return fmt.Errorf("query [%s] exists", key)
		}
	}
	a.Queries[key] = value
	return nil
}

//DelQuery 删除一个Query
func (a *APIArgs)DelQuery(key string){
	delete(a.Queries, key)
}

//DelHeader 删除一个Header
func (a *APIArgs)DelHeader(key string){
	delete(a.Queries, key)
}