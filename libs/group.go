package libs

import (
	"fmt"
)

//APIGroup api group
type APIGroup struct {
	Name string `json:"name"`
	// APIs []*API  `json:"apis"`
	APIs []*API `json:"apis"`
	Auth string `json:"auth"`
}

//NewAPIGroup create a new APIGroup
func NewAPIGroup(name string) *APIGroup {
	return &APIGroup{
		APIs: []*API{},
		Name: name,
		Auth: "",
	}
}

//AddAPI 添加一个API
func (g *APIGroup)AddAPI(a *API)error {
	for _, api := range g.APIs {
		if api.Name == a.Name {
			return fmt.Errorf("api [%s] is exits", api.Name)
		}
	}
	g.APIs = append(g.APIs, a)
	return nil
}

//AddMethod 添加一个method
func (g *APIGroup)AddMethod(name string, method string, args *APIMethod)error {
	for _, a := range g.APIs {
		if a.Name == name {
			return a.AddMethod(args)
		}
	}
	api := NewAPI(name)
	api.AddMethod(args)
	g.AddAPI(api)
	return nil
}

//GetMethods 获取method列表
func (g *APIGroup)GetMethods(name string)[]string{
	res := []string{}
	for _, a := range g.APIs {
		if a.Name == name {
			return a.GetMethods()
		}
	}
	return res
}