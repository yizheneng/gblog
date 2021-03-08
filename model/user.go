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

// 添加用户
func CreateUser(data *User) error {
	data.Password = ScryptPassword(data.Password)
	return db.Create(data).Error
}

// 获取用户信息
func GetUserInfo(username string) (*User, error) {
	var user User
	err := db.Model(&User{}).Select("id,user_name,created_at,updated_at,deleted_at,role,email,phone").Where("user_name = ?", username).First(&user).Error
	return &user, err
}

// 获取用户列表
func GetUsers(pageSize int, pageNum int) (users []User, total uint64, err error) {
	total = 0
	err = db.Model(&User{}).Select("*").Count(&total).Error
	if err != nil {
		return
	}

	err = db.Model(&User{}).Select("id,user_name,created_at,updated_at,deleted_at,role,email,phone").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil {
		return
	}

	err = db.Model(&users).Count(&total).Error
	return
}

// 更新用户信息
func UpdateUserInfo(username string, email string, phone string) (err error) {
	err = db.Model(&User{}).Where("user_name = ?", username).Select("phone", "email").Update(&User{Phone: phone, Email: email}).Error
	return
}

// func DeleteUser() {

// }

// 检查后端登录密码
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

// 检查前端登录密码
func CheckPassword(username string, password string) error {
	if !CheckUser(username) {
		return errors.New("User does not exist")
	}

	var user User
	db.Where("user_name = ?", username).First(&user)

	if user.ID <= 0 {
		return errors.New("User does not exist")
	}

	if user.Password == "" {
		return nil
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return err
	}

	return nil
}

// 修改密码
func ChangePassword(username string, password string, newPassword string) error {
	var user User
	db.Where("user_name = ?", username).First(&user)

	if user.Role == 0 {
		if CheckPasswordBackEnd(username, password) != nil {
			return errors.New("Password error")
		}
	} else {
		if CheckPassword(username, password) != nil {
			return errors.New("Password error")
		}
	}

	user.Password = ScryptPassword(newPassword)
	return db.Model(&User{}).Select("password").Where("id = ?", user.ID).Updates(user).Error
}

// func (this *User) BeforeSave(_ *gorm.DB) error {
// 	this.Password = ScryptPassword(this.Password)
// 	return nil
// }

// func (this *User) BeforeUpdate(_ *gorm.DB) error {
// 	this.Password = ScryptPassword(this.Password)
// 	return nil
// }

// 加密密码
func ScryptPassword(password string) string {
	const cost = 10
	hashPw, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return ""
	}

	return string(hashPw)
}
