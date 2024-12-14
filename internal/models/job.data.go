package models

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type JobData struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

func sleepSomeTime() string {
	sleepTime := time.Duration(rand.Intn(60)) * time.Second
	message := fmt.Sprintf("Sleeping for %v", sleepTime)
	fmt.Printf("About to process: %s\n", message)
	time.Sleep(sleepTime)
	return message
}

func (j *JobData) Bytes() []byte {
	fmt.Printf("%s:%s\n", j.Name, j.Message)
	res := sleepSomeTime()
	j = &JobData{Name: "I am awake", Message: res}
	b, _ := json.Marshal(j)
	return b
}
