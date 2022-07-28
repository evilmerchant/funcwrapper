package funcwrapper

import (
	"net/http"
	"time"
)

type FuncResponse struct {
	Outputs     map[string]HttpTriggerResponse `json:"Outputs"`
	Logs        []string                       `json:"Logs"`
	ReturnValue string                         `json:"ReturnValue"`
}

type HttpTriggerResponse struct {
	Body       string            `json:"body"`
	StatusCode int               `json:"statusCode"`
	Headers    map[string]string `json:"headers"`
}

type FunctionAppSystem struct {
	MethodName string    `json:"MethodName"`
	UtcNow     time.Time `json:"UtcNow"`
}

type Data = map[string]interface{}
type Metadata struct {
	System FunctionAppSystem `json:"sys"`
}
type FunctionAppRequest struct {
	Data     Data     `json:"Data"`
	Metadata Metadata `json:"Metadata"`
}
type HttpTriggerRequest struct {
	Url     string      `json:"Url"`
	Method  string      `json:"Method"`
	Headers http.Header `json:"Headers"`
	Body    string      `json:"Body"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
