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

//APIMethod Api conf
type APIMethod struct {
	Method string `json:"method"`
	Responese map[string]interface{}   `json:"responese"`
	Request    map[string]interface{}   `json:"request"`
	Headers   map[string]string  `json:"headers"`
	Queries   map[string]string  `json:"queries"`
}

//API Api conf
type API struct {
	Name string `json:"name"`
	Methods map[string]*APIMethod	`json:"methods"`
	Auth string	`json:"auth"`
} 

//NewAPI create a new api
func NewAPI(name string) *API {
	return &API{
		Name: name,
		Methods: make(map[string]*APIMethod),
		Auth: "",
	}
}

//AddMethod 为API添加一个method
func (api *API)AddMethod(method string, args *APIMethod)error{
	for m := range api.Methods {
		if m == method {
			return fmt.Errorf("[%s] method is exists", method)
		}
	}
	api.Methods[strings.ToUpper(method)] = args
	return nil
}

//GetMethod 获取一个对应Method的配置
func (api *API)GetMethod(method string)(*APIMethod, error){
	if args, ok := api.Methods[method]; ok {
		return args, nil
	}
	return nil, fmt.Errorf("404 not found")
}

//AddHeader 添加一个header
func (a *APIMethod)AddHeader(key string, value string) error {
	for k := range a.Headers {
		if key == k {
			return fmt.Errorf("header [%s] exists", k)
		}
	}
	a.Headers[key] = value
	return nil
}

//AddQuery 添加一个Query
func (a *APIMethod)AddQuery(key string, value string)error{
	for k := range a.Queries {
		if k == key {
			return fmt.Errorf("query [%s] exists", key)
		}
	}
	a.Queries[key] = value
	return nil
}

//DelQuery 删除一个Query
func (a *APIMethod)DelQuery(key string){
	delete(a.Queries, key)
}

//DelHeader 删除一个Header
func (a *APIMethod)DelHeader(key string){
	delete(a.Queries, key)
}

//GetMethods 获取method列表
func (api *API)GetMethods()[]string{
	res := []string{}
	for m := range api.Methods {
		res = append(res, m)
	}
	return res
}