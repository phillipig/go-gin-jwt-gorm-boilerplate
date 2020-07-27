package Repositories

import (
	"fmt"
	"go-api/Databases"
	"go-api/Models"

	_ "github.com/go-sql-driver/mysql"
)

//GetAllUsers Fetch all user data
func GetAllUsers(user *[]Models.User) (err error) {
	if err = Databases.Mysql.Find(user).Error; err != nil {
		return err
	}
	return nil
}

//CreateUser ... Insert New data
func CreateUser(user *Models.User) (err error) {
	if err = Databases.Mysql.Create(user).Error; err != nil {
		return err
	}
	return nil
}

//GetUserByID ... Fetch only one user by Id
func GetUserByID(user *Models.User, id string) (err error) {
	if err = Databases.Mysql.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

//UpdateUser ... Update user
func UpdateUser(user *Models.User, id string) (err error) {
	fmt.Println(user)
	Databases.Mysql.Save(user)
	return nil
}

//DeleteUser ... Delete user
func DeleteUser(user *Models.User, id string) (err error) {
	Databases.Mysql.Where("id = ?", id).Delete(user)
	return nil
}
