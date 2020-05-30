package main

import (
	"RabbitmqDemo/RabbitMQ"
	"strconv"
	"time"
)

func main()  {
	rabbitMQOne := RabbitMQ.NewRabbitMQTopic("exImoocTopic","imooc.topic.one")
	rabbitMQTwo := RabbitMQ.NewRabbitMQTopic("exImoocTopic","imooc.topic.two")
	for i:=0;i<=100 ;i++  {
		rabbitMQOne.PublishgTopic("hello imooc.topic.one"+strconv.Itoa(i))
		rabbitMQTwo.PublishgTopic("hello imooc.topic.two"+strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
