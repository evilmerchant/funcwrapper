package funcwrapper

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type handler struct {
	inner http.Handler
}

func Handler(innerHandler http.Handler) *handler {
	return &handler{inner: innerHandler}
}

func (h *handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	var input FunctionAppRequest
	content, err := ioutil.ReadAll(req.Body)
	if err != nil {
		createErrorResponse(err, res)
		return
	}

	err = json.Unmarshal(content, &input)
	if err != nil {
		createErrorResponse(err, res)
		return
	}

	data, err := json.Marshal(input.Data["req"])
	if err != nil {
		createErrorResponse(err, res)
		return
	}

	var httpTrigger HttpTriggerRequest
	err = json.Unmarshal(data, &httpTrigger)
	if err != nil {
		createErrorResponse(err, res)
		return
	}

	req.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(httpTrigger.Body)))
	for k := range req.Header {
		delete(req.Header, k)
	}

	for key, value := range httpTrigger.Headers {
		for _, v := range value {
			req.Header.Add(key, v)
		}
	}
	req.ContentLength = int64(len(httpTrigger.Body))

	rawUrl, err := url.Parse(httpTrigger.Url)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	rawUrl.Scheme = ""
	rawUrl.Host = ""

	req.URL = rawUrl
	req.Method = httpTrigger.Method
	writer := NewWriter()

	h.inner.ServeHTTP(writer, req)

	createHttpResponse(input.Metadata, *writer, res)

}
