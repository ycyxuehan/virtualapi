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
	return err
}

//Save save configuration to file
func (c *Configuration)Save(file string)error {
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
	return err
}

//AddService create a service and add it
func (c *Configuration)AddService(svc *Service){
	c.Services = append(c.Services, svc)
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