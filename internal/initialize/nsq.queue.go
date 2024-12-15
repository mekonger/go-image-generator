package initialize

import (
	"encoding/json"
	"fmt"
	"github.com/mekonger/go-image-generator/internal/models"
	"github.com/nsqio/go-nsq"
	"log"
)

type NSQQueue struct {
	rets     chan string
	consumer *nsq.Consumer
	producer *nsq.Producer
}

func (q *NSQQueue) InitNSQConsumer() {
	//wg := &sync.WaitGroup{}
	//wg.Add(1)

	cfg := nsq.NewConfig()
	c, err := nsq.NewConsumer("My_NSQ_Topic", "My_NSQ_Channel", cfg)
	if err != nil {
		panic(err)
	}
	q.consumer = c

	c.AddHandler(nsq.HandlerFunc(func(m *nsq.Message) error {
		var v *models.JobData
		if err := json.Unmarshal(m.Body, &v); err != nil {
			return err
		}

		q.rets <- v.Message
		log.Println(fmt.Sprintf("Received a message: %+v", *v))
		return nil
		//log.Println("NSQ message received:")
		//log.Println(string(m.Body))
		//return nil
	}))

	err = c.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		log.Panic("Could not connect")
	}
	log.Println("Awaiting messages from NSQ topic \"My NSQ Topic\"...")
	//wg.Wait()
}

func (q *NSQQueue) InitNSQProducer() {
	cfg := nsq.NewConfig()
	p, err := nsq.NewProducer("127.0.0.1:4150", cfg)
	if err != nil {
		log.Panic(err)
	}
	q.producer = p
	err = p.Publish("My_NSQ_Topic", []byte("sample NSQ message"))
	if err != nil {
		log.Panic(err)
	}
}

func (q *NSQQueue) SendMessage(data *models.JobData) {
	msg, err := json.Marshal(data)
	if err != nil {
		log.Panic(err)
	}

	err = q.producer.Publish("My_NSQ_Topic", msg)
	if err != nil {
		log.Panic(err)
	}
}
