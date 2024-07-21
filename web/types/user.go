package types

type User struct {
	Name string
	Age  int64
}

type UserResponse struct {
	*ApiResponse
	User
}
