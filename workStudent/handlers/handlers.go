package handlers

import (
	structerr "CountStud/structerr"
	SimpleWork "CountStud/workStudent/database/simpleWork"
	usersSt "CountStud/workStudent/student"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type HTTPhandler struct {
	conn *pgx.Conn
}

type studentsAll struct {
	students []usersSt.Student
}

func NewHttpHandlers(conn *pgx.Conn) *HTTPhandler {
	return &HTTPhandler{
		conn: conn,
	}
}

var newErr structerr.Err

func (s *HTTPhandler) HandlerCreateStudent(c *gin.Context) {
	student := usersSt.Student{}

	ctxFromGin := c.Request.Context()

	if err := c.ShouldBindJSON(student); err != nil {
		newErr = structerr.Err{
			Message: err.Error(),
		}
		c.JSON(400, gin.H{"error": newErr})
		return
	}
	err := SimpleWork.InsertRow(ctxFromGin, s.conn, &student)

	if err != nil {
		newErr = structerr.Err{
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": newErr})
		return
	}
	c.JSON(http.StatusOK, gin.H{"error": student})
}

func (s *HTTPhandler) HandlerGetStudentsID(c *gin.Context) {

}

func (s *HTTPhandler) HandlerGetStudentID(c *gin.Context) {

	getIdString := c.Param("id")
	getIdUUID, err := uuid.Parse(getIdString)
	if err != nil {
		newErr = structerr.Err{
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": newErr})
		return
	}

	ctxFromGin := c.Request.Context()

	student, err := SimpleWork.GetStudentByID(ctxFromGin, s.conn, getIdUUID)
	if err != nil {
		newErr = structerr.Err{
			Message: err.Error(),
		}
		c.JSON(http.StatusNotFound, gin.H{"error": newErr})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": student})
}

func (s *HTTPhandler) HandlerGetAllStudents(c *gin.Context) {

	ctxFromGin := c.Request.Context()

	studnets, err := SimpleWork.GetAllStudent(ctxFromGin, s.conn)
	if err != nil {
		newErr = structerr.Err{
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": newErr})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": studnets})
}

func (s *HTTPhandler) HandlerDeleteStudent(c *gin.Context) {
	getIdString := c.Param("id")
	getIdUUID, err := uuid.Parse(getIdString)
	if err != nil {
		newErr = structerr.Err{
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": newErr})
		return
	}

	ctxFromGin := c.Request.Context()

	student, err := SimpleWork.GetStudentByID(ctxFromGin, s.conn, getIdUUID)
	if err != nil {
		newErr = structerr.Err{
			Message: err.Error(),
		}
		c.JSON(http.StatusNotFound, gin.H{"error": newErr})
		return
	}

	err = SimpleWork.DeleteRow(ctxFromGin, s.conn, student)
	if err != nil {
		errSt := structerr.NewErr(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": errSt})
		return
	}
}
