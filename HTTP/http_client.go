package HTTP

import (
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

func (http HTTPClient) GET(url string, headers *Headers) ([]byte, error) {

	request := NewRequest(url, "GET")
	request.WriteHeader("User-Agent", http.UserAgent)
	request.WriteHeader("Accept", "*/*")
	request.WriteHeader("Host", request.GetHostname())
	request.WriteHeader("Connection", "close")

	if headers != nil {
		for key, value := range *headers {
			request.WriteHeader(key, value)
		}
	}

	return http.Send(request)
}

func (http HTTPClient) POST(url string, headers *Headers, body []byte) ([]byte, error) {

	request := NewRequest(url, "POST")
	request.WriteHeader("User-Agent", http.UserAgent)
	request.WriteHeader("Accept", "*/*")
	request.WriteHeader("Host", request.GetHostname())
	request.WriteHeader("Connection", "close")

	if body != nil {
		request.WriteBody(body)
	}

	if headers != nil {
		for key, value := range *headers {
			request.WriteHeader(key, value)
		}
	}

	return http.Send(request)

}
