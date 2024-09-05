package kafka

import (
	"fmt"
	"github.com/IBM/sarama"
	"github.com/hespecial/gin-mall/global"
	"go.uber.org/zap"
)

type Consumer struct {
	Consumer sarama.Consumer
}

func NewKafkaConsumer(brokers []string) (*Consumer, error) {
	consumer, err := sarama.NewConsumer(brokers, nil)
	if err != nil {
		return nil, err
	}

	return &Consumer{Consumer: consumer}, nil
}

func (kc *Consumer) ConsumeMessages(topic string, partition int32, offset int64) error {
	partitionConsumer, err := kc.Consumer.ConsumePartition(topic, partition, offset)
	if err != nil {
		global.Log.Error(fmt.Sprintf("Failed to start consumer for partition %d", partition), zap.Error(err))
	}
	defer func() {
		_ = partitionConsumer.Close()
	}()

	for message := range partitionConsumer.Messages() {
		global.Log.Info(fmt.Sprintf("Consumed message with key %s and value %s from topic %s\n", string(message.Key), string(message.Value), message.Topic))
	}

	return nil
}

func (kc *Consumer) Close() error {
	return kc.Consumer.Close()
}
