package HTTP

import (
	"fmt"
	"io/ioutil"
	"net"
)

type HTTPClient struct {
	Protocol  string
	UserAgent string
}

func NewClient() *HTTPClient {
	return &HTTPClient{
		Protocol:  "HTTP/1.0",
		UserAgent: "Go/1.0",
	}
}


func (http HTTPClient) Send(request *Request) ([]byte, error) {

	buf, err := request.Serialize()
	if err != nil {
		return nil, err
	}

	fmt.Println(buf)
	conn, err := net.Dial("tcp", request.GetAddress())
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
	request := NewRequest(url, "GET")
	request.WriteHeader("User-Agent", http.UserAgent)
	request.WriteHeader("Accept", "*/*")
	request.WriteHeader("Host", request.GetHostname())
	request.WriteHeader("Connection", "close")
	return http.Send(request)
}
