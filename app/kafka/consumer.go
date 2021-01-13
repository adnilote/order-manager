package kafka

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	"github.com/adnilote/order-manager/app/config"
	"github.com/sirupsen/logrus"
	confluent "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

const (
	StatusConnected    = "connected"
	StatusDisconnected = "disconnected"
	StatusPending      = "pending"
)

type MessageHandler interface {
	Handle(*confluent.Message) error
}

type Response struct {
	Data json.RawMessage `json:"data"`
	UUID string          `json:"uuid"`
}

//Потребитель Kafka
type Consumer struct {
	consumer  *confluent.Consumer
	handler   MessageHandler
	ctx       context.Context
	cancelAll context.CancelFunc
	wg        *sync.WaitGroup
	status    string
	config    config.KafkaConsumer
}

//Создание экземпляра Потребителя
func NewConsumer(conf config.KafkaConsumer, address string, executor MessageHandler) (*Consumer, error) {
	kafkaconf := confluent.ConfigMap{
		"bootstrap.servers":               address,
		"group.id":                        conf.GroupID,
		"go.events.channel.enable":        false,
		"enable.auto.commit":              true,
		"go.application.rebalance.enable": true,
	}

	consumer, err := confluent.NewConsumer(&kafkaconf)
	if err != nil {
		return nil, fmt.Errorf("cannot create new consumer: %w", err)
	}
	myConsumer := Consumer{
		consumer: consumer,
		wg:       &sync.WaitGroup{},
		handler:  executor,
		status:   StatusDisconnected,
		config:   conf,
	}
	return &myConsumer, nil
}

func (consumer *Consumer) GetOffset() (int64, int64, error) {
	low, high, err := consumer.consumer.QueryWatermarkOffsets(consumer.config.Topic, 0, 3000)
	return low, high, err
}

//Запуск confluent единожды
func (consumer *Consumer) Start(ctx context.Context) error {
	consumer.ctx, consumer.cancelAll = context.WithCancel(ctx)

	err := consumer.consumer.SubscribeTopics([]string{consumer.config.Topic}, nil)
	if err != nil {
		return fmt.Errorf("cannot subscribe to the topics %w", err)
	}

	consumer.status = StatusPending
	consumer.wg.Add(1)
	go consumer.consumeClaim()

	return nil
}

func (consumer *Consumer) GetStatus() string {
	return consumer.status
}

func (consumer *Consumer) Close() {
	consumer.cancelAll()
	consumer.wg.Wait()
}

func (consumer *Consumer) consumeClaim() {
	defer func() {
		err := consumer.consumer.Close()
		if err != nil {
			logrus.WithError(err).WithField("topic", consumer.config.Topic).Error("failed stopping consumer topic")
		}
		consumer.status = StatusDisconnected
		consumer.wg.Done()
		logrus.WithField("topics", consumer.config.Topic).Info("Stop consuming")
	}()
	log := logrus.WithField("topic", consumer.config.Topic)
	log.Info("Start comsuming topic")

	consumer.status = StatusConnected
	for {
		select {
		case <-consumer.ctx.Done():
			return
		default:
			msg := consumer.consumer.Poll(5000)
			if msg == nil {
				continue
			}

			switch message := msg.(type) {
			case *confluent.Message:
				log = log.WithField("message", string(message.Value))
				err := consumer.handler.Handle(message)
				if err == nil {
					_, err := consumer.consumer.CommitMessage(message)
					if err != nil {
						log.WithError(err).Error("Error commiting msg")
					}
				}
			case confluent.TopicPartition:
				log.Debugf("kafka paritions: %+v", message)
			case confluent.AssignedPartitions:
				partitions := consumer.SetOffset(message)
				err := consumer.consumer.Assign(partitions)
				if err != nil {
					logrus.Error("error assign: ", err)
				}

			case confluent.Error:
				log.WithError(errors.New(message.String())).Error("Got error from confluent")

			case confluent.RevokedPartitions:
				err := consumer.consumer.Unassign()
				if err != nil {
					logrus.Error("error unassign: ", err)
				}
			default:
			}
		}
	}
}

func (consumer *Consumer) SetOffset(partitions confluent.AssignedPartitions) []confluent.TopicPartition {
	var result []confluent.TopicPartition
	var part confluent.TopicPartition

	for _, partition := range partitions.Partitions {
		part = partition

		switch consumer.config.Offset {
		case "beginning":
			part.Offset = confluent.OffsetBeginning
		case "latest":
			part.Offset = confluent.OffsetStored
		case "newest":
			part.Offset = confluent.OffsetEnd
		case "tail":
			part.Offset = confluent.OffsetTail(1)
		}

		result = append(result, part)
	}

	return result
}
