package user

import (
	"time"

	"github.com/rommms07/idream-erp/core/pb/user_schema"
	"github.com/rommms07/idream-erp/core/source"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

func init() {
	source.GormMigrator.Add(&User{}).Add(&UserAuthToken{})
}

type User struct {
	Id                              uint64 `gorm:"primaryKey"`
	Uname                           string `gorm:"unique"`
	FirstName, MiddleName, LastName string
	Suffix                          uint64
	Email                           string `gorm:"unique"`
	Mobile                          string `gorm:"unique"`
	PictureUrl                      string
	Gender                          user_schema.UserGender
	Type                            user_schema.UserType
	Fbid                            uint64 `eorm:"unique"`
	Birthdate                       time.Time
	CreatedAt                       time.Time
	Uflags                          uint64
	State                           user_schema.UserState
}

type UserAuthToken struct {
	UserId          uint64 `gorm:"uniqueIndex"`
	AuthTokenString string `gorm:"unique"`
	Type            user_schema.TokenType
	ExpiresIn       time.Time
}

func (authToken *UserAuthToken) Owner() *User {
	db := source.Source[gorm.DB]()
	user := &User{}

	db.Model(User{Id: authToken.UserId}).First(user)
	return user
}

func (u *User) Proto() *user_schema.User {
	return &user_schema.User{
		Id:    u.Id,
		Fbid:  u.Fbid,
		Uname: u.Uname,

		FullName: &user_schema.UserFullname{
			FirstName:  u.FirstName,
			LastName:   u.LastName,
			MiddleName: u.MiddleName,
			Suffix:     u.Suffix,
		},

		PictureUrl: u.PictureUrl,
		Gender:     u.Gender,
		Type:       u.Type,

		Birthdate: timestamppb.New(u.Birthdate),
		CreatedAt: timestamppb.New(u.CreatedAt),

		State:  u.State,
		Uflags: u.Uflags,
	}
}

func (u *UserAuthToken) Proto() *user_schema.UserAuthToken {
	return &user_schema.UserAuthToken{
		UserId:    u.UserId,
		AuthToken: u.AuthTokenString,
		Type:      u.Type,
		ExpiresIn: timestamppb.New(u.ExpiresIn),
	}
}
