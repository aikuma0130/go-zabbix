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
	Token    string
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

func NewZabbixRequest(method string, params interface{}) *ZabbixRequest {
	r := new(ZabbixRequest)
	r.Jsonrpc = "2.0"
	r.Method = method
	r.Params = params
	r.ID = 1
	return r
}

var Client Zabbix

func (zabbix *Zabbix) Login() error {
	u := zabbix.URL + "/api_jsonrpc.php"
	params := map[string]interface{}{
		"user":     zabbix.User,
		"password": zabbix.Password,
	}
	data := NewZabbixRequest("user.login", params)
	jsonBytes, _ := json.Marshal(&data)

	req, _ := http.NewRequest("POST", u, bytes.NewBuffer(jsonBytes))
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

	zabbix.Token = response.Result.(string)

	return nil
}

func (zabbix *Zabbix) Do(param interface{}) (ZabbixResponse, error) {
	var response ZabbixResponse
	return response, nil
}
