package models

import (
	"server/pkg/config"

	"gorm.io/gorm"
)

type Likes struct{
	gorm.Model
    UserID int `json:"userid"`
    PostID int `json:"postid"`
	
}

func init()  {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Likes{})
}


func (liked_by *Likes) LikePost() *Likes{
	db.Create(&liked_by) 
	db.Model(&Posts{}).Where("ID = ?", liked_by.PostID).Update("likes", gorm.Expr("likes + ?", 1))

	return liked_by
}