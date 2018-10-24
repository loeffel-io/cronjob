package main

import (
	"github.com/loeffel-io/cronjob"
	"log"
)

func main() {
	testCronjob := cronjob.Cronjob{
		Interval: "0 */1 * * * *", // @every 1m
		Call: func() {
			log.Println("cronjob is running ..")
		},
	}

	if _, err := testCronjob.Start(); err != nil {
		log.Fatal(err)
	}
}
