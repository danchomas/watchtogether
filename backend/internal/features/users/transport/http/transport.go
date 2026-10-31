package users_transport_http

type UsersHTTPHandler struct {
	userService UserService
}

type UserService interface {
}

func NewUsersHttpHandler(
	usersService UserService,
) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		userService: usersService,
	}
}
