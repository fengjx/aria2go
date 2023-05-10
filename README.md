# aria2go

aria2 go client

## usage

install

```bash
go get github.com/fengjx/aria2go
```

```go
aria2Cli := aria2go.NewClient("1024", "http://localhost:6800/jsonrpc")
opt := &Options{
    Dir: "/path/to/download/dir",
}
resp, err := aria2Cli.AddDownload("https://github.com/fengjx/java-hot-reload-agent/releases/download/hot-reload-agent-all-1.1.0/hot-reload-agent-bin.zip", opt)
```

## doc

```
package aria2go // import "github.com/fengjx/aria2go"

TYPES

func NewClient(token string, jsonRpcUrl string) *Client
    NewClient create aria2 client

func (cli *Client) AddDownload(url string, opt *Options) (*http.Response, error)
    AddDownload add download task
```

