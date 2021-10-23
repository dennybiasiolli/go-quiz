package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email      string `gorm:"not null; size:255"`
	FirstName  string `gorm:"not null; size:255"`
	LastName   string `gorm:"not null; size:155"`
	IsActive   bool   `gorm:"not null; default:true"`
	PictureUrl string `gorm:""`
	Locale     string `gorm:"not null; size:15; default:en"`
}

type TokenRefreshInput struct {
	Refresh string `json:"refresh" xml:"refresh" form:"refresh" validate:"required"`
}

type JwtUserInfo struct {
	Email      string `json:"email"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	FullName   string `json:"full_name"`
	PictureUrl string `json:"picture_url"`
	Locale     string `json:"locale"`
}
type JwtCustomClaims struct {
	TokenType string      `json:"token_type"`
	UserId    uint        `json:"user_id"`
	UserInfo  JwtUserInfo `json:"user_info"`
	jwt.RegisteredClaims
}

type GoogleUserInfo struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	FamilyName    string `json:"family_name"`
	GivenName     string `json:"given_name"`
	Name          string `json:"name"`
	Locale        string `json:"locale"`
	Picture       string `json:"picture"`
	VerifiedEmail bool   `json:"verified_email"`
}
