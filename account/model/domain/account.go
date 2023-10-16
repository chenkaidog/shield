package domain

type Account struct {
	AccountID string
	Username  string
	Status    AccountStatus
}

type AccountStatus string

const (
	AccountStatusValid   AccountStatus = "valid"
	AccountStatusInvalid AccountStatus = "invalid"
)

type AccountCreateReq struct {
	Username string
	Password string
}

type AccountPswResetReq struct {
	AccountID string
	Password  string
}

type AccountPswUpdateReq struct {
	AccountID   string
	Password    string
	NewPassword string
}

type AccountStatusUpdateReq struct {
	AccountID string
	Status    AccountStatus
}

type AccountQueryReq struct {
	AccountID string
}
