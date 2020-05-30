package main

import (
	"RabbitmqDemo/RabbitMQ"
)

func main()  {
	rabbitmq := RabbitMQ.NewRabbitMQTopic("exImoocTopic","#")
	rabbitmq.ReceiveTopic()
}
