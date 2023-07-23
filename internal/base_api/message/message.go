package baseAPIMessage

// 该模块对mq的消息进行消费:
// 1.作为客户端连接websocket，进行转发消息
// 2.将消息存储在mongodb中

import (
	"Spider/pkg/rabbitmq"
)

func main() {
	consumer := &rabbitmq.Consumer{}
	consumer.Consumer_Init("send", "addr")
	consumer.Run()
}
