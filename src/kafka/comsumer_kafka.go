package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
)

// go run flag_.go -b back -d true
// consumer kafka data
func main() {

	addr := flag.String("kafka_addr", "", "kafka_addr")
	topic := flag.String("kafka_topic", "", "kafka_topic")
	flag.Parse()

	if len(*addr) == 0 || len(*topic) == 0 {
		panic(`please input params  example " --kafka_addr 10.0.0.159:9092 --kafka_topic default"`)
	}

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	consumer, err := sarama.NewConsumer([]string{*addr}, config)
	if err != nil {

		panic(err)
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	partitionConsumer, err := consumer.ConsumePartition(*topic, 0, -1)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	consumed := 0
ConsumerLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			log.Printf("Consumed message offset %d\n", msg.Offset)
			log.Printf("Consumed message : %s\n", string(msg.Value))
			consumed++
		case <-signals:
			break ConsumerLoop
		}
	}

	log.Printf("Consumed: %d\n", consumed)
}
