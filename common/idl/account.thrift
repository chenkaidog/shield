include "base.thrift"

namespace go kaidog.shield.account

enum AccountStatus {
    valid = 1
    invalid = 2 
}

struct Account {
    1: required string accountID
    2: required string username
    3: required AccountStatus status
}

struct AccountQueryReq {
    1: required string accountID (vt.pattern="^\\w{8,128}$")
    
    255: required base.BaseReq base
}

struct AccountQueryResp {
    1: optional Account account

    255: required base.BaseResp base
}

struct AccountCreateReq {
    1: required string username (vt.pattern="^\\w{8,64}$")
    2: required string password (vt.pattern="^\\w{8,128}$")
   
    255: required base.BaseReq base
}

struct AccountCreateResp {
    1: optional string accountID
    255: required base.BaseResp base
}

struct AccountPasswordUpdateReq {
    1: required string accountID (vt.pattern="^\\w{8,128}$")
    2: required string password (vt.pattern="^\\w{8,128}$")
    3: required string newPassword (vt.pattern="^[0-9a-zA-Z.,_]{8,128}$")

    255: required base.BaseReq base
}

struct AccountPasswordUpdateResp {
    255: required base.BaseReq base
}

struct AccountPasswordResetReq {
    1: required string accountID (vt.pattern="^\\w{8,128}$")
    2: required string password (vt.pattern="^[0-9a-zA-Z.,_]{8,128}$")
  
    255: required base.BaseReq base
}

struct AccountPasswordResetResp { 
    255: required base.BaseResp base 
}

struct AccountStatusUpdateReq {
    1: required string accountID (vt.pattern="^\\w{8,128}$")
    2: required AccountStatus status (vt.in="AccountStatus.valid", vt.in="AccountStatus.invalid")
   
    255: required base.BaseReq base
}

struct AccountStatusUpdateResp {
    255: required base.BaseResp base
}

enum LoginStatus {
    success = 1
    fail = 2
}

struct LoginRecord {
    1: required string accountID
    2: required string ipv4
    3: required string device
    4: required LoginStatus status
    5: required string reason
    6: required i64 loginAt
}

struct AccountLoginReq {
    1: required string username (vt.pattern="^\\w{8,128}$")
    2: required string password (vt.pattern="^\\w{8,128}$")
    3: required string ipv4 (vt.pattern="((2(5[0-5]|[0-4]\\d))|[0-1]?\\d{1,2})(\\.((2(5[0-5]|[0-4]\\d))|[0-1]?\\d{1,2})){3}")
    4: required string device (vt.pattern="^\\w{1,128}$")
    255: required base.BaseReq base
}

struct AccountLoginResp {
    1: optional string AccountID
    255: required base.BaseResp base
}

struct LoginRecordQueryReq {
    1: required string accountID (vt.pattern="^\\w{8,128}$")
    2: required i64 page (vt.ge = "1", vt.le = "999")
    3: required i64 size (vt.ge = "1", vt.le = "999")

    255: required base.BaseReq base
}

struct LoginRecordQueryResp {
    1: optional list<LoginRecord> recordList
    2: optional i64 page 
    3: optional i64 size 
    4: optional i64 total

    255: required base.BaseResp base 
}

enum Gender {
    male = 1
    female = 2
    others = 3
}

struct User {
    1: required string accountID
    2: required string userID
    3: required string name
    4: required Gender gender
    5: required string phone
    6: required string email
    7: required string description
    
    200: required i64 createdAt
    201: required i64 updatedAt
}

struct UserCreateReq {
    1: required string accountID (vt.pattern="^\\w{8,128}$")
    2: required string name (vt.pattern="^[\\p{Han}a-zA-Z\\s]{1,128}$")
    3: required Gender gender (vt.in="Gender.male", vt.in="Gender.female", vt.in="Gender.others")
    4: required string phone (vt.pattern="^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\\d{8}$")
    5: required string email (vt.pattern="^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$")
    6: required string description (vt.max_size="256")

    255: required base.BaseResp base
}

struct UserCreateResp {
    1: optional string userID
    255: required base.BaseResp base
}

struct UserQueryReq {
    1: optional string userID (vt.pattern="^\\w{8,128}$")
    2: optional string accountID (vt.pattern="^\\w{8,128}$")

    255: required base.BaseResp base
}

struct UserQueryResp {
    1: optional User user

    255: required base.BaseResp base
}

struct UserUpdateReq {
    1: required string userID (vt.pattern="^\\w{8,128}$")
    2: optional string name (vt.pattern="^[\\p{Han}a-zA-Z\\s]{1,128}$")
    3: optional Gender gender (vt.in="Gender.male", vt.in="Gender.female", vt.in="Gender.others")
    4: optional string phone (vt.pattern="^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\\d{8}$")
    5: optional string email (vt.pattern="^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$")
    6: optional string description (vt.max_size="256")

    255: required base.BaseResp base
}

struct UserUpdateResp {
    255: required base.BaseResp base
}

service AccountService {
    AccountQueryResp QueryAccount(1: AccountQueryReq req)
    AccountCreateResp CreateAccount(1: AccountCreateReq req)
    AccountPasswordUpdateResp UpdateAccountPassword(1: AccountPasswordUpdateReq req)
    AccountPasswordResetResp ResetAccountPassword(1: AccountPasswordResetReq req)
    AccountStatusUpdateResp UpdateAccountStatus(1: AccountStatusUpdateReq req)
    AccountLoginResp Login(1: AccountLoginReq req)
    LoginRecordQueryResp QueryLoginRecord(1: LoginRecordQueryReq req)

    UserCreateResp CreateUser(1: UserCreateReq req)
    UserQueryResp QueryUser(1: UserQueryReq req)
    UserUpdateResp UpdateUser(1: UserUpdateReq req)
}