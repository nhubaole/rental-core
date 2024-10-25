package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	kafka "github.com/segmentio/kafka-go"
)

var (
	kafkaProducer *kafka.Writer
)

const (
	kafkaURL = "14.225.255.85:9092" // Ensure this is correct
)

// For consumer
func GetKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
		CommitInterval: time.Second,
		StartOffset:    kafka.FirstOffset,
	})
}

// For producer
func GetKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

type Stock struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

func newStock(msg, msgType string) *Stock {
	s := Stock{}
	s.Message = msg
	s.Type = msgType
	return &s
}

func actionStock(c *gin.Context) {
	s := newStock(c.Query("msg"), c.Query("type"))
	body := make(map[string]interface{})
	body["action"] = "action"
	body["info"] = s

	jsonBody, _ := json.Marshal(body)
	message := kafka.Message{
		Key:   []byte("action"),
		Value: []byte(string(jsonBody)),
	}

	err := kafkaProducer.WriteMessages(context.Background(), message)

	if err != nil {
		c.JSON(500, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"err": "",
		"msg": "successfully",
	})
}

func RegisterConsumer(id int) {
	kafkaGroupID := "my-group-id" // Ensure a valid group ID is used
	reader := GetKafkaReader(kafkaURL, "test", kafkaGroupID)
	defer reader.Close()

	fmt.Printf("Consumer (%d) is running:\n", id)
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("Consumer (%d) err: %v\n", id, err.Error())
		}

		fmt.Printf("Consumer %d received message from topic: %v\n", id, m.Topic)
	}
}

func main() {
	r := gin.Default()
	kafkaProducer = GetKafkaWriter(kafkaURL, "test")
	defer kafkaProducer.Close()
	r.POST("action/stock", actionStock)

	go RegisterConsumer(1)
	go RegisterConsumer(2)

	r.Run(":8001")
}
