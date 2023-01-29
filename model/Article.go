package model

import (
	"gin_blog/utils/err_msg"
	"gorm.io/gorm"
)

type Article struct {
	Category Category `gorm:"foreignKey:Cid"`
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Desc    string `gorm:"type:varchar(200)" json:"desc"`
	Content string `gorm:"type:longtext" json:"content"`
	Img     string `gorm:"type:varchar(100)" json:"img"`
}

func CreateArticle(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return err_msg.ERROR
	}
	return err_msg.SUCCESS
}

func GetCateArticle(cid, pageSize, pageNum int) ([]Article, int64) {
	var cateArticleList []Article
	total := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).
		Where("cid = ?", cid).Find(&cateArticleList).RowsAffected
	return cateArticleList, total
}

func GetArticleInfo(id int) (Article, int) {
	var article Article
	err := db.Where("id = ?", id).Preload("Category").First(&article).Error
	if err != nil {
		return article, err_msg.ERROR_ARTICLE_NOT_EXIST
	}
	return article, err_msg.SUCCESS
}

func GetArticle(pageSize, pageNum int) ([]Article, int64) {
	var articleList []Article
	total := db.Select("article.id, title, img, created_at, updated_at, `desc`, category.name").
		Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("created_at DESC").Joins("Category").
		Find(&articleList).RowsAffected
	return articleList, total
}

func SearchArticle(title string, pageSize, pageNum int) ([]Article, int64) {
	var articleList []Article
	total := db.Select("article.id, title, img, created_at, updated_at, `desc`, category.name").
		Order("created_at DESC").Joins("Category").Where("title LIKE ?", title+"%").
		Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articleList).RowsAffected
	return articleList, total
}

func EditArticle(id int, data *Article) int {
	data.ID = uint(id)
	err := db.Model(&data).Updates(&data).Select("title, cid, `desc`, content, img").Error
	if err != nil {
		return err_msg.ERROR
	}
	return err_msg.SUCCESS
}

func DeleteArticle(id int) int {
	var article Article
	err := db.Where("id = ?", id).Delete(&article).Error
	if err != nil {
		return err_msg.ERROR
	}
	return err_msg.SUCCESS
}
