package libs

import ()

//Service api server config
type Service struct {
	Name string	`json:"name"`
	Port int `json:"port"`
	Prefix string `json:"prefix"`
	Groups []Group `json:"groups"`
	APIs []API `json:"apis"`
	Description string `json:"description"`
	Enabled bool `json:"enabled"`
}

//NewService create a new service
func NewService(name string)Service {
	if name == "" {
		name = "default"
	}
	svc := Service{
		Name: name,
	}
	return svc
}

