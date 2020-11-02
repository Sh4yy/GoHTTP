package HTTP

type Headers map[string]string

type HTTP interface {
	GET(url string, headers *Headers) ([]byte, error)
	POST(url string, headers *Headers, body []byte) ([]byte, error)
}
