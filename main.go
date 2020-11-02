package main

import (
	"flag"
	"fmt"
	"github.com/sh4yy/GoHTTP/HTTP"
	"github.com/sh4yy/GoHTTP/Utils"
)



func main() {

	var headers Utils.ListFlags
	url := flag.String("url", "", "Target URL")
	method := flag.String("method", "GET", "Request Method")
	body := flag.String("body", "", "Request Body")
	flag.Var(&headers, "header", "Request Headers")
	flag.Parse()

	if *url == "" {
		panic("missing url")
	}

	request := HTTP.NewRequest(*url, *method)
	err := request.WriteRawHeaders(headers)
	if err != nil {
		panic(err)
	}

	request.WriteHeader("Connection", "close")
	request.WriteStringBody(*body)

	client := HTTP.NewClient()
	res, err := client.Send(request)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(res))

}
