package Repositories

import (
	"fmt"
	"go-api/Databases"
	"go-api/Models"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllUsers(user *[]Models.User) (err error) {
	if err = Databases.Mysql.Find(user).Error; err != nil {
		return err
	}
	return nil
}

func CreateUser(user *Models.User) (err error) {
	if err = Databases.Mysql.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByID(user *Models.User, id string) (err error) {
	if err = Databases.Mysql.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *Models.User, id string) (err error) {
	fmt.Println(user)
	Databases.Mysql.Save(user)
	return nil
}

func DeleteUser(user *Models.User, id string) (err error) {
	Databases.Mysql.Where("id = ?", id).Delete(user)
	return nil
}

func LoginUser(user *Models.User, login string) (err error) {
	if err = Databases.Mysql.Where("login=?", login).First(user).Error; err != nil {
		return err
	}
	return nil
}
