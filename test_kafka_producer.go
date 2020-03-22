package main

import (
	"fmt"
	"github.com/Shopify/sarama" //引用github的sarama，实际存储在golib里
	"os"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true // must be true to be used in a SyncProducer

	msg := &sarama.ProducerMessage{}
	msg.Topic = "my_test_michaelxu"
	msg.Value = sarama.StringEncoder("adamxiao|" + os.Args[1]) // msg序列化

	//client, err := sarama.NewSyncProducer([]string{"100.116.126.51:9092"}, config) //创建一个生产者，并把kafka后端传进去
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config) //创建一个生产者，并把kafka后端传进去
	if err != nil {
		fmt.Println("producer close, err:", err)
		return
	}
	defer client.Close()

	pid, offset, err := client.SendMessage(msg) //发送消息给kafka并返回offset
	if err != nil {
		fmt.Println("send message failed,", err)
		return
	}

	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}
