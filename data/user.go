package data

import (
	"time"
)

type User struct {
	Id        int       `gorm:"primaryKey"`
	Uuid      string    `gorm:"type:varchar(64);unique;not null"`
	Name      string    `gorm:"type:varchar(255)"`
	Email     string    `gorm:"type:varchar(255);unique;not null"`
	Password  string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"not null"`
}

type Session struct {
	Id        int       `gorm:"primaryKey"`
	Uuid      string    `gorm:"type:varchar(64);unique;not null"`
	Email     string    `gorm:"type:varchar(255)"`
	UserId    int       `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
}

// Create a new session for an existing user
func (user *User) CreateSession() (session Session, err error) {
	session = Session{
		Uuid:      createUUID(),
		Email:     user.Email,
		UserId:    user.Id,
		CreatedAt: time.Now(),
	}
	err = Db.Create(&session).Error
	return
}

// Get the session for an existing user
func (user *User) Session() (session Session, err error) {
	err = Db.Where("user_id = ?", user.Id).First(&session).Error
	return
}

// Check if session is valid in the database
func (session *Session) Check() (valid bool, err error) {
	err = Db.Where("uuid = ?", session.Uuid).First(&session).Error
	if err != nil {
		valid = false
		return
	}
	valid = session.Id != 0
	return
}

// Delete session from database
func (session *Session) DeleteByUUID() (err error) {
	err = Db.Where("uuid = ?", session.Uuid).Delete(&Session{}).Error
	return
}

// Get the user from the session
func (session *Session) User() (user User, err error) {
	err = Db.Where("id = ?", session.UserId).First(&user).Error
	return
}

// Delete all sessions from database
func SessionDeleteAll() (err error) {
	err = Db.Exec("DELETE FROM sessions").Error
	return
}

// Create a new user, save user info into the database
func (user *User) Create() (err error) {
	user.Uuid = createUUID()
	user.CreatedAt = time.Now()
	user.Password = Encrypt(user.Password)
	err = Db.Create(&user).Error
	return
}

// Delete user from database
func (user *User) Delete() (err error) {
	err = Db.Delete(&user).Error
	return
}

// Update user information in the database
func (user *User) Update() (err error) {
	err = Db.Save(&user).Error
	return
}

// Delete all users from database
func UserDeleteAll() (err error) {
	err = Db.Exec("DELETE FROM users").Error
	return
}

// Get all users in the database and returns it
func Users() (users []User, err error) {
	err = Db.Find(&users).Error
	return
}

// Get a single user given the email
func UserByEmail(email string) (user User, err error) {
	err = Db.Where("email = ?", email).First(&user).Error
	return
}

// Get a single user given the UUID
func UserByUUID(uuid string) (user User, err error) {
	err = Db.Where("uuid = ?", uuid).First(&user).Error
	return
}
