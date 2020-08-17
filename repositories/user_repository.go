package repositories

import (
	"fmt"
	"go-api/databases"
	"go-api/models"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

var once sync.Once
var instance *UserRepository

func NewUserRepository() *UserRepository {
	once.Do(func() {
		instance = &UserRepository{
			db: databases.NewMysql(),
		}
	})
	return instance
}

func (rep *UserRepository) ReadAll(user *[]models.User) (err error) {
	if err = rep.db.Find(user).Error; err != nil {
		return err
	}
	return nil
}

func (rep *UserRepository) Create(user *models.User) (err error) {
	if err = rep.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (rep *UserRepository) ReadByID(user *models.User, id string) (err error) {
	if err = rep.db.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (rep *UserRepository) Update(user *models.User, id string) (err error) {
	fmt.Println(user)
	rep.db.Save(user)
	return nil
}

func (rep *UserRepository) Delete(user *models.User, id string) (err error) {
	rep.db.Where("id = ?", id).Delete(user)
	return nil
}

func (rep *UserRepository) LoginUser(user *models.User, login string) (err error) {
	if err = rep.db.Where("login=?", login).First(user).Error; err != nil {
		return err
	}
	return nil
}
