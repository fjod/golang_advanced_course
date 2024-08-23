package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fjod/golang_advanced_course/internal"
	data "github.com/fjod/golang_advanced_course/internal/Data"
	"net/http"
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
			send(g, server)

		case c := <-chc10s:

			fmt.Println("отправка counter ", c.GetName(), " ", c.GetValue())
			send(c, server)

		default:
			time.Sleep(sleepDur)
		}
	}
}

func send(m data.IMetric, server string) {
	fmt.Println("пробуем что-то отправить")
	s := fmt.Sprintf("http://%v/update/", server)
	var j = m.ToJSON()
	fmt.Println(j)
	jsonData, err := json.Marshal(j)
	if err != nil {
		fmt.Println(err)
	}
	resp, err := http.Post(s, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Status)
	fmt.Println("пробуем закрыть body")
	err = resp.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
}
