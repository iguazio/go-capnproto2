// Code generated by capnpc-go. DO NOT EDIT.

package books

import (
	capnp "github.com/iguazio/go-capnproto2"
	text "github.com/iguazio/go-capnproto2/encoding/text"
	schemas "github.com/iguazio/go-capnproto2/schemas"
)

type Book struct{ capnp.Struct }

// Book_TypeID is the unique identifier for the type Book.
const Book_TypeID = 0x8100cc88d7d4d47c

func NewBook(s *capnp.Segment) (Book, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Book{st}, err
}

func NewRootBook(s *capnp.Segment) (Book, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Book{st}, err
}

func ReadRootBook(msg *capnp.Message) (Book, error) {
	root, err := msg.RootPtr()
	return Book{root.Struct()}, err
}

func (s Book) String() string {
	str, _ := text.Marshal(0x8100cc88d7d4d47c, s.Struct)
	return str
}

func (s Book) Title() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Book) HasTitle() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Book) TitleBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Book) SetTitle(v string) error {
	return s.Struct.SetText(0, v)
}

func (s Book) PageCount() int32 {
	return int32(s.Struct.Uint32(0))
}

func (s Book) SetPageCount(v int32) {
	s.Struct.SetUint32(0, uint32(v))
}

// Book_List is a list of Book.
type Book_List struct{ capnp.List }

// NewBook creates a new list of Book.
func NewBook_List(s *capnp.Segment, sz int32) (Book_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return Book_List{l}, err
}

func (s Book_List) At(i int) Book { return Book{s.List.Struct(i)} }

func (s Book_List) Set(i int, v Book) error { return s.List.SetStruct(i, v.Struct) }

func (s Book_List) String() string {
	str, _ := text.MarshalList(0x8100cc88d7d4d47c, s.List)
	return str
}

// Book_Promise is a wrapper for a Book promised by a client call.
type Book_Promise struct{ *capnp.Pipeline }

func (p Book_Promise) Struct() (Book, error) {
	s, err := p.Pipeline.Struct()
	return Book{s}, err
}

const schema_85d3acc39d94e0f8 = "x\xda\x12\x88w`2d\xdd\xcf\xc8\xc0\x10(\xc2\xca" +
	"\xb6\xbf\xe6\xca\x95\xeb\x1dg\x1a\x03y\x18\x19\xff\xffx" +
	"0e\xee\xe15\x97[\x19X\x19\xd9\x19\x18\x04\x8fv" +
	"\x09\x9e\x05\xd1'\xcb\x19t\xff'\xe5\xe7g\x17\xeb%" +
	"'2\x16\xe4\x15X9\xe5\xe7g30\x0402\x06" +
	"r0\xb300\xb0020\x08j\x1a10\x04\xaa" +
	"03\x06\x1a0122\x8a0\x82\xc4t\x83\x18\x18" +
	"\x02u\x98\x19\x03-\x98\x18\xe5K2KrR\x19y" +
	"\x18\x98\x18y\x18\x18\xff\x17$\xa6\xa7:\xe7\x97\xe61" +
	"0\x960\xb2001\xb200\x02\x02\x00\x00\xff\xff" +
	"F\xa9$\xae"

func init() {
	schemas.Register(schema_85d3acc39d94e0f8,
		0x8100cc88d7d4d47c)
}
