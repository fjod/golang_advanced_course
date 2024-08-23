package internal

import (
	"fmt"
	internal "github.com/fjod/golang_advanced_course/internal/Data"
)

type Storage struct {
	StorageOperations memStorage
}

func NewStorage() *Storage {
	st := &Storage{
		StorageOperations: memStorage{},
	}
	st.StorageOperations.Init()
	return st
}

type memStorage struct {
	gauges   gaugeStorage
	counters counterStorage
}

type counterStorage struct {
	data map[string]internal.Counter
}

type gaugeStorage struct {
	data map[string]internal.Gauge
}

type StorageOperations interface {
	Add(any interface{}, name string) error
	Init()
	KeyExists(string) bool
	AddOrEdit(any interface{}, name string) error
	GetValue(name string, metricType string) (string, error)
	GetJSONValue(name string, metricType string) (internal.Metrics, error)
	Print() map[string]string
}

func (r *memStorage) Add(any interface{}, name string) error {
	_, ok := any.(int64)
	if ok {
		c := &internal.Counter{
			Name:  name,
			Val:   any.(int64),
			State: internal.NotSent,
		}
		r.counters.data[name] = *c
		return nil
	}
	_, ok = any.(float64)
	if ok {
		g := &internal.Gauge{
			Name: name,
			Val:  any.(float64),
		}
		r.gauges.data[name] = *g
		return nil
	}

	err := fmt.Errorf("wrong type")
	return err
}

func (r *memStorage) Init() {
	r.gauges.data = make(map[string]internal.Gauge)
	r.counters.data = make(map[string]internal.Counter)
}

func (r *memStorage) KeyExists(k string) bool {
	for gauge := range r.gauges.data {
		if gauge == k {
			return true
		}
	}
	for counter := range r.counters.data {
		if counter == k {
			return true
		}
	}
	return false
}

func (r *memStorage) AddOrEdit(any interface{}, name string) error {
	cval, ok := any.(int64)
	if ok {
		counter, counterFound := r.counters.data[name]
		if counterFound {
			delete(r.counters.data, name)
			counter.Val += cval
			counter.State = internal.NotSent
			r.counters.data[name] = counter
			return nil
		}
		c := &internal.Counter{
			Name:  name,
			Val:   cval,
			State: internal.NotSent,
		}
		r.counters.data[name] = *c
		return nil
	}
	fval, ok := any.(float64)
	if ok {
		gauge, gaugeFound := r.gauges.data[name]
		if gaugeFound {
			delete(r.gauges.data, name)
			gauge.Val = fval
			gauge.State = internal.NotSent
			r.gauges.data[name] = gauge
			return nil
		}

		g := &internal.Gauge{
			Name: name,
			Val:  fval,
		}
		r.gauges.data[name] = *g
		return nil
	}

	err := fmt.Errorf("wrong type")
	return err
}

func (r *memStorage) GetValue(name string, metricType string) (string, error) {
	if metricType == "counter" {
		c, ok := r.counters.data[name]
		if ok {
			return fmt.Sprintf("%v", c.Val), nil
		}
	} else if metricType == "gauge" {
		g, ok := r.gauges.data[name]
		if ok {
			return fmt.Sprintf("%v", g.Val), nil
		}
	}
	err := fmt.Errorf("key not found or not supported type")
	return "", err
}

func (r *memStorage) GetJSONValue(name string, metricType string) (internal.Metrics, error) {
	if metricType == "counter" {
		fmt.Println("inside GetJSONValue")
		fmt.Println(name)
		fmt.Println(metricType)
		fmt.Println(r.counters.data)
		c, ok := r.counters.data[name]
		if ok {
			return internal.Metrics{
				ID:    c.Name,
				Delta: c.Val,
				MType: "counter",
			}, nil
		}
	} else if metricType == "gauge" {
		g, ok := r.gauges.data[name]
		if ok {
			return internal.Metrics{
				ID:    g.Name,
				Value: g.Val,
				MType: "gauge",
			}, nil
		}
	}
	err := fmt.Errorf("key not found or not supported type")
	return internal.Metrics{}, err
}

func (r *memStorage) Print() map[string]string {
	ret := make(map[string]string)
	for gName, gauge := range r.gauges.data {
		ret[gName] = fmt.Sprintf("%v", gauge.Val)
	}
	for cName, counter := range r.counters.data {
		ret[cName] = fmt.Sprintf("%v", counter.Val)
	}
	return ret
}
