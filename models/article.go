package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	Model
	TagID      int    `json:"tag_id" gorm:"index"`
	Tag        Tag    `json:"tag"`
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func ExistArticleByID(id int) bool {
	var article Article
	maps := CommonMaps()
	maps["id"] = id
	db.Select("id").Where(maps).First(&article)
	if article.ID > 0 {
		return true
	}
	return false
}

func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

func GetArticles(offset int, limit int, maps interface{}) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(offset).Limit(limit).Find(&articles)
	return
}

func GetArticle(id int) (article Article) {
	maps := CommonMaps()
	maps["id"] = id
	db.Where(maps).First(&article)
	db.Model(&article).Related(&article.Tag)
	return
}

func EditArticle(id int, data interface{}) bool {
	maps := CommonMaps()
	maps["id"] = id
	db.Model(&Article{}).Where(maps).Updates(data)
	return true
}

func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     STATE_ONLINE,
	})
	return true
}

func DeleteArticle(id int) bool {
	maps := CommonMaps()
	maps["id"] = id

	// 物理删除
	//db.Where(maps).Delete(Article{})

	// 逻辑删除
	data := GetArticle(id)
	data.State = STATE_OFFLINE
	db.Model(&Article{}).Where(maps).Updates(data)
	return true
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("Ctime", time.Now().Unix())
	return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("Mtime", time.Now().Unix())
	return nil
}
