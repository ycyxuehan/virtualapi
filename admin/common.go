package admin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"bing89.com/virtualapi/libs"
	"github.com/gin-gonic/gin"
)
//错误代码
const (
	Success = 0
	ParamError = 99990
	ObjectNotFound = 99991
	ReadPostDataError = 99992
	ParseDataError = 99994
	AddObjectError = 99993
)


//AddGroup 添加一个APIgroup
//@route [POST] /admin/group?service=[string]
func AddGroup(c *gin.Context){
	name := c.Query("service")
	if name == "" {
		ServeJSON(c, ParamError, "service名称为空", nil)
		return
	}
	svc := libs.Config.GetService(name)
	if svc == nil {
		ServeJSON(c, ObjectNotFound, "service不存在", nil)
		return
	}
	group := c.Request.Form.Get("group")
	if group == "" {
		ServeJSON(c, ParamError, "group名称为空", nil)
		return
	}
	err := svc.AddGroup(libs.NewAPIGroup(group))
	if err != nil {
		ServeJSON(c, AddObjectError, err, nil)
		return
	}
	ServeJSON(c, Success, "success", nil)
}

//AddAPI 添加一个API
//@route [POST] /admin/api?group=[string]&service=[string]
func AddAPI(c *gin.Context){
	name := c.Query("service")
	if name == "" {
		ServeJSON(c, ParamError, "service名称为空", nil)
		return
	}
	data, err := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	if err != nil {
		ServeJSON(c, ReadPostDataError, err, nil)
		return
	}
	svc := libs.Config.GetService(name)
	if svc == nil {
		ServeJSON(c, ObjectNotFound, "service不存在", nil)
		return
	}
	api := libs.NewAPI("")
	err = json.Unmarshal(data, api)
	if err != nil {
		ServeJSON(c, ParseDataError, err, nil)
		return
	}
	group := c.Query("group")
	err = svc.AddAPI(api, group)
	if err != nil {
		ServeJSON(c, AddObjectError, err, nil)
		return
	}
	ServeJSON(c, Success, "success", nil)
}

//GetAPIs 获取API列表
//@route [GET] /admin/api?group=[string]&service=[string]
func GetAPIs(c *gin.Context){

}

//GetService 获取API列表
//@route [GET] /admin/service/:name
func GetService(c *gin.Context){
	service := c.Param("name")
	svc := libs.Config.GetService(service)
	ServeJSON(c, Success, "success", svc)
}

//GetServices 获取API列表
//@route [GET] /admin/services
func GetServices(c *gin.Context){
	svcs := libs.Config.GetServiceNames()
	ServeJSON(c, Success, "success", &svcs)
}

//AddService 添加一个服务
//@route [POST] /admin/service
func AddService(c *gin.Context){
	data, err := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	if err != nil {
		ServeJSON(c, ReadPostDataError, err, nil)
		return
	}
	svc := libs.NewService("")
	err = json.Unmarshal(data, svc)
	if err != nil {
		ServeJSON(c, ParseDataError, err, nil)
		return
	}
	libs.Config.AddService(svc)
	ServeJSON(c, Success, "success", nil)
}

//GetMethods 获取API的method列表
//@route [GET] /admin/api/:name/methods?service=[string]&group=[string]
func GetMethods(c *gin.Context){
	api := c.Param("name")
	service := c.Query("service")
	group := c.Query("group")
	svc := libs.Config.GetService(service)
	if svc == nil {
		ServeJSON(c, ObjectNotFound, "service不存在", nil)
		return
	}
	methods := svc.GetMethods(api, group)
	ServeJSON(c, Success, "success", methods)
}

//ServeJSON 为指定上下文返回JSON数据
func ServeJSON(c *gin.Context, code uint, message interface{}, data interface{}){
	msg := libs.ResponseMessage{
		Code: code,
		Message: fmt.Sprintf("%v", message),
		Data: data,
	}
	c.JSON(200, &msg)
}