package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Tag struct {
	Model
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetTag(maps interface{}) (tag Tag) {
	db.Where(maps).First(&tag)
	return
}

func GetTags(page int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(page).Limit(pageSize).Find(&tags)
	return
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func ExistTagByName(name string) bool {
	maps := CommonMaps()
	maps["name"] = name

	var tag Tag
	db.Select("id").Where(maps).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func AddTag(name string, createdBy string) bool {
	newTag := Tag{
		Name:      name,
		State:     STATE_ONLINE,
		CreatedBy: createdBy,
	}
	db.Create(&newTag)
	return true
}

func ExistTagByID(id int) bool {
	maps := CommonMaps()
	maps["id"] = id
	var tag Tag
	db.Select("id").Where(maps).First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}

func DeleteTag(id int) bool {
	maps := CommonMaps()
	maps["id"] = id
	db.Where(maps).Delete(&Tag{})
	return true
}

func EditTag(id int, data interface{}) bool {
	maps := CommonMaps()
	maps["id"] = id
	db.Model(&Tag{}).Where(maps).Updates(data)
	return true
}

func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("Ctime", time.Now().Unix())
	return nil
}
func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("Mtime", time.Now().Unix())
	return nil
}
