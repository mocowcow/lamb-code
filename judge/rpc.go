package judge

import (
	"context"
	"encoding/json"
	"fmt"
	"lamb-code/config"
	"lamb-code/playground"
	"log"
	"math/rand"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func playgroudRPC(code string, inputs []string) (res []string, err error) {
	url := fmt.Sprintf("amqp://%s:%s@%s:%s", config.GetString("mq.user"), config.GetString("mq.pw"), config.GetString("mq.host"), config.GetString("mq.port"))
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // noWait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	corrId := randomString(32)
	data := playground.PlaygroundArgs{Code: code, Inputs: inputs}
	bytes, _ := json.Marshal(data)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"", // exchange
		config.GetString("service.playground.rpc.queue"), // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			CorrelationId: corrId,
			ReplyTo:       q.Name,
			Body:          bytes,
		})
	failOnError(err, "Failed to publish a message")

	for d := range msgs {
		if corrId == d.CorrelationId {
			json.Unmarshal(d.Body, &res)
			failOnError(err, "Failed to convert body to integer")
			break
		}
	}

	return
}
