package lru

import (
	"reflect"
	"testing"
)

type String string

func (s String) Len() int {
	return len(s)
}

func TestGet(t *testing.T) {
	lru := New(int64(0), nil)
	lru.Add("key1", String("1234"))
	if v, ok := lru.Get("key1"); !ok || string(v.(String)) != "1234" {
		t.Fatalf("cache hit failed: key1 = 1234")
	}
	if _, ok := lru.Get("key2"); ok {
		t.Fatalf("cache miss key2 failed")
	}
}

func TestRemoveoldest(t *testing.T) {
	k1, k2, k3 := "k1", "k2", "k3"
	v1, v2, v3 := "v1", "v2", "v3"
	cap := len(k1 + k2 + v1 + v2)
	lru := New(int64(cap), nil)
	lru.Add(k1, String(v1))
	lru.Add(k2, String(v2))
	lru.Add(k3, String(v3))

	if _, ok := lru.Get("k1"); ok {
		t.Fatalf("removeoldest failed")
	}
}

func TestOnEvicted(t *testing.T) {
	deleted := []string{}

	lru := New(int64(10), func(key string, val Value) {
		deleted = append(deleted, key)
	})

	lru.Add("k1", String("01234"))
	lru.Add("k2", String("56789"))
	lru.Add("k3", String("v3"))
	lru.Add("k4", String("v4"))

	expect := []string{"k1", "k2"}
	if !reflect.DeepEqual(expect, deleted) {
		t.Fatalf("onEvicted test failed")
	}
}
