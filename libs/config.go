package libs

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

//Config 配置
var Config Configuration

//Configuration virtual api server configuration
type Configuration struct {
	filePath string
	Services []*Service `json:"services"`
}

//Load load configuration from file
func (c *Configuration)Load(file string)error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, c)
	if err == nil {
		c.filePath = file
	}
	return err
}

//Save save configuration to file
func (c *Configuration)Save()error {
	data, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		return err
	}
	// f, err := os.Open(file)
	// if err != nil {
	// 	return err
	// }
	// defer f.Close()
	err = ioutil.WriteFile(c.filePath, data, os.ModePerm)
	return err
}

//SaveAs save configuration to file
func (c *Configuration)SaveAs(file string)error {
	data, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		return err
	}
	// f, err := os.Open(file)
	// if err != nil {
	// 	return err
	// }
	// defer f.Close()
	err = ioutil.WriteFile(file, data, os.ModePerm)
	if err == nil {
		c.filePath = file
	}
	return err
}


//AddService create a service and add it
func (c *Configuration)AddService(svc *Service)error{
	c.Services = append(c.Services, svc)
	return c.Save()
}

//GetService 通过service名称获取service
func (c *Configuration)GetService(name string)*Service {
	for _, svc := range c.Services {
		if svc.Name == name {
			return svc
		}
	}
	return nil
}

//GetServiceNames 获取所有service名称
func (c *Configuration)GetServiceNames()[]string {
	res := []string{}
	for _, svc := range c.Services {
		res = append(res, svc.Name)
	}
	return res
}

//AddGroup 为一个服务添加group并保存
func (c *Configuration)AddGroup(service string, group string)error{
	svc := c.GetService(service)
	g := NewAPIGroup(group)
	err := svc.AddGroup(g)
	if err != nil {
		return err
	}
	return c.Save()
}

//AddAPI 为一个服务添加API
func (c *Configuration)AddAPI(svc string, group string, api string) error {
	a := NewAPI(api)
	s := c.GetService(svc)
	err := s.AddAPI(a, group)
	if err != nil {
		return err
	}
	return c.Save()
}

//AddMethod 为指定服务的API添加一个method
func (c *Configuration)AddMethod(svc string, group string, api string, method *APIMethod)error{
	s := c.GetService(svc)
	a, err := s.GetAPI(api, group)
	if err != nil {
		return err
	}
	err = a.AddMethod(method)
	if err != nil {
		return err
	}
	return c.Save()
}