package user

import (
	"time"

	"github.com/rommms07/idream-erp/core/pb"
)

type User struct {
	Id                                 uint64 `gorm:"primaryKey"`
	Uname                              string `gorm:"unique"`
	First_name, Middle_name, Last_name string
	Email                              string `gorm:"unique"`
	Mobile                             string `gorm:"unique"`
	Picture_url                        string
	Gender                             pb.UserGender
	Type                               pb.UserType
	Fbid                               uint64 `gorm:"unique"`
	Created_at                         time.Time
}

type UserFacebookAccessToken struct {
	User_id      uint64 `gorm:"uniqueIndex"`
	Access_token string
	Type         pb.TokenType
	Expires_in   time.Time
}
