package http

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/valyala/fasthttp"
)

func readContent(resp *http.Response) ([]byte, error) {
	if resp == nil {
		return nil, errors.New("resp is nil")
	}

	contents, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	return contents, err
}

func TestServer(t *testing.T) {
	l, _ := net.Listen("tcp", "127.0.0.1:0")
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := []byte("hello")
		w.WriteHeader(http.StatusOK)
		w.Header().Add("x-count", fmt.Sprintf("%v", len(data)))
		w.Write(data)
	})

	s := &http.Server{Handler: handler}
	go s.Serve(l)

	resp, err := http.Get("http://" + l.Addr().String())
	log.Println(err, resp.Body)
	log.Println(resp.Header)

	bs, _ := readContent(resp)
	log.Printf("content: %s", string(bs))
}

func TestListenAndServe(t *testing.T) {
	server := &http.Server{Addr: ":10010"}
	errC := make(chan error)
	go func() { errC <- server.ListenAndServe() }()

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	select {
	case <-ctx.Done():
		log.Println("Hello world")
	case <-errC:
		log.Println("Hello error")
	}
}

// this blocks
/*
func TestListenAndServeToHang(t *testing.T) {
	server := &http.Server{Addr: ":10010"}
	err := server.ListenAndServe()

	log.Println(err)
}
*/

func TestResponseHeader(t *testing.T) {
	myHandler := func(w http.ResponseWriter, r *http.Request) {
		str := strings.Repeat("data", 1024*1024*8)
		data := []byte(str)
		w.Header().Add("Content-Length", fmt.Sprintf("%d", len(data)))
		w.Write(data)
	}
	server := &http.Server{
		Handler: http.HandlerFunc(myHandler),
	}
	err := server.ListenAndServe()

	log.Println("err: " + err.Error())
}

func TestAuthServer(t *testing.T) {
	// t.Skip()
	myHandler := func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		data := fmt.Sprintf("user: %v, pass: %v, ok: %v", user, pass, ok)
		w.Header().Add("Content-Length", fmt.Sprintf("%d", len(data)))
		w.Write([]byte(data))
	}
	var err error
	var server *http.Server
	readyCh := make(chan struct{}, 1)
	go func() {
		server = &http.Server{
			Addr:    "localhost:10011",
			Handler: http.HandlerFunc(myHandler),
		}
		close(readyCh)
		err = server.ListenAndServe()
	}()

	if err != nil {
		t.Fatalf("start server err: %v", err)
	}

	select {
	case <-readyCh:
	}

	// 1. default http client
	strUrl := fmt.Sprintf("http://user:password@%v", server.Addr)
	log.Printf("request: %v", strUrl)
	resp, err := http.Get(strUrl)
	bs, _ := readContent(resp)
	log.Printf("content: %s", string(bs))

	// 2. fasthttp client
	fClient := &fasthttp.Client{}
	auth, pureURL, _ := extractAuth(strUrl)
	fReq := fasthttp.AcquireRequest()
	fReq.SetRequestURI(pureURL)
	fReq.Header.Set("Authorization", "Basic "+auth)
	fResp := fasthttp.AcquireResponse()
	err = fClient.Do(fReq, fResp)
	statusCode := fResp.StatusCode()
	body := fResp.Body()
	log.Printf("pure url: %v", pureURL)
	log.Printf("fasthttp response -- statusCode: %v, body: %v, err: %v", statusCode, string(body), err)
	//
	// u, _ = url.Parse(strUrl)
	// log.Printf("parsed request: %v", st.String())
	// req.Header.Set("Authorization", "Basic "+basicAuth(username, password))

	time.Sleep(10 * time.Second)
}

func TestExtractAuth(t *testing.T) {
	addr := "localhost:8001"
	strUrl := fmt.Sprintf("http://user:password@%v", addr)
	log.Println(extractAuth(strUrl))
}

// extractAuth - extract auth info from url in form of `http://user:password@localhost`
// return - (auth, pureURL, error)
func extractAuth(rawurl string) (string, string, error) {
	scheme, rest, err := getscheme(rawurl)
	if err != nil {
		return "", rawurl, err
	}

	if !strings.HasPrefix(rest, "///") && strings.HasPrefix(rest, "//") {
		var authority string
		authority, rest = split(rest[2:], "/", false)
		i := strings.LastIndex(authority, "@")
		if i < 0 {
			return "", rawurl, nil
		}

		userinfo := authority[:i]
		if !strings.Contains(userinfo, ":") {
			return "", "", errors.New("net/url: invalid userinfo - " + userinfo)
		}
		username, password := split(userinfo, ":", true)
		if username == "" || password == "" {
			return "", "", errors.New("net/url: invalid userinfo - empty username or password")
		}

		pureURL := "//"
		if scheme != "" {
			pureURL = scheme + "://"
		}
		pureURL += authority[(i+1):] + rest
		return basicAuth(username, password), pureURL, nil
	}

	return "", rawurl, nil
}

// Maybe rawurl is of the form scheme:path.
// (Scheme must be [a-zA-Z][a-zA-Z0-9+-.]*)
// If so, return scheme, path; else return "", rawurl.
// NOTE: copy from `src/net/url/url.go`
func getscheme(rawurl string) (scheme, path string, err error) {
	for i := 0; i < len(rawurl); i++ {
		c := rawurl[i]
		switch {
		case 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z':
		// do nothing
		case '0' <= c && c <= '9' || c == '+' || c == '-' || c == '.':
			if i == 0 {
				return "", rawurl, nil
			}
		case c == ':':
			if i == 0 {
				return "", "", errors.New("missing protocol scheme")
			}
			return rawurl[:i], rawurl[i+1:], nil
		default:
			// we have encountered an invalid character,
			// so there is no valid scheme
			return "", rawurl, nil
		}
	}
	return "", rawurl, nil
}

// Maybe s is of the form t c u.
// If so, return t, c u (or t, u if cutc == true).
// If not, return s, "".
// NOTE: copy from `src/net/url/url.go`
func split(s string, c string, cutc bool) (string, string) {
	i := strings.Index(s, c)
	if i < 0 {
		return s, ""
	}
	if cutc {
		return s[:i], s[i+len(c):]
	}
	return s[:i], s[i:]
}

// See 2 (end of page 4) https://www.ietf.org/rfc/rfc2617.txt
// "To receive authorization, the client sends the userid and password,
// separated by a single colon (":") character, within a base64
// encoded string in the credentials."
// It is not meant to be urlencoded.
// NOTE: copy from `src/net/http/client.go`
func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
