package cronjob

import (
	"gopkg.in/robfig/cron.v2"
)

type Cronjobs struct {
	Cronjobs []Cronjob
}

type Cronjob struct {
	Interval string
	Call     func()
}

// Setup starts cronjobs
func Setup(cronjobs Cronjobs) error {
	for _, cronjob := range cronjobs.Cronjobs {
		cronSetup := cron.New()
		_, err := cronSetup.AddFunc(cronjob.Interval, cronjob.Call)

		if err != nil {
			return err
		}

		cronSetup.Start()
	}

	return nil
}
