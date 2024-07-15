package network

import "github.com/gin-gonic/gin"

type Network struct {
	engine *gin.Engine
}

func NewNetwork() *Network {
	n := &Network{
		engine: gin.New(),
	}
	newUserRouter(n)
	return n
}

func (n *Network) ServerStart(port string) error {
	return n.engine.Run(port)
}

func (n *Network) registerGET(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.GET(path, handler...)
}

func (n *Network) registerPOST(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.POST(path, handler...)
}

func (n *Network) registerPUT(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.PUT(path, handler...)
}

func (n *Network) registerDELETE(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.DELETE(path, handler...)
}
