package initialize

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hespecial/gin-mall/config"
	"github.com/hespecial/gin-mall/global"
	"github.com/hespecial/gin-mall/internal/model"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"time"
)

func InitMySQL() *gorm.DB {
	master := global.Config.MySQL["master"]
	slave := global.Config.MySQL["slave"]
	masterDsn := getDSN(master)
	slaveDsn := getDSN(slave)

	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       masterDsn, // DSN data source name
		DefaultStringSize:         256,       // string 类型字段的默认长度
		DisableDatetimePrecision:  true,      // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,      // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,      // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,     // 根据版本自动配置
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(fmt.Sprintf("Fatal error gorm Open: %v", err))
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)  // 设置连接池，空闲
	sqlDB.SetMaxOpenConns(100) // 打开
	sqlDB.SetConnMaxLifetime(time.Second * 30)

	slowQueryLog(db)
	gormRateLimiter(db, rate.NewLimiter(500, 1000))

	err = db.Use(dbresolver.Register(dbresolver.Config{
		Sources:           []gorm.Dialector{mysql.Open(masterDsn)},
		Replicas:          []gorm.Dialector{mysql.Open(slaveDsn)},
		Policy:            dbresolver.RandomPolicy{},
		TraceResolverMode: true,
	}))
	if err != nil {
		panic(fmt.Sprintf("Fatal error gorm Use: %v", err))
	}

	err = migrate(db)
	if err != nil {
		panic(fmt.Sprintf("Fatal error gorm AutoMigrate: %v", err))
	}

	return db
}

func getDSN(m *config.MySQL) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.Username,
		m.Password,
		m.Host,
		m.Port,
		m.Database,
	)
}

func migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&model.User{},
		&model.Category{},
		&model.Product{}, &model.ProductImage{},
		&model.Cart{}, &model.CartItem{},
		// &model.Order{}, &model.Address{},
		// &model.Notice{},
		// &model.SkillProduct{},
		// &model.SkillProduct2MQ{},
	)
	return err
}

// 慢查询日志
func slowQueryLog(db *gorm.DB) {
	err := db.Callback().Query().Before("*").Register("slow_query_start", func(d *gorm.DB) {
		now := time.Now()
		d.Set("start_time", now)
	})
	if err != nil {
		panic(fmt.Sprintf("Fatal error gorm start SlowQueryLog: %v", err))
	}

	err = db.Callback().Query().After("*").Register("slow_query_end", func(d *gorm.DB) {
		now := time.Now()
		start, ok := d.Get("start_time")
		if ok {
			duration := now.Sub(start.(time.Time))
			// 一般认为 200 Ms 为Sql慢查询
			if duration > time.Millisecond*200 {
				global.Log.Warn("Slow SQL", zap.String("sql", d.Statement.SQL.String()))
			}
		}
	})
	if err != nil {
		panic(fmt.Sprintf("Fatal error gorm end SlowQueryLog: %v", err))
	}
}

var GormTooManyRequestError = errors.New("gorm: too many request")

// Gorm限流器 此限流器不能终止GORM查询链
func gormRateLimiter(db *gorm.DB, r *rate.Limiter) {
	err := db.Callback().Query().Before("*").Register("RateLimitGormMiddleware", func(d *gorm.DB) {
		if !r.Allow() {
			_ = d.AddError(GormTooManyRequestError)
			global.Log.Error(GormTooManyRequestError.Error())
			return
		}
	})
	if err != nil {
		panic(fmt.Sprintf("Fatal error gorm RateLimitGormMiddleware: %v", err))
	}
}
