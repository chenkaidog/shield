package repos

import "shield/gateway/biz/repos/redis"

func Init() {
	redis.InitReis()
}