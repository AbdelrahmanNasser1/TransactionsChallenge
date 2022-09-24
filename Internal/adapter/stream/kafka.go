package stream

import (
	"FirstWeek/Config"
	"FirstWeek/Transaction"
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func SaveToKafka(config Config.Configurations, transaction Transaction.Transaction) bool {
	fmt.Println("Save to kafka server...")
	jsonString, err := json.Marshal(transaction)
	transactionString := string(jsonString)
	fmt.Println(transactionString)
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": config.Kafka.URL})
	if err != nil {
		panic(err)
		return false
	}
	if p == nil {
		return false
	}
	topic := config.Kafka.Topic
	for _, word := range []string{string(transactionString)} {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)

	}
	return true
}
func ReceiveFromKafka(config Config.Configurations) {

	fmt.Println("Start receiving from Kafka")
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": config.Kafka.URL,
		"group.id":          "group-id-1",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{config.Kafka.Topic}, nil)

	for {
		msg, err := c.ReadMessage(0)

		if err == nil {
			fmt.Printf("Received from Kafka %s: %s\n", msg.TopicPartition, string(msg.Value))
			event := string(msg.Value)
			fmt.Println(event)
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			break
		}
	}

	c.Close()

}
