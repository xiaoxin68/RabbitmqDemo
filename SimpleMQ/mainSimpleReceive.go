package main

import (
	"RabbitmqDemo/RabbitMQ"
	"fmt"
)

func main()  {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("imoocSimple")
	rabbitmq.ConsumeSimple()
	fmt.Println("消费成功！")
}
