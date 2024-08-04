package internal

import (
	internal "github.com/fjod/golang_advanced_course/internal/Data"
	"testing"
)

func TestMemStorageAdd(t *testing.T) {
	storage := &memStorage{
		counters: counterStorage{data: make(map[internal.Counter]bool)},
		gauges:   gaugeStorage{data: make(map[internal.Gauge]bool)},
	}

	t.Run("add counter", func(t *testing.T) {
		err := storage.Add(int64(42), "test_counter")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		counter, ok := storage.counters.data[internal.Counter{Name: "test_counter", Val: 42}]
		if !ok {
			t.Error("counter not found in storage")
		} else if counter {
			t.Error("counter value should be false")
		}
	})

	t.Run("add gauge", func(t *testing.T) {
		err := storage.Add(float64(3.14), "test_gauge")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		gauge, ok := storage.gauges.data[internal.Gauge{Name: "test_gauge", Val: 3.14}]
		if !ok {
			t.Error("gauge not found in storage")
		} else if gauge {
			t.Error("gauge value should be false")
		}
	})

	t.Run("add invalid type", func(t *testing.T) {
		err := storage.Add("invalid", "test_invalid")
		if err == nil {
			t.Error("expected error for invalid type, but got nil")
		}
	})
}

func TestMemStorageKeyExists(t *testing.T) {
	storage := &memStorage{
		counters: counterStorage{data: make(map[internal.Counter]bool)},
		gauges:   gaugeStorage{data: make(map[internal.Gauge]bool)},
	}

	t.Run("key exists in counters", func(t *testing.T) {
		storage.counters.data[internal.Counter{Name: "test_counter", Val: 42}] = false
		if !storage.KeyExists("test_counter") {
			t.Error("expected key to exist in counters")
		}
	})

	t.Run("key exists in gauges", func(t *testing.T) {
		storage.gauges.data[internal.Gauge{Name: "test_gauge", Val: 3.14}] = false
		if !storage.KeyExists("test_gauge") {
			t.Error("expected key to exist in gauges")
		}
	})

	t.Run("key does not exist", func(t *testing.T) {
		if storage.KeyExists("non_existent_key") {
			t.Error("expected key to not exist")
		}
	})

	t.Run("empty storage", func(t *testing.T) {
		emptyStorage := &memStorage{
			counters: counterStorage{data: make(map[internal.Counter]bool)},
			gauges:   gaugeStorage{data: make(map[internal.Gauge]bool)},
		}
		if emptyStorage.KeyExists("any_key") {
			t.Error("expected key to not exist in empty storage")
		}
	})
}
