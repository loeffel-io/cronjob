# Cronjobs

[![GoDoc](https://godoc.org/github.com/loeffel-io/cronjob?status.svg)](https://godoc.org/github.com/loeffel-io/cronjob)
[![Go Report Card](https://goreportcard.com/badge/github.com/loeffel-io/cronjob)](https://goreportcard.com/report/github.com/loeffel-io/cronjob)
[![GitHub license](https://img.shields.io/github/license/Naereen/StrapDown.js.svg)](https://github.com/Naereen/StrapDown.js/blob/master/LICENSE)

Simple package to setup some cronjobs based on [gopkg.in/robfig/cron.v2](https://godoc.org/gopkg.in/robfig/cron.v2)

### Useful links

- [CRON Expression Format](https://godoc.org/gopkg.in/robfig/cron.v2#hdr-CRON_Expression_Format)
- [crontab.guru](https://crontab.guru)

### Example

```go
package main

import (
	"log"
	"time"
	"github.com/loeffel-io/cronjob"
)

// Cronjobs config
var cronjobs = cronjob.Cronjobs{
	Cronjobs: []cronjob.Cronjob{
		{
			Interval: "0 */1 * * * *", // @every 1m
			Call: func() {
				log.Println("cronjob is running ..")
			},
		},
		{
			Interval: "0 */3 * * * *", // @every 3m
			Call: func() {
				log.Println("cronjob is running ..")
			},
		},
	},
}

func init() {
	// Setup cronjobs
	cronjob.Setup(cronjobs)
}
```