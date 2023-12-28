package system

import (
	"MEIS-server/controller/system"
	sysModel "MEIS-server/model/system"
	"MEIS-server/utils"
	"context"
	"fmt"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// 14
const initOrderUser = initOrderRole + 1

type initUser struct {
}

// 自动运行
func init() {
	system.RegisterInit(initOrderUser, &initUser{})
}

// 自动创建表格
func (i *initUser) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&sysModel.SysUser{})
}

// 检查表格是否已经创建
func (i *initUser) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&sysModel.SysUser{})
}

// 获取表格名字 并且作为键值存在缓存中
func (i initUser) InitializerName() string {
	return sysModel.SysUser{}.TableName()
}

// 初始化数据
func (i *initUser) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	password := utils.BcryptHash("123456")
	adminPassword := utils.BcryptHash("181022")

	entities := []sysModel.SysUser{
		{
			UUID:     uuid.NewV4(),
			Username: "admin",
			Password: adminPassword,
			NickName: "Mr.Mango",
			Email:    "333333333@qq.com",
			Roles: []sysModel.SysRole{
				{
					RoleId: 777,
				},
			},
		},
		{
			UUID:     uuid.NewV4(),
			Username: "zzk",
			Password: password,
			NickName: "康大少",
			Email:    "333333333@qq.com",
			Roles: []sysModel.SysRole{
				{
					RoleId: 666,
				},
			},
		},
	}
	if err = db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, sysModel.SysUser{}.TableName()+"表数据初始化失败!")
	}
	// 存入
	next = context.WithValue(ctx, i.InitializerName(), entities)
	// 获取
	rolesEntities, ok := ctx.Value(initRole{}.InitializerName()).([]sysModel.SysRole)
	fmt.Println("rolesEntities", rolesEntities)

	if !ok {
		return next, errors.Wrap(system.ErrMissingDependentContext, "创建 [用户-权限] 关联失败, 未找到权限表初始化数据")
	}
	return next, err
}

// 查询数据是否插入
func (i *initUser) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	var record sysModel.SysUser
	if errors.Is(db.Where("username = ?", "admin").
		Preload("Roles").First(&record).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return len(record.Roles) > 0
}
