package libs

import (
	"fmt"
)

//APIGroup api group
type APIGroup struct {
	// Name string `json:"name"`
	// APIs []*API  `json:"apis"`
	APIs map[string]*API `json:"apis"`
	Auth string `json:"auth"`
}

//NewAPIGroup create a new APIGroup
func NewAPIGroup() *APIGroup {
	return &APIGroup{
		APIs: make(map[string]*API),
		Auth: "",
	}
}

//AddAPI 添加一个API
func (g *APIGroup)AddAPI(location string, a *API)error {
	for l := range g.APIs {
		if l == location {
			return fmt.Errorf("api [%s] is exits", l)
		}
	}
	g.APIs[location] = a
	return nil
}

//AddMethod 添加一个method
func (g *APIGroup)AddMethod(location string, method string, args *APIArgs)error {
	for l := range g.APIs {
		if l == location {
			return g.APIs[location].AddMethod(method, args)
		}
	}
	api := NewAPI()
	api.AddMethod(method, args)
	g.AddAPI(location, api)
	return nil
}