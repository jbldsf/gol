package service

import (
	"io"
	"net/http"
	"net/url"
)

type country struct {
	Area       int    `json:"area"`
	ID         string `json:"id"`
	Name       string `json:"name"`
	Population int    `json:"population"`
}

func Country(r *http.Request, callback func(statusCode int, response []byte)) {
	switch r.Method {
	case "POST":
		(&country{}).post(r.Body, callback)
	case "GET":
		(&country{}).get(r.URL.Query(), callback)
	case "PUT":
		(&country{}).put(r.Body, callback)
	case "DELETE":
		(&country{}).delete(r.URL.Query(), callback)
	default:
		callback(http.StatusMethodNotAllowed, []byte("{}"))
	}
}

func (c *country) post(body io.ReadCloser, callback func(statusCode int, response []byte)) {}

func (c *country) get(query url.Values, callback func(statusCode int, response []byte)) {}

func (c *country) put(body io.ReadCloser, callback func(statusCode int, response []byte)) {}

func (c *country) delete(query url.Values, callback func(statusCode int, response []byte)) {}
