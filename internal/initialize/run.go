package initialize

import (
	"fmt"
	"github.com/jdxyw/generativeart"
	"github.com/jdxyw/generativeart/arts"
	"github.com/jdxyw/generativeart/common"
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
}

func genrateArt() {
	c := generativeart.NewCanva(600, 400)
	c.SetBackground(common.NavajoWhite)
	c.FillBackground()
	c.SetLineWidth(1.0)
	c.SetLineColor(common.Orange)
	c.Draw(arts.NewColorCircle(30))
	c.ToPNG("circle.png")
}

func Run() {
	config.LoadConfig()

	r := InitRouter()

	//go runTasks()
	go genrateArt()

	err := r.Run(":8060")
	if err != nil {
		panic(err)
	}
}
