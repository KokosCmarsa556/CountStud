package handlers

import (
	SimpleWork "CountStud/database/simpleWork"
	structerr "CountStud/structerr"
	usersSt "CountStud/student"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type HTTPhandler struct {
	conn *pgx.Conn
}

type studentsAll struct {
	students []usersSt.User
}

func NewHttpHandlers(conn *pgx.Conn) *HTTPhandler {
	return &HTTPhandler{
		conn: conn,
	}
}

var newErr structerr.Err

func (s *HTTPhandler) HandlerCreateStudent(c *gin.Context) {
	student := usersSt.User{}

	ctxFromGin := c.Request.Context()

	if err := c.ShouldBindJSON(student); err != nil {
		newErr = structerr.Err{
			Message: err.Error(),
		}
		c.JSON(400, gin.H{"error": newErr})
		return
	}

	if student.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password is required"})
		return
	}

	hash, errCrypto := bcrypt.GenerateFromPassword([]byte(student.Password), bcrypt.DefaultCost)
	if errCrypto != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errCrypto})
		return
	}

	student.Password = string(hash)
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
