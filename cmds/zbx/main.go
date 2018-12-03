package main

import (
	"fmt"

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
	fmt.Printf("%v", zabbix.Client.Token)
}
