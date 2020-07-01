package main

import "bing89.com/virtualapi/libs"



func main(){
	api := libs.NewAPI("/demo")
	resp1 := make(map[string]interface{})
	resp1["name"] = "demo"
	resp1["age"] = 12
	resp1["ccc"] = "a7"
	header1 := make(map[string]string)
	header1["Content-Type"] = "application/json"
	method1 := libs.APIMethod{
		Response: resp1,
		Headers: header1,
		Method: "GET",
	}
	api.AddMethod("GET", &method1)
	resp2 := make(map[string]interface{})
	resp2["name"] = "tom"
	resp2["age"] = 32
	resp2["ccc"] = "x6"
	method2 := libs.APIMethod{
		Response: resp2,
		Headers: header1,
		Method: "POST",
	}
	api.AddMethod("POST", &method2)

	group := libs.NewAPIGroup("/api/v1")
	group.AddAPI(api)

	service := libs.NewService("demo")
	service.Port = 3000
	service.AddGroup(group)
	service.AddAPI(api, "")

	config := libs.Configuration{}
	config.AddService(service)
	service1 := libs.NewService("default")
	service1.Port = 3001
	service1.AddAPI(api, "")
	config.AddService(service1)
	config.Save("/git/virtualapi/config.json")
}