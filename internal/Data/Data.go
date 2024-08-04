package data

import (
	"runtime"
)

type IMetric interface {
	GetName() string
	GetValue() any
}

type Gauge struct {
	Name string
	Val  float64
}

type Counter struct {
	Name string
	Val  int64
}

func (g Gauge) GetName() string {
	return g.Name
}

func (c Counter) GetName() string {
	return c.Name
}

func (g Gauge) GetValue() any {
	return g.Val
}

func (c Counter) GetValue() any {
	return c.Val
}

type GetAndSend func(string, runtime.MemStats, chan<- Gauge)

var MemMetrics = map[string]GetAndSend{
	"Alloc": func(name string, rtm runtime.MemStats, ch chan<- Gauge) {
		ch <- Gauge{
			Name: name,
			Val:  float64(rtm.Alloc),
		}
	},
	"BuckHashSys": func(name string, rtm runtime.MemStats, ch chan<- Gauge) {
		ch <- Gauge{
			Name: name,
			Val:  float64(rtm.BuckHashSys),
		}
	},
	"Frees": func(name string, rtm runtime.MemStats, ch chan<- Gauge) {
		ch <- Gauge{
			Name: name,
			Val:  float64(rtm.Frees),
		}
	},
	"GCCPUFraction": func(name string, rtm runtime.MemStats, ch chan<- Gauge) {
		ch <- Gauge{
			Name: name,
			Val:  rtm.GCCPUFraction,
		}
	},
	"GCSys": func(name string, rtm runtime.MemStats, ch chan<- Gauge) {
		ch <- Gauge{
			Name: name,
			Val:  float64(rtm.GCSys),
		}
	},
	"HeapAlloc": func(name string, rtm runtime.MemStats, ch chan<- Gauge) {
		ch <- Gauge{
			Name: name,
			Val:  float64(rtm.HeapAlloc),
		}
	},
	"HeapIdle": func(name string, rtm runtime.MemStats, ch chan<- Gauge) {
		ch <- Gauge{
			Name: name,
			Val:  float64(rtm.HeapIdle),
		}
	},
	"HeapInuse": func(name string, rtm runtime.MemStats, ch chan<- Gauge) {
		ch <- Gauge{
			Name: name,
			Val:  float64(rtm.HeapInuse),
		}
	},
	"HeapObjects": func(name string, rtm runtime.MemStats, ch chan<- Gauge) {
		ch <- Gauge{
			Name: name,
			Val:  float64(rtm.HeapObjects),
		}
	},
	"HeapReleased": func(name string, rtm runtime.MemStats, ch chan<- Gauge) {
		ch <- Gauge{
			Name: name,
			Val:  float64(rtm.HeapReleased),
		}
	},
	"HeapSys": func(name string, rtm runtime.MemStats, ch chan<- Gauge) {
		ch <- Gauge{
			Name: name,
			Val:  float64(rtm.HeapSys),
		}
	},
	"LastGC": func(name string, rtm runtime.MemStats, ch chan<- Gauge) {
		ch <- Gauge{
			Name: name,
			Val:  float64(rtm.LastGC),
		}
	},
	"Lookups": func(name string, rtm runtime.MemStats, ch chan<- Gauge) {
		ch <- Gauge{
			Name: name,
			Val:  float64(rtm.Lookups),
		}
	},
	"MCacheInuse": func(name string, rtm runtime.MemStats, ch chan<- Gauge) {
		ch <- Gauge{
			Name: name,
			Val:  float64(rtm.MCacheInuse),
		}
	},
	"MCacheSys": func(name string, rtm runtime.MemStats, ch chan<- Gauge) {
		ch <- Gauge{
			Name: name,
			Val:  float64(rtm.MCacheSys),
		}
	},
	"MSpanInuse": func(name string, rtm runtime.MemStats, ch chan<- Gauge) {
		ch <- Gauge{
			Name: name,
			Val:  float64(rtm.MSpanInuse),
		}
	},
	"MSpanSys": func(name string, rtm runtime.MemStats, ch chan<- Gauge) {
		ch <- Gauge{
			Name: name,
			Val:  float64(rtm.MSpanSys),
		}
	},
	"Mallocs": func(name string, rtm runtime.MemStats, ch chan<- Gauge) {
		ch <- Gauge{
			Name: name,
			Val:  float64(rtm.Mallocs),
		}
	},
	"NextGC": func(name string, rtm runtime.MemStats, ch chan<- Gauge) {
		ch <- Gauge{
			Name: name,
			Val:  float64(rtm.NextGC),
		}
	},
	"NumForcedGC": func(name string, rtm runtime.MemStats, ch chan<- Gauge) {
		ch <- Gauge{
			Name: name,
			Val:  float64(rtm.NumForcedGC),
		}
	},
	"NumGC": func(name string, rtm runtime.MemStats, ch chan<- Gauge) {
		ch <- Gauge{
			Name: name,
			Val:  float64(rtm.NumGC),
		}
	},
	"OtherSys": func(name string, rtm runtime.MemStats, ch chan<- Gauge) {
		ch <- Gauge{
			Name: name,
			Val:  float64(rtm.OtherSys),
		}
	},
	"PauseTotalNs": func(name string, rtm runtime.MemStats, ch chan<- Gauge) {
		ch <- Gauge{
			Name: name,
			Val:  float64(rtm.PauseTotalNs),
		}
	},
	"StackInuse": func(name string, rtm runtime.MemStats, ch chan<- Gauge) {
		ch <- Gauge{
			Name: name,
			Val:  float64(rtm.StackInuse),
		}
	},
	"StackSys": func(name string, rtm runtime.MemStats, ch chan<- Gauge) {
		ch <- Gauge{
			Name: name,
			Val:  float64(rtm.StackSys),
		}
	},
	"Sys": func(name string, rtm runtime.MemStats, ch chan<- Gauge) {
		ch <- Gauge{
			Name: name,
			Val:  float64(rtm.Sys),
		}
	},
	"TotalAlloc": func(name string, rtm runtime.MemStats, ch chan<- Gauge) {
		ch <- Gauge{
			Name: name,
			Val:  float64(rtm.TotalAlloc),
		}
	},
}
