package rabbitmq

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

var queueName string = "DemoQue"

func ProducerQueue(c *gin.Context) {
	requestedParam := make(map[string]int)
	c.BindJSON(&requestedParam)
	// var s string
	// for key, val := range requestedParam {
	// 	s += fmt.Sprintf("%s==>%s  ", key, val)
	// }
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	for i := 1; i < requestedParam["limit"]; i++ {
		body := "PID:" + strconv.Itoa(i)
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		failOnError(err, "Failed to publish a message")
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Pushed to queue",
	})

}
