package main

import (
	"RabbitmqDemo/RabbitMQ"
	"strconv"
	"time"
)

func main()  {
	rabbitMQOne := RabbitMQ.NewRabbitMQRouting("exImooc","imooc_one")
	rabbitMQTwo := RabbitMQ.NewRabbitMQRouting("exImooc","imooc_two")
	for i:=0;i<=100 ;i++  {
		rabbitMQOne.PublishgRouting("hello imooc_one"+strconv.Itoa(i))
		rabbitMQTwo.PublishgRouting("hello imooc_two"+strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
