package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T){
	cache := NewCache(time.Millisecond)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddGetCache(t *testing.T){
	cache := NewCache(time.Millisecond)

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("val1"),
		},
		{
			inputKey: "key2",
			inputVal: []byte("val2"),
		},
		{
			inputKey: "key3",
			inputVal: []byte("val3"),
		},
	}

	for _, cas := range cases {
		cache.Add(cas.inputKey, []byte(cas.inputVal))
		actualValue, ok := cache.Get(cas.inputKey)
		if !ok {
			t.Errorf("%v not found", cas.inputKey)
			return 
		}

		if string(actualValue) != string(cas.inputVal) {
			t.Errorf("values do not match: %v vs %v", actualValue, cas.inputVal)
			return
		}
	}
}

func TestReapCache(t *testing.T) {
	const interval = 5 * time.Millisecond
	const waitTime = interval + 5*time.Millisecond

	cache := NewCache(interval)
	cache.Add("key1", []byte("val1"))

	_, ok := cache.Get("key1")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("key1")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}