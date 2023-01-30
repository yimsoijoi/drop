package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/yimsoijoi/drop/internal/domain"
	error2 "github.com/yimsoijoi/drop/lib/error"
	"log"
	"net/http"
)

func (restful restfulDelivery) GetTextHandler(c *gin.Context) {
	code := c.Query("key")
	log.Println(code)
	text, err := restful.usecase.GetText(c, &code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, error2.RestfulError{Err: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.RestfulResponse{Text: *text})
}
