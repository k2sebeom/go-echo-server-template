package schema

type User struct {
	Model

	SocialCredential *SocialCredential `gorm:"foreignKey:UserId"`
	PasswordCredential *PasswordCredential `gorm:"foreignKey:UserId"`
}
