package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fjod/golang_advanced_course/internal"
	data "github.com/fjod/golang_advanced_course/internal/Data"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

var janitor = &sync.Mutex{}

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

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func send(m data.IMetric, server string) {
	janitor.Lock()
	code := generateRandomString(10)

	s := fmt.Sprintf("http://%v/update/", server)
	var j = m.ToJSON()
	fmt.Printf("пробуем что-то отправить %v %v\n", j, code)
	jsonData, err := json.Marshal(j)
	if err != nil {
		fmt.Printf("ошибка жсон %v\n", code)
		fmt.Println(err, code)
	}
	fmt.Println("отправляю")
	resp, err := http.Post(s, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("ошибка Post %v\n", code)
		fmt.Println(err, code)
		janitor.Unlock()
		return
	}
	fmt.Printf("ошибка Close %v", code)
	err = resp.Body.Close()
	if err != nil {
		fmt.Println(err, code)
	}
	janitor.Unlock()
}
