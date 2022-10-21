package service

import (
	"DistributeSyetemDemo/registry"
	"context"
	"fmt"
	"log"
	"net/http"
)

// Start 启动服务
func Start(ctx context.Context,host,port string,
	reg registry.Registration,
	registerHandlersFunc func())(context.Context,error){
	//注册web服务
	registerHandlersFunc()
	//启动web服务
	ctx = startService(ctx, reg.ServiceName,host,port)
	//注册服务
	err := registry.RegisterService(reg)
	if err != nil{
		return ctx, nil
	}

	return ctx,nil
}

// startService 启动web服务
func startService(ctx context.Context,
	serviceName registry.ServiceName,
	host,port string)context.Context{
	ctx,cancel := context.WithCancel(ctx)

	var srv http.Server
	srv.Addr =":" +port

	go func(){
		log.Println(srv.ListenAndServe())
		err := registry.ShutdownService(fmt.Sprintf("http://%s:%s",host,port))
		if err != nil{
			log.Println(err)
		}
		cancel()
	}()

	go func(){
		fmt.Printf("%v started. Press any key to stop. \n",serviceName)
		//手动停止服务
		var s string
		fmt.Scanln(&s)
		err := registry.ShutdownService(fmt.Sprintf("http://%s:%s",host,port))
		if err != nil{
			log.Println(err)
		}
		srv.Shutdown(ctx)
		cancel()
	}()
	return ctx
}