package user

import "time"

type User struct {
	Id               int
	Name             string
	Occupation       string
	Email            string
	Password         string
	AvatarFileName   string
	Role             string
	CreatedAt        time.Time
	UpdatedAt 		 time.Time
}