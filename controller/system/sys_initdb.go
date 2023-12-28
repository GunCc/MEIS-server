package system

import (
	"MEIS-server/global"
	"MEIS-server/model/system/request"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"sort"

	"gorm.io/gorm"
)

// SubInitializer 提供 source/*/init() 使用的接口，每个 initializer 完成一个初始化过程
type SubInitializer interface {
	InitializerName() string
	MigrateTable(ctx context.Context) (next context.Context, err error)   // 迁移表
	InitializeData(ctx context.Context) (next context.Context, err error) // 初始化数据
	TableCreated(ctx context.Context) bool                                // 创建表
	DataInserted(ctx context.Context) bool                                // 数据插入
}

// TypedDBInitHandler 执行传入的 initializer
type TypedDBInitHandler interface {
	// 创建数据库并且初始化
	EnsureDB(ctx context.Context, conf *request.InitDB) (context.Context, error) // 建库，失败属于 fatal error，因此让它 panic
	WriteConfig(ctx context.Context) error                                       // 回写配置
	InitTables(ctx context.Context, inits initSlice) error                       // 建表 handler
	InitData(ctx context.Context, inits initSlice) error                         // 建数据 handler
}

// 错误提示
var (
	ErrDBTypeMismatch          = errors.New("db type mismatch")                   // 数据库类型不匹配
	ErrMissingDBContext        = errors.New("missing db in context")              // 找不到数据库
	ErrMissingDependentContext = errors.New("missing dependent value in context") // 上下文找不到
)

// 初始化提示
const (
	Mysql           = "mysql"
	InitSuccess     = "\n[%v] --> 初始数据成功!\n"
	InitDataExist   = "\n[%v] --> %v 的初始数据已存在!\n"
	InitDataFailed  = "\n[%v] --> %v 初始数据失败! \nerr: %+v\n"
	InitDataSuccess = "\n[%v] --> %v 初始数据成功!\n"
)

const (
	InitOrderSystem   = 10
	InitOrderInternal = 1000
	InitOrderExternal = 100000
)

// 组合一个顺序字段 以供排序
type orderedInitializer struct {
	order int
	SubInitializer
}

// initSlice 供 initialize 排序依赖的时用
type initSlice []*orderedInitializer

var (
	initializers initSlice
	// 缓存
	cache map[string]*orderedInitializer
)

// 注册要初始化的函数
func RegisterInit(order int, i SubInitializer) {
	if initializers == nil {
		initializers = initSlice{}
	}

	if cache == nil {
		cache = map[string]*orderedInitializer{}
	}

	name := i.InitializerName()
	if _, existed := cache[name]; existed {
		panic(fmt.Sprintf("Name conflict on %s", name))
	}

	ni := orderedInitializer{order, i}
	initializers = append(initializers, &ni)

	cache[name] = &ni
}

// ----------------------- 控制器 ------------------------

type InitDBController struct {
}

// Init 创建数据库并初始化
func (i *InitDBController) InitDB(conf request.InitDB) (err error) {
	ctx := context.TODO()
	if len(initializers) == 0 {
		return errors.New("无可用初始化过程，请检查初始化是否已执行完成")
	}

	// 保证有依赖的在后面执行
	sort.Sort(&initializers)

	var initHandler TypedDBInitHandler

	switch conf.DBType {
	case "mysql":
		initHandler = NewMysqlInitHandler()
		ctx = context.WithValue(ctx, "dbtype", "mysql")
	default:
		initHandler = NewMysqlInitHandler()
		ctx = context.WithValue(ctx, "dbtype", "mysql")
	}

	ctx, err = initHandler.EnsureDB(ctx, &conf)
	if err != nil {
		return err
	}

	// 从ctx中获取gorm.DB 开始初始化表和数据
	db := ctx.Value("db").(*gorm.DB)
	global.MEIS_DB = db

	if err = initHandler.InitTables(ctx, initializers); err != nil {
		return errors.New("InitTables Error：" + err.Error())
	}

	fmt.Println("initializers", initializers)
	if err = initHandler.InitData(ctx, initializers); err != nil {
		return errors.New("InitData Error：" + err.Error())
	}

	if err = initHandler.WriteConfig(ctx); err != nil {
		return errors.New("WriteConfig Error：" + err.Error())
	}

	// 完成初始化后清空
	initializers = initSlice{}
	cache = map[string]*orderedInitializer{}
	return nil
}

// 创建数据库
func createDatabase(dsn string, driver string, createSql string) error {
	// 原生sql
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}

	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)

	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec(createSql)
	return err
}

// 创建表
func createTables(ctx context.Context, inits initSlice) error {
	// 创建一个可以取消的上下文
	next, cancel := context.WithCancel(ctx)
	defer func(c func()) { c() }(cancel)
	for _, init := range inits {
		if init.TableCreated(next) {
			continue
		}
		if n, err := init.MigrateTable(next); err != nil {
			return err
		} else {
			next = n
		}
	}
	return nil
}

// sort 所要实现的类
func (i initSlice) Len() int {
	return len(i)
}

func (a initSlice) Less(i, j int) bool {
	return a[i].order < a[j].order
}

func (a initSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
