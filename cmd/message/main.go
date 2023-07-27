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
	"log"

	"github.com/gorilla/websocket"
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

	// 连接websocket
	wsConn, _, err := websocket.DefaultDialer.Dial("ws://192.168.45.128:8848/ws?clientID=1&clientUUID=9da5c08b-6e3f-401b-ae5e-394cb5504cbf&platformID=0", nil)
	if err != nil {
		log.Println("ws conn error")
	}
	defer conn.Close()

	// // 发送消息
	// err = wsConn.WriteMessage(websocket.TextMessage, []byte("Hello, world!"))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // 读取消息
	// messageType, p, err := conn.ReadMessage()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// 编写业务代码
	go func() {
		for d := range msgs {
			log.Printf("MQ Received a message: %s", d.Body)
			// 1.将消息反序列化(这里先直接转发，后期需要添加seq)
			// var messageToMQ pkgPublic.MessageToMQ
			// json.Unmarshal(d.Body, &messageToMQ)
			// 解析消息后，查看RecvID是否在线

			// 2.发送给wss
			err := wsConn.WriteMessage(websocket.TextMessage, d.Body)
			if err != nil {
				log.Println("ws send error")
			}
			// 3. 存储在mongodb中(先不写)
			log.Printf("Done")
			d.Ack(false) // 必须在成功处理此交付后调用 Delivery.Ack,如果autoAck为true则不需要. 参数 true表示回复当前信道所有未回复的ack，用于批量确认。false表示回复当前条目
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}

// 2.发送给WS
// 3.存储在mongoDB中
