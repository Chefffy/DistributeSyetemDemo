package main

import (
	"DistributeSyetemDemo/log"
	"DistributeSyetemDemo/registry"
	"DistributeSyetemDemo/service"
	"context"
	"fmt"
	stlog "log"
)

// 日志服务
func main(){
	log.Run("./distributed.log")
	host, port := "localhost","4000"
	serviceAddress := fmt.Sprintf("http://%s:%s",host,port)

	r := registry.Registration{
		ServiceName: registry.LogService,
		ServiceURL: serviceAddress,
		RequiredServices: make([]registry.ServiceName,0),
		ServiceUpdateURL: serviceAddress + "/services",
		HeartbeatURL: serviceAddress + "/heartbeat",
	}
	ctx,err := service.Start(
		context.Background(),
		host,
		port,
		r,
		log.RegisterHandlers,
		)
	if err != nil{
		stlog.Fatal(err)
	}

	<- ctx.Done()

	fmt.Println("Shutting down log service.")
}
