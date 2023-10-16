include "base.thrift"

namespace go kaidog.shield.rbac

enum UserRoleStatus {
    valid = 1
    invalid = 2
    expired = 3
}

struct UserRoleRelation {
    1: required string RelationID 
    2: required string userID
    3: required string roleID
    4: required UserRoleStatus status
    6: required i64 expireAt
    7: required i64 createdAt
    8: required i64 updatedAt
}

struct UserRoleRelationCreateReq {
    1: required string userID (vt.pattern="^\\w{8,128}$")
    2: required string roleID (vt.pattern="^\\w{8,128}$")
    3: required i64 expireAt (vt.gt="0")
    5: required string opeatorID (vt.pattern="^\\w{8,128}$")

    255: required base.BaseReq base 
}

struct UserRoleRelationCreateResp {
    255: required base.BaseResp base
}

struct UserRoleRelationUpdateReq {
    1: required string RelationID (vt.pattern="^\\w{8,128}$")
    2: optional i64 expireAt (vt.gt="0")
    3: required string opeatorID (vt.pattern="^\\w{8,128}$")

    255: required base.BaseReq base 
}

struct UserRoleRelationUpdateResp {
    255: required base.BaseResp base
}

struct UserRoleStatusUpdateReq {
    1: required string RelationID (vt.pattern="^\\w{8,128}$")
    2: optional UserRoleStatus status (vt.in="UserRoleStatus.valid", vt.in="UserRoleStatus.invalid")
    3: required string opeatorID (vt.pattern="^\\w{8,128}$")

    255: required base.BaseReq base 
}

struct UserRoleStatusUpdateResp {
    255: required base.BaseResp base
}

struct UserRoleRelationDeleteReq {
    1: required string RelationID (vt.pattern="^\\w{8,128}$")
    2: required string opeatorID (vt.pattern="^\\w{8,128}$")
    255: required base.BaseReq base 
}

struct UserRoleRelationDeleteResp {
    255: required base.BaseResp base
}

struct UserRoleRelationQueryReq {
    1: optional string userID (vt.pattern="^\\w{8,128}$")
    2: optional string roleID (vt.pattern="^\\w{8,128}$")
    
    10: required i64 page (vt.gt="0", vt.le="999")
    11: required i64 size (vt.gt="0", vt.le="999")

    255: required base.BaseReq base 
}

struct UserRoleRelationQueryResp {
    1: optional list<UserRoleRelation> RelationList
    2: optional i64 page
    3: optional i64 size 
    4: optional i64 total

    255: required base.BaseResp base
}

enum RoleStatus {
    valid = 1
    invalid = 2
}

enum RoleLevel {
    public = 1
    internal = 2
    secret = 3
    confidential = 4
}

struct Role {
    1: required string roleID 
    2: required string domain
    3: required string parentRoleID
    4: required i64 maximumApplicant 
    5: required i64 maxValidDay
    6: required string name
    7: required RoleStatus status 
    8: required RoleLevel level 
    9: required string description
    10: required string ownerID 
    11: optional list<string> managerIDList 
    12: required i64 createdAt
    13: required i64 updatedAt
}

struct PermissionInRole {
    1: required string resource (vt.max_size="128")
    2: required string action (vt.pattern="^\\w{8,128}$")
    3: required PermissionStatus status (vt.in="PermissionStatus.allow", vt.in="PermissionStatus.deny")
}

struct RoleCreateReq {
    1: optional string parentRoleID (vt.pattern="^\\w{8,128}$")
    2: required string domain (vt.pattern="^[\\w.]{1,64}$")
    3: required i64 maximumApplicant (vt.not_in="0")
    4: required i64 maxValidDay (vt.not_in="0")
    5: required string name (vt.pattern="^[\\p{Han}a-zA-Z\\s]{1,128}$")
    6: required RoleStatus status (vt.in="RoleStatus.valid", vt.in="RoleStatus.invalid")
    7: required RoleLevel level (vt.in="RoleLevel.public", vt.in="RoleLevel.internal", vt.in="RoleLevel.secret", vt.in="RoleLevel.confidential")
    8: required string description (vt.max_size="256")
    9: required string ownerID (vt.pattern="^\\w{8,128}$")
    10: optional list<string> managerIDList (vt.elem.pattern="^\\w{8,128}$")
    11: required string opeatorID (vt.pattern="^\\w{8,128}$")
    12: optional list<PermissionInRole> permissionList (vt.elem.skip="false")

