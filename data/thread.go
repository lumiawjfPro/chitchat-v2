package data

import (
	"time"

	"gorm.io/gorm"
)

type Thread struct {
	Id        int       `gorm:"primaryKey;autoIncrement"`
	Uuid      string    `gorm:"type:varchar(64);not null;unique"`
	Topic     string    `gorm:"type:text"`
	UserId    int       `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
}

type Post struct {
	Id        int       `gorm:"primaryKey;autoIncrement"`
	Uuid      string    `gorm:"type:varchar(64);not null;unique"`
	Body      string    `gorm:"type:text"`
	UserId    int       `gorm:"not null"`
	ThreadId  int       `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
}

var Db *gorm.DB

func SetDB(DB *gorm.DB) {
	Db = DB
}

// format the CreatedAt date to display nicely on the screen
func (thread *Thread) CreatedAtDate() string {
	return thread.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

func (post *Post) CreatedAtDate() string {
	return post.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

// get the number of posts in a thread
func (thread *Thread) NumReplies() (count int64, err error) {
	err = Db.Model(&Post{}).Where("thread_id = ?", thread.Id).Count(&count).Error
	return
}

// get posts to a thread
func (thread *Thread) Posts() (posts []Post, err error) {
	err = Db.Where("thread_id = ?", thread.Id).Find(&posts).Error
	return
}

// Create a new thread
func (user *User) CreateThread(topic string) (conv Thread, err error) {
	conv = Thread{
		Uuid:      createUUID(),
		Topic:     topic,
		UserId:    user.Id,
		CreatedAt: time.Now(),
	}
	err = Db.Create(&conv).Error
	return
}

// Create a new post to a thread
func (user *User) CreatePost(conv Thread, body string) (post Post, err error) {
	post = Post{
		Uuid:      createUUID(),
		Body:      body,
		UserId:    user.Id,
		ThreadId:  conv.Id,
		CreatedAt: time.Now(),
	}
	err = Db.Create(&post).Error
	return
}

// Get all threads in the database and returns it
func Threads() (threads []Thread, err error) {
	err = Db.Order("created_at desc").Find(&threads).Error
	return
}

// Get a thread by the UUID
func ThreadByUUID(uuid string) (conv Thread, err error) {
	err = Db.Where("uuid = ?", uuid).First(&conv).Error
	return
}

// Get the user who started this thread
func (thread *Thread) User() (user User) {
	_ = Db.Where("id = ?", thread.UserId).First(&user).Error
	return
}

// Get the user who wrote the post
func (post *Post) User() (user User) {
	_ = Db.Where("id = ?", post.UserId).First(&user).Error
	return
}
