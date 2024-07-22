package network

import (
	"sync"
	"web/service"
	"web/types"

	"github.com/gin-gonic/gin"
)

var (
	userRouterInit     sync.Once // 외부에서 2번이상 호출되도 최초의 호출에서만 실행됨
	userRouterInstance *userRouter
)

type userRouter struct {
	router  *Network
	service *service.UserService
}

func newUserRouter(router *Network, userService *service.UserService) *userRouter {
	userRouterInit.Do(func() {
		userRouterInstance = &userRouter{
			router:  router,
			service: userService,
		}

		router.registerPOST("/user", userRouterInstance.create)
		router.registerGET("/user", userRouterInstance.get)
		router.registerPUT("/user", userRouterInstance.update)
		router.registerDELETE("/user", userRouterInstance.delete)
	})

	return userRouterInstance
}

func (u *userRouter) create(c *gin.Context) {
	var req types.CreateUserRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		u.router.failedResponse(c, &types.CreateUserResponse{
			ApiResponse: types.NewApiResponse(-1, err.Error()),
		})
	} else if err = u.service.Create(&types.User{Name: req.Name, Age: req.Age}); err != nil {
		u.router.failedResponse(c, &types.CreateUserResponse{
			ApiResponse: types.NewApiResponse(-1, err.Error()),
		})
	} else {
		u.router.okResponse(c, &types.CreateUserResponse{
			ApiResponse: types.NewApiResponse(1, "Success"),
		})
	}
}

func (u *userRouter) get(c *gin.Context) {
	u.router.okResponse(c, &types.GetUserResponse{
		ApiResponse: types.NewApiResponse(1, "Success"),
		Users:       u.service.Get(),
	})
}

func (u *userRouter) update(c *gin.Context) {
	var req types.UpdateUserRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		u.router.failedResponse(c, &types.UpdateUserResponse{
			ApiResponse: types.NewApiResponse(-1, err.Error()),
		})
	} else if err = u.service.Update(req.Name, req.UpdateAge); err != nil {
		u.router.failedResponse(c, &types.UpdateUserResponse{
			ApiResponse: types.NewApiResponse(-1, err.Error()),
		})
	} else {
		u.router.okResponse(c, &types.UpdateUserResponse{
			ApiResponse: types.NewApiResponse(1, "Success"),
		})
	}
}

func (u *userRouter) delete(c *gin.Context) {
	var req types.DeleteUserRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		u.router.failedResponse(c, &types.DeleteUserResponse{
			ApiResponse: types.NewApiResponse(-1, err.Error()),
		})
	} else if err = u.service.Delete(req.Name); err != nil {
		u.router.failedResponse(c, &types.DeleteUserResponse{
			ApiResponse: types.NewApiResponse(-1, err.Error()),
		})
	} else {
		u.router.okResponse(c, &types.DeleteUserResponse{
			ApiResponse: types.NewApiResponse(1, "Success"),
		})
	}
}
