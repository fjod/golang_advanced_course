package internal

import (
	internal "github.com/fjod/golang_advanced_course/internal/Data"
	"math/rand"
	"runtime"
	"time"
)

func Monitor(intervalSecs int, chg chan<- internal.Gauge, chc chan<- internal.Counter) {
	var rtm runtime.MemStats
	t := time.Duration(intervalSecs) * time.Second
	sleepDur := time.Duration(100) * time.Millisecond
	ticker10 := time.NewTicker(t)
	pollCount := 0
	for {
		select {
		case <-ticker10.C:
			{
				runtime.ReadMemStats(&rtm)
				for name, f := range internal.MemMetrics {
					f(name, rtm, chg)
				}
				pollCount += intervalSecs
				chc <- internal.Counter{
					Name: "PollCount",
					Val:  int64(pollCount),
				}
				chg <- internal.Gauge{
					Name: "RandomValue",
					Val:  rand.Float64(),
				}
			}
		default:
			time.Sleep(sleepDur)
		}
	}
}
