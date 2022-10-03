package stream

import (
	"FirstWeek/Config"
	"FirstWeek/Transaction"
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"time"
)

func KafkaProducer(transaction *Transaction.Transaction, configurations Config.Configurations) {
	conn, _ := kafka.DialLeader(context.Background(), "tcp", configurations.Kafka.URL, configurations.Kafka.Topic, 0)
	conn.SetWriteDeadline(time.Now().Add(time.Second * 10))
	obj, _ := json.Marshal(&transaction)
	conn.WriteMessages(kafka.Message{Value: []byte(obj)})
}

func KafkaConsumer(configurations Config.Configurations) {
	config := kafka.ReaderConfig{
		Brokers:  []string{configurations.Kafka.URL},
		Topic:    configurations.Kafka.Topic,
		MaxBytes: 1e6,
	}
	reader := kafka.NewReader(config)
	for {
		message, error := reader.ReadMessage(context.Background())
		if error != nil {
			fmt.Println(time.Now().String()+":: Error happened during calling kafka server %v", error)
			continue
		}
		fmt.Println(time.Now().String() + "::message of transaction consumed:: " + string(message.Value))
	}
}
