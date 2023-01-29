package model

import (
	"encoding/base64"
	"gin_blog/utils/err_msg"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=3,max=20" label:"用戶名"`
	Password string `gorm:"type:varchar(50);not null" json:"password" validate:"required,min=6,max=20" label:"密碼"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色碼"`
}

func (u *User) BeforeCreate(_ *gorm.DB) error {
	u.Password = ScryptPw(u.Password)
	//u.Role = 2
	return nil
}

func ScryptPw(password string) string {
	salt := []byte{12, 59, 8, 47, 129, 63, 73, 230}

	hashPW, err := scrypt.Key([]byte(password), salt, 1<<15, 8, 1, 32)
	if err != nil {
		log.Fatal(err)
	}

	bHashPW := base64.StdEncoding.EncodeToString(hashPW)
	return bHashPW
}

// CheckUser 查詢用戶是否存在
func CheckUser(name string) (code int) {
	var user User
	db.Select("id").Where("user_name = ?", name).First(&user)
	if user.ID > 0 {
		return err_msg.ERROR_USERNAME_USED
	}
	return err_msg.SUCCESS
}

func CreateUser(data *User) int {
	err := db.Create(data).Error
	if err != nil {
		return err_msg.ERROR
	}
	return err_msg.SUCCESS
}

func GetUsers(username string, pageSize int, pageNum int) ([]User, int64) {
	var users []User

	if username != "" {
		total := db.Select("id, user_name, role, created_at").Where("user_name LIKE ?", username+"%").
			Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).RowsAffected
		return users, total
	}

	total := db.Select("id, user_name, role, created_at").Limit(pageSize).Offset((pageNum - 1) * pageSize).
		Find(&users).RowsAffected

	return users, total
}

func EditUser(id int, data *User) int {
	data.ID = uint(id)
	err := db.Model(data).Updates(data).Error
	if err != nil {
		return err_msg.ERROR
	}
	return err_msg.SUCCESS
}

func DeleteUser(id int) int {
	var user User
	err := db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return err_msg.ERROR
	}
	return err_msg.SUCCESS
}

func CheckLogin(userName, password string) (User, int) {
	var user User

	err := db.Where("user_name = ?", userName).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return user, err_msg.ERROR_USER_NOT_EXIST
	}

	if ScryptPw(password) != user.Password {
		return user, err_msg.ERROR_PASSWORD_WRONG
	}

	if user.Role != 1 {
		return user, err_msg.ERROR_USER_NO_RIGHT
	}

	return user, err_msg.SUCCESS
}
