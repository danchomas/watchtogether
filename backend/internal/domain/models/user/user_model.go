package user

import "time"

type User struct {
	Id           string
	Nickname     string
	Email        string
	PasswordHash string
	AvatarUrl    string
	CreatedAt    time.Time
}

func (u User) IsGuest() bool {
	return u.PasswordHash != ""
}
