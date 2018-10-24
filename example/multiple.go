package main

import (
	"github.com/loeffel-io/cronjob"
	"log"
	"time"
)

var cronjobs = cronjob.Cronjobs{
	Cronjobs: []cronjob.Cronjob{
		{
			Interval: "0 */1 * * * *", // @every 1m
			Call: func() {
				test()
			},
		},
		{
			Interval: "0 */3 * * * *", // @every 3m
			Call: func() {
				test()
			},
		},
	},
}

func test() {
	log.Println("cronjob is running ..")
}

func init() {
	// Setup cronjobs
	cronjob.Setup(cronjobs)
}

func main() {
	for {
		log.Println("endless loop is running ..")
		time.Sleep(5 * time.Second)
	}
}
