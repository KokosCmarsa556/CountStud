package handlers

import (
	usersSt "CountStud/User"
	structerr "CountStud/structerr"

	"github.com/gin-gonic/gin"
)

type HTTPhandler struct {
	Student *usersSt.User
}

func NewHttpHandlers(u *usersSt.User) *HTTPhandler {
	return &HTTPhandler{
		Student: u,
	}
}

/*
pattern: /student
method: POST
info: JSON in HTTP requst body

succeed:
  - status code:   201 Created
  - response body: JSON represent created task

failed:
  - status code:   400, 409, 500, ...
  - response body: JSON with error + time
*/

func (s *HTTPhandler) HandleCreateStudent(c *gin.Context) {
	var student HTTPhandler
	var newErr structerr.Err

	if err := c.ShouldBind(&student); err != nil {
		newErr = structerr.Err{
			Message: err.Error(),
			HasErr:  true,
		}
		c.JSON(400, newErr)
		return
	}

}
