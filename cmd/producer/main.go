package main

import "github.com/confluentinc/confluent-kafka-go/kafka"
import "log"
import "fmt"

func main() {
	deliveryChan := make(chan kafka.Event)
	producer := NewKafkaProducer()
	//Fixando uma key a mensagem sempre cairá na mesma partição
	Publish("Transferindo 55 reais", "teste", producer, []byte("transferencia"), deliveryChan)
	go DeliveryReport(deliveryChan) //async
	producer.Flush(3000)
}

func NewKafkaProducer() *kafka.Producer{
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "go-kafka_kafka_1:9092",
		"delivery.timeout.ms": "0",
		"acks":"1",
		"enable.idempotence": "false",
	}
	p, err := kafka.NewProducer(configMap)
	if err != nil {
		log.Println(err.Error())
	}

	return p

}

func Publish(msg string, topic string, producer *kafka.Producer, key []byte, deliveryChan chan kafka.Event) error {
	message := &kafka.Message{
		Value: []byte(msg),
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key: key,
	}

	err:= producer.Produce(message, deliveryChan)
	if err != nil {
		return err
	}

	return nil
}

func DeliveryReport(deliveryChan chan kafka.Event){
	for e := range deliveryChan {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Println("Erro ao enviar")
			} else {
				fmt.Println("Mensagem enviada:", ev.TopicPartition)
				//Aqui abaixo poderiamos anotar no banco que a mensagem foi processada
			}
	
		}
	}
}
