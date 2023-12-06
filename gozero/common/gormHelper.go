package common

import (
	"context"
	"errors"
	"go-cms/common/utils"
	"time"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbConf struct {
	DNS             string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxIdleTime int
	ConnMaxLifetime int
	//SlowThresholdMillisecond int64
}

type GormLogger struct {
	SlowThreshold time.Duration
	Model         string // pro dev
}

func NewGormLogger(mode string) *GormLogger {
	return &GormLogger{
		SlowThreshold: 200 * time.Millisecond, // 一般超过200毫秒就算慢查所以不使用配置进行更改
		Model:         mode,
	}
}

var _ logger.Interface = (*GormLogger)(nil)

func (l *GormLogger) LogMode(lev logger.LogLevel) logger.Interface {
	return &GormLogger{}
}
func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	//logc.Infof(ctx, msg, data)
	logc.Errorf(ctx, msg, data)
}
func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	//logx.WithContext(ctx).Errorf(msg, data)
	logc.Errorf(ctx, msg, data)
}
func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	//logx.WithContext(ctx).Errorf(msg, data)
	logc.Errorf(ctx, msg, data)
}
func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	// 获取运行时间
	elapsed := time.Since(begin)
	// 获取 SQL 语句和返回条数
	sql, rows := fc()
	// 通用字段
	logFields := []logx.LogField{
		logx.Field("sql", sql),
		logx.Field("time", utils.MicrosecondsStr(elapsed)),
		logx.Field("rows", rows),
	}
	// Gorm 错误
	if err != nil {
		// 记录未找到的错误使用 warning 等级
		if errors.Is(err, gorm.ErrRecordNotFound) {
			//logx.WithContext(ctx).Infow("Database ErrRecordNotFound", logFields...)
			logc.Errorw(ctx, "DB未找到", logFields...)
		} else {
			// 其他错误使用 error 等级
			logFields = append(logFields, logx.Field("catch error", err))
			logc.Errorw(ctx, "DB错误", logFields...)
			//logx.WithContext(ctx).Errorw("Database Error", logFields...)
		}
	} else {
		// 非生产模式下，记录所有 SQL 请求
		if l.Model != service.ProMode {
			//logx.WithContext(ctx).Infow("Database Query", logFields...)
			logc.Errorw(ctx, "DB查询", logFields...)
		}
		//	logc.Infow(ctx, "DB查询", logFields...)
	}
	// 慢查询日志
	if l.SlowThreshold != 0 && elapsed > l.SlowThreshold {
		//logx.WithContext(ctx).Sloww("数据库 Slow Log", logFields...)
		logc.Sloww(ctx, "DB慢查询", logFields...)
	}

}
