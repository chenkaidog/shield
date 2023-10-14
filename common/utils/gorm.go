package utils

import (
	"context"
	"errors"
	"shield/common/logs"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/mattn/go-sqlite3"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GormLogger struct {
	SlowThreshold time.Duration
	LogLevel      logger.LogLevel
}

func (l *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	l.LogLevel = level

	return l
}

func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	logs.CtxInfo(ctx, "GORM LOG %s %+v", msg, data)
}

func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	logs.CtxWarn(ctx, "GORM LOG %s %+v", msg, data)
}

func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	logs.CtxError(ctx, "GORM LOG %s %+v", msg, data)
}

func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel > logger.Silent {
		costDuration := time.Since(begin)
		cost := float64(costDuration.Nanoseconds()/1e4) / 100.0
		switch {

		// err hapends and log level is greater than 'Error'. if we shold ignore data not found err
		case err != nil && l.LogLevel >= logger.Error && !errors.Is(err, gorm.ErrRecordNotFound) && !IsEntryDuplicateErr(err):
			sql, _ := fc()
			logs.CtxError(ctx, "GORM LOG: %s, Err: %s, Cost: %.2fms", sql, err.Error(), cost)

		// slow SQL exec hapends and level is greater than 'Warn'
		case l.LogLevel >= logger.Warn && costDuration > l.SlowThreshold && l.SlowThreshold > 0:
			sql, rows := fc()
			logs.CtxWarn(ctx, "GORM LOG SLOW SQL: %s, Rows: %d, Cost: %.2fms, Limit: %s", sql, rows, cost, l.SlowThreshold)

		// normal SQL record
		case l.LogLevel >= logger.Info:
			sql, rows := fc()
			logs.CtxInfo(ctx, "GORM LOG SQL: %s, Rows: %d, Cost: %.2fms", sql, rows, cost)
		}
	}
}

func IsEntryDuplicateErr(err error) bool {
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		if mysqlErr.Number == 1062 {
			return true
		}
	}

	if sqliteErr, ok := err.(sqlite3.Error); ok {
		if sqliteErr.Code == sqlite3.ErrConstraint {
			return true
		}
	}

	return false
}
