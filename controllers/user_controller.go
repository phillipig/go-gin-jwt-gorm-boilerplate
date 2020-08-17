package controllers

import (
	"fmt"
	"go-api/models"
	"go-api/repositories"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	rep *repositories.UserRepository
}

var once sync.Once
var instance *UserController

func NewUserController() *UserController {
	once.Do(func() {
		instance = &UserController{
			rep: repositories.NewUserRepository(),
		}
	})
	return instance
}

func (ctrl *UserController) ReadAll(c *gin.Context) {
	var user []models.User
	err := ctrl.rep.ReadAll(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func (ctrl *UserController) ReadByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	err := ctrl.rep.ReadByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func (ctrl *UserController) Create(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	user.Senha, _ = hashPassword(user.Senha)
	err := ctrl.rep.Create(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func (ctrl *UserController) Update(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	err := ctrl.rep.ReadByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}
	c.BindJSON(&user)
	user.Senha, _ = hashPassword(user.Senha)
	err = ctrl.rep.Update(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func (ctrl *UserController) Delete(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	err := ctrl.rep.Delete(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
