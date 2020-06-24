package libs

import ()

//APIGroup api group
type APIGroup struct {
	Name string	`json:"name"`
	APIs []API `json:"apis"`
	Auth string	`json:"auth"`
}

//NewAPIGroup create a new APIGroup
func NewAPIGroup(name string)APIGroup{
	if name == "" {
		name = "/group"
	}
	g := APIGroup{
		Name: name,
	}
	return g
}

