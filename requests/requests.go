package requests

import (
	"context"

	"github.com/carlmjohnson/requests"
)

type Request struct {
	Uri    string
	Route  string
	Params map[string]string

	protocol string

	emptyParams map[string]string
}

func Initialize(prot, uri string) *Request {
	return &Request{
		protocol: prot,
		Uri:      uri,
		Route:    "",
		Params:   make(map[string]string),

		emptyParams: make(map[string]string), // don't allocate new memory for each request
	}
}

func (req *Request) SetUri(uri string) {
	req.Uri = uri
}

func (req *Request) Get(route string) (string, error) {
	return req.GetRequest(route, req.emptyParams)
}

func (req *Request) GetRequest(route string, params map[string]string) (string, error) {
	var resp string

	domain := req.Uri

	if len(route) > 0 {
		domain += "/" + route
	}

	r := requests.
		URL(domain).
		Scheme(req.protocol)

	if len(params) > 0 {
		for key, value := range params {
			r = r.Param(key, value)

		}
	}

	err := r.CheckStatus(200).
		ToString(&resp).
		Fetch(context.Background())

	if err != nil {
		println(err.Error())
	} else {
		println(resp)
	}

	return resp, err
}
