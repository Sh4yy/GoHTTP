package HTTP

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	netUrl "net/url"
)

var (
	clrf = []byte("\r\n")
)

type HTTPClient struct {
	Protocol  string
	UserAgent string
}

type Request struct {
	Address string
	Path    string
	Method  string
	Headers map[string]string
	Body    []byte
}

func NewClient() *HTTPClient {
	return &HTTPClient{
		Protocol:  "HTTP/1.1",
		UserAgent: "GoClient",
	}
}

func (http HTTPClient) constructRequest(request *Request) (*bytes.Buffer, error) {

	var buf bytes.Buffer

	// write headers
	_, err := fmt.Fprintf(&buf, "%s %s %s", request.Method, request.Path, http.Protocol)
	if err != nil {
		return nil, err
	}
	buf.Write(clrf)

	for key, value := range request.Headers {
		fmt.Fprintf(&buf, "%s: %s", key, value)
		buf.Write(clrf)
	}

	// terminate headers
	buf.Write(clrf)

	// write Body
	if len(request.Body) != 0 {
		buf.Write(request.Body)
	}

	return &buf, nil

}

func (http HTTPClient) Send(request *Request) ([]byte, error) {

	buf, err := http.constructRequest(request)
	if err != nil {
		return nil, err
	}

	fmt.Println(buf)
	conn, err := net.Dial("tcp", request.Address)
	if err != nil {
		return nil, err
	}

	_, err = conn.Write(buf.Bytes())
	if err != nil {
		return nil, err
	}

	res, err := ioutil.ReadAll(conn)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func (http HTTPClient) GET(url string) ([]byte, error) {

	parsedURL, err := netUrl.Parse(url)
	if err != nil {
		return nil, err
	}

	rawPath := parsedURL.RawPath
	if rawPath == "" {
		rawPath = "/"
	}

	request := Request{
		Address: parsedURL.Host,
		Path:    rawPath,
		Method:  "GET",
		Headers: map[string]string{
			"User-Agent": http.UserAgent,
			"Accept":     "*/*",
			"Host":       parsedURL.Host,
		},
	}

	return http.Send(&request)

}
