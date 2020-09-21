package types

import (
	"fmt"
	"strconv"
	"strings"
	"unsafe"
)

type BaseType interface {
	Name() string
	SizeOf() uint
	String() string
	Native() interface{}
}

type Int struct {
	data int64
}

type String struct {
	data string
}

type Boolean struct {
	data bool
}

type Collection struct {
	data []BaseType
}

/* -----------------int------------------ */

func (i *Int) Name() string {
	return "integer"
}

func (i *Int) SizeOf() uint {
	return uint(unsafe.Sizeof(i.data))
}

func (i *Int) String() string {
	return strconv.FormatInt(i.data, 10)
}

func (i *Int) Native() interface{} {
	return i.data
}

func NewInt(data int64) *Int {
	return &Int{data: data}
}

/* -----------------str----------------- */

func (s *String) Name() string {
	return "str"
}

func (s *String) SizeOf() uint {
	return uint(len(s.data))
}

func (s *String) String() string {
	return s.data
}

func (s *String) Native() interface{} {
	return s.data
}

func NewString(data string) *String {
	return &String{data: data}
}

/* ----------------collection---------------- */

func (c *Collection) Name() string {
	return "collection"
}

func (c *Collection) SizeOf() uint {
	if l := uint(len(c.data)); l == 0 {
		return 0
	} else {
		var total uint = 0
		for _, item := range c.data {
			total += item.SizeOf()
		}
		return total
	}
}

func (c *Collection) String() string {
	if len(c.data) == 0 {
		return "[]"
	} else {
		var b strings.Builder
		b.WriteString("[ ")
		for i := 0; i < len(c.data)-1; i++ {
			_, _ = fmt.Fprintf(&b, "%s, ", c.data[i].String())
		}
		_, _ = fmt.Fprintf(&b, "%s ]", c.data[len(c.data)-1].String())
		return b.String()
	}
}

func serializeCollection(collection *Collection) []interface{} {
	var list []interface{}
	for _, val := range collection.data {
		if c, ok := val.(*Collection); ok {
			list = append(list, serializeCollection(c))
		} else {
			list = append(list, val.Native())
		}
	}
	return list
}

func (c *Collection) Native() interface{} {
	return serializeCollection(c)
}

func NewCollection(data []BaseType) *Collection {
	return &Collection{data: data}
}

/* ----------------boolean---------------- */

func (b *Boolean) Name() string {
	return "boolean"
}

func (b *Boolean) SizeOf() uint {
	return 1
}

func (b *Boolean) String() string {
	if b.data {
		return "true"
	} else {
		return "false"
	}
}

func (b *Boolean) Native() interface{} {
	return b.data
}

func NewBoolean(data bool) *Boolean {
	return &Boolean{data: data}
}
