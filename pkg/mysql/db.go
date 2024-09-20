package mysql

import (
	"fmt"
	"github.com/caohui123/goweb/internal/app/model"
	"github.com/caohui123/goweb/pkg/config"
	lg "github.com/caohui123/goweb/pkg/logger"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"sync"
	"time"
)

var MysqlDb *gorm.DB
var once sync.Once

func InitMysql(cfg *config.DBConfig) *gorm.DB {
	once.Do(func() {
		db, err := gorm.Open(mysql.New(mysql.Config{
			DSN:                       cfg.DSN, // DSN data source name
			DefaultStringSize:         256,     // string 类型字段的默认长度
			DisableDatetimePrecision:  true,    // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
			DontSupportRenameIndex:    true,    // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
			DontSupportRenameColumn:   true,    // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
			SkipInitializeWithVersion: false,   // 根据版本自动配置
		}), GetGormConfig(cfg))
		MysqlDb = db
		if err != nil {
			log.Fatalln(err)
		}
	})
	MysqlDb.AutoMigrate(&model.User{})
	return MysqlDb
}

// 得到配置日志文件
func GetGormConfig(cfg *config.DBConfig) *gorm.Config {
	newLogger := logger.New(
		//log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		NewMyWriter(), //记录在文件
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,       // 禁用彩色打印
		},
	)
	if cfg.LogMode {
		return &gorm.Config{
			Logger: newLogger,
		}
	}
	return &gorm.Config{}

}

// 定义自己的Writer
type MyWriter struct {
	mlog *zap.Logger
}

// 实现gorm/logger.Writer接口
func (m *MyWriter) Printf(format string, v ...interface{}) {
	logstr := fmt.Sprintf(format, v...)
	//利用loggus记录日志
	m.mlog.Info(logstr)
}

func NewMyWriter() *MyWriter {
	log := lg.GetDbLogger()
	return &MyWriter{mlog: log}
}
