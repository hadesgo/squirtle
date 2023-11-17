package models

import (
	orm "squirtle/database"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

var Users []User

func init() {
	orm.Eloquent.AutoMigrate(&User{})
}

func (user User) Insert() (id uint, err error) {
	result := orm.Eloquent.Create(&user)
	id = user.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

func (user *User) Users() (users []User, err error) {
	if err = orm.Eloquent.Find(&users).Error; err != nil {
		return
	}
	return
}

func (user *User) Update(id uint) (updateUser User, err error) {
	if err = orm.Eloquent.Select([]string{"id", "username"}).First(&updateUser, id).Error; err != nil {
		return
	}
	if err = orm.Eloquent.Model(&updateUser).Updates(&user).Error; err != nil {
		return
	}
	return
}

func (user *User) Delete(id uint) (result User, err error) {
	if err = orm.Eloquent.Select([]string{"id"}).First(&user, id).Error; err != nil {
		return
	}

	if err = orm.Eloquent.Delete(&user).Error; err != nil {
		return
	}
	result = *user
	return
}
