package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/yimsoijoi/drop/internal/domain"
	error2 "github.com/yimsoijoi/drop/lib/error"
	"net/http"
)

func (restful restfulDelivery) CreatTextHandler(c *gin.Context) {
	model := new(domain.CreateText)
	if err := c.BindJSON(model); err != nil {
		c.JSON(http.StatusBadRequest, error2.RestfulError{Err: err.Error()})
		return
	}
	ctx := c.Request.Context()
	code, err := restful.usecase.CreateText(ctx, &model.Text)
	if err != nil {
		c.JSON(http.StatusInternalServerError, error2.RestfulError{Err: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, *code)
}
