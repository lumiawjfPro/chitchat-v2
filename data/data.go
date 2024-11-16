package data

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"log"
)

//var DB *gorm.DB
//
//func init() {
//	dsn := "host=/var/run/postgresql user=postgres password=892289 dbname=chitchat sslmode=disable"
//	// dsn := "user=postgres password=892289 dbname=chitchat  sslmode=disable"
//	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
//	if err != nil {
//		log.Fatal("Failed to connect to database:", err)
//	}
//
//	err = DB.AutoMigrate(&User{}, &Session{}, &Thread{}, &Post{})
//	if err != nil {
//		log.Fatal(err)
//	}
//}

// create a random UUID with from RFC 4122 adapted from http://github.com/nu7hatch/gouuid
func createUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}

	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

// hash plaintext with SHA-1
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}
