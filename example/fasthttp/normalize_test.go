package fasthttp_test

import (
	"net"
	"testing"
	"time"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttputil"
)

func TestFasthttp_PathNormalize(t *testing.T) {
	t.Parallel()

	ln := fasthttputil.NewInmemoryListener()

	s := &fasthttp.Server{
		Handler: func(ctx *fasthttp.RequestCtx) {
			uri := ctx.URI()
			uri.DisablePathNormalizing = true
			ctx.Response.Header.Set("received-uri", string(uri.FullURI()))
		},
	}

	serverStopCh := make(chan struct{})
	go func() {
		if err := s.Serve(ln); err != nil {
			t.Errorf("unexpected error: %s", err)
		}
		close(serverStopCh)
	}()

	c := &fasthttp.Client{
		Dial: func(addr string) (net.Conn, error) {
			return ln.Dial()
		},
		DisablePathNormalizing: false,
	}

	// urlWithEncodedPath := "http://example.com/encoded/Y%2BY%2FY%3D/stuff"
	urlWithEncodedPath := "http://192.168.4.252:80/ESWHD_HD_NAT_16675_0_8106614643628609163/root]audio102/7957627962.ts?origin_url=aHR0cDovL2Nkbk1lYy1uYXAtMzA0LmNsYy1jaWYtZGNmLnhjci5jb21jYXN0Lm5ldC9FU1dIRF9IRF9OQVRfMTY2NzVfMF84MTA2NjE0NjQzNjI4NjA5MTYzL3Jvb3RfYXVkaW8xMDIvNzk1NzYyNzk2Mi50cw==&period_id=dVlwm4vNRJWkMwR3UKz8Qg&rep_id=root_audio102&start_pts=7957627962&start_time=2020-11-05T01:58:38Z&stream_id=8106614643628609163&no-redirect=true"

	var req fasthttp.Request
	fasthttp.AcquireRequest()
	req.SetRequestURI(urlWithEncodedPath)
	var resp fasthttp.Response
	for i := 0; i < 5; i++ {
		if err := c.DoTimeout(&req, &resp, time.Second); err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		hv := resp.Header.Peek("received-uri")
		if string(hv) != urlWithEncodedPath {
			t.Fatalf("request uri was normalized: %q. Expecting %q", hv, urlWithEncodedPath)
		}
	}

	if err := ln.Close(); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	select {
	case <-serverStopCh:
	case <-time.After(time.Second):
		t.Fatalf("timeout")
	}
}
