package stream

import (
	"FirstWeek/Config"
	"FirstWeek/Internal/Models"
	"FirstWeek/Internal/Services"
	"FirstWeek/Transaction"
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"time"
)

var kafkaConfig = kafka.ReaderConfig{}

func InitializeKafka(configurations *Config.Configurations) {
	kafkaConfig.Brokers = []string{configurations.Kafka.URL}
	kafkaConfig.Topic = configurations.Kafka.Topic
	kafkaConfig.MaxBytes = configurations.Kafka.MaxByte
}

func KafkaProducer(transaction *Transaction.Transaction, configurations Config.Configurations) {
	conn, _ := kafka.DialLeader(context.Background(), "tcp", configurations.Kafka.URL, configurations.Kafka.Topic, 0)
	conn.SetWriteDeadline(time.Now().Add(time.Second * 10))
	obj, _ := json.Marshal(&transaction)
	conn.WriteMessages(kafka.Message{Value: []byte(obj)})
	fmt.Println(time.Now().String() + "::message of transaction produced:: " + string(obj))
}

func KafkaConsumer(ts Services.IService) {
	reader := kafka.NewReader(kafkaConfig)
	var transaction Models.TransactionModel
	for {
		message, error := reader.ReadMessage(context.Background())
		if error != nil {
			fmt.Println(time.Now().String()+":: Error happened during calling kafka server %v", error)
			continue
		}
		fmt.Println(time.Now().String() + "::message of transaction consumed:: " + string(message.Value))
		json.Unmarshal(message.Value, &transaction)
		transaction.Status = true
		_, err := ts.Update(context.Background(), transaction)
		if err != nil {
			fmt.Println("failed to update status of transaction")
		}
	}
}
