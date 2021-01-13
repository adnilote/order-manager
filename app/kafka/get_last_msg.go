package kafka

import (
	"context"
	"fmt"

	"github.com/adnilote/order-manager/app/config"
	"github.com/sirupsen/logrus"
	confluent "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

const (
	ErrEmptyTopic = "empty topic"
)

type LastMessageConsumer struct {
	consumer *confluent.Consumer
	config   config.KafkaConsumer
}

// NewLastMessageConsumer returns consumer if topic is not empty
func NewLastMessageConsumer(config config.KafkaConsumer) (*LastMessageConsumer, error) {
	var err error
	cons := &LastMessageConsumer{
		config: config,
	}

	kafkaconf := confluent.ConfigMap{
		"bootstrap.servers":               config.Address,
		"group.id":                        config.GroupID,
		"go.events.channel.enable":        false,
		"enable.auto.commit":              false,
		"go.application.rebalance.enable": true,
	}

	cons.consumer, err = confluent.NewConsumer(&kafkaconf)
	if err != nil {
		return nil, fmt.Errorf("cannot create new consumer: %w", err)
	}

	ok := cons.isLastMessageAvailable()
	if !ok {
		return nil, fmt.Errorf(ErrEmptyTopic)
	}

	return cons, nil
}

func (cons *LastMessageConsumer) isLastMessageAvailable() bool {
	low, high, err := cons.consumer.QueryWatermarkOffsets(cons.config.Topic, 0, 3000)
	if err != nil {
		return false
	}
	if (high == 0) || (high == low) {
		return false
	}
	return true
}

// GetLastMessage get last message from kafka topic
// if topic is not empty
func (cons *LastMessageConsumer) GetLastMessage(ctx context.Context) (*confluent.Message, error) {

	err := cons.consumer.SubscribeTopics([]string{cons.config.Topic}, nil)
	if err != nil {
		return nil, fmt.Errorf("cannot subscribe to the topics %w", err)
	}

	defer func() {
		err := cons.consumer.Close()
		if err != nil {
			logrus.WithError(err).WithField("topic", cons.config.Topic).Error("failed stopping consumer topic")
		}
	}()
	for {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("context done")

		default:
			ev := cons.consumer.Poll(5000)
			if ev == nil {
				continue
			}

			switch message := ev.(type) {
			case confluent.AssignedPartitions:
				logrus.Debug("assign partitions")

				parts := []confluent.TopicPartition{}

				for _, partition := range message.Partitions {
					if *partition.Topic == cons.config.Topic {
						partition.Offset = confluent.OffsetTail(1)
						parts = append(parts, partition)
					}

				}

				err := cons.consumer.Assign(parts)
				if err != nil {
					continue
				}

			case *confluent.Message:
				return message, err

			case confluent.Error:
				logrus.WithField("message", message).Warning("error from kafka")
				continue
			default:
				logrus.WithField("message", message).Debug("msg from kafka")
				continue
			}
		}

	}
}
