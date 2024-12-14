package initialize

import (
	"context"
	"fmt"
	"github.com/golang-queue/queue"
	"github.com/mekonger/go-image-generator/config"
	"math/rand"
	"time"
)

func sleepSomeTime() string {
	sleepTime := time.Duration(rand.Intn(60)) * time.Second
	message := fmt.Sprintf("Sleeping for %v", sleepTime)
	fmt.Printf("About to process: %s\n", message)
	time.Sleep(sleepTime)
	return message
}

func job(i int, rets chan string) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		sleepSomeTime()
		rets <- fmt.Sprintf("Hello commander, I'm handling the job: %02d", +i)
		return nil
	}
}

func runTasks() {
	taskN := 100
	rets := make(chan string, taskN)

	q := queue.NewPool(5)
	defer q.Release()

	for i := 0; i < taskN; i++ {
		go func() {
			err := q.QueueTask(job(i, rets))
			if err != nil {
				panic(err)
			}
		}()
	}

	for i := 0; i < taskN; i++ {
		fmt.Println("message: ", <-rets)
		time.Sleep(20 * time.Second)
	}
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
