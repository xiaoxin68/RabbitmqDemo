package main

import (
	"RabbitmqDemo/RabbitMQ"
)

func main()  {
	rabbitmq := RabbitMQ.NewRabbitMQTopic("exImoocTopic","imooc.*.two")
	rabbitmq.ReceiveTopic()
}
