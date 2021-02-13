package main

import (
	"fmt"

	"github.com/streadway/amqp"
	//"github.com/wade-sam/fyp-backup-server/client_scan"
	//"github.com/wade-sam/fyp-backup-server/pkg/"
	//"project/pkg/connections"
)

func main() {
	//backup.ConnectFileScan()
	//backup.ConnectIncrementalBackup()
	//client_scan.ConnectFileScan()
	//client_scan.ConnectIncrementalBackup()
	fmt.Println("Hello World!")
	//conn := connections.DBstart()
	//collection := conn.Connection.Database(conn.Config.Database).Collection("client_collection")
	//fmt.Println(collection)
	//conn := RabitCreateConnection()
	//ch := RabbitCreateDefaultChannel(conn)
	//go DefaultConsumer()
	//DefaultProducer(ch, "Helllo World")
	//DefaultConsumer()

	fmt.Println("succesfully connected to RabbitMQ instance")
}

func DefaultProducer(channel amqp.Channel, msg string) {
	for i := 0; i < 100; i++ {
		message := fmt.Sprintf("%s%d", msg, i)
		channel.Publish(
			"amq.topic",
			"host1",
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(message),
			},
		)
		fmt.Println("Printed message")
	}

}

func DefaultConsumer() {
	conn := RabitCreateConnection()
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	msgs, err := ch.Consume(
		"new-config",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	forever := make(chan bool)
	for d := range msgs {
		fmt.Printf("Received Message: %s\n", d.Body)
	}
	<-forever

}

func RabitCreateConnection() amqp.Connection {
	conn, err := amqp.Dial("amqp://guest:guest@192.168.1.210:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	//defer conn.Close()
	return *conn
}

func RabbitCreateDefaultChannel(conns amqp.Connection) amqp.Channel {
	ch, err := conns.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	//defer ch.Close()
	q, err := ch.QueueDeclare(
		"host1",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(q)
	return *ch
}
