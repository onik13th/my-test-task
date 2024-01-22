package apiserver

import (
	"github.com/gin-gonic/gin"
	"github.com/onik13th/my-test-task/internal/app/model"
	"github.com/onik13th/my-test-task/internal/app/store"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strconv"
)

type server struct {
	router *gin.Engine
	logger *logrus.Logger
	store  store.Store
}

func newServer(store store.Store) *server {
	s := &server{
		router: gin.Default(),
		logger: logrus.New(),
		store:  store,
	}

	s.logger.SetFormatter(&logrus.TextFormatter{})
	s.logger.SetOutput(os.Stdout)
	s.logger.SetLevel(logrus.DebugLevel)

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)

}

func (s *server) configureRouter() {
	api := s.router.Group("/api")
	{
		users := api.Group("/users")
		{
			users.POST("/", s.handleUsersCreate)
			users.GET("/", s.handleFoundAllUsers)
			users.GET("/:id", s.handleFoundById)
			users.PUT("/:id", s.handleUpdateUser)
			users.DELETE("/:id", s.handleRemoveUser)
		}
	}
}

func (s *server) handleUsersCreate(c *gin.Context) {
	s.logger.Debug("Handle users create request")
	var user *model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		s.logger.Debug("Failed to bind JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	age, err := GetAge(user.Name)
	if err != nil {
		s.logger.Debug("Failed to get age:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get age"})
		return
	}

	gender, err := GetGender(user.Name)
	if err != nil {
		s.logger.Debug("Failed to get gender:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get gender"})
		return
	}

	nationalities, err := GetNationalities(user.Name)
	if err != nil {
		s.logger.Debug("Failed to get nationalities:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get nationalities"})
		return
	}

	user.Age = age
	user.Gender = gender
	user.Nationalities = nationalities

	createdUser, err := s.store.User().Create(user)
	if err != nil {
		s.logger.Error("Failed to create user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created", "User": createdUser})
}

func (s *server) handleFoundAllUsers(c *gin.Context) {
	s.logger.Debug("Handle found all users request")
	users, err := s.store.User().FindAll()
	if err != nil {
		s.logger.Debug("Failed to find all users")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err}) // надо бы обработать ошибки...
	} else {
		c.IndentedJSON(http.StatusOK, users)
	}
}

func (s *server) handleFoundById(c *gin.Context) {
	s.logger.Debug("Handle found by id request")
	linesId := c.Param("id")
	id, err := strconv.Atoi(linesId)
	if err != nil {
		s.logger.Error("Failed to parse id:", err)
	}

	user, err := s.store.User().FindById(id)
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()}) // обработать
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}

func (s *server) handleUpdateUser(c *gin.Context) {
	s.logger.Debug("Handle update user request")
	var updateUser *model.User
	if err := c.ShouldBindJSON(&updateUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err}) // надо бы обработать ошибки...
		return
	}

	linesId := c.Param("id")
	id, err := strconv.Atoi(linesId)
	if err != nil {
		s.logger.Error("Failed to parse id:", err)
	}

	newUpdateUser, err := s.store.User().Update(updateUser, id)
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated", "User": newUpdateUser})
}

func (s *server) handleRemoveUser(c *gin.Context) {
	s.logger.Debug("Handle remove user request")
	linesId := c.Param("id")
	id, err := strconv.Atoi(linesId)
	if err != nil {
		s.logger.Error("Failed to parse id:", err)
	}

	err = s.store.User().Remove(id)
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.Status(http.StatusNoContent)
}
