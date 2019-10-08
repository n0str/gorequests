package http_session

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

type Header map[string]string

type HTTPSession struct {
	CookieHeader string
	Cookies      []*http.Cookie
}

type Response struct {
	Raw  *http.Response
	Body []byte
}

func New() *HTTPSession {
	return &HTTPSession{}
}

func (httpSession *HTTPSession) EasyRequest(method string, urlString string, vs ...interface{}) *Response {
	r, _ := httpSession.Request(method, urlString, vs)
	return r
}

func (httpSession *HTTPSession) Request(method string, urlString string, vs ...interface{}) (*Response, error) {
	req := &http.Request{
		Method: method,
		Header: make(http.Header),
		Proto:  "HTTP/1.1",
	}
	u, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}
	req.URL = u

	for _, v := range vs {
		switch vv := v.(type) {
		case Header:
			for key, value := range vv {
				req.Header.Add(key, value)
			}
		case http.Header:
			for key, values := range vv {
				for _, value := range values {
					req.Header.Add(key, value)
				}
			}
		case *http.Cookie:
			req.AddCookie(vv)
		case http.Cookie:
			req.AddCookie(&vv)
		}
	}

	if _, ok := req.Header["Cookie"]; !ok {
		req.Header.Set("Cookie", httpSession.CookieHeader)
	}

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	cookie := response.Header.Get("Set-Cookie")
	if cookie != "" {
		httpSession.CookieHeader = cookie
		httpSession.Cookies = response.Cookies()
	}

	resp := &Response{Raw: response}
	resp.Body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
