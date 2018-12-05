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

func (zabbix *Zabbix) Do(request *ZabbixRequest) (*ZabbixResponse, error) {
	if zabbix.Token != "" {
		request.Auth = zabbix.Token
	}
	u := zabbix.URL + "/api_jsonrpc.php"
	jsonBytes, _ := json.Marshal(request)
	req, _ := http.NewRequest("POST", u, bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json-rpc")
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var response ZabbixResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (zabbix *Zabbix) Login() error {
	params := map[string]interface{}{
		"user":     zabbix.User,
		"password": zabbix.Password,
	}
	data := NewZabbixRequest("user.login", params)

	var response *ZabbixResponse
	response, err := zabbix.Do(data)
	if err != nil {
		return err
	}

	zabbix.Token = response.Result.(string)

	return nil
}

func (zabbix *Zabbix) Logout() error {
	return nil
}
