package core

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"MEIS-server/core/internal"
	"MEIS-server/global"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// viper 初始化，主要是来读取配置文件
// 优先级：命令行 》 环境变量 》 默认值
func Viper(path ...string) *viper.Viper {
	var config string

	// 获取环境变量
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" { // 判断命令行参数是否为空
			if configEnv := os.Getenv(internal.ConfigEnv); configEnv == "" { // 判断 internal.ConfigEnv 常量存储的环境变量是否为空
				switch gin.Mode() {
				case gin.DebugMode:
					config = internal.ConfigDefaultFile
				case gin.ReleaseMode:
					config = internal.ConfigReleaseFile
				case gin.TestMode:
					config = internal.ConfigTestFile
				}
				fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, config)
			} else { // internal.ConfigEnv 常量存储的环境变量不为空 将值赋值于config
				config = configEnv
				fmt.Printf("您正在使用%s环境变量,config的路径为%s\n", internal.ConfigEnv, config)
			}

		} else { // 命令行参数不为空 将值赋值于config
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%s\n", config)
		}
	} else { // 函数传递的可变参数的第一个值赋值于config
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%s\n", config)
	}

	viper := viper.New()

	// 设置文件
	viper.SetConfigFile(config)

	// 设置文件类型
	viper.SetConfigType("yaml")

	// 读取文件
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("文件错误：%s\n", err))
	}

	// 查看文件
	viper.WatchConfig()

	// 监听文件修改
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("文件修改为：%s \n", in.Name)
		// 序列化配置文件
		if err = viper.Unmarshal(&global.MEIS_CONFIG); err != nil {
			fmt.Println("错误：", err)
		}
	})

	if err = viper.Unmarshal(&global.MEIS_CONFIG); err != nil {
		fmt.Println("错误：", err)

	}

	// root 适配性 根据root位置去找到对应迁移位置,保证root路径有效
	global.MEIS_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	return viper
}
