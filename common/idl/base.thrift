namespace go base

struct BaseReq {
    1: required string traceID
    2: required string caller

    255: required map<string, string> extra
}

struct BaseResp {
    1: required bool success
    2: required i32 code
    3: required string msg
    255: required map<string, string> extra
}