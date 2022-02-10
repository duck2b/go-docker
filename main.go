package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"net/http"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")
	fmt.Fprintf(w, "hello world "+name+": age ="+age)
}

func main() {
	KafkaProduct()

	http.HandleFunc("/", SayHello)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// KafkaProduct kafka 生产者
func KafkaProduct()  {
	config := sarama.NewConfig()

	// 设置
	// ack应答机制
	config.Producer.RequiredAcks = sarama.WaitForAll

	// 发送分区
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	// 回复确认
	config.Producer.Return.Successes = true

	// 构造消息
	msg1 := &sarama.ProducerMessage{
		Topic: "kafeidou",
		Value: sarama.StringEncoder("哈哈哈 "),
	}

	var msg []*sarama.ProducerMessage
	msg = append(msg, msg1)

	// 链接kafka
	client , err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)

	if err != nil {
		fmt.Println("producer closed,err:", err)
	}

	defer client.Close()

	err = client.SendMessages(msg)

	if err != nil {
		fmt.Println("send msg failed,err:", err)
	}


}

