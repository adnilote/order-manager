package kafka

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

// Create topics and set retention policy
func CreateAndConfigTopics(address string, topicConfig []kafka.TopicSpecification) {

	admin, err := kafka.NewAdminClient(
		&kafka.ConfigMap{
			"bootstrap.servers": address,
		},
	)
	if err != nil {
		logrus.WithError(err).Fatal("Error creating admin client")
		return
	}

loop:
	for {
		// create topics
		createTopicResult, err := admin.CreateTopics(
			context.Background(),
			topicConfig,
		)
		if err != nil {
			logrus.WithError(err).Error("Error creating topics")
			return
		}

		// check topics configuration succeed
		done := true
		alterConfig := []kafka.ConfigResource{}

		for i := 0; i < len(topicConfig); i++ {
			if createTopicResult[i].Error.Code() == kafka.ErrTopicAlreadyExists {
				resources := kafka.ConfigResource{
					Type: kafka.ResourceTopic,
					Name: topicConfig[i].Topic,
					Config: kafka.StringMapToConfigEntries(
						topicConfig[i].Config,
						kafka.AlterOperationSet),
				}
				alterConfig = append(alterConfig, resources)
				done = false
			} else if createTopicResult[i].Error.Code() != kafka.ErrNoError {
				logrus.WithField("topic", createTopicResult[i].Topic).WithError(
					createTopicResult[i].Error,
				).Error("Error creating topic")
				done = false
				break
			}
		}
		if done {
			break loop
		}

		// if topic exists, change its config
		done = true
		if len(alterConfig) > 0 {
			alterConfigResult, err := admin.AlterConfigs(context.Background(), alterConfig)
			if err != nil {
				logrus.WithError(err).Error("Error altering topics config")
				return
			}
			for i := 0; i < len(alterConfig); i++ {
				if alterConfigResult[i].Error.Code() != kafka.ErrNoError {
					logrus.WithField("topic", alterConfigResult[i].Name).
						WithError(alterConfigResult[i].Error).
						Error("Error altering topics config")
					done = false
				}
			}
		}
		if done {
			break loop
		}

		time.Sleep(time.Minute)
	}

	logrus.Info("Topics have been created")

}
