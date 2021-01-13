package kafka

import (
	"encoding/json"
	"fmt"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type Producer struct {
	value *kafka.Producer
}

func NewProducer(address string) (*Producer, error) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": address,
	})
	if err != nil {
		return nil, fmt.Errorf("cannot create new producer: %w", err)
	}
	p := Producer{producer}
	go p.readEvents()
	return &p, nil
}

func (producer *Producer) readEvents() {
	ch := producer.value.Events()
	for msg := range ch {
		switch e := msg.(type) {
		case *kafka.Message:
			if e.TopicPartition.Error != nil {
				logrus.WithField("data", string(e.Value)).
					WithError(e.TopicPartition.Error).Error("Got error from kafka producer")
			}
		case kafka.Error:
			logrus.WithError(errors.New(e.String())).Error("Got error from kafka producer")
		default:
		}
	}
}

func (producer Producer) Send(body interface{}, topic, key string, customUUID ...string) error {
	log := logrus.WithFields(logrus.Fields{
		"body":  body,
		"topic": topic,
	})

	msg, err := json.Marshal(body)
	if err != nil {
		log.WithError(err).Error("Marshal response failed")
		return errors.Wrap(err, "Marshal response failed")
	}

	err = producer.value.Produce(
		&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny,
			},
			Value: msg,
			Key:   []byte(key),
		}, nil)
	if err != nil {
		return errors.Wrap(err, "Send message to kafka failed")
	}
	return nil
}

func (producer Producer) Close() {
	producer.value.Close()
}

func (producer Producer) Send2(body proto.Message, topic, key string, customUUID ...string) error {
	log := logrus.WithFields(logrus.Fields{
		"body":  body,
		"topic": topic,
	})

	jpb := jsonpb.Marshaler{}
	msg, err := jpb.MarshalToString(body)
	if err != nil {
		log.WithError(err).Error("Marshal response failed")
		return errors.Wrap(err, "Marshal response failed")
	}

	err = producer.value.Produce(
		&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny,
			},
			Value: []byte(msg),
			Key:   []byte(key),
		}, nil)
	if err != nil {
		return errors.Wrap(err, "Send message to kafka failed")
	}
	return nil
}
