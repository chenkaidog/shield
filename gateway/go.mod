module shield/gateway

go 1.20

replace (
	github.com/apache/thrift => github.com/apache/thrift v0.13.0
	shield/common => ../common
)

require (
	github.com/apache/thrift v0.19.0
	github.com/cloudwego/hertz v0.7.3
	github.com/cloudwego/kitex v0.8.0
	github.com/gorilla/securecookie v1.1.2
	github.com/hertz-contrib/csrf v0.1.1
	github.com/mssola/user_agent v0.6.0
	shield/common v0.0.0-00010101000000-000000000000
)

require gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect

require (
	github.com/andeya/ameda v1.5.3 // indirect
	github.com/andeya/goutil v1.0.1 // indirect
	github.com/bufbuild/protocompile v0.7.1 // indirect
	github.com/bytedance/go-tagexpr/v2 v2.9.11 // indirect
	github.com/bytedance/gopkg v0.0.0-20230728082804-614d0af6619b // indirect
	github.com/bytedance/sonic v1.10.2 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20230717121745-296ad89f973d // indirect
	github.com/chenzhuoyu/iasm v0.9.1 // indirect
	github.com/choleraehyq/pid v0.0.17 // indirect
	github.com/cloudwego/configmanager v0.2.0 // indirect
	github.com/cloudwego/dynamicgo v0.1.6 // indirect
	github.com/cloudwego/fastpb v0.0.4 // indirect
	github.com/cloudwego/frugal v0.1.12 // indirect
	github.com/cloudwego/localsession v0.0.2 // indirect
	github.com/cloudwego/netpoll v0.5.1 // indirect
	github.com/cloudwego/thriftgo v0.3.4 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/fatih/structtag v1.2.0 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/gomodule/redigo v1.8.9 // indirect
	github.com/google/pprof v0.0.0-20231212022811-ec68065c825e // indirect
	github.com/google/uuid v1.5.0 // indirect
	github.com/gorilla/context v1.1.2 // indirect
	github.com/gorilla/sessions v1.2.2 // indirect
	github.com/hertz-contrib/sessions v1.0.3
	github.com/iancoleman/strcase v0.3.0 // indirect
	github.com/jhump/protoreflect v1.15.4 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.6 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/gls v0.0.0-20220109145502-612d0167dce5 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/nyaruka/phonenumbers v1.3.0 // indirect
	github.com/oleiade/lane v1.0.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/redis/go-redis/v9 v9.3.1
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	github.com/tidwall/gjson v1.17.0 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	golang.org/x/arch v0.6.0 // indirect
	golang.org/x/net v0.19.0 // indirect
	golang.org/x/sync v0.5.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231212172506-995d672761c0 // indirect
	google.golang.org/grpc v1.60.0 // indirect
	google.golang.org/protobuf v1.31.1-0.20231027082548-f4a6c1f6e5c1 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/yaml.v3 v3.0.1
)
