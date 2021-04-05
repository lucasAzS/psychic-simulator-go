package kafka

import (
	"encoding/json"
	"log"
	"os"
	"time"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/lucasAzS/psychic-octo-garbanzo/application/route"
	"github.com/lucasAzS/psychic-octo-garbanzo/infra/kafka"
)

//{"client":"1","routeId":"1"}
//{"client":"2","routeId":"2"}
//{"client":"3","routeId":"3"}
func Produce(msg *ckafka.Message){
	producer := kafka.NewKafkaProducer()
	route := route.NewRoute()

	json.Unmarshal(msg.Value, &route)
	route.LoadPositions()
	positions, err := route.ExportJsonPositions()
	if err != nil {
		log.Println(err.Error())
	}
	for _, p := range positions {
		kafka.Publish(p, os.Getenv("KafkaProduceTopic"), producer)
		time.Sleep(time.Millisecond * 500)
	}

}