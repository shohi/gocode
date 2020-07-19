package toxi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/Shopify/toxiproxy"
	"github.com/Shopify/toxiproxy/toxics"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
}

func doGet(reqUrl string) {
	resp, err := http.Get(reqUrl)
	if err != nil {
		panic(err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	content, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("=====> Get content: %v, err: %v\n",
		string(content), err)

}

func TestToxiServer(t *testing.T) {
	srv := newTestServer(18080)
	defer srv.Close()

	doGet("http://localhost:18080/hello")
}

func TestToxicProxy_Basic(t *testing.T) {
	proxy := toxiproxy.NewProxy()
	proxy.Name = "http-proxy"
	proxy.Listen = ":18081"
	proxy.Upstream = ":18080"
	proxy.Start()
	defer proxy.Stop()

	srv := newTestServer(18080)
	defer srv.Close()

	doGet("http://localhost:18081/hello")
}

func ToxicToJson(name, typeName, stream string, toxic toxics.Toxic) io.Reader {
	data := map[string]interface{}{
		"name":       name,
		"type":       typeName,
		"stream":     stream,
		"attributes": toxic,
	}
	request, err := json.Marshal(data)
	if err != nil {
		msg := fmt.Sprintf("Failed to marshal toxic for api (1): %v", toxic)
		panic(msg)
	}

	return bytes.NewReader(request)
}

func TestToxicProxy_Timeout(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	proxy := toxiproxy.NewProxy()
	proxy.Name = "http-proxy"
	proxy.Listen = ":18081"
	proxy.Upstream = ":18080"
	proxy.Start()

	tx := &toxics.TimeoutToxic{
		Timeout: 1000, // millisecond
	}
	reader := ToxicToJson("timeout-up", "timeout", "downstream", tx)
	proxy.Toxics.AddToxicJson(reader)
	defer proxy.Stop()

	srv := newTestServer(18080)
	defer srv.Close()

	doGet("http://localhost:18081/hello")
}
