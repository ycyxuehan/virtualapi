package libs

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

//Service api server config
type Service struct {
	Name        string  `json:"name"`
	Port        int     `json:"port"`
	Prefix      string  `json:"prefix"`
	Groups      map[string]*APIGroup `json:"groups"`
	APIs        map[string]*API   `json:"apis"`
	Description string  `json:"description"`
	Enabled     bool    `json:"enabled"`
}

//NewService create a new service
func NewService(name string) *Service {
	if name == "" {
		name = "default"
	}
	svc := Service{
		Name: name,
		Groups: make(map[string]*APIGroup),
		APIs: make(map[string]*API),
		Enabled: true,
		Port: 3000,
	}
	return &svc
}

//AddGroup 添加一个API组
func (s *Service)AddGroup(group *APIGroup)error{
	for g := range s.Groups {
		if g == group.Name {
			return fmt.Errorf("group [%s] 已存在", group.Name)
		}
	}
	s.Groups[group.Name] = group
	return nil
}

//AddOwnAPI 添加一个独立API
func (s *Service)AddOwnAPI(api *API)error{
	for g := range s.APIs {
		if g == api.Name {
			return fmt.Errorf("api [%s] 已存在", api.Name)
		}
	}
	s.APIs[api.Name] = api
	return nil
}
//GetAPI 获得匹配URL的API配置
func (s *Service)GetAPI(URL string)(*API, error){
	if URL == "" {
		return nil, fmt.Errorf("404 not found")
	}
	if s.Prefix != "" && strings.Index(URL, s.Prefix) < 0 {
		return nil, fmt.Errorf("404 not found")
	}
	for group := range s.Groups {
		if strings.Index(URL, strings.Join([]string{s.Prefix, group}, "")) > -1 {
			//组匹配，匹配API
			for api := range s.Groups[group].APIs {
				//完全匹配
				if strings.Index(URL, strings.Join([]string{s.Prefix, group, api}, "")) > -1 {
					return s.Groups[group].APIs[api], nil
				}
			}
			
		}
	}
	//没有匹配的group，匹配独立API
	for api := range s.APIs {
		//完全匹配
		if strings.Index(URL, strings.Join([]string{s.Prefix, api}, "")) > -1 {
			return s.APIs[api], nil
		}
	}
	return nil, fmt.Errorf("404 not found")
}

func (s *Service)ServeHTTP(w http.ResponseWriter, r *http.Request){
	msg := ResponseMessage{}
	if r.RequestURI == "/favicon.ico" {
		msg.Code = 404
		msg.Message = "404 not found"
		w.Write(msg.ToBytes())
		return
	}
	fmt.Printf("%s [%s] %s\n", time.Now().Format("2000-01-02 15:04:05"), r.Method, r.RequestURI)
	api, err := s.GetAPI(r.RequestURI)
	if err != nil {
		msg.Code = 404
		msg.Message = err.Error()
		w.Write(msg.ToBytes())
		return
	}
	args, err := api.GetMethod(r.Method)
	if err != nil {
		msg.Code = 404
		msg.Message = err.Error()
		w.Write(msg.ToBytes())
		return
	}
	msg.Data = args.Responese
	w.Write(msg.ToBytes())
}


//Run 运行service
func (s *Service)Run(done chan Message)error{
	fmt.Printf("为服务%s监听端口%d\n", s.Name, s.Port)
	done <- Message{s.Name, true}
	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", s.Port), s)
	done <- Message{s.Name, false}
	return err
}

//IsGroupExists service中是否存在API组
func (s *Service)IsGroupExists(group string)bool{
	for g := range s.Groups {
		if g == group {
			return true
		}
	}
	return false
}
//AddGroupAPI 添加一个API到group
func (s *Service)AddGroupAPI(group string, api *API)error{
	if s.IsGroupExists(group) {
		return s.Groups[group].AddAPI(api)
	}
	g := NewAPIGroup(group)
	g.AddAPI(api)
	s.Groups[group] = g
	return nil
}

//AddAPI 添加一个API
func (s *Service)AddAPI(api *API, group string)error{
	if group == "" {
		return s.AddOwnAPI(api)
	}
	return s.AddGroupAPI(group, api)
}

//GetMethods 获取method列表
func (s *Service)GetMethods(api string, group string)[]string{
	res := []string{}
	if group == "" {
		if g, ok:= s.Groups[group]; ok {
			return g.GetMethods(api)
		}
	}else {
		if a, ok := s.APIs[api]; ok {
			return a.GetMethods()
		}
	}
	return res
}