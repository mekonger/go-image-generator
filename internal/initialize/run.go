package initialize

import (
	"fmt"
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

func job(i int, rets chan string) {
	sleepSomeTime()
	rets <- fmt.Sprintf("Handle the job: %d", i+1)
}

func runTasks() {
	taskN := 100
	rets := make(chan string, taskN)

	for i := 0; i < taskN; i++ {
		go job(i, rets)
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
