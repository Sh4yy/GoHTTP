package Utils

import (
	"errors"
	"strings"
)

type ListFlags []string

func (list *ListFlags) String() string {
	return "my string representation"
}

func (list *ListFlags) Set(value string) error {
	*list = append(*list, value)
	return nil
}

func ParseHeader(header string) (string, string, error) {

	headerSlice := strings.SplitN(header, " ", 2)
	if len(headerSlice) != 2 {
		return "", "", errors.New("invalid header")
	}

	key := strings.Replace(headerSlice[0], ":", "", 1)
	value := headerSlice[1]
	return key, value, nil

}