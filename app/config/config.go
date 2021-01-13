package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config required for dashboard service
var Config struct {
	ServiceName string `yaml:"service_name"`
	HTTP        HTTP   `yaml:"http"`
	SentryDSN   string `yaml:"sentry_dsn"`
	Redis       Redis  `yaml:"redis"`
	Kafka       Kafka  `yaml:"kafka"`
}

type Kafka struct {
	Address       string        `yaml:"address"`
	OrderConsumer KafkaConsumer `yaml:"order_consumer"`
	OrderProducer KafkaProducer `yaml:"order_producer"`
}

type KafkaConsumer struct {
	Address        string `yaml:"address"`
	Topic          string `yaml:"topic"`
	GroupID        string `yaml:"group_id"`
	SessionTimeout int    `yaml:"session_timeout"`
	Offset         string `yaml:"offset"`
}

type KafkaProducer struct {
	Address     string `yaml:"address"`
	Topic       string `yaml:"topic"`
	RetentionMs string `yaml:"retention_ms"`
}

type Redis struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	Database int    `yaml:"database"`
}

type HTTP struct {
	Port int `yaml:"port"`
}

// LoadConfig loads config
func LoadConfig(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(file, &Config)
	if err != nil {
		return err
	}

	return nil
}
