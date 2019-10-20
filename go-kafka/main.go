package main

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var demoTopic = "demo-topic"

/*
Run: go run main.go

Output:
2019/10/20 16:31:41 start
2019/10/20 16:31:41 sending msg ..
2019/10/20 16:31:43 Msg delivered, topic:demo-topic, offset:1, key:msg1, value:Hello Go
2019/10/20 16:31:43 Try to recv msg ..
2019/10/20 16:31:46 Msg received, topic:demo-topic, offset:0, key:msg1, value:Hello Go
2019/10/20 16:31:46 end
*/
func main() {
	log.Println("start")

	DemoProducer()
	DemoConsumer()

	log.Println("end")

}

// DemoProducer send msg to kafka
func DemoProducer() {
	config := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	}

	producer, err := kafka.NewProducer(config)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	// send msg
	log.Println("sending msg ..")
	key := []byte("msg1")
	value := []byte("Hello Go")
	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &demoTopic, Partition: kafka.PartitionAny},
		Value:          value,
		Key:            key,
	}
	err = producer.Produce(msg, nil)
	if err != nil {
		panic(err)
	}
	producer.Flush(2000)

	// confirm msg delivery
	e := <-producer.Events()
	if msg, ok := e.(*kafka.Message); ok {
		if msg.TopicPartition.Error != nil {
			log.Panicf("Fail to deliver msg, err:%v", err)
		}
		log.Printf("Msg delivered, topic:%s, offset:%s, key:%s, value:%s", *msg.TopicPartition.Topic, msg.TopicPartition.Offset, string(msg.Key), string(msg.Value))
	}
}

// DemoConsumer recv msg from kafka
func DemoConsumer() {
	config := &kafka.ConfigMap{
		"bootstrap.servers":  "localhost:9092",
		"group.id":           "demo-group",
		"enable.auto.commit": "false",
		"auto.offset.reset":  "earliest",
	}

	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		panic(err)
	}

	err = consumer.Subscribe(demoTopic, nil)

	// recv msg
	log.Printf("Try to recv msg ..")
	msg, err := consumer.ReadMessage(-1)
	if err != nil {
		panic(err)
	}
	log.Printf("Msg received, topic:%s, offset:%s, key:%s, value:%s", *msg.TopicPartition.Topic, msg.TopicPartition.Offset, string(msg.Key), string(msg.Value))
}
