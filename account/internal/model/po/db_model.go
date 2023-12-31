package po

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	AccountID string `gorm:"column:account_id" json:"account_id"`
	Username  string `gorm:"column:username" json:"username"`
	Password  string `gorm:"column:password" json:"password"`
	Salt      string `gorm:"column:salt" json:"salt"`
	Status    string `gorm:"column:status" json:"status"`
}

func (Account) TableName() string {
	return "account"
}

func NewAccount() *Account {
	return &Account{}
}

type User struct {
	gorm.Model
	UserID      string `gorm:"column:user_id" json:"user_id"`
	AccountID   string `gorm:"column:account_id" json:"account_id"`
	Name        string `gorm:"column:name" json:"name"`
	Gender      string `gorm:"column:gender" json:"gender"`
	Phone       string `gorm:"column:phone" json:"phone"`
	Email       string `gorm:"column:email" json:"email"`
	Description string `gorm:"column:description" json:"description"`
}

func (User) TableName() string {
	return "user"
}

func NewUser() *User {
	return &User{}
}

type LoginRecord struct {
	gorm.Model
	AccountID string    `gorm:"column:account_id" json:"account_id"`
	LoginAt   time.Time `gorm:"column:login_at" json:"login_at"`
	IPv4      string    `gorm:"column:ipv4" json:"ipv4"`
	Device    string    `gorm:"column:device" json:"device"`
	Reason    string    `gorm:"column:reason" json:"reason"`
	Status    string    `gorm:"column:status" json:"status"`
}

func (LoginRecord) TableName() string {
	return "login_record"
}

func NewLoginRecord() *LoginRecord {
	return &LoginRecord{}
}
