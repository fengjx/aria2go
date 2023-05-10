package aria2go

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client aira2 client
type Client struct {
	token      string
	jsonRpcUrl string
	httpClient *http.Client
}

type Options struct {
	Out       string `json:"out,omitempty"`
	Dir       string `json:"dir,omitempty"`
	AllProxy  string `json:"all-proxy,omitempty"`
	Referer   string `json:"referer,omitempty"`
	UserAgent string `json:"user-agent,omitempty"`
}

// NewClient create aria2 client
func NewClient(token string, jsonRpcUrl string) *Client {
	httpClient := &http.Client{
		Timeout: time.Second * 3,
	}
	cli := &Client{
		token:      token,
		jsonRpcUrl: jsonRpcUrl,
		httpClient: httpClient,
	}
	return cli
}

// Result rpc result
// {"id":"dmuITXlBkCRenRIr","jsonrpc":"2.0","result":"956724f4b4bad3b0"}
type Result struct {
	Id      string `json:"id,omitempty"`
	Jsonrpc string `json:"jsonrpc,omitempty"`
	Result  string `json:"result,omitempty"`
}

func (cli *Client) DoRequest(method string, params []interface{}) (*Result, error) {
	id := RandString(16)
	token := fmt.Sprintf("token:%s", cli.token)
	paramArr := []interface{}{
		token,
	}
	for _, param := range params {
		paramArr = append(paramArr, param)
	}
	data := map[string]interface{}{
		"id":      id,
		"jsonrpc": "2.0",
		"method":  method,
		"params":  paramArr,
	}
	bys, _ := json.Marshal(data)
	resp, err := cli.httpClient.Post(cli.jsonRpcUrl, "application/json", bytes.NewReader(bys))
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	res := &Result{}
	err = json.Unmarshal(body, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// AddDownload add download task
//
//	{
//	   "jsonrcp": "2.0",
//	   "id": "someID",
//	   "method": "aria2.addUri",
//	   "params": [
//	       "token:tokenString",
//	       [
//	           "https://github.com/fengjx/java-hot-reload-agent/releases/download/hot-reload-agent-all-1.1.0/hot-reload-agent-bin.zip"
//	       ],
//	       {
//	           "dir": "/path/to/dir"
//	       }
//	   ]
//	}
func (cli *Client) AddDownload(url string, opt *Options) (*Result, error) {
	params := []interface{}{
		[]string{url},
		opt,
	}
	return cli.DoRequest("aria2.addUri", params)
}
