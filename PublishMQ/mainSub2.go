package main

import (
	"RabbitmqDemo/RabbitMQ"
)

func main()  {
	rabbitmq := RabbitMQ.NewRabbitMQPubSub("newProduct")
	rabbitmq.ReceiveSub()
}
