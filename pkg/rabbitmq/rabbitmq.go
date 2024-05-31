package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

func OpenChannel() *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	return ch
}

func Consume(ch *amqp.Channel, out chan<- amqp.Delivery) error {
	msgs, err := ch.Consume(
		"minhafila",
		"go-consumer",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	for m := range msgs {
		out <- m
	}

	return nil
}
