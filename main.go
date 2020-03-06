package main

import (
	"gin_scaffold/public"
	"gin_scaffold/router"
	"github.com/379763376/gcommon/lib"
	"os"
	"os/signal"
	"syscall"
)

func main()  {
	lib.InitModule("./conf/dev/",[]string{"base","mysql","redis",})
	defer lib.Destroy()
	public.InitMysql()
	public.InitValidate()
	router.HttpServerRun()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	router.HttpServerStop()
}