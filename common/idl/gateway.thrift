namespace go kaidog.shield.gateway

enum AccountStatus {
    valid = 1
    invalid = 2 
}

enum LoginStatus {
    success = 1
    fail = 2
}

enum Gender {
    male = 1
    female = 2
    others = 3
}

struct LoginReq {
    1: required string username (api.body="username");
    2: required string password (api.body="password");
}

struct LoginResp {
    1: i32 code (api.body="code");
    2: bool success (api.body="success");
    3: string msg (api.body="msg");

    4: string accountID (api.body="account_id");
}

struct LogoutReq {
    
}

struct LogoutResp {
    1: i32 code (api.body="code");
    2: bool success (api.body="success");
    3: string msg (api.body="msg");
}

struct SelfUserInfoQueryReq {
    
}

struct UserInfo {
    1: string accountID (api.body="account_id");
    2: string username (api.body="username");
    3: AccountStatus status (api.body="status");
    4: string userID (api.body="user_id");
    5: string name (api.body="name");
    6: Gender gender (api.body="gender");
    7: string phone (api.body="phone");
    8: string email (api.body="email");
    9: string description (api.body="description");
    10: i64 createdAt (api.body="created_at");
    11: i64 updatedAt (api.body="updated_at");
}

struct SelfUserInfoQueryResp {
    1: i32 code (api.body="code");
    2: bool success (api.body="success");
    3: string msg (api.body="msg");

    4: UserInfo info (api.body="info");
}

struct LoginRecord {
    1: string accountID (api.body="account_id");
    2: string ipv4 (api.body="ipv4");
    3: string device (api.body="device");
    4: LoginStatus status (api.body="status");
    5: string reason (api.body="reason");
    6: i64 loginAt (api.body="login_at");
}

struct SelfLoginRecordQueryReq {
     
}

struct SelfLoginRecordQueryResp{
    1: i32 code (api.body="code");
    2: bool success (api.body="success");
    3: string msg (api.body="msg");

    4: list<LoginRecord> LoginRecord (api.body="login_record");
    5: i64 total (api.body="total");
    6: i64 page (api.body="page");
    7: i64 size  (api.body="size");
}

struct PasswordUpdateReq {
     1: string oldPassword (api.body="old_password");
     2: string newPassword (api.body="new_password");
}

struct PasswordUpdateResp {
    1: i32 code (api.body="code");
    2: bool success (api.body="success");
    3: string msg (api.body="msg");
}

service UserService {
    LoginResp Login(1: LoginReq req) (api.post="/login");
    LogoutResp Logout(1: LogoutReq req) (api.post="/operator/logout");
    SelfUserInfoQueryResp QuerySelfUserInfo(1: SelfUserInfoQueryReq req) (api.get="/operator/user/user_info/query");
    SelfLoginRecordQueryResp QuerySelfLoginRecord(1: SelfLoginRecordQueryReq req) (api.get="/operator/user/login_record/query");
    PasswordUpdateResp UpdatePassword(1: PasswordUpdateReq req) (api.post="/operator/user/password/update");
}

struct Account { 
    1: required string AccountID (api.body="account_id");
    2: required string Username (api.body="username");
    3: required AccountStatus Status (api.body="status");
}

struct AccountCreateReq {
    1: required string Username (api.body="username");
    2: required string Password (api.body="password");
}

struct AccountCreateResp {
    1: i32 Code (api.body="code");
    2: bool Success (api.body="success");
    3: string Msg (api.body="msg");

    4: string AccountID (api.body="account_id");
}

struct AccountQueryReq {
    1: required i64 page (api.query="page");
    2: required i64 size (api.query="size");
}

struct AccountQueryResp {
    1: i32 code (api.body="code");
    2: bool success (api.body="success");
    3: string msg (api.body="msg");

    4: list<Account> accountList (api.body="account_list");
    5: i64 total (api.body="total");
    6: i64 page (api.body="page");
    7: i64 size  (api.body="size");
}

struct UserInfoQueryReq {
    1: list<string> AccountIdList (api.body="account_id_list");
}

struct UserInfoQueryResp {
    1: i32 code (api.body="code");
    2: bool success (api.body="success");
    3: string msg (api.body="msg");

    4: list<UserInfo> userList (api.body="user_list");
}

struct LoginRecordQueryReq {
    1: required string accountID (api.body="account_id");
}

struct LoginRecordQueryResp{
    1: i32 code (api.body="code");
    2: bool success (api.body="success");
    3: string msg (api.body="msg");

    4: list<LoginRecord> LoginRecord (api.body="login_record");
    5: i64 total (api.body="total");
    6: i64 page (api.body="page");
    7: i64 size  (api.body="size");
}

struct UserCreateReq {
    1: required string accountID (api.body="account_id");
    2: required string name (api.body="name");
    3: Gender gender (api.body="gender");
    4: required string phone (api.body="phone");
    5: required string email (api.body="email");
    6: required string description (api.body="description");
}

struct UserCreateResp {
    1: i32 code (api.body="code");
    2: bool success (api.body="success");
    3: string msg (api.body="msg");

    4: string userID (api.body="user_id");
}

struct UserInfoUpdateReq {
    1: required string userID (api.body="user_id");
    2: optional string name (api.body="name");
    3: optional Gender gender (api.body="gender");
    4: optional string phone (api.body="phone");
    5: optional string email (api.body="email");
    6: optional string description (api.body="description"); 
}

struct UserInfoUpdateResp {
    1: i32 code (api.body="code");
    2: bool success (api.body="success");
    3: string msg (api.body="msg");
}

struct PasswordRestReq {
    1: required string accountID (api.body="account_id");
    2: required string newPassword (api.body="new_password"); 
}

struct PasswordRestResp {
    1: i32 code (api.body="code");
    2: bool success (api.body="success");
    3: string msg (api.body="msg");
}

struct AccountStatusSwitchReq {
    1: required string accountID (api.body="account_id");
    2: required AccountStatus status (api.body="status");
}

struct AccountStatusSwitchResp {
    1: i32 code (api.body="code");
    2: bool success (api.body="success");
    3: string msg (api.body="msg");
}

service AdminService {
    AccountCreateResp CreateAccount(1: AccountCreateReq req) (api.post="/operator/admin/account/create");
    AccountQueryResp QueryAccount(1: AccountQueryReq req) (api.get="/operator/admin/account/query");
    UserInfoQueryResp QueryUserInfo(1: UserInfoQueryReq req) (api.get="/operator/admin/user_info/query");
    LoginRecordQueryResp QueryLoginRecord(1: LoginRecordQueryReq req) (api.get="/operator/admin/login_record/query");
    UserCreateResp CreateUser(1: UserCreateReq req) (api.post="/operator/admin/user_info/create");
    UserInfoUpdateResp UpdateUserInfo(1: UserInfoUpdateReq req) (api.post="/operator/admin/user_info/update");
    PasswordRestResp ResetPassword(1: PasswordRestReq req) (api.post="/operator/admin/password/reset");
    AccountStatusSwitchResp SwitchAccountStatus(1: AccountStatusSwitchReq req) (api.post="/operator/admin/account_status/change");
}