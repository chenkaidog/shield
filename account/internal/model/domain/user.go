package domain

import "time"

type User struct {
	UserID      string
	AccountID   string
	Name        string
	Gender      Gender
	Phone       string
	Email       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Gender string

const (
	GenderMale   Gender = "male"
	GenderFemale Gender = "female"
	GenderOthers Gender = "others"
)

type UserCreateReq struct {
	AccountID   string
	Name        string
	Gender      Gender
	Phone       string
	Email       string
	Description string
}

type UserUpdateReq struct {
	UserID      string
	Name        string
	Gender      Gender
	Phone       string
	Email       string
	Description string
}

type UserQueryReq struct {
	UserID    string
	AccountID string
}
