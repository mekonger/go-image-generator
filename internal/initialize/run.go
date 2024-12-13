package initialize

import "github.com/mekonger/go-image-generator/config"

func Run() {
	config.LoadConfig()

	r := InitRouter()

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
