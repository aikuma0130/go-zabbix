package zabbix

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Zabbix struct {
	URL      string
	User     string
	Password string
	Auth     string
}

type ZabbixResponse struct {
	Result interface{} `json:result`
}

type ZabbixRequest struct {
	Jsonrpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	ID      int         `json:"id"`
	Auth    string      `json:"auth,omitempty"`
}

var Client Zabbix

func (zabbix *Zabbix) Login() error {
	u := zabbix.URL + "/api_jsonrpc.php"
	data := ZabbixRequest{}
	data.Jsonrpc = "2.0"
	data.Method = "user.login"
	data.ID = 1
	//data.Auth = ""
	data.Params = map[string]interface{}{
		"user":     zabbix.User,
		"password": zabbix.Password,
	}

	jsonStr, _ := json.Marshal(&data)

	req, _ := http.NewRequest("POST", u, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json-rpc")
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var response ZabbixResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return err
	}

	zabbix.Auth = response.Result.(string)

	return nil
}
