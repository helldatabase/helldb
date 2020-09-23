package store

import (
	"sync"
	"testing"

	. "helldb/engine/types"
)

var store = Init()

var data = map[string]BaseType{
	"age":         NewInt(69),
	"name":        NewString("Manan"),
	"environment": NewString("remain"),
	"dwarf": NewCollection([]BaseType{
		NewString("hello-world"),
		NewString("this-is-pt-2"),
		NewString("i-ve-run-out"),
		NewString("involve"),
		NewString("statement"),
		NewString("plant"),
		NewString("international"),
		NewString("south"),
		NewString("something"),
		NewString("major"),
		NewString("remain"),
	}),
	"posts": NewCollection([]BaseType{
		NewInt(31),
		NewCollection([]BaseType{
			NewBoolean(false),
			NewBoolean(true),
			NewInt(310),
		}),
		NewString("hello-world"),
		NewString("plant"),
		NewString("international"),
		NewString("south"),
		NewString("something"),
	}),
}

func put(key string, value BaseType, store *Store, wg *sync.WaitGroup) {
	defer wg.Done()
	store.Put(key, value)
}

func initStore() {
	var wg sync.WaitGroup
	for key, val := range data {
		wg.Add(1)
		go put(key, val, store, &wg)
	}
	wg.Wait()
}

func TestStore_Len(t *testing.T) {
	if store.Len() == 0 {
		initStore()
	}
	if store.Len() != uint64(len(data)) {
		t.Errorf("store.Len() returned %d; expected=%d", store.Len(), uint64(len(data)))
	}
}

func TestStore_JSON(t *testing.T) {
	if store.Len() == 0 {
		initStore()
	}
	if store.JSON() != dataJson {
		t.Errorf("store.JSON() returned %s. expected=%s", store.JSON(), dataJson)
	}
}

func TestStore_Del(t *testing.T) {
	if store.Len() == 0 {
		initStore()
	}
	keys := []string{"age"}
	if store.Del(keys)[0].Native() != true {
		t.Error("key `age` does not exist in store")
	}
	if store.Get(keys)[0] != nil {
		t.Error("key `age` was not deleted from store")
	}

	if store.Del([]string{"abcd"})[0].Native() != false {
		t.Error("invalid key `abcd` deleted from store")
	}
}

func TestStore_Get(t *testing.T) {

	var wg sync.WaitGroup

	if store.Len() != 0 {
		initStore()
	}

	for key, val := range data {
		wg.Add(1)
		go compare(t, key, val, store, &wg)
	}
	wg.Wait()

}

func compare(t *testing.T, key string, val BaseType, store *Store, wg *sync.WaitGroup) bool {
	defer wg.Done()
	if vals := store.Get([]string{key}); len(vals) != 0 {
		if vals[0] != val {
			t.Errorf("expected value %v, got=%v", val.String(), vals[0].String())
			return false
		}
	} else {
		t.Error("store didn't return any values")
		return false
	}
	return true
}

const dataJson = "{\"age\":69,\"dwarf\":[\"hello-world\",\"this-is-pt-2\",\"i-ve-run-out\",\"involve\",\"statement\"," +
	"\"plant\",\"international\",\"south\",\"something\",\"major\",\"remain\"],\"environment\":\"remain\",\"name\":" +
	"\"Manan\",\"posts\":[31,[false,true,310],\"hello-world\",\"plant\",\"international\",\"south\",\"something\"]}"
