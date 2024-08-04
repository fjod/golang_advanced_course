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
	data map[internal.Counter]bool
}

type gaugeStorage struct {
	data map[internal.Gauge]bool
}

type StorageOperations interface {
	Add(any interface{}, name string) error
	Init()
	KeyExists(string) bool
}

func (r *memStorage) Add(any interface{}, name string) error {
	_, ok := any.(int64)
	if ok {
		c := &internal.Counter{
			Name: name,
			Val:  any.(int64),
		}
		r.counters.data[*c] = false
		return nil
	}
	_, ok = any.(float64)
	if ok {
		g := &internal.Gauge{
			Name: name,
			Val:  any.(float64),
		}
		r.gauges.data[*g] = false
		return nil
	}

	err := fmt.Errorf("wrong type")
	return err
}

func (r *memStorage) Init() {
	r.gauges.data = make(map[internal.Gauge]bool)
	r.counters.data = make(map[internal.Counter]bool)
}

func (r *memStorage) KeyExists(k string) bool {
	for gauge := range r.gauges.data {
		if gauge.Name == k {
			return true
		}
	}
	for counter := range r.counters.data {
		if counter.Name == k {
			return true
		}
	}
	return false
}
