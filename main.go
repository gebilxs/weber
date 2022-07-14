package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"weber/controllers"
	"weber/dao/mysql"
	"weber/dao/redis"
	"weber/logger"
	"weber/pkg/snowflake"
	"weber/routes"
	"weber/setting"

	"github.com/spf13/viper"

	"go.uber.org/zap"
)

func main() {
	//1.加载配置文件
	if err := setting.Init(); err != nil {
		fmt.Printf("init setting error: %v\n", err)
		return
	}
	//2.初始化日志

	if err := logger.Init(setting.Conf.LogConfig); err != nil {
		fmt.Printf("init logger error: %v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success")
	//3.初始化mysql
	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql error: %v\n", err)
		return
	}
	defer mysql.Close()
	//4.初始化Rides连接
	if err := redis.Init(setting.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis error: %v\n", err)
		return
	}

	defer redis.Close()

	//初始化雪花函数
	if err := snowflake.Init(setting.Conf.StartTime, setting.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake error: %v\n", err)
		return
	}

	//注册翻译机
	if err := controllers.InitTrans("zh"); err != nil {
		fmt.Printf("init	Validator trans failed,err:%v\n", err)
		return
	}
	//5.注册路由

	r := routes.Setup()
	//6.启动服务（优雅关机）
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("port")),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	//quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	//signal.Notify(setting.Quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	signal.Notify(setting.Quit) //任意即可
	//go End(setting.Quit)
	<-setting.Quit // 阻塞在此，当接收到上述两种信号时才会往下执行
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

//存在supervior 的时候不适用
func End(ch chan os.Signal) {
	time.Sleep(10 * time.Second)
	ch <- syscall.SIGINT
}
