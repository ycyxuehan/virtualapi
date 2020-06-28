package libs

import (
	"fmt"
)

//APIGroup api group
type APIGroup struct {
	Name string `json:"name"`
	// APIs []*API  `json:"apis"`
	APIs map[string]*API `json:"apis"`
	Auth string `json:"auth"`
}

//NewAPIGroup create a new APIGroup
func NewAPIGroup(name string) *APIGroup {
	return &APIGroup{
		APIs: make(map[string]*API),
		Name: name,
		Auth: "",
	}
}

//AddAPI 添加一个API
func (g *APIGroup)AddAPI(a *API)error {
	for l := range g.APIs {
		if l == a.Name {
			return fmt.Errorf("api [%s] is exits", l)
		}
	}
	g.APIs[a.Name] = a
	return nil
}

//AddMethod 添加一个method
func (g *APIGroup)AddMethod(name string, method string, args *APIArgs)error {
	for l := range g.APIs {
		if l == name {
			return g.APIs[name].AddMethod(method, args)
		}
	}
	api := NewAPI(name)
	api.AddMethod(method, args)
	g.AddAPI(api)
	return nil
}

//GetMethods 获取method列表
func (g *APIGroup)GetMethods(name string)[]string{
	res := []string{}
	if a, ok := g.APIs[name]; ok {
		return a.GetMethods()
	}
	return res
}