package types

type User struct {
	Name string
	Age  int64
}

type GetUserResponse struct {
	*ApiResponse
	Users []*User
}

type CreateUserRequest struct {
	Name string `json "name" binding:"required"`
	Age  int64  `json "age" binding:"required"`
}

type CreateUserResponse struct {
	*ApiResponse
}

type UpdateUserRequest struct {
	Name      string `json "name" binding:"required"`
	UpdateAge int64  `json "age" binding:"required"`
}
type UpdateUserResponse struct {
	*ApiResponse
}

type DeleteUserRequest struct {
	Name string `json "name" binding:"required"`
}
type DeleteUserResponse struct {
	*ApiResponse
}
