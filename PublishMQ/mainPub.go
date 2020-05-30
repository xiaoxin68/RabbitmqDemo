package main

import (
	"RabbitmqDemo/RabbitMQ"
	"strconv"
	"time"
)

func main()  {
	rabbitMQ := RabbitMQ.NewRabbitMQPubSub("newProduct")
	for i:=0;i<=100 ;i++  {
		rabbitMQ.PublishgPub("订阅模式生产第"+strconv.Itoa(i)+"条数据")
		time.Sleep(time.Second)
	}
}
