package HTTP

import (
	"bytes"
	"fmt"
	"github.com/sh4yy/GoHTTP/Utils"
	"net"
	NetURL "net/url"
	"strconv"
)

var (
	PortMap = map[string]string{
		"http":  "80",
		"https": "443",
	}
	CLRF = []byte("\r\n")
)

type Request struct {
	Protocol string
	URL		 string
	Method   string
	Headers  map[string]string
	Body     []byte
}

func NewRequest(url, method string) *Request {
	return &Request{
		Protocol: "HTTP/1.1",
		URL: url,
		Method: method,
		Headers: map[string]string{},
	}
}

func (request *Request) WriteBody(body []byte) {
	request.Body = body
	request.WriteHeader("Content-Length", strconv.Itoa(len(body)))
}

func (request *Request) WriteStringBody(body string) {
	if body != "" {
		request.WriteBody([]byte(body))
	}
}

func (request *Request) WriteHeader(key, value string) {
	request.Headers[key] = value
}

func (request *Request) WriteRawHeader(header string) error {
	key, value, error := Utils.ParseHeader(header)
	if error != nil {
		return error
	}
	request.WriteHeader(key, value)
	return nil
}

func (request *Request) WriteRawHeaders(headers []string) error {
	for _, header := range headers {
		err := request.WriteRawHeader(header)
		if err != nil {
			return err
		}
	}
	return nil
}

// urlParser parses url and returns (hostname, port, path, scheme)
func urlParser(url string) (string, string, string, string) {
	parsedURL, err := NetURL.Parse(url)
	if err != nil {
		panic(err)
	}

	hostname := parsedURL.Hostname()
	port := parsedURL.Port()
	path := parsedURL.RawPath
	scheme := parsedURL.Scheme
	return hostname, port, path, scheme
}

func (request Request) GetHostname() string {
	hostname, _, _, _ := urlParser(request.URL)
	return hostname
}

func (request Request) GetPort() string {
	_, port, _, scheme := urlParser(request.URL)
	if port == "" {
		port = PortMap[scheme]
	}
	return port
}

func (request Request) GetPath() string {
	_, _, path, _ := urlParser(request.URL)
	if path == "" {
		path = "/"
	}
	return path
}

func (request Request) GetAddress() string {
	return net.JoinHostPort(request.GetHostname(), request.GetPort())
}

func (request Request) Serialize() (*bytes.Buffer, error) {

	var buf bytes.Buffer

	// write headers
	_, err := fmt.Fprintf(&buf, "%s %s %s", request.Method, request.GetPath(), request.Protocol)
	if err != nil {
		return nil, err
	}
	buf.Write(CLRF)

	for key, value := range request.Headers {
		fmt.Fprintf(&buf, "%s: %s", key, value)
		buf.Write(CLRF)
	}

	// terminate headers
	buf.Write(CLRF)

	// write Body
	if len(request.Body) != 0 {
		buf.Write(request.Body)
	}

	return &buf, nil

}