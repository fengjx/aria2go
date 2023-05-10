package aria2go

import "testing"

var testCli = NewClient("1024", "http://localhost:6800/jsonrpc")

func TestAddDownload(t *testing.T) {
	opt := &Options{
		Dir: "/Users/fengjianxin/Downloads",
	}
	resp, err := testCli.AddDownload("https://github.com/fengjx/java-hot-reload-agent/releases/download/hot-reload-agent-all-1.1.0/hot-reload-agent-bin.zip", opt)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf(resp.Status)
}
