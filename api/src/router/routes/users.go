package routes

import (
	"api/src/controllers"
	"net/http"
)

var routes = []Route{
	{
		URI:       "/users",
		Method:    http.MethodPost,
		Function:  controllers.CreateUser,
		NeedsAuth: false,
	},
	{
		URI:       "/users",
		Method:    http.MethodGet,
		Function:  controllers.GetUsers,
		NeedsAuth: false,
	},
	{
		URI:       "/users/{userID}",
		Method:    http.MethodGet,
		Function:  controllers.GetUser,
		NeedsAuth: false,
	},
	{
		URI:       "/users/{userID}",
		Method:    http.MethodPut,
		Function:  controllers.UpdateUser,
		NeedsAuth: false,
	},
	{
		URI:       "/users/{userID}",
		Method:    http.MethodDelete,
		Function:  controllers.DeleteUser,
		NeedsAuth: false,
	},
}
