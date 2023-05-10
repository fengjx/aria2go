# aria2go
aria2 go client

## doc

```
package aria2go // import "github.com/fengjx/aria2go"

TYPES

func NewClient(token string, jsonRpcUrl string) *Client
    NewClient create aria2 client

func (cli *Client) AddDownload(url string, opt *Options) (*http.Response, error)
    AddDownload add download task
```

