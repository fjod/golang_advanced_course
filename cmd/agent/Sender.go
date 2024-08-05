package main

import (
	"fmt"
	"github.com/fjod/golang_advanced_course/internal"
	data "github.com/fjod/golang_advanced_course/internal/Data"
	"net/http"
	"strings"
	"time"
)

func SendMetrics(server string, reportInterval int, pollInterval int) {
	chg10s := make(chan data.Gauge)
	chc10s := make(chan data.Counter)
	sleepDur := time.Duration(100) * time.Millisecond
	go internal.CollectMetrics(pollInterval, reportInterval, chg10s, chc10s)
	for {
		select {
		case g := <-chg10s:

			fmt.Println("отправка gauge ", g.GetName(), " ", g.GetValue())
			send("gauge", g, server)

		case c := <-chc10s:

			fmt.Println("отправка counter ", c.GetName(), " ", c.GetValue())
			send("counter", c, server)

		default:
			time.Sleep(sleepDur)
		}
	}
}

func send(name string, m data.IMetric, server string) {
	s := fmt.Sprintf("http://%v/update/%v/%v/%v", server, name, m.GetName(), m.GetValue())
	resp, err := http.Post(s, "text/plain", strings.NewReader(""))
	if err != nil {
		fmt.Println(err)
	}
	err = resp.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
}
