package domain

import "time"

type LoginRecord struct {
	AccountID string
	LoginAt   time.Time
	IPv4      string
	Device    string
	Reason    string
	Status    string
}

type LoginStatus string

const (
	LoginStatusSuccess LoginStatus = "success"
	LoginStatusFail    LoginStatus = "fail"
)

type LoginReq struct {
	Username string
	Password string
	IPv4     string
	Device   string
}

type LoginRecordQueryReq struct {
	AccountID string
	Page      int64
	Size      int64
}
