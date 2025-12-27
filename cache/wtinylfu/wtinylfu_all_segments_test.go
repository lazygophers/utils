package wtinylfu

import (
	"fmt"
	"testing"
)

func TestKeysAllSegmentsPopulated(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("window", 1)

	cache.Put("prob1", 10)
	cache.Get("prob1")

	cache.Put("prot1", 100)
	cache.Get("prot1")
	cache.Get("prot1")

	cache.Put("prob2", 20)
	cache.Get("prob2")

	cache.Put("prot2", 200)
	cache.Get("prot2")

	keys := cache.Keys()

	if len(keys) < 3 {
		t.Errorf("Expected at least 3 keys, got %d", len(keys))
	}

	keySet := make(map[string]bool)
	for _, key := range keys {
		keySet[key] = true
	}

	if !keySet["window"] {
		t.Error("Expected 'window' key to be present")
	}
	if !keySet["prob1"] && !keySet["prob2"] {
		t.Error("Expected at least one probation key to be present")
	}
	if !keySet["prot1"] && !keySet["prot2"] {
		t.Error("Expected at least one protected key to be present")
	}
}

func TestValuesAllSegmentsPopulated(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("window", 1)

	cache.Put("prob1", 10)
	cache.Get("prob1")

	cache.Put("prot1", 100)
	cache.Get("prot1")
	cache.Get("prot1")

	cache.Put("prob2", 20)
	cache.Get("prob2")

	cache.Put("prot2", 200)
	cache.Get("prot2")

	values := cache.Values()

	if len(values) < 3 {
		t.Errorf("Expected at least 3 values, got %d", len(values))
	}

	valueSet := make(map[int]bool)
	for _, val := range values {
		valueSet[val] = true
	}

	if !valueSet[1] {
		t.Error("Expected value 1 to be present")
	}
	if !valueSet[10] && !valueSet[20] {
		t.Error("Expected at least one probation value to be present")
	}
	if !valueSet[100] && !valueSet[200] {
		t.Error("Expected at least one protected value to be present")
	}
}

func TestItemsAllSegmentsPopulated(t *testing.T) {
	cache, err := New[string, int](100)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("window", 1)

	cache.Put("prob1", 10)
	cache.Get("prob1")

	cache.Put("prot1", 100)
	cache.Get("prot1")
	cache.Get("prot1")

	cache.Put("prob2", 20)
	cache.Get("prob2")

	cache.Put("prot2", 200)
	cache.Get("prot2")

	items := cache.Items()

	if len(items) < 3 {
		t.Errorf("Expected at least 3 items, got %d", len(items))
	}

	if val, ok := items["window"]; !ok || val != 1 {
		t.Error("Expected 'window' key with value 1 to be present")
	}
	if _, ok := items["prob1"]; !ok {
		if _, ok := items["prob2"]; !ok {
			t.Error("Expected at least one probation key to be present")
		}
	}
	if _, ok := items["prot1"]; !ok {
		if _, ok := items["prot2"]; !ok {
			t.Error("Expected at least one protected key to be present")
		}
	}
}

func TestDemoteFromProtectedFull(t *testing.T) {
	cache, err := New[string, int](30)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 20; i++ {
		key := fmt.Sprintf("prot%d", i)
		cache.Put(key, i)
		for j := 0; j < 8; j++ {
			cache.Get(key)
		}
	}

	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("prob%d", i)
		cache.Put(key, i+1000)
		cache.Get(key)
		cache.Get(key)
	}

	stats := cache.Stats()
	if stats.Size > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", stats.Size, cache.Cap())
	}
}

func TestEvictFromWindowAdmit(t *testing.T) {
	cache, err := New[string, int](20)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("freq", 1)
	for i := 0; i < 10; i++ {
		cache.Get("freq")
	}

	cache.Put("victim", 2)
	cache.Put("temp", 999)
	cache.Put("victim", 2)

	cache.Put("new", 3)

	stats := cache.Stats()
	if stats.Size > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", stats.Size, cache.Cap())
	}
}

func TestEvictFromWindowNoAdmit(t *testing.T) {
	cache, err := New[string, int](20)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 30; i++ {
		key := fmt.Sprintf("key%d", i)
		cache.Put(key, i)
	}

	stats := cache.Stats()
	if stats.Size > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", stats.Size, cache.Cap())
	}
}

func TestEvictFromProbationFull(t *testing.T) {
	cache, err := New[string, int](20)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 15; i++ {
		key := fmt.Sprintf("prob%d", i)
		cache.Put(key, i)
		cache.Get(key)
	}

	stats := cache.Stats()
	if stats.Size > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", stats.Size, cache.Cap())
	}
}

func TestEvictFromProtectedFull(t *testing.T) {
	cache, err := New[string, int](20)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 25; i++ {
		key := fmt.Sprintf("prot%d", i)
		cache.Put(key, i)
		for j := 0; j < 10; j++ {
			cache.Get(key)
		}
	}

	stats := cache.Stats()
	if stats.Size > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", stats.Size, cache.Cap())
	}
}

func TestEvictPrioritizeProbation(t *testing.T) {
	cache, err := New[string, int](20)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("prob%d", i)
		cache.Put(key, i)
		cache.Get(key)
	}

	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("prot%d", i)
		cache.Put(key, i+100)
		for j := 0; j < 5; j++ {
			cache.Get(key)
		}
	}

	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("window%d", i)
		cache.Put(key, i+200)
	}

	stats := cache.Stats()
	if stats.Size > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", stats.Size, cache.Cap())
	}
}

func TestEvictPrioritizeProtected(t *testing.T) {
	cache, err := New[string, int](20)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 15; i++ {
		key := fmt.Sprintf("prot%d", i)
		cache.Put(key, i)
		for j := 0; j < 5; j++ {
			cache.Get(key)
		}
	}

	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("window%d", i)
		cache.Put(key, i+100)
	}

	stats := cache.Stats()
	if stats.Size > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", stats.Size, cache.Cap())
	}
}

func TestEvictPrioritizeWindow(t *testing.T) {
	cache, err := New[string, int](20)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 20; i++ {
		key := fmt.Sprintf("window%d", i)
		cache.Put(key, i)
	}

	stats := cache.Stats()
	if stats.Size > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", stats.Size, cache.Cap())
	}
}

func TestDemoteFromProtectedEmptyProbation(t *testing.T) {
	cache, err := New[string, int](30)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 25; i++ {
		key := fmt.Sprintf("prot%d", i)
		cache.Put(key, i)
		for j := 0; j < 10; j++ {
			cache.Get(key)
		}
	}

	stats := cache.Stats()
	if stats.Size > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", stats.Size, cache.Cap())
	}
}

func TestDemoteFromProtectedFullProbation(t *testing.T) {
	cache, err := New[string, int](30)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	for i := 0; i < 20; i++ {
		key := fmt.Sprintf("prot%d", i)
		cache.Put(key, i)
		for j := 0; j < 8; j++ {
			cache.Get(key)
		}
	}

	for i := 0; i < 8; i++ {
		key := fmt.Sprintf("prob%d", i)
		cache.Put(key, i+1000)
		cache.Get(key)
		cache.Get(key)
	}

	stats := cache.Stats()
	if stats.Size > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", stats.Size, cache.Cap())
	}
}

func TestEvictFromWindowCompete(t *testing.T) {
	cache, err := New[string, int](20)
	if err != nil {
		t.Fatalf("Failed to create cache: %v", err)
	}

	cache.Put("freq", 1)
	for i := 0; i < 10; i++ {
		cache.Get("freq")
	}

	for i := 0; i < 15; i++ {
		key := fmt.Sprintf("prob%d", i)
		cache.Put(key, i)
		cache.Get(key)
	}

	cache.Put("victim", 2)
	cache.Put("temp", 999)
	cache.Put("victim", 2)

	stats := cache.Stats()
	if stats.Size > cache.Cap() {
		t.Errorf("Cache size %d exceeds capacity %d", stats.Size, cache.Cap())
	}
}
