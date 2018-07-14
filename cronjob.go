package cronjob

import (
	"gopkg.in/robfig/cron.v2"
	"github.com/getsentry/raven-go"
	"log"
)

type Cronjobs struct {
	Cronjobs []Cronjob
}

type Cronjob struct {
	Interval string
	Call     func()
}

func SetupCronjobs(cronjobs Cronjobs) {
	for _, cronjob := range cronjobs.Cronjobs {
		cronSetup := cron.New()
		_, err := cronSetup.AddFunc(cronjob.Interval, cronjob.Call)

		if err != nil {
			raven.CaptureError(err, nil)
			log.Fatal(err.Error())
		}

		cronSetup.Start()
	}
}