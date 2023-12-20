package playground

import (
	"context"
	"encoding/json"
	"fmt"
	"lamb-code/config"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func RunPRCServer() {
	url := fmt.Sprintf("amqp://%s:%s@%s:%s", config.GetString("mq.user"), config.GetString("mq.pw"), config.GetString("mq.host"), config.GetString("mq.port"))
	conn, err := amqp.Dial(url)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to connect to RabbitMQ")
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("Failed to open a channel")
		return
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		config.GetString("service.playground.rpc.queue"), // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		fmt.Println("Failed to declare a queue")
		return
	}

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		fmt.Println("Failed to set QoS")
		return
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		fmt.Println("Failed to register a consumer")
		return
	}

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		for d := range msgs {
			fmt.Println("recieved a msg")
			var data playgroudRPCInput
			err = json.Unmarshal(d.Body, &data)
			if err != nil {
				fmt.Println("Failed to convert body to json")
				return
			}

			// get language strategy concrete
			langStrategy := langMap[data.Lang]
			if langStrategy == nil {
				langStrategy = langMap["invalid"]
			}

			// run playround
			outputs := langStrategy.Run(data.Code, data.Inputs)
			res, _ := json.Marshal(outputs)
			err = ch.PublishWithContext(ctx,
				"",        // exchange
				d.ReplyTo, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					CorrelationId: d.CorrelationId,
					Body:          res,
				})
			if err != nil {
				fmt.Println("Failed to publish a message")
				return
			}

			d.Ack(false)
		}
	}()

	fmt.Println(" [*] Awaiting RPC requests")
	var forever chan struct{}
	<-forever
}
