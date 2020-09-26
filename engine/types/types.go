package types

import (
	"fmt"
	"strconv"
	"strings"
	"unsafe"
)

// BaseType makes up the most basic building blocks of all
// primitive data types for HellDB. There is no differentiation
// between composite or self sufficient types.
//
// The Name function is a meta function that returns the string
// identifier for a type - in case we need a strongly typed interpreter
// that supports operations for reducing or mapping over keys.
//
// SizeOf returns an unsigned integer for the number of bytes taken
// up by a data member - it's statically defined for some types (ex. Boolean),
// grows with size for some (ex. String, Collection) and is platform dependent
// for others (ex. Int).
//
// String is another meta function returns a native string
// representation for a type useful for taking snapshots (currently
// unsupported) before shutdown or at random intervals.
//
// Native returns a native Go type representation for a BaseType. Compound
// data types can also be represented such as heterogeneous arrays (Collection)
// since they all implement the BaseType interface and required methods.
type BaseType interface {
	Name() string
	SizeOf() uint
	String() string
	Native() interface{}
}

// Int is a struct that implements BaseType for signed 64
// bit integers.
type Int struct {
	Data int64 `json:"int"`
}

func (i *Int) Native() interface{} { return i.Data }
func (i *Int) Name() string        { return "integer" }
func NewInt(data int64) *Int       { return &Int{Data: data} }
func (i *Int) SizeOf() uint        { return uint(unsafe.Sizeof(i.Data)) }
func (i *Int) String() string      { return strconv.FormatInt(i.Data, 10) }

type String struct {
	Data string `json:"string"`
}

func (s *String) Name() string        { return "str" }
func (s *String) String() string      { return s.Data }
func (s *String) Native() interface{} { return s.Data }
func (s *String) SizeOf() uint        { return uint(len(s.Data)) }
func NewString(data string) *String   { return &String{Data: data} }

type Boolean struct {
	Data bool `json:"bool"`
}

func (b *Boolean) SizeOf() uint        { return 1 }
func (b *Boolean) Native() interface{} { return b.Data }
func (b *Boolean) Name() string        { return "boolean" }
func NewBoolean(data bool) *Boolean    { return &Boolean{Data: data} }

func (b *Boolean) String() string {
	if b.Data {
		return "true"
	} else {
		return "false"
	}
}

type Collection struct {
	Data []BaseType `json:"collection"`
}

func (c *Collection) Name() string              { return "collection" }
func (c *Collection) Native() interface{}       { return serializeCollection(c) }
func NewCollection(data []BaseType) *Collection { return &Collection{Data: data} }

func serializeCollection(collection *Collection) []interface{} {
	var list []interface{}
	for _, val := range collection.Data {
		if c, ok := val.(*Collection); ok {
			list = append(list, serializeCollection(c))
		} else {
			list = append(list, val.Native())
		}
	}
	return list
}

func (c *Collection) SizeOf() uint {
	if l := uint(len(c.Data)); l == 0 {
		return 0
	} else {
		var total uint = 0
		for _, item := range c.Data {
			total += item.SizeOf()
		}
		return total
	}
}

func (c *Collection) String() string {
	if len(c.Data) == 0 {
		return "[]"
	} else {
		var b strings.Builder
		b.WriteString("[ ")
		for i := 0; i < len(c.Data)-1; i++ {
			_, _ = fmt.Fprintf(&b, "%s, ", c.Data[i].String())
		}
		_, _ = fmt.Fprintf(&b, "%s ]", c.Data[len(c.Data)-1].String())
		return b.String()
	}
}
