package telegram

import (
	"net/http"
	"net/url"
	"path"
	"strconv"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

func New(host string, token string) Client {
	return Client{
		host:     host,
		basePath: newBasePath(token),
		client:   http.Client{}
	}
}

func newBasePath(token string) {
	return "bot" + token
}

func (c Client*) Updates(offset int, limit int) (Update[], err)  {
	q := url.Values()
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))

	//do request
}

func (c Client*) doRequest(method string, query url.Values) ([]byte, err) {
	u:=url.URL {
		Scheme: "https",
		Host: c.host,
		Path: path.Join(c.basePath, method),
	}

	req, err = http.NewRequest("GET")
}

func (c Client*) SendMessage() {

}
