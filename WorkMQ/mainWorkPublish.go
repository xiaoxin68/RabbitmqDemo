package main

import (
	"RabbitmqDemo/RabbitMQ"
	"fmt"
	"strconv"
	"time"
)

func main()  {
	rabbitMQ := RabbitMQ.NewRabbitMQSimple("imoocWork")
	for i:=0;i<=100 ;i++  {
		rabbitMQ.PublishSimple("hello mooc!"+strconv.Itoa(i))
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}
