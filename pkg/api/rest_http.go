package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Api struct {
	Engine *gin.Engine
	Group  gin.RouterGroup
}

func NewApi() *Api {
	router := gin.Default()
	Group := router.Group("/v1")
	return &Api{
		Engine: router,
		Group:  *Group,
	}
}

func (a *Api) InitGin(port int) error {
	var err error
	if port != 0 {
		err = a.Engine.Run(fmt.Sprintf(":%d", port))
	} else {
		err = a.Engine.Run(":8080")
	}
	return err
}
