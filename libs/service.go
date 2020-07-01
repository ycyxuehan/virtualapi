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
	Groups      []*APIGroup `json:"groups"`
	APIs        []*API   `json:"apis"`
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
		Groups: []*APIGroup{},
		APIs: []*API{},
		Enabled: true,
		Port: 3000,
	}
	return &svc
}

//AddGroup 添加一个API组
func (s *Service)AddGroup(group *APIGroup)error{
	for _, g := range s.Groups {
		if g.Name == group.Name {
			return fmt.Errorf("group [%s] 已存在", group.Name)
		}
	}
	s.Groups = append(s.Groups, group)
	return nil
}

//AddOwnAPI 添加一个独立API
func (s *Service)AddOwnAPI(api *API)error{
	for _, a := range s.APIs {
		if a.Name == api.Name {
			return fmt.Errorf("api [%s] 已存在", api.Name)
		}
	}
	s.APIs = append(s.APIs, api)
	return nil
}
//MatchAPI 匹配URL的API
func (s *Service)MatchAPI(URL string)(*API, error){
	if URL == "" {
		return nil, fmt.Errorf("404 not found")
	}
	if s.Prefix != "" && strings.Index(URL, s.Prefix) < 0 {
		return nil, fmt.Errorf("404 not found")
	}
	for _, group := range s.Groups {
		if strings.Index(URL, strings.Join([]string{s.Prefix, group.Name}, "")) > -1 {
			//组匹配，匹配API
			for _, api := range group.APIs {
				//完全匹配
				if strings.Index(URL, strings.Join([]string{s.Prefix, group.Name, api.Name}, "")) > -1 {
					return api, nil
				}
			}
			
		}
	}
	//没有匹配的group，匹配独立API
	for _, api := range s.APIs {
		//完全匹配
		if strings.Index(URL, strings.Join([]string{s.Prefix, api.Name}, "")) > -1 {
			return api, nil
		}
	}
	return nil, fmt.Errorf("404 not found")
}

//GetAPI 获得匹配URL的API配置
func (s *Service)GetAPI(name string, group string)(*API, error){
	if group == ""{
		for _, api := range s.APIs {
			if api.Name == name {
				return api, nil
			}
		}
	}
	for _, g := range s.Groups {
		for _, api := range g.APIs {
			if api.Name == name {
				return api, nil
			}
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
	fmt.Printf("%s [%s] %s\n", time.Now().Format("2006-01-02 15:04:05"), r.Method, r.RequestURI)
	api, err := s.MatchAPI(r.RequestURI)
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
	msg.Data = args.Response
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
	for _, g := range s.Groups {
		if g.Name == group {
			return true
		}
	}
	return false
}
//GetGroup 获取指定group
func (s *Service)GetGroup(name string)(*APIGroup, error){
	for _, g := range s.Groups {
		if g.Name == name {
			return g, nil
		}
	}
	return nil, fmt.Errorf("group not found")
}
//AddGroupAPI 添加一个API到group
func (s *Service)AddGroupAPI(group string, api *API)error{
	g, err := s.GetGroup(group)
	if err != nil {
		return err
	}
	return g.AddAPI(api)
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
	a, err := s.GetAPI(api, group)
	if err!=nil {
		return res
	}
	return a.GetMethods()
}