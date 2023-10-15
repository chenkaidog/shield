namespace go base

struct BaseReq {
    1: string logID
    2: string traceID
    3: string spanID

    255: map<string, string> extra
}

struct BaseResp {
    1: bool success
    2: i32 code
    3: string msg
    255: map<string, string> extra
}