package libs

import ()

//Configuration virtual api server configuration
type Configuration struct {
	Services []Service `json:"services"`
}

//Load load configuration from file
func (c *Configuration)Load(file string)error {

	return nil
}

//Save save configuration to file
func (c *Configuration)Save(file string)error {
	return nil
}

//NewService create a service and add it
func (c *Configuration)NewService(){

}