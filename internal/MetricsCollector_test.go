package internal

import (
	data "github.com/fjod/golang_advanced_course/internal/Data"
	"testing"
)

func TestAppendMetric(t *testing.T) {
	t.Run("different metrics placed in one storage", func(t *testing.T) {
		storages := make(map[int]Storage)

		metric1 := &data.Counter{Name: "metric1", Val: int64(10)}
		err := AppendMetric(metric1, storages)

		if err != nil {
			t.Errorf("unexpected error: %v , cant add good metric", err)
		}

		metric2 := &data.Counter{Name: "metric2", Val: int64(20)}
		err = AppendMetric(metric2, storages)

		if err != nil {
			t.Errorf("unexpected error: %v , cant add good metric", err)
		}

		if len(storages) != 1 {
			t.Errorf("must be one storage, got %v", len(storages))
		}
	})
	t.Run("same metrics placed in different storages", func(t *testing.T) {
		storages := make(map[int]Storage)

		metric1 := &data.Counter{Name: "metric1", Val: int64(10)}
		err := AppendMetric(metric1, storages)

		if err != nil {
			t.Errorf("unexpected error: %v , cant add good metric", err)
		}

		metric2 := &data.Counter{Name: "metric1", Val: int64(20)}
		err = AppendMetric(metric2, storages)

		if err != nil {
			t.Errorf("unexpected error: %v , cant add good metric", err)
		}

		if len(storages) != 2 {
			t.Errorf("must be two storages, got %v", len(storages))
		}

		if len(storages[0].StorageOperations.counters.data) != 1 {
			t.Errorf("must be one item in first storage, got %v", len(storages[0].StorageOperations.counters.data))
		}
		for _, counter := range storages[0].StorageOperations.counters.data {
			if counter.Name != "metric1" || counter.Val != 10 {
				t.Errorf("expected metric1, 10, got %v, %v", counter.Name, counter.Val)
			}
		}

		if len(storages[1].StorageOperations.counters.data) != 1 {
			t.Errorf("must be one item in second storage, got %v", len(storages[1].StorageOperations.counters.data))
		}
		for _, counter := range storages[1].StorageOperations.counters.data {
			if counter.Name != "metric1" || counter.Val != 20 {
				t.Errorf("expected metric1, 20, got %v, %v", counter.Name, counter.Val)
			}
		}
	})

	t.Run("many same metrics placed in different storages", func(t *testing.T) {
		storages := make(map[int]Storage)

		for i := 0; i < 10; i++ {
			metric1 := &data.Counter{Name: "metric1", Val: int64(10)}
			err := AppendMetric(metric1, storages)
			if err != nil {
				t.Errorf("unexpected error: %v , cant add good metric", err)
			}
		}

		if len(storages) != 10 {
			t.Errorf("must be 10 storages, got %v", len(storages))
		}
	})
}
func TestClean(t *testing.T) {
	t.Run("remove storage with sent gauges", func(t *testing.T) {
		storages := make(map[int]Storage)
		storages[0] = Storage{
			StorageOperations: memStorage{
				gauges: gaugeStorage{data: map[string]data.Gauge{
					"gauge1": {Name: "gauge1", Val: 1.0, State: data.Sent},
					"gauge2": {Name: "gauge2", Val: 2.0, State: data.Sent},
				}},
			},
		}
		storages[1] = Storage{
			StorageOperations: memStorage{
				gauges: gaugeStorage{data: map[string]data.Gauge{
					"gauge3": {Name: "gauge3", Val: 2.0, State: data.NotSent},
				}},
			},
		}

		Clean(&storages)

		if len(storages) != 1 {
			t.Errorf("expected one storage after cleaning, got %d", len(storages))
		}
	})

	t.Run("no storages to clean", func(t *testing.T) {
		storages := make(map[int]Storage)
		Clean(&storages)
		if len(storages) != 0 {
			t.Errorf("expected no storages after cleaning, got %d", len(storages))
		}
	})

	t.Run("no removal when storage was not sent", func(t *testing.T) {
		storages := make(map[int]Storage)
		storages[0] = Storage{
			StorageOperations: memStorage{
				counters: counterStorage{data: map[string]data.Counter{
					"counter1": {Name: "counter1", Val: 2.0, State: data.NotSent},
				}},
			},
		}

		Clean(&storages)

		if len(storages) != 1 {
			t.Errorf("expected one storage after cleaning, got %d", len(storages))
		}
	})
}