    255: required base.BaseReq base 
}

struct RoleCreateResp {

    255: required base.BaseResp base
}

struct RoleUpdateReq {
    1: required string roleID (vt.pattern="^\\w{8,128}$")
    2: optional i64 maximumApplicant (vt.not_in="0")
    3: optional i64 maxValidDay (vt.not_in="0")
    4: optional string name (vt.pattern="^[\\p{Han}a-zA-Z\\s]{1,128}$")
    5: optional RoleLevel level (vt.in="RoleLevel.public", vt.in="RoleLevel.internal", vt.in="RoleLevel.secret", vt.in="RoleLevel.confidential")
    6: optional string description (vt.max_size="256")
    7: required string opeatorID (vt.pattern="^\\w{8,128}$")

    255: required base.BaseReq base
}

struct RoleUpdateResp {

    255: required base.BaseResp base
}

struct RoleStatusUpdateReq {
    1: required string roleID (vt.pattern="^\\w{8,128}$")
    2: optional RoleStatus status (vt.in="RoleStatus.valid", vt.in="RoleStatus.invalid")
    3: required string opeatorID (vt.pattern="^\\w{8,128}$")

    255: required base.BaseReq base
}

struct RoleStatusUpdateResp {

    255: required base.BaseResp base
}

struct RoleDeleteReq {
    1: required string roleID (vt.pattern="^\\w{8,128}$")
    2: required string opeatorID (vt.pattern="^\\w{8,128}$")
    255: required base.BaseReq base
}

struct RoleDeleteResp {
    255: required base.BaseResp base
}

struct RoleQueryReq {
    1: optional string roleID (vt.pattern="^\\w{8,128}$")
    2: optional string parentRoleID (vt.pattern="^\\w{8,128}$")
    3: optional string domain (vt.pattern="^[\\w.]{1,64}$")

    10: required i64 page (vt.gt="0", vt.le="999")
    11: required i64 size (vt.gt="0", vt.le="999")

    255: required base.BaseReq base
}

struct RoleQueryResp {
    1: optional list<Role> roleList
    2: optional i64 page
    3: optional i64 size
    4: optional i64 total

    255: required base.BaseResp base
}

struct RoleOwnerChangeReq {
    1: required string userID (vt.pattern="^\\w{8,128}$")
    2: required string roleID (vt.pattern="^\\w{8,128}$")
    3: required string opeatorID (vt.pattern="^\\w{8,128}$")

    255: required base.BaseReq base
}

struct RoleOwnerChangeResp {
    255: required base.BaseResp base
}

struct RoleManagerAppendReq {
    1: required string userID (vt.pattern="^\\w{8,128}$")
    2: required string roleID (vt.pattern="^\\w{8,128}$")
    3: required string opeatorID (vt.pattern="^\\w{8,128}$")

    255: required base.BaseReq base
}

struct RoleManagerAppendResp {
    255: required base.BaseResp base
}

struct RoleManagerRemoveReq {
    1: required string userID (vt.pattern="^\\w{8,128}$")
    2: required string roleID (vt.pattern="^\\w{8,128}$")
    3: required string opeatorID (vt.pattern="^\\w{8,128}$")

    255: required base.BaseReq base
}

struct RoleManagerRemoveResp {
    255: required base.BaseResp base
}

enum PermissionStatus {
    allow = 1
    deny = 2
}

struct Permission {
    1: required string permissionID
    2: required string roleID
    3: required string resource
    4: required string action 
    5: required PermissionStatus status
    7: required i64 createdAt
    8: required i64 updatedAt
}

struct PermissionCreateReq {
    1: required string roleID (vt.pattern="^\\w{8,128}$")
    2: required string resource (vt.max_size="128")
    3: required string action (vt.pattern="^\\w{8,128}$")
    4: required PermissionStatus status (vt.in="PermissionStatus.allow", vt.in="PermissionStatus.deny")
    5: required string opeatorID (vt.pattern="^\\w{8,128}$")

    255: required base.BaseReq base
}

struct PermissionCreateResp {

    255: required base.BaseResp base
}

struct PermissionUpdateReq {
    1: required string permissionID (vt.pattern="^\\w{8,128}$")
    2: optional string resource (vt.max_size="128")
    3: optional string action (vt.pattern="^\\w{8,128}$")
    4: required string opeatorID (vt.pattern="^\\w{8,128}$")

