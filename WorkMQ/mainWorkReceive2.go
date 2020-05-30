package main

import (
	"RabbitmqDemo/RabbitMQ"
	"fmt"
)

func main()  {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("imoocWork")
	rabbitmq.ConsumeSimple()
	fmt.Println("消费成功！")
}
