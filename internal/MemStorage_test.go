package internal

import (
	internal "github.com/fjod/golang_advanced_course/internal/Data"
	"reflect"
	"testing"
)

func TestMemStorageAdd(t *testing.T) {
	storage := &memStorage{
		counters: counterStorage{data: make(map[string]internal.Counter)},
		gauges:   gaugeStorage{data: make(map[string]internal.Gauge)},
	}

	t.Run("add counter", func(t *testing.T) {
		err := storage.Add(int64(42), "test_counter")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		_, ok := storage.counters.data["test_counter"]
		if !ok {
			t.Error("counter not found in storage")
		}
	})

	t.Run("add gauge", func(t *testing.T) {
		err := storage.Add(3.14, "test_gauge")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		_, ok := storage.gauges.data["test_gauge"]
		if !ok {
			t.Error("gauge not found in storage")
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
		counters: counterStorage{data: make(map[string]internal.Counter)},
		gauges:   gaugeStorage{data: make(map[string]internal.Gauge)},
	}

	t.Run("key exists in counters", func(t *testing.T) {
		storage.counters.data["test_counter"] = internal.Counter{Name: "test_counter", Val: 42, State: internal.NotSent}
		if !storage.KeyExists("test_counter") {
			t.Error("expected key to exist in counters")
		}
	})

	t.Run("key exists in gauges", func(t *testing.T) {
		storage.gauges.data["test_gauge"] = internal.Gauge{Name: "test_gauge", Val: 3.14, State: internal.NotSent}
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
			counters: counterStorage{data: make(map[string]internal.Counter)},
			gauges:   gaugeStorage{data: make(map[string]internal.Gauge)},
		}
		if emptyStorage.KeyExists("any_key") {
			t.Error("expected key to not exist in empty storage")
		}
	})
}

func TestMemStorageAddOrEdit(t *testing.T) {
	storage := &memStorage{
		counters: counterStorage{data: make(map[string]internal.Counter)},
		gauges:   gaugeStorage{data: make(map[string]internal.Gauge)},
	}

	t.Run("add new counter", func(t *testing.T) {
		err := storage.AddOrEdit(int64(42), "new_counter")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		_, ok := storage.counters.data["new_counter"]
		if !ok {
			t.Error("counter not found in storage")
		}
	})

	t.Run("edit existing counter", func(t *testing.T) {
		storage.counters.data["existing_counter"] = internal.Counter{Name: "existing_counter", Val: 10, State: internal.NotSent}
		err := storage.AddOrEdit(int64(5), "existing_counter")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		counter, ok := storage.counters.data["existing_counter"]
		if !ok {
			t.Error("counter not found in storage")
		} else if counter.Val != 15 {
			t.Error("counter value should be false")
		}
	})

	t.Run("add new gauge", func(t *testing.T) {
		err := storage.AddOrEdit(3.14, "new_gauge")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		_, ok := storage.gauges.data["new_gauge"]
		if !ok {
			t.Error("gauge not found in storage")
		}
	})

	t.Run("edit existing gauge", func(t *testing.T) {
		storage.gauges.data["existing_gauge"] = internal.Gauge{Name: "existing_gauge", Val: 10.0, State: internal.NotSent}
		err := storage.AddOrEdit(5.0, "existing_gauge")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		gauge, ok := storage.gauges.data["existing_gauge"]
		if !ok {
			t.Error("gauge not found in storage")
		} else if gauge.Val != 5.0 {
			t.Error("gauge value should be false")
		}
	})

	t.Run("invalid type", func(t *testing.T) {
		err := storage.AddOrEdit("invalid", "invalid_type")
		if err == nil {
			t.Error("expected error for invalid type, but got nil")
		}
	})
}

func TestMemStorageGetValue(t *testing.T) {
	storage := &memStorage{
		counters: counterStorage{data: map[string]internal.Counter{
			"test_counter": {Name: "existing_counter", Val: 42, State: internal.NotSent},
		}},
		gauges: gaugeStorage{data: map[string]internal.Gauge{
			"test_gauge": {Name: "existing_gauge", Val: 3.14, State: internal.NotSent},
		}},
	}

	t.Run("get counter value", func(t *testing.T) {
		value, err := storage.GetValue("test_counter", "counter")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if value != "42" {
			t.Errorf("expected value '42', got '%s'", value)
		}
	})

	t.Run("get gauge value", func(t *testing.T) {
		value, err := storage.GetValue("test_gauge", "gauge")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if value != "3.14" {
			t.Errorf("expected value '3.14', got '%s'", value)
		}
	})

	t.Run("key not found", func(t *testing.T) {
		_, err := storage.GetValue("non_existent_key", "counter")
		if err == nil {
			t.Error("expected error for non-existent key, but got nil")
		}
	})

	t.Run("invalid metric type", func(t *testing.T) {
		_, err := storage.GetValue("test_counter", "invalid_type")
		if err == nil {
			t.Error("expected error for invalid metric type, but got nil")
		}
	})

	t.Run("empty storage", func(t *testing.T) {
		emptyStorage := &memStorage{
			counters: counterStorage{data: make(map[string]internal.Counter)},
			gauges:   gaugeStorage{data: make(map[string]internal.Gauge)},
		}
		_, err := emptyStorage.GetValue("any_key", "counter")
		if err == nil {
			t.Error("expected error for empty storage, but got nil")
		}
	})
}

func TestMemStoragePrint(t *testing.T) {
	t.Run("empty storage", func(t *testing.T) {
		storage := &memStorage{
			counters: counterStorage{data: make(map[string]internal.Counter)},
			gauges:   gaugeStorage{data: make(map[string]internal.Gauge)},
		}
		result := storage.Print()
		if len(result) != 0 {
			t.Errorf("expected empty map, but got %v", result)
		}
	})

	t.Run("non-empty storage", func(t *testing.T) {
		storage := &memStorage{
			counters: counterStorage{data: map[string]internal.Counter{
				"counter1": {Name: "counter1", Val: 42, State: internal.NotSent},
				"counter2": {Name: "counter1", Val: 100, State: internal.NotSent},
			}},
			gauges: gaugeStorage{data: map[string]internal.Gauge{
				"gauge1": {Name: "counter1", Val: 3.14, State: internal.NotSent},
				"gauge2": {Name: "counter1", Val: 10.0, State: internal.NotSent},
			}},
		}
		expected := map[string]string{
			"counter1": "42",
			"counter2": "100",
			"gauge1":   "3.14",
			"gauge2":   "10",
		}
		result := storage.Print()
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("unexpected result, got %v, expected %v", result, expected)
		}
	})
}
