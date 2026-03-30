package handlers

import (
	structerr "CountStud/structerr"
	SimpleWork "CountStud/workStudent/database/simpleWork"
	usersSt "CountStud/workStudent/student"
	userdto "CountStud/workStudent/userDTO"
	worktable "CountStud/workUsers/database/workTable"
	User "CountStud/workUsers/users"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
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

func (s *HTTPhandler) createJWT(user User.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.Id,
		"role":    user.Role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

var newErr structerr.Err

func (s *HTTPhandler) HandlerCreateAdmin(c *gin.Context) {

	ctxGin := c.Request.Context()
	user := User.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		newErr = structerr.Err{
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": newErr})
		return
	}
	user.Role = "Admin"
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ошибка хэширования"})
		return
	}
	user.Password = string(hash)

	if err := worktable.InsertRow(ctxGin, s.conn, &user); err != nil {
		newErr = structerr.Err{
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": newErr})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (s *HTTPhandler) HandlerCreateUser(c *gin.Context) {

	ctxGin := c.Request.Context()
	user := User.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		newErr = structerr.Err{
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": newErr})
		return
	}
	user.Role = "Teacher"
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ошибка хэширования"})
		return
	}
	user.Password = string(hash)

	if err := worktable.InsertRow(ctxGin, s.conn, &user); err != nil {
		newErr = structerr.Err{
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": newErr})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (s *HTTPhandler) HandlerEntrance(c *gin.Context) {
	ctxGin := c.Request.Context()
	userDTO := userdto.UserDTO{}

	if err := c.ShouldBindJSON(&userDTO); err != nil {
		newErr = structerr.Err{
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": newErr})
		return
	}

	user, err := worktable.GetUser(ctxGin, s.conn, userDTO.Email)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "пользователь не найден"})
		return
	}
	if errPass := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDTO.Password)); errPass != nil {
		newErr = *structerr.NewErr(errPass.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"error": "неверный пароль"})
		return
	}

	token, err := s.createJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ошибка генерации токена"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user": gin.H{
			"id":    user.Id,
			"email": user.Email,
			"name":  user.Name,
			"role":  user.Role,
		},
	})
}

func (s *HTTPhandler) HandlerCreateStudent(c *gin.Context) {
	student := usersSt.Student{}

	ctxGin := c.Request.Context()

	if err := c.ShouldBindJSON(&student); err != nil {
		newErr = structerr.Err{
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": newErr})
		return
	}
	err := SimpleWork.InsertRow(ctxGin, s.conn, &student)

	if err != nil {
		newErr = structerr.Err{
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": newErr})
		return
	}
	c.JSON(http.StatusOK, gin.H{"error": student})
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

	ctxGin := c.Request.Context()

	student, err := SimpleWork.GetStudentByID(ctxGin, s.conn, getIdUUID)
	if err != nil {
		newErr = structerr.Err{
			Message: err.Error(),
		}
		c.JSON(http.StatusNotFound, gin.H{"error": newErr})
		return
	}

	c.JSON(http.StatusOK, student)
}

func (s *HTTPhandler) HandlerGetAllStudents(c *gin.Context) {

	ctxGin := c.Request.Context()

	studnets, err := SimpleWork.GetAllStudent(ctxGin, s.conn)
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

	ctxGin := c.Request.Context()

	student, err := SimpleWork.GetStudentByID(ctxGin, s.conn, getIdUUID)
	if err != nil {
		newErr = structerr.Err{
			Message: err.Error(),
		}
		c.JSON(http.StatusNotFound, gin.H{"error": newErr})
		return
	}

	err = SimpleWork.DeleteRow(ctxGin, s.conn, student)
	if err != nil {
		errSt := structerr.NewErr(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": errSt})
		return
	}
}
