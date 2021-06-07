package main

import (
	pb "laracom/demo-service/proto/demo"
	"log"
	"context"
	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(micro.Name("laracom.demo.cli"))
	service.Init()

	client := pb.NewDemoServiceClient("laracom.demo.service", service.Client())

	rep, err := client.SayHello(context.TODO()), &pb.DemoRequest{Name: "zero"})

	if err != nil {
		log.Fatalf("服务调用失败：%v", err)
	}

	log.Println(rsp.Text)
}
