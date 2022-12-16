package models

import (
	"server/pkg/config"

	"gorm.io/gorm"
)

var db *gorm.DB

type Users struct{
	gorm.Model
	ID int           `gorm:"primaryKey;autoIncrement"`
	Fullname string `form:"Fullname" json:"fullname"`
    Username string `gorm:"unique" form:"Username" json:"username"`
    Password string `form:"Password" json:"password"`
    Email    string  `gorm:"unique" form:"Email" json:"email"`
	Profile string 	`form:"Image" json:"image"`
	Bio		string 	`form:"Bio" json:"bio"`
	Posts []Posts  	`gorm:"foreignKey:UserID;references:ID"`
	LikedPosts []Likes `gorm:"foreignKey:PostID;references:ID"`
}

func init()  {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Users{})
}

func (user *Users) RegisterUser() *Users{
	db.Create(&user) 
	return user
}

func (u *Users) LoginUser() *Users{
	db.First(&u, "username = ?", u.Username)
	return u
}

func GetAllUsers() []Users{
	var Users []Users
	db.Find(&Users)
    return Users
}

func GetPostsOfUser(id uint64) ([]Result,[]Result,Users){
	var posts []Result
	var liked_posts []Result
	var user_data Users
	db.First(&user_data, "id = ?", id)
	db.Raw("SELECT posts.id,fullname,username,content,email,profile,likes,user_id,bio,posts.created_at from posts JOIN users ON posts.user_id = users.id WHERE users.id = ?",id).Scan(&posts)
	db.Raw("SELECT * from posts JOIN likes ON posts.id = likes.post_id JOIN users ON posts.user_id = users.id WHERE likes.user_id = ?",id).Scan(&liked_posts)
	return posts ,liked_posts,user_data
}

func (user *Users) UpdateUser() *Users {
	var new_info *Users
	db.Model(&user).Where("ID = ?",user.ID).Updates(Users{Fullname: user.Fullname, Username:user.Username,Bio: user.Bio})
	db.First(&new_info,user.ID)
	return new_info
}

func DeleteUser(id string) Users{
    var deleteUser Users
	db.Unscoped().Delete(&Users{}, id)
	return deleteUser
}