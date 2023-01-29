package model

import (
	"gin_blog/utils/err_msg"
)

type Category struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

func CheckCategory(name string) int {
	var category Category
	db.Select("id").Where("name = ?", name).First(&category)
	if category.ID > 0 {
		return err_msg.ERROR_CATENAME_USED
	}
	return err_msg.SUCCESS
}

func CreateCategory(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return err_msg.ERROR
	}
	return err_msg.SUCCESS
}

func GetCategory(pageSize, pageNum int) ([]Category, int64) {
	var categorys []Category
	total := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).
		Find(&categorys).RowsAffected
	return categorys, total
}

func EditCategory(id int, data *Category) int {
	data.ID = uint(id)
	err := db.Model(data).Updates(data).Error
	if err != nil {
		return err_msg.ERROR
	}
	return err_msg.SUCCESS
}

func DeleteCategory(id int) int {
	var category Category
	err := db.Where("id = ?", id).Delete(&category).Error
	if err != nil {
		return err_msg.ERROR
	}
	return err_msg.SUCCESS
}
