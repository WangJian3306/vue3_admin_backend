package main

import (
	"context"
	"flag"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"vue3_admin/dao/mysql"
	"vue3_admin/logger"
	"vue3_admin/pkg/snowflake"
	"vue3_admin/pkg/translation"
	"vue3_admin/router"
	"vue3_admin/settings"
)

// @title vue3_admin 项目接口文档
// @version 1.0
// @description 硅谷甄选项目后端

// @contact.name WangJian
// @contact.url https://github.com/wangjian3306
// @contact.email 929000201@qq.com

// @host 127.0.0.1:10086
// @BasePath /
func main() {

	// 1. 加载配置
	var configPath string
	flag.StringVar(&configPath, "f", "config.yaml", "配置文件路径")
	flag.Parse()
	if err := settings.Init(configPath); err != nil {
		fmt.Printf("init settings failed, err:%v\n", err)
		return
	}

	// 2. 初始化日志
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success...")

	// 3. 初始化MySQL连接
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close()

	// 4. 初始化雪花ID
	if err := snowflake.Init(settings.Conf.SnowFlake.StartTime, settings.Conf.SnowFlake.MachineID); err != nil {
		fmt.Printf("init failed, err:%v \n", err)
		return
	}

	// 5. 初始化gin框架内置校验器的翻译器
	if err := translation.InitTrans("zh"); err != nil {
		fmt.Printf("init validator trans failed, err:%v \n", err)
		return
	}

	// 6. 注册路由
	r := router.Setup(settings.Conf.Mode)

	// 7. 启动服务
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", settings.Conf.Port),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("Server start failed", zap.Error(err))
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}

	zap.L().Info("Server exiting")

}
