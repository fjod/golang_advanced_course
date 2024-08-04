package internal

import (
	"fmt"
	internal "github.com/fjod/golang_advanced_course/internal/Data"
	"sync"
	"time"
)

var storages = make(map[int]Storage)
var janitor = &sync.Mutex{}

func CollectMetrics(pollinterval2s int, chg10s chan<- internal.Gauge, chc10s chan<- internal.Counter) {
	chg2s := make(chan internal.Gauge)
	chc2s := make(chan internal.Counter)
	st := NewStorage()
	storages[len(storages)] = *st
	ticker10 := time.NewTicker(10 * time.Second)
	ticker60 := time.NewTicker(60 * time.Second)
	go Monitor(pollinterval2s, chg2s, chc2s)
	for {
		select {
		case g := <-chg2s:
			err := AppendMetric(g, storages)
			if err != nil {
				return
			}
		case c := <-chc2s:
			err := AppendMetric(c, storages)
			if err != nil {
				return
			}
		case <-ticker10.C:
			{
				janitor.Lock()
				fmt.Println("отправка метрик из CollectMetrics")
				// отправляем все что насобирали
				for _, s := range storages {
					for g := range s.StorageOperations.gauges.data {
						s.StorageOperations.gauges.data[g] = false
						chg10s <- g
					}
					for c := range s.StorageOperations.counters.data {
						s.StorageOperations.counters.data[c] = false
						chc10s <- c
					}
				}
				janitor.Unlock()
			}
		case <-ticker60.C:
			{
				janitor.Lock()
				// чистим то что отправили
				Clean(&storages)
				janitor.Unlock()
			}
		}
	}
}

func Clean(s *map[int]Storage) {
	for i, storage := range *s {
		for _, b := range storage.StorageOperations.gauges.data {
			// удаляем весь набор метрик, если хотя бы одну из них уже отправили
			// все метрики не проверяются из-за мютекса
			if b {
				delete(*s, i)
			}
			break
		}
	}
}

func AppendMetric(c internal.IMetric, s map[int]Storage) error {
	if len(s) == 0 {
		st := NewStorage()
		s[0] = *st
		err := (*st).StorageOperations.Add(c.GetValue(), c.GetName())
		if err != nil {
			return err
		}
		return nil
	}
	last := s[len(s)-1]
	if last.StorageOperations.KeyExists(c.GetName()) {
		st := NewStorage()
		s[len(s)] = *st
	}
	last = s[len(s)-1]
	err := last.StorageOperations.Add(c.GetValue(), c.GetName())
	if err != nil {
		return err
	}
	return nil
}

func SaveMetric(c internal.IMetric, s Storage) error {
	err := s.StorageOperations.AddOrEdit(c.GetValue(), c.GetName())
	if err != nil {
		return err
	}
	return nil
}
