package kafka

import (
	"github.com/IBM/sarama"
)

type Producer struct {
	Producer sarama.SyncProducer
}

func NewKafkaProducer(brokers []string) (*Producer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &Producer{Producer: producer}, nil
}

func (kp *Producer) SendMessage(topic, key, value string) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(value),
	}
	_, _, err := kp.Producer.SendMessage(msg)
	return err
}

func (kp *Producer) Close() error {
	return kp.Producer.Close()
}
