package kafka

import (
	"github.com/IBM/sarama"
	"log"
)

var brokers = []string{"127.0.0.1:9092"}

func NewProducer() *sarama.SyncProducer {
	conf := sarama.NewConfig()
	conf.Producer.Partitioner = sarama.NewRandomPartitioner
	conf.Producer.RequiredAcks = sarama.WaitForLocal
	conf.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, conf)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return &producer
}

func NewConsumer() *sarama.Consumer {
	consumer, err := sarama.NewConsumer(brokers, nil)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return &consumer
}
