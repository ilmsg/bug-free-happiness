package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/spf13/viper"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("kafka.servers:", viper.GetStringSlice("kafka.servers"))
	fmt.Println("kafka.group:", viper.GetString("kafka.group"))

	servers := viper.GetStringSlice("kafka.servers")
	group := viper.GetString("kafka.group")
	consumer, err := sarama.NewConsumerGroup(servers, group, nil)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	topics := []string{
		viper.GetString("kafka.topic1"),
		viper.GetString("kafka.topic2"),
	}
	fmt.Printf("topics: %+v\n", topics)

	ctx := context.Background()
	eventHandler := NewEventHandler()
	handler := NewTopicConsumerGroupHandler(eventHandler)

	fmt.Println("Consumer start.")
	for {
		consumer.Consume(ctx, topics, handler)
	}
}

type EventHandler interface {
	Handle(topic string, eventBytes []byte)
}

type topicHandler struct {
}

func (*topicHandler) Handle(topic string, eventBytes []byte) {
	switch topic {
	case viper.GetString("kafka.topic1"):
		var event protoreflect.ProtoMessage
		err := proto.Unmarshal(eventBytes, event)
		if err != nil {
			log.Println(err)
			return
		}
		// log.Printf("event handler: [%v] %+v", topic, event)
		log.Printf("-----------------------------------------------\n")
		log.Printf("Topic: %v\n", topic)
		log.Printf("Value: %+v\n", event)
	case viper.GetString("kafka.topic2"):
		var event protoreflect.ProtoMessage
		err := proto.Unmarshal(eventBytes, event)
		if err != nil {
			log.Println(err)
			return
		}
		// log.Printf("event handler: [%v] %+v", topic, event)
		log.Printf("-----------------------------------------------\n")
		log.Printf("Topic: %v\n", topic)
		log.Printf("Value: %+v\n", event)
	default:
		log.Println("no event handler:", topic)
	}
}

func NewEventHandler() EventHandler {
	return &topicHandler{}
}

type TopicConsumerGroupHandler struct {
	eventHandler EventHandler
}

func (*TopicConsumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (conn *TopicConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		conn.eventHandler.Handle(msg.Topic, msg.Value)
		session.MarkMessage(msg, "")
	}
	return nil
}

func (*TopicConsumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func NewTopicConsumerGroupHandler(eventHandler EventHandler) sarama.ConsumerGroupHandler {
	return &TopicConsumerGroupHandler{eventHandler}
}
