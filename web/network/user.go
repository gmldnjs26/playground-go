package network

import (
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	userRouterInit     sync.Once // 외부에서 2번이상 호출되도 최초의 호출에서만 실행됨
	userRouterInstance *userRouter
)

type userRouter struct {
	router *Network
}

func newUserRouter(router *Network) *userRouter {
	userRouterInit.Do(func() {
		userRouterInstance = &userRouter{
			router: router,
		}

		router.registerPOST("/user", userRouterInstance.create)
		router.registerGET("/user", userRouterInstance.get)
		router.registerPUT("/user", userRouterInstance.update)
		router.registerDELETE("/user", userRouterInstance.delete)
	})

	return userRouterInstance
}

func (u *userRouter) create(c *gin.Context) {

}

func (u *userRouter) get(c *gin.Context) {

}

func (u *userRouter) update(c *gin.Context) {

}

func (u *userRouter) delete(c *gin.Context) {

}