    255: required base.BaseReq base
}

struct PermissionUpdateResp {
    255: required base.BaseResp base
}

struct PermissionStatusUpdateReq {
    1: required string permissionID (vt.pattern="^\\w{8,128}$")
    2: optional PermissionStatus status (vt.in="PermissionStatus.allow", vt.in="PermissionStatus.deny", vt.in="PermissionStatus.deleted")
    3: required string opeatorID (vt.pattern="^\\w{8,128}$")

    255: required base.BaseReq base
}

struct PermissionStatusUpdateResp {
    255: required base.BaseResp base
}

struct PermissionDeleteReq {
    1: required string permissionID (vt.pattern="^\\w{8,128}$")
    2: required string opeatorID (vt.pattern="^\\w{8,128}$")

    255: required base.BaseReq base
}

struct PermissionDeleteResp {
    255: required base.BaseResp base
}

struct PermissionQueryReq {
    1: required string roleID (vt.pattern="^\\w{8,128}$")
    
    255: required base.BaseReq base
}

struct PermissionQueryResp {
    1: optional list<Permission> permissionList

    255: required base.BaseResp base
}

enum DataType {
    unknown = 0
    userRole = 1
    role =2 
    permission = 3
}

enum Operation {
    unknown = 0
    create = 1
    update = 2 
    delete = 3 
}

struct OperationRecord {
    1: required string recordID 
    2: required DataType dataType
    3: required string dataID
    4: required Operation operation 
    5: required string operatorID
    6: required string traceID
    7: required string previousValue
    8: required string currentValue
    9: required i64 createdAt 
}

struct OperationRecordQueryReq {
    1: optional string recordID
    2: optional string dataID
    3: optional string operatorID

    10: required i64 page (vt.gt="0", vt.le="999")
    11: required i64 size (vt.gt="0", vt.le="999")
    12: required bool createdAsc

    255: required base.BaseReq base
}

struct OperationRecordQueryResp {
    1: optional list<OperationRecord> recordList
    2: optional i64 page 
    3: optional i64 size
    4: optional i64 total

    255: required base.BaseResp base
}

struct AccessReq {
    1: required string userID (vt.pattern="^\\w{8,128}$")
    2: required string domain (vt.pattern="^[\\w.]{1,64}$")
    3: required string resource (vt.max_size="128")
    4: required string action (vt.pattern="^\\w{8,128}$")

    255: required base.BaseReq base
}

struct AccessResp {
    1: required bool licensed

    255: required base.BaseResp base
}

service RbacService {
    UserRoleRelationCreateResp CreateUserRoleRelation(1: UserRoleRelationCreateReq req)
    UserRoleRelationUpdateResp UpdateUserRoleRelation(1: UserRoleRelationUpdateReq req)
    UserRoleRelationDeleteResp DeleteUserRoleRelation(1: UserRoleRelationDeleteReq req)
    UserRoleRelationQueryResp QueryUserRoleRelation(1: UserRoleRelationQueryReq req)
    UserRoleStatusUpdateResp UpdateUserRoleStatus(1: UserRoleStatusUpdateReq req)

    RoleCreateResp CreateRole(1: RoleCreateReq req)
    RoleUpdateResp UpdateRole(1: RoleUpdateReq req)
    RoleDeleteResp DeleteRole(1: RoleDeleteReq req)
    RoleQueryResp QueryRole(1: RoleQueryReq req)
    RoleOwnerChangeResp ChangeRoleOwner(1: RoleOwnerChangeReq req)
    RoleManagerAppendResp AppendRoleManager(1: RoleManagerAppendReq req)
    RoleManagerAppendReq RemoveRoleManager(1: RoleManagerRemoveReq req)
    RoleStatusUpdateResp UpdateRoleStatus(1: RoleStatusUpdateReq req)

    PermissionCreateResp CreatePermission(1: PermissionCreateReq req)
    PermissionQueryResp QueryPermission(1: PermissionQueryReq req)
    PermissionUpdateResp UpdatePermission(1: PermissionUpdateReq req)
    PermissionDeleteResp DeletePermission(1: PermissionDeleteReq req)
    PermissionStatusUpdateResp UpdatePermissionStatus(1: PermissionStatusUpdateReq req)

    OperationRecordQueryResp QueryOperationRecord(1: OperationRecordQueryReq req)

    AccessResp Access(1: AccessReq req)
}
