package funcwrapper

import "net/http"

type ResponseWriter struct {
	body       []byte
	statusCode int
	header     http.Header
}

func (rw *ResponseWriter) Header() http.Header {
	return rw.header
}

func (rw *ResponseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
}

func (rw *ResponseWriter) Write(data []byte) (int, error) {
	rw.body = data
	return 0, nil
}

func NewWriter() *ResponseWriter {
	return &ResponseWriter{
		header: http.Header{},
	}
}
