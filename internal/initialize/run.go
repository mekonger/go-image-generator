package initialize

import (
	"fmt"
	"github.com/mekonger/go-image-generator/config"
	"github.com/mekonger/go-image-generator/internal/models"
)

func runTasks() {
	taskN := 100
	rets := make(chan string, taskN)

	q := NSQQueue{rets: rets}
	q.InitNSQConsumer()
	q.InitNSQProducer()

	for i := 0; i < taskN; i++ {
		go func(i int) {
			q.SendMessage(&models.JobData{
				Name:    "Sleeping Gophers",
				Message: fmt.Sprintf("Hello commander, I am handling the job: %02d", +i),
			})
		}(i)
	}

	/*q := queue.NewPool(30, queue.WithFn(func(ctx context.Context, m core.QueuedMessage) error {
		v, _ := m.(*models.JobData)
		err := json.Unmarshal(m.Bytes(), &v)
		if err != nil {
			return err
		}

		rets <- "Hello, " + v.Name + ", " + v.Message
		return nil
	}))
	defer q.Release()

	for i := 0; i < taskN; i++ {
		go func(i int) {
			err := q.Queue(&models.JobData{
				Name:    "Sleeping Gophers",
				Message: fmt.Sprintf("Hello commander, I am handling the job: %02d", +i),
			})
			if err != nil {
				return
			}
		}(i)
	}

	for i := 0; i < taskN; i++ {
		fmt.Println("message: ", <-rets)
		time.Sleep(20 * time.Millisecond)
	}*/
}

func Run() {
	config.LoadConfig()

	r := InitRouter()

	go runTasks()

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
