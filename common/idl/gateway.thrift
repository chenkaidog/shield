namespace go kaidog.shield.gateway

struct BaseResp {
    1: i32 code (api.body="code");
    2: bool success (api.body="success");
    3: string msg (api.body="msg");
}

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
    1: string accountID (api.body="account_id");
}

struct LogoutReq {
    1: required string accountID (api.body="account_id");
}

struct UserInfoQueryReq {
    1: required string accountID (api.query="account_id");
}

struct UserInfoQueryResp {
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

struct LoginRecord {
    1: string accountID (api.body="account_id");
    2: string ipv4 (api.body="ipv4");
    3: string device (api.body="device");
    4: LoginStatus status (api.body="status");
    5: string reason (api.body="reason");
    6: i64 loginAt (api.body="login_at");
}

struct LoginRecordQueryReq {
     1: required string accountID (api.query="account_id");
}

struct LoginRecordQueryResp{
    1: list<LoginRecord> LoginRecord (api.body="login_record");
    2: i64 total (api.body="total");
    3: i64 page (api.body="page");
    4: i64 size  (api.body="page");
}

struct PasswordUpdateReq {
     1: string accountID (api.body="account_id");
     2: string oldPassword (api.body="old_password");
     3: string newPassword (api.body="new_password");
}

service UserService {
    BaseResp Login(1: LoginReq req) (api.post="/user/login");
    BaseResp Logout(1: LogoutReq req) (api.post="/user/logout");
    BaseResp QueryUserInfo(1: UserInfoQueryReq req) (api.get="/user/query_user_info");
    BaseResp QueryLoginRecord(1: LoginRecordQueryReq req) (api.get="/user/query_login_record");
    BaseResp UpdatePassword(1: PasswordUpdateReq req) (api.post="/user/update_password");
}

struct Account { 
    1: required string accountID (api.body="account_id");
    2: required string username (api.body="username");
    3: required AccountStatus status (api.body="status");
}

struct AccountCreateReq {
    1: required string username (api.body="username");
    2: required string password (api.body="password");
}

struct AccountCreateResp {
    1: string accountID (api.body="account_id");
}

struct AccountQueryReq {
    1: required i64 page (api.query="page");
    2: required i64 size (api.query="size");
}

struct AccountQueryResp {
    1: list<Account> accountList (api.body="account_list");
    2: i64 total (api.body="total");
    3: i64 page (api.body="page");
    4: i64 size  (api.body="page");
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
    1: string userID (api.body="user_id");
}

struct UserInfoUpdateReq {
    1: required string userID (api.body="user_id");
    2: optional string name (api.body="name");
    3: optional Gender gender (api.body="gender");
    4: optional string phone (api.body="phone");
    5: optional string email (api.body="email");
    6: optional string description (api.body="description"); 
}

struct PasswordRestReq {
    1: required string accountID (api.body="account_id");
    2: required string newPassword (api.body="new_password"); 
}

struct AccountStatusSwitchReq {
    1: required string accountID (api.body="account_id");
    2: required AccountStatus status (api.body="status");
}

service AdminService {
    BaseResp CreateAccount(1: AccountCreateReq req) (api.post="/admin/create_account");
    BaseResp QueryAccount(1: AccountQueryReq req) (api.get="/admin/query_account");
    BaseResp CreateUser(1: UserCreateReq req) (api.post="/admin/create_user");
    BaseResp UpdateUserInfo(1: UserInfoUpdateReq req) (api.post="/admin/update_user");
    BaseResp ResetPassword(1: PasswordRestReq req) (api.post="/admin/rest_password");
    BaseResp SwitchAccountStatus(1: AccountStatusSwitchReq req) (api.post="/admin/switch_account_status");
}