package users

import (
	"net/http"

	"github.com/borislav-rangelov/gonet/request"
	"github.com/gorilla/mux"

	"github.com/borislav-rangelov/gonet/response"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var users []User

func init() {
	users = append(users,
		User{FirstName: "John", LastName: "Doe"},
		User{FirstName: "Jane", LastName: "Doe"})
}

// GetUsersPage returns a page of users
func GetUsersPage(resp http.ResponseWriter, req *http.Request) {
	page := request.MapPage(mux.Vars(req))
	start := page.GetPage() * page.GetSize()
	end := start + page.GetSize()

	if end > len(users) {
		end = len(users)
	}

	var content []interface{}
	if page.GetSize() > 0 && start < len(users) {
		content = append(content, users[start:end])
	}

	result := response.BuildPage(&content, page, int64(len(users)))

	response.WriteJSON(http.StatusOK, result, resp)
}
