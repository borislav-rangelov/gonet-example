package home

import (
	"net/http"

	"github.com/borislav-rangelov/gonet/response"
)

// HelloHandler Responds with hello
func HelloHandler(resp http.ResponseWriter, req *http.Request) {
	body := map[string]string{"body": "Hello!"}
	response.WriteJSON(http.StatusOK, body, resp)
}
