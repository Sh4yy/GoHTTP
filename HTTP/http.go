package HTTP

type HTTP interface {
	GET(url string) ([]byte, error)
}
