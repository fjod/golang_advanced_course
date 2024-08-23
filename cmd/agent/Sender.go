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
	chg10s := make(chan data.Gauge, 1)
	chc10s := make(chan data.Counter, 1)
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
	s := fmt.Sprintf("http://%v/update/", server)
	var j = m.ToJSON()
	fmt.Printf("пробуем что-то отправить %v\n", j)
	jsonData, err := json.Marshal(j)
	if err != nil {
		fmt.Printf("ошибка жсон")
		fmt.Println(err)
	}
	fmt.Printf("отправляю")
	resp, err := http.Post(s, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("ошибка1")
		fmt.Println(err)
	}
	fmt.Printf("закрываю")
	err = resp.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
}
