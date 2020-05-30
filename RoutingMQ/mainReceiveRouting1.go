package main

import (
	"RabbitmqDemo/RabbitMQ"
)

func main()  {
	rabbitmq := RabbitMQ.NewRabbitMQRouting("exImooc","imooc_one")
	rabbitmq.ReceiveRouting()
}
