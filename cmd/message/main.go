package main

// // 该模块对mq的消息进行消费:
// // 1.作为客户端连接websocket，进行转发消息
// // 2.将消息存储在mongodb中

// import (
// 	baseAPIMessage "SpiderIM/internal/base_api/message"
// )

// func main() {
// 	baseAPIMessage.Start()
// }

// 这里没有设计好，一个message应该是可以连接多个ws服务端，还可以连接多个mq，并行消费队列中的数据。
// 可能是goruntine没用好，封装的consumer不正常，先把跑起来写业务代码，后面再封装
import (
	"bytes"
	"log"
	"time"

	"github.com/streadway/amqp"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"send", // name
		true,   // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	FailOnError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count 服务器将在收到确认之前将那么多消息传递给消费者。
		0,     // prefetch size  服务器将尝试在收到消费者的确认之前至少将那么多字节的交付保持刷新到网络
		false, // 当 global 为 true 时，这些 Qos 设置适用于同一连接上所有通道上的所有现有和未来消费者。当为 false 时，Channel.Qos 设置将应用于此频道上的所有现有和未来消费者
	)
	FailOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	// 编写业务代码
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

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}
