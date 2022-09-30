package stream

import (
	"FirstWeek/Config"
	"FirstWeek/Transaction"
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
	"os/signal"
	"syscall"
)

func SaveToKafka(config Config.Configurations, transaction Transaction.Transaction) bool {
	fmt.Println("Save to kafka server...")
	jsonString, err := json.Marshal(transaction)
	transactionString := string(jsonString)
	fmt.Println(transactionString)
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": config.Kafka.URL,
		"security.protocol": "SASL_SSL",
		"sasl.mechanisms":   "PLAIN",
		"sasl.username":     config.Kafka.UserName,
		"sasl.password":     config.Kafka.Password})
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
		"security.protocol": "SASL_SSL",
		"sasl.mechanisms":   "PLAIN",
		"sasl.username":     config.Kafka.UserName,
		"sasl.password":     config.Kafka.Password})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{config.Kafka.Topic}, nil)
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	run := true
	for run == true {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev, err := c.ReadMessage(10)
			if err != nil {
				// Errors are informational and automatically handled by the consumer
				continue
			}
			fmt.Printf("Consumed event from topic %s: key = %-10s value = %s\n",
				*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
		}
	}
	c.Close()
}
