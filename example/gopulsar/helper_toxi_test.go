package gopulsar

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/Shopify/toxiproxy"
	"github.com/Shopify/toxiproxy/toxics"
	"github.com/apache/pulsar-client-go/pulsar"
)

// copy from toxics/toxic_test.go
// Usage:
// 	    reader := ToxicToJson(
// 					"bandwidth-up",
// 					"bandwidth",
// 					"upstream",
// 					&toxics.BandwidthToxic{
// 						Rate: 100 * 1000,
// 					}))
//
// 		proxy.Toxics.AddToxicJson(reader)
//
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

// Available type name
// - noop
// - slicer
// - timeout
// - bandwidth
// - latency
// - limit_data
func getToxicType(tx toxics.Toxic) string {
	switch tx.(type) {
	case *toxics.NoopToxic:
		return "noop"
	case *toxics.SlicerToxic:
		return "slicer"
	case *toxics.TimeoutToxic:
		return "timeout"
	case *toxics.BandwidthToxic:
		return "bandwidth"
	case *toxics.LatencyToxic:
		return "latency"
	case *toxics.LimitDataToxic:
		return "limit_data"
	default:
		panic("Unknown toxic type")
	}
}

func getStreamDirect(stream string) string {
	tmp := strings.ToLower(stream)

	if strings.HasPrefix(tmp, "down") {
		return "downstream"
	}

	return "upstream"
}

func NewTestProxy(name, upstream string) *toxiproxy.Proxy {
	proxy := toxiproxy.NewProxy()

	proxy.Name = name
	proxy.Listen = "localhost:0"
	proxy.Upstream = upstream

	return proxy
}

type ToxicOptions struct {
	Name   string
	Stream string // upstream/downstream
	Toxic  toxics.Toxic
}

func toxicToReader(opts ToxicOptions) io.Reader {
	toxicType := getToxicType(opts.Toxic)
	direct := getStreamDirect(opts.Stream)

	return ToxicToJson(opts.Name, toxicType, direct, opts.Toxic)
}

func newTestProxyWithToxic(name, upstream string, opts ToxicOptions) *toxiproxy.Proxy {
	proxy := NewTestProxy(name, upstream)
	proxy.Start()
	// defer proxy.Stop()

	if opts.Toxic == nil {
		return proxy
	}

	reader := toxicToReader(opts)
	_, err := proxy.Toxics.AddToxicJson(reader)
	if err != nil {
		panic(err)
	}

	return proxy
}

func newTestClientViaProxy(clientOpts ClientOptions, txOpts ToxicOptions) (*toxiproxy.Proxy, pulsar.Client) {
	proxy := newTestProxyWithToxic("pulsar-proxy",
		"localhost:6650", txOpts)

	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               fmt.Sprintf("pulsar://%s", proxy.Listen),
		ConnectionTimeout: clientOpts.DialTimeout,
		OperationTimeout:  clientOpts.OpTimeout,
	})

	if err != nil {
		panic(err)
	}

	return proxy, client
}

// update or add toxic
// FIXME: not work, only for same type
func updateProxyToxic(proxy *toxiproxy.Proxy, opts ToxicOptions) {
	var err error
	defer func() {
		if err != nil {
			panic(err)
		}
	}()

	reader := toxicToReader(opts)

	w := proxy.Toxics.GetToxic(opts.Name)
	if w == nil {
		_, err = proxy.Toxics.AddToxicJson(reader)
		return
	}

	if reflect.TypeOf(w.Toxic) == reflect.TypeOf(opts.Toxic) {
		_, err = proxy.Toxics.UpdateToxicJson(opts.Name, reader)
		return
	}

	// remove old and add new
	err = proxy.Toxics.RemoveToxic(opts.Name)
	if err != nil {
		return
	}

	_, err = proxy.Toxics.AddToxicJson(reader)
}

func removeProxyToxic(proxy *toxiproxy.Proxy, opts ToxicOptions) {
	err := proxy.Toxics.RemoveToxic(opts.Name)
	if err != nil {
		panic(err)
	}
}

func nopToxicOptions() ToxicOptions {
	return ToxicOptions{
		Name:   "nop",
		Stream: "up",
		Toxic:  nil,
	}
}
