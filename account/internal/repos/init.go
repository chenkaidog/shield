package repos

import "shield/account/internal/repos/mysql"

func Init() {
	mysql.Init()
}
