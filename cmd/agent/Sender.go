package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fjod/golang_advanced_course/internal"
	data "github.com/fjod/golang_advanced_course/internal/Data"
	"io/ioutil"
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
			send2(g, server)

		case c := <-chc10s:
			fmt.Println("отправка counter ", c.GetName(), " ", c.GetValue())
			send2(c, server)

		default:
			time.Sleep(sleepDur)
		}
	}
}

/*
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	func generateRandomString(length int) string {
		rand.Seed(time.Now().UnixNano())
		b := make([]byte, length)
		for i := range b {
			b[i] = charset[rand.Intn(len(charset))]
		}
		return string(b)
	}
*/
func send2(m data.IMetric, server string) {
	fmt.Printf(time.DateTime)
	s := fmt.Sprintf("http://%v/update/", server)
	var j = m.ToJSON()
	jsonData, err := json.Marshal(j)
	client := &http.Client{}
	req, err := http.NewRequest("POST", s, bytes.NewBuffer(jsonData))
	req.Close = true
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		// whatever
		fmt.Printf("ошибка client.Do \n")
		fmt.Println(err)
		return
	}

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		// Whatever
		fmt.Printf("ошибка ioutil.ReadAll \n")
		fmt.Println(err)
		return
	}

	err = resp.Body.Close()
	if err != nil {
		fmt.Printf("ошибка Close")
		fmt.Println(err)
	}
}

/*
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

	err = resp.Body.Close()
	if err != nil {
		fmt.Printf("ошибка Close %v", code)
		fmt.Println(err, code)
	}
	janitor.Unlock()
}
*/
