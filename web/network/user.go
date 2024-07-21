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
	service *service.User
}

func newUserRouter(router *Network, userService *service.User) *userRouter {
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
	c.JSON(200, "success")
}

func (u *userRouter) get(c *gin.Context) {
	u.router.okResponse(c, &types.UserResponse{
		ApiResponse: types.NewApiResponse(1, "Success"),
		User:        types.User{},
	})
}

func (u *userRouter) update(c *gin.Context) {

}

func (u *userRouter) delete(c *gin.Context) {

}
