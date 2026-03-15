package handlers

import (
	SimpleWork "CountStud/database/simpleWork"
	structerr "CountStud/structerr"
	usersSt "CountStud/student"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type HTTPhandler struct {
	conn *pgx.Conn
}

func NewHttpHandlers(conn *pgx.Conn) *HTTPhandler {
	return &HTTPhandler{
		conn: conn,
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
var studentErr structerr.Err

func (s *HTTPhandler) HandlerCreateStudent(c *gin.Context) {
	student := &usersSt.User{}

	ginCtx := c.Request.Context()

	if err := c.ShouldBindJSON(student); err != nil {
		studentErr = Err.New(err.Error())
		c.JSON(400, studentErr)
		return
	}

	err := SimpleWork.InsertRow(ginCtx, s.conn, student)

	if err != nil {
		studentErr = Err.New(err.Error())
		c.JSON(500, studentErr)
		return
	}

	// JWT token
	c.JSON(http.StatusOK, student)
}

/*
pattern: /tasks
method:  GET
info:    -

succeed:
  - status code: 200 Ok
  - response body: JSON represented found tasks

failed:
  - status code: 400, 500, ...
  - response body: JSON with error + time
*/

func (s *HTTPhandler) HandlerGetStudentID(c *gin.Context) {

	getIdString := c.Param("id")
	getIdUUID, err := uuid.Parse(getIdString)
	if err != nil {
		newErr = structerr.Err{
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, newErr)
		return
	}

	ctxFromGin := c.Request.Context()

	student, err := SimpleWork.GetStudentByID(ctxFromGin, s.conn, getIdUUID)
	if err != nil {
		newErr = structerr.Err{
			Message: err.Error(),
		}
		c.JSON(http.StatusNotFound, newErr)
		return
	}

	c.JSON(http.StatusOK, student)
	return
}
