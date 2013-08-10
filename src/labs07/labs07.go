package labs07

import "reflect"
import "unsafe"

type BigStruct struct {
	next *BigStruct
	C01  int
	C02  int
	C03  int
	C04  int
	C05  int
	C06  int
	C07  int
	C08  int
	C09  int
	C10  int
	C11  int
	C12  int
	C13  int
	C14  int
	C15  int
	C16  int
	C17  int
	C18  int
	C19  int
	C20  int
	C21  int
	C22  int
	C23  int
	C24  int
	C25  int
	C26  int
	C27  int
	C28  int
	C29  int
	C30  int
}

// operator
const OP_EQ = 0 // ==

// type
const TP_INT = 0 // int

type Query struct {
	conditions []*Condition
}

type Condition struct {
	op     int
	tp     int
	offset uintptr
	value  uint64
}

func (q *Query) Match(n *BigStruct) bool {
	for _, c := range q.conditions {
		if c.Match(n) == false {
			return false
		}
	}
	return true
}

func (c *Condition) Match(n *BigStruct) bool {
	switch c.op {
	case OP_EQ:
		var b = unsafe.Pointer(uintptr(unsafe.Pointer(n)) + c.offset)
		switch c.tp {
		case TP_INT:
			return int(c.value) == *(*int)(b)
		}
	}
	return false
}

func NewQuery(name string, op int, value int) *Query {
	var t = reflect.TypeOf(BigStruct{})
	var f, _ = t.FieldByName(name)

	return &Query{
		[]*Condition{
			&Condition{
				op:     op,
				tp:     TP_INT,
				offset: f.Offset,
				value:  uint64(value),
			},
		},
	}
}
