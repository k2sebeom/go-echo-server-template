package schema

type PasswordCredential struct {
	Model

	UserId uint `gorm:"column:user_id"`

	Email string `gorm:"column:email"`
	PasswordHash string `gorm:"column:password_hash"`
}

type SocialProvider string

const (
	KAKAO SocialProvider = "KAKAO"
	NAVER SocialProvider = "NAVER"
	GOOGLE SocialProvider = "GOOGLE"
)

type SocialCredential struct {
	Model

	UserId uint `gorm:"column:user_id"`

	IdentifierHash string `gorm:"column:identifier_hash"`
	Provider SocialProvider `gorm:"column:provider"`
}