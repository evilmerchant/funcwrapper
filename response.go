package funcwrapper

import (
	"encoding/json"
	"net/http"
)

func writeResponse(v any, writer http.ResponseWriter) {
	marshalled, _ := json.Marshal(v)
	writer.Header().Add("content-type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(marshalled)
}

func createHttpResponse(metadata Metadata, inner ResponseWriter, writer http.ResponseWriter) {

	response := FuncResponse{}
	response.Outputs = make(map[string]HttpTriggerResponse)
	responseHeaders := make(map[string]string)
	for k, h := range inner.header {
		for _, h2 := range h {
			responseHeaders[k] = h2
		}
	}
	triggerResponse := HttpTriggerResponse{
		Headers:    responseHeaders,
		Body:       string(inner.body),
		StatusCode: inner.statusCode,
	}
	response.Outputs["res"] = triggerResponse
	writeResponse(response, writer)
}

func createErrorResponse(err error, writer http.ResponseWriter) {
	response := FuncResponse{}
	response.Outputs = make(map[string]HttpTriggerResponse)
	resp := &ErrorResponse{Error: err.Error()}
	content, _ := json.Marshal(resp)

	triggerResponse := HttpTriggerResponse{
		Body:       string(content),
		StatusCode: http.StatusInternalServerError,
	}
	response.Outputs["res"] = triggerResponse
	writeResponse(response, writer)
}
