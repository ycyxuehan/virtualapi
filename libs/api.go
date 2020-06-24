package libs

import "fmt"

//API Api conf
type API struct {
	Name string	`json:"name"`
	Method string `json:"method"`
	Responese string `json:"responese"`
	Request string	`json:"request"`
	Headers []Header	`json:"headers"`
	Queries	[]Query	`json:"queries"`
}

//NewAPI create a new api
func NewAPI(name string)API{
	if name == "" {
		name = "/"
	}
	a := API{
		Name: name,
	}
	return a
}

//AddHeader 添加一个header
func (a *API)AddHeader(h Header)error{
	for _, header := range a.Headers {
		if header.Key == h.Key {
			return fmt.Errorf("header [%s] exists", h.Key)
		}
	}
	a.Headers = append(a.Headers, h)
	return nil
}