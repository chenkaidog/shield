package mysql

import (
	"database/sql"
	"fmt"
	"shield/account/internal/config"
	"shield/common/utils/gorm_utils"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var gormDB *gorm.DB

func Init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.GetMySQLConf().Username,
		config.GetMySQLConf().Password,
		config.GetMySQLConf().IP,
		config.GetMySQLConf().Port,
		config.GetMySQLConf().DBName,
	)

	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	gormDB, err = gorm.Open(
		mysql.New(mysql.Config{Conn: sqlDB}),
		&gorm.Config{
			SkipDefaultTransaction: true,
			Logger: &gorm_utils.GormLogger{
				SlowThreshold: 2 * time.Second,
				LogLevel:      logger.Info,
			},
		})
	if err != nil {
		panic(err)
	}
}

func GetGormDB() *gorm.DB {
	return gormDB
}

type Dal interface {
	GetGormDB() *gorm.DB
	SetGomDB(gormDB *gorm.DB)
}

type defaultDal struct {
	gormDB *gorm.DB
}

func (dal *defaultDal) GetGormDB() *gorm.DB {
	return dal.gormDB
}

func (dal *defaultDal) SetGomDB(gormDB *gorm.DB) {
	dal.gormDB = gormDB
}

func NewDefaultDal(gormDB ...*gorm.DB) Dal {
	dal := &defaultDal{}

	if len(gormDB) > 0 {
		dal.SetGomDB(gormDB[0])
	} else {
		dal.SetGomDB(GetGormDB())
	}

	return dal
}
