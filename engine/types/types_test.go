package types

import (
	"testing"
	"unsafe"
)

func TestNewString(t *testing.T) {
	s := NewString("hello world")
	if s.SizeOf() != 11 {
		t.Errorf("expecting s.SizeOf() == 11. got=%d", s.SizeOf())
	} else if s.String() != "hello world" {
		t.Errorf("expecting s.String() as `hello world`. got=%s", s.String())
	}
}

func TestNewInt(t *testing.T) {
	i := NewInt(31442059)
	if i.SizeOf() != uint(unsafe.Sizeof(uint64(0))) {
		t.Errorf("expecting s.SizeOf() == %d. got=%d",
			unsafe.Sizeof(uint64(0)), i.SizeOf())
	} else if i.String() != "31442059" {
		t.Errorf("expecting s.String() as `31442059`. got=%s", i.String())
	}
}

func TestNewBoolean(t *testing.T) {
	var b BaseType
	b = NewBoolean(true)
	if b.SizeOf() != uint(1) {
		t.Errorf("expecting b.SizeOf() == %d. got=%d", 1, b.SizeOf())
	} else if b.String() != "true" {
		t.Errorf("expecting s.String() as `true`. got=%s", b.String())
	}
}

func TestString(t *testing.T) {
	var data = map[string]BaseType{
		"age":         NewInt(69),
		"bool":        NewBoolean(false),
		"name":        NewString("Manan"),
		"environment": NewString("remain"),
		"empty":       NewCollection([]BaseType{}),
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
	for _, val := range data {
		val.Name()
		val.Native()
		val.String()
		val.SizeOf()
	}
}
