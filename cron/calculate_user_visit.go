package cron

import (
	"time"

	"github.com/Ptt-official-app/go-openbbsmiddleware/schema"

	"github.com/sirupsen/logrus"
)

// RetryCalculateUserVisit make loop job to call CalculateUserVisit per 10 mins
func RetryCalculateUserVisit() {
	for {
		logrus.Infof("RetryCalculateUserVisit: to calculate user visit")
		CalculateUserVisit()
		logrus.Infof("RetryCalculateUserVisit: to sleep 10 mins")
		time.Sleep(10 * time.Minute)
	}
}

// CalculateUserVisit get user visit count from db
// and set to redis
func CalculateUserVisit() {
	count, err := schema.CalculateAllUserVisitCounts()
	if err != nil {
		logrus.Printf("get error in calculate user visit count %v", err)
	}
	// set to redis
	err = schema.SetUserVisitCount(count)
	if err != nil {
		logrus.Printf("set to redis  %v", err)
	}
}
