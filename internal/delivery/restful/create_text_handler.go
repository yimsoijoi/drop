package restful

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/yimsoijoi/drop/internal/delivery/restful/model"
	errorLibrary "github.com/yimsoijoi/drop/lib/error"
)

func (restful restfulDelivery) CreatTextHandler(c *gin.Context) {
	model := new(model.CreateTextModel)
	if err := c.BindJSON(model); err != nil {
		c.JSON(http.StatusBadRequest, errorLibrary.RestfulError{Err: err.Error()})
		return
	}
	ctx := c.Request.Context()
	code, err := restful.usecase.CreateText(ctx, &model.Text)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorLibrary.RestfulError{Err: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, *code)
}
