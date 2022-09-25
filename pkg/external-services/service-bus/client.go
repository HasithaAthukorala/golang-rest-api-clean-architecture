package service_bus

import (
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"golang-rest-api-clean-architecture/pkg/config"
	"golang-rest-api-clean-architecture/pkg/entities"
	"time"
)

type ServiceBusClient interface {
	Publish(company *entities.Company)
}

type serviceBusClient struct {
	amqpChannel *amqp.Channel
	amqpQueue   amqp.Queue
}

func New(cfg *config.Config, stopCh <-chan struct{}) ServiceBusClient {
	conn, err := amqp.Dial(cfg.ServiceBusConnectionString)
	if err != nil {
		logrus.Fatalf("could not connect with the service bus: %v", err)
	}

	go func(conn *amqp.Connection) {
		<-stopCh
		err := conn.Close()
		if err != nil {
			logrus.Fatalf("could not close the service bus connection properly: %v", err)
		}
	}(conn)

	channel, err := conn.Channel()
	if err != nil {
		logrus.Fatalf("failed to open a channel in the service bus: %v", err)
	}

	go func(ch *amqp.Channel) {
		<-stopCh
		err := ch.Close()
		if err != nil {
			logrus.Fatalf("failed to close the channel in the service bus: %v", err)
		}
	}(channel)

	queue, err := channel.QueueDeclare(
		"companies",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logrus.Fatalf("failed to declre the queue in the service bus: %v", err)
	}
	return &serviceBusClient{
		amqpChannel: channel,
		amqpQueue:   queue,
	}
}

func (client *serviceBusClient) Publish(company *entities.Company) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	companyJson, err := json.Marshal(company)
	if err != nil {
		logrus.Errorf("failed to marshal the company struct: %v", err)
	}

	err = client.amqpChannel.PublishWithContext(ctx,
		"",
		client.amqpQueue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        companyJson,
		})
	if err != nil {
		logrus.Errorf("failed to publish the company details to the service bus: %v", err)
	}
}
