package main

import (
	"fmt"
	"github.com/fjod/golang_advanced_course/internal"
	data "github.com/fjod/golang_advanced_course/internal/Data"
	"net/http"
	"os"
	"strconv"
)

func findWordsBetweenSlashes(s string) []string {
	var words []string
	var word string
	for _, c := range s {
		if c == '/' {
			if word != "" {
				words = append(words, word)
				word = ""
			}
		} else {
			word += string(c)
		}
	}
	if word != "" {
		words = append(words, word)
	}
	return words
}

func mainPage(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		q := req.URL.RequestURI()
		words := findWordsBetweenSlashes(q)
		if len(words) < 3 {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if words[0] != "update" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if words[1] == "gauge" {
			n, err := strconv.ParseFloat(words[3], 64)
			if err == nil {
				g := data.Gauge{
					Name: words[2],
					Val:  n,
				}
				err = internal.AppendMetric(g, storages)
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
				}
				fmt.Println("приняли gauge ", g.GetName(), " ", g.GetValue())
				w.WriteHeader(http.StatusOK)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if words[1] == "counter" {
			n, err := strconv.ParseInt(words[3], 10, 64)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
			c := data.Counter{
				Name: words[2],
				Val:  n,
			}
			err = internal.AppendMetric(c, storages)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
			fmt.Println("приняли counter ", c.GetName(), " ", c.GetValue())
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
}

var storages = make(map[int]internal.Storage) // не надо хранить это в памяти

func main() {
	fmt.Println("server")
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Failed to get executable path:", err)
	} else {
		fmt.Println("Executable:", exePath)
	}

	err = http.ListenAndServe(`localhost:8080`, http.HandlerFunc(mainPage))
	if err != nil {
		panic(err)
	}
}
