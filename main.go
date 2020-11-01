package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/sh4yy/GoTestHTTP/HTTP"
)

func ParseURLArg() (string, error) {

	var url string
	flag.StringVar(&url, "url", "", "URL to test against")
	flag.Parse()

	if url == "" {
		return "", errors.New("empty url")
	}

	fmt.Println(url)
	return url, nil

}

func main() {

	url, err := ParseURLArg()
	if err != nil {
		panic(err)
	}

	client := HTTP.NewClient()
	res, err := client.GET(url)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(res))

}
