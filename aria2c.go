package aria2go

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func (cli *Client) DoRequest(method string, params []interface{}) (*http.Response, error) {
	id := time.Now().UnixNano()
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
	return cli.httpClient.Post(cli.jsonRpcUrl, "application/json", bytes.NewReader(bys))
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
//	           "http://m.gettywallpapers.com/wp-content/uploads/2020/01/Wallpaper-Naruto-2.jpg"
//	       ],
//	       {
//	           "out": "test.jpg"
//	       }
//	   ]
//	}
func (cli *Client) AddDownload(url string, opt *Options) (*http.Response, error) {
	params := []interface{}{
		[]string{url},
		opt,
	}
	return cli.DoRequest("aria2.addUri", params)
}
