package main

import (
	"bytes"
	"log"
	"time"

	"github.com/streadway/amqp"
)

type Consumer struct {
	topic   string
	addr    string
	channel *amqp.Channel
}

func (c *Consumer) Consumer_Init(addr string, topic string) {
	c.addr = addr
	c.topic = topic
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Println("Failed to connect to RabbitMQ")
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Println("Failed to open a channel")
	}

	q, err := ch.QueueDeclare(
		topic, // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Println("Failed to declare a queue")
	}

	err = ch.Qos(
		1,     // prefetch count 服务器将在收到确认之前将那么多消息传递给消费者。
		0,     // prefetch size  服务器将尝试在收到消费者的确认之前至少将那么多字节的交付保持刷新到网络
		false, // 当 global 为 true 时，这些 Qos 设置适用于同一连接上所有通道上的所有现有和未来消费者。当为 false 时，Channel.Qos 设置将应用于此频道上的所有现有和未来消费者
	)
	if err != nil {
		log.Println("Failed to set QoS")
	}
	c.channel = ch
}

// 这里需要修改一下，Consumer业务代码暂时写在这里
func (c *Consumer) Run() {

	msgs, err := c.channel.Consume(
		c.topic, // queue
		"",      // consumer
		false,   // auto-ack
		false,   // exclusive
		false,   // no-local
		false,   // no-wait
		nil,     // args
	)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			// func Count(s, sep [] byte ) int  计算s中sep的非重叠实例数。如果sep是空切片，则Count返回1+s中UTF-8编码的代码点数
			dotCount := bytes.Count(d.Body, []byte(".")) // 返回 . 的个数
			t := time.Duration(dotCount)                 // 表示为 int64 纳秒计数
			time.Sleep(t * time.Second)                  //将当前 goroutine 暂停至少持续时间d
			log.Printf("Done")
			d.Ack(false) // 必须在成功处理此交付后调用 Delivery.Ack,如果autoAck为true则不需要. 参数 true表示回复当前信道所有未回复的ack，用于批量确认。false表示回复当前条目
		}
	}()
}
