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
var newErr structerr.Err

func (s *HTTPhandler) HandlerCreateStudent(c *gin.Context) {
	student := &usersSt.User{}

	ginCtx := c.Request.Context()

	if err := c.ShouldBindJSON(student); err != nil {
		newErr = structerr.Err{
			Message: err.Error(),
			HasErr:  true,
		}
		c.JSON(400, newErr)
		return
	}

	student.Id = uuid.New()

	err := SimpleWork.InsertRow(ginCtx, s.conn, student)

	if err != nil {
		newErr = structerr.Err{
			Message: err.Error(),
			HasErr:  true,
		}
		c.JSON(500, newErr)
		return
	}
	c.JSON(http.StatusOK, student)
}

func (s *HTTPhandler) HandlerGetStudentsID(c *gin.Context) {

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

	ginCtx := c.Request.Context()

	student, err := SimpleWork.GetStudentByID(ginCtx, s.conn, getIdUUID)
	if err != nil {
		newErr = structerr.Err{
			Message: err.Error(),
		}
		c.JSON(http.StatusNotFound, newErr)
		return
	}

	c.JSON(http.StatusOK, student)
}

func (s *HTTPhandler) HandlerGetAllStudents(c *gin.Context) {
	ginCtx := c.Request.Context()

	students, err := SimpleWork.GetAllStudent(ginCtx, s.conn)
	if err != nil {
		newErr = structerr.Err{
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, newErr)
		return
	}

	c.JSON(http.StatusOK, students)
}

func (s *HTTPhandler) HandlerDeleteStudent(c *gin.Context) {
	getIdString := c.Param("id")
	getIdUUID, err := uuid.Parse(getIdString)
	if err != nil {
		newErr = structerr.Err{
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, newErr)
		return
	}

	ginCtx := c.Request.Context()

	student, err := SimpleWork.GetStudentByID(ginCtx, s.conn, getIdUUID)
	if err != nil {
		newErr = structerr.Err{
			Message: err.Error(),
		}
		c.JSON(http.StatusNotFound, newErr)
		return
	}

	deleteErr := SimpleWork.DeleteRow(ginCtx, s.conn, student)
	if deleteErr != nil {
		errSt := structerr.NewErr(deleteErr.Error())
		c.JSON(http.StatusBadRequest, errSt)
		return
	}
	emptyStudnet := usersSt.User{}
	c.JSON(http.StatusOK, emptyStudnet)
}
