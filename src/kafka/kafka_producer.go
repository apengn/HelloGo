package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/Shopify/sarama"
)

// producer kafka
func Producer() {
	addr := flag.String("kafka_addr", "", "kafka_addr")
	topic := flag.String("kafka_topic", "", "kafka_topic")
	flag.Parse()

	if len(*addr) == 0 || len(*topic) == 0 {
		panic(`please input params  example "--kafka_addr 10.0.0.159:9092 --kafka_topic default"`)
	}
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	msg := &sarama.ProducerMessage{}
	msg.Topic = *topic
	client, err := sarama.NewSyncProducer([]string{*addr}, config)
	if err != nil {
		fmt.Println("producer close, err:", err)
		return
	}

	for {
		fmt.Println("please input producer data:")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		msg.Value = sarama.StringEncoder(input.Bytes())
		pid, offset, err := client.SendMessage(msg)
		if err != nil {
			fmt.Println("send message failed,", err)
			continue
		}
		fmt.Printf("pid:%v offset:%v\n", pid, offset)
	}

	client.Close()

}

func main() {
	Producer()
}
