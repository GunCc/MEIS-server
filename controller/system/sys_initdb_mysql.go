package system

import (
	"MEIS-server/config"
	"MEIS-server/global"
	"MEIS-server/model/system/request"
	"MEIS-server/utils"
	"context"
	"errors"
	"fmt"

	"github.com/gookit/color"
	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlInitHandler struct{}

func NewMysqlInitHandler() *MysqlInitHandler {
	return &MysqlInitHandler{}
}

// 创建数据库并且初始化
func (m *MysqlInitHandler) EnsureDB(ctx context.Context, conf *request.InitDB) (next context.Context, err error) {
	if s, ok := ctx.Value("dbtype").(string); !ok || s != "mysql" {
		return ctx, ErrDBTypeMismatch
	}

	dsn := conf.MysqlEmptyDsn()

	c := conf.ToMysqlConfig()
	next = context.WithValue(ctx, "config", c)

	// 创建数据库
	createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", c.DbName)

	if err = createDatabase(dsn, "mysql", createSql); err != nil {
		return nil, err
	}

	var db *gorm.DB

	fmt.Println("c.Dsn()", c.Dsn(), dsn)
	if db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       c.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: true,    // 根据版本自动配置
	}), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}); err != nil {
		return ctx, err
	}
	next = context.WithValue(next, "db", db)

	return next, err
}

// mysql 回写配置
func (m *MysqlInitHandler) WriteConfig(ctx context.Context) (err error) {

	c, ok := ctx.Value("config").(config.Mysql)
	if !ok {
		return errors.New("mysql config invalid")
	}
	global.MEIS_CONFIG.System.DbType = "mysql"
	global.MEIS_CONFIG.Mysql = c
	global.MEIS_CONFIG.JWT.SigningKey = uuid.NewV4().String()
	cs := utils.StructToMap(global.MEIS_CONFIG)

	fmt.Println("StructToMap出来的结果：", cs)
	for k, v := range cs {
		global.MEIS_Viper.Set(k, v)
	}
	return global.MEIS_Viper.WriteConfig()
}

// 创建表
func (m *MysqlInitHandler) InitTables(ctx context.Context, inits initSlice) (err error) {
	return createTables(ctx, inits)
}

// 创建数据
func (m *MysqlInitHandler) InitData(ctx context.Context, inits initSlice) (err error) {
	next, cancel := context.WithCancel(ctx)

	defer func(c func()) { c() }(cancel)
	for _, init := range inits {
		if init.DataInserted(next) {
			color.Info.Printf(InitDataExist, Mysql, init.InitializerName())
			continue
		}

		if n, err := init.InitializeData(next); err != nil {
			color.Info.Printf(InitDataFailed, Mysql, init.InitializerName(), err)
			return err
		} else {
			next = n
			color.Info.Printf(InitDataSuccess, Mysql, init.InitializerName())
		}
	}

	color.Info.Printf(InitSuccess, Mysql)
	return nil
}
