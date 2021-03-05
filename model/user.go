package model

import (
	"errors"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model

	UserName string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=12"`
	Password string `gorm:"type:varchar(100);not null" json:"password" validate:"required,min=4,max=120"`
	Email    string `gorm:"type:varchar(100);not null" json:"email"`
	Phone    string `gorm:"type:varchar(100);not null" json:"phone"`
	Role     int    `gorm:"type:int;not null" json:"role" validate:"min=0,max=2"`
}

// 检查用户是否存在
func CheckUser(username string) bool {
	var user User
	db.Select("id").Where("user_name = ?", username).First(&user)
	if user.ID > 0 {
		return true
	}
	return false
}

func CheckUserWithId() {

}

func CreateUser(data *User) error {
	return db.Create(data).Error
}

func GetUserInfo() {

}

func GetUsers() {

}

func EditUser() {

}

func ChangePassword() {

}

func DeleteUser() {

}

func CheckPasswordBackEnd(username string, password string) error {
	if !CheckUser(username) {
		return errors.New("User does not exist")
	}

	var user User
	db.Where("user_name = ?", username).First(&user)

	if user.Role != 0 {
		return errors.New("No permission")
	}

	if user.ID <= 0 {
		return errors.New("User does not exist")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return err
	}

	return nil
}

func CheckPassword(username string, password string) error {
	if !CheckUser(username) {
		return errors.New("User does not exist")
	}

	var user User
	db.Where("user_name = ?", username).First(&user)

	if user.ID <= 0 {
		return errors.New("User does not exist")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return err
	}

	return nil
}

func (this *User) BeforeSave(_ *gorm.DB) error {
	this.Password = ScryptPassword(this.Password)
	return nil
}

func (this *User) BeforeUpdate(_ *gorm.DB) error {
	this.Password = ScryptPassword(this.Password)
	return nil
}

func ScryptPassword(password string) string {
	const cost = 10
	hashPw, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return ""
	}

	return string(hashPw)
}
