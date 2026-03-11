package handlers

import (
	usersSt "CountStud/User"
	SimpleWork "CountStud/database/SimpleWork"
	structerr "CountStud/structerr"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type HTTPhandler struct {
	Student *usersSt.User
	conn    *pgx.Conn
}

func NewHttpHandlers(u *usersSt.User, conn *pgx.Conn) *HTTPhandler {
	return &HTTPhandler{
		Student: u,
		conn:    conn,
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
	student := &usersSt.User{}
	var newErr structerr.Err
	ctxFromGin := c.Request.Context()

	if err := c.ShouldBindJSON(student); err != nil {
		newErr = structerr.Err{
			Message: err.Error(),
			HasErr:  true,
		}
		c.JSON(400, newErr)
		return
	}

	student.Id = uuid.New()

	err := SimpleWork.InsertRow(ctxFromGin, s.conn, student)

	if err != nil {
		newErr = structerr.Err{
			Message: err.Error(),
			HasErr:  true,
		}
		c.JSON(500, newErr)
		return
	}
	c.JSON(201, student)
}
