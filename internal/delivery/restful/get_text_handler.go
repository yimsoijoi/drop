package restful

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/yimsoijoi/drop/internal/delivery/restful/model"
	errorLibrary "github.com/yimsoijoi/drop/lib/error"
)

func (restful restfulDelivery) GetTextHandler(c *gin.Context) {
	code := c.Query("key")
	text, err := restful.usecase.GetText(c, &code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorLibrary.RestfulError{Err: err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.ResponseModel{Text: *text})
}
