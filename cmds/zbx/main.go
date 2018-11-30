package main

import (
	"github.com/aikuma0130/go-zabbix"
)

func main() {
	zabbix.Client.URL = "http://localhost/zabbix"
	zabbix.Client.Login()
}
