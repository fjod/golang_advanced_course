package main

import (
	"fmt"
	"github.com/fjod/golang_advanced_course/internal"
	data "github.com/fjod/golang_advanced_course/internal/Data"
	"net/http"
	"strings"
	"time"
)

func SendMetrics() {
	chg_10s := make(chan data.Gauge)
	chc_10s := make(chan data.Counter)
	sleepDur := time.Duration(100) * time.Millisecond
	go internal.CollectMetrics(2, chg_10s, chc_10s)
	for {
		select {
		case g := <-chg_10s:

			fmt.Println("отправка gauge ", g.GetName(), " ", g.GetValue())
			send("gauge", g)

		case c := <-chc_10s:

			fmt.Println("отправка counter ", c.GetName(), " ", c.GetValue())
			send("counter", c)

		default:
			time.Sleep(sleepDur)
		}
	}
}

func send(name string, m data.IMetric) {
	s := fmt.Sprintf("http://localhost:8080/update/%v/%v/%v", name, m.GetName(), m.GetValue())
	_, err := http.Post(s, "text/plain", strings.NewReader(""))
	if err != nil {
		fmt.Println(err)
	}
}
