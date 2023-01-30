package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/yimsoijoi/drop/internal/common"
	"github.com/yimsoijoi/drop/internal/domain"
)

type restfulDelivery struct {
	usecase domain.Usecase
}

func NewRestfulDelivery(engine *gin.Engine, usecase domain.Usecase) {
	restful := &restfulDelivery{
		usecase: usecase,
	}

	{
		engine.POST(common.RestfulPathDrop, restful.CreatTextHandler)
		engine.GET("/:key"+common.RestfulPathDrop, restful.GetTextHandler)
	}
}
