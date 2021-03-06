package main

import (
	"github.com/aikuma0130/go-zabbix"
)

func main() {
	zabbix.Client.URL = "http://localhost:8081"
	zabbix.Client.User = "Admin"
	zabbix.Client.Password = "zabbix"
	err := zabbix.Client.Login()
	if err != nil {
		return
	}
	defer zabbix.Client.Logout()

	params := map[string]interface{}{
		"output": "extend",
		"filter": map[string]interface{}{
			"host": []string{"Template OS Linux", "Template OS Windows"}}}

	req := zabbix.NewZabbixRequest("template.get", params)
	//var response *zabbix.ZabbixResponse
	//response, err = zabbix.Client.Do(req)
	_, err = zabbix.Client.Do(req)
	if err != nil {
		return
	}

	return
	//fmt.Printf("%v", response.Result)
}
