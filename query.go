package zabbix

import (
	"fmt"
	"net/url"
)

type Zabbix struct {
	URL string
}

type ZabbixResponse struct {
	Result map[string]interface{} `json:result`
}

var Client Zabbix

func (zabbix Zabbix) Login() (ZabbixResponse, error) {
	vals := make(url.Values)
	fmt.Println(zabbix.URL)
	var res ZabbixResponse
	res.Result = map[string]interface{}{
		"result": "ok",
	}

	return res, nil
}
