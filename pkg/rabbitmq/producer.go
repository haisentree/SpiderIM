package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

type Producer struct {
	topic   string
	addr    string
	channel *amqp.Channel
}

func (p *Producer) Producer_Init(addr string, topic string) {
	p.addr = addr
	p.topic = topic
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Println("Failed to connect to RabbitMQ")
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Println("Failed to open a channel")
	}

	q, err := ch.QueueDeclare(
		"work queues", // name
		true,          // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		log.Println("Failed to declare a queue")
	}

	p.channel = ch
}

func (p *Producer) SendMessage(msg []byte) error {
	err := p.channel.Publish(
		"",      // exchange  这里使用默认交换机,消息是通过交换机传给队列
		p.topic, // routing key
		false,   // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         msg,
		})
	if err != nil {
		log.Println("mq producer send message error!")
	}

	return err
}
