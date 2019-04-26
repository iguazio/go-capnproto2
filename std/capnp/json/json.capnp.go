// Code generated by capnpc-go. DO NOT EDIT.

package json

import (
	math "math"
	strconv "strconv"
	capnp "github.com/iguazio/go-capnproto2"
	text "github.com/iguazio/go-capnproto2/encoding/text"
	schemas "github.com/iguazio/go-capnproto2/schemas"
)

type JsonValue struct{ capnp.Struct }
type JsonValue_Which uint16

const (
	JsonValue_Which_null    JsonValue_Which = 0
	JsonValue_Which_boolean JsonValue_Which = 1
	JsonValue_Which_number  JsonValue_Which = 2
	JsonValue_Which_string_ JsonValue_Which = 3
	JsonValue_Which_array   JsonValue_Which = 4
	JsonValue_Which_object  JsonValue_Which = 5
	JsonValue_Which_call    JsonValue_Which = 6
)

func (w JsonValue_Which) String() string {
	const s = "nullbooleannumberstring_arrayobjectcall"
	switch w {
	case JsonValue_Which_null:
		return s[0:4]
	case JsonValue_Which_boolean:
		return s[4:11]
	case JsonValue_Which_number:
		return s[11:17]
	case JsonValue_Which_string_:
		return s[17:24]
	case JsonValue_Which_array:
		return s[24:29]
	case JsonValue_Which_object:
		return s[29:35]
	case JsonValue_Which_call:
		return s[35:39]

	}
	return "JsonValue_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// JsonValue_TypeID is the unique identifier for the type JsonValue.
const JsonValue_TypeID = 0x8825ffaa852cda72

func NewJsonValue(s *capnp.Segment) (JsonValue, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1})
	return JsonValue{st}, err
}

func NewRootJsonValue(s *capnp.Segment) (JsonValue, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1})
	return JsonValue{st}, err
}

func ReadRootJsonValue(msg *capnp.Message) (JsonValue, error) {
	root, err := msg.RootPtr()
	return JsonValue{root.Struct()}, err
}

func (s JsonValue) String() string {
	str, _ := text.Marshal(0x8825ffaa852cda72, s.Struct)
	return str
}

func (s JsonValue) Which() JsonValue_Which {
	return JsonValue_Which(s.Struct.Uint16(0))
}
func (s JsonValue) SetNull() {
	s.Struct.SetUint16(0, 0)

}

func (s JsonValue) Boolean() bool {
	if s.Struct.Uint16(0) != 1 {
		panic("Which() != boolean")
	}
	return s.Struct.Bit(16)
}

func (s JsonValue) SetBoolean(v bool) {
	s.Struct.SetUint16(0, 1)
	s.Struct.SetBit(16, v)
}

func (s JsonValue) Number() float64 {
	if s.Struct.Uint16(0) != 2 {
		panic("Which() != number")
	}
	return math.Float64frombits(s.Struct.Uint64(8))
}

func (s JsonValue) SetNumber(v float64) {
	s.Struct.SetUint16(0, 2)
	s.Struct.SetUint64(8, math.Float64bits(v))
}

func (s JsonValue) String_() (string, error) {
	if s.Struct.Uint16(0) != 3 {
		panic("Which() != string_")
	}
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s JsonValue) HasString_() bool {
	if s.Struct.Uint16(0) != 3 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s JsonValue) String_Bytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s JsonValue) SetString_(v string) error {
	s.Struct.SetUint16(0, 3)
	return s.Struct.SetText(0, v)
}

func (s JsonValue) Array() (JsonValue_List, error) {
	if s.Struct.Uint16(0) != 4 {
		panic("Which() != array")
	}
	p, err := s.Struct.Ptr(0)
	return JsonValue_List{List: p.List()}, err
}

func (s JsonValue) HasArray() bool {
	if s.Struct.Uint16(0) != 4 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s JsonValue) SetArray(v JsonValue_List) error {
	s.Struct.SetUint16(0, 4)
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewArray sets the array field to a newly
// allocated JsonValue_List, preferring placement in s's segment.
func (s JsonValue) NewArray(n int32) (JsonValue_List, error) {
	s.Struct.SetUint16(0, 4)
	l, err := NewJsonValue_List(s.Struct.Segment(), n)
	if err != nil {
		return JsonValue_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s JsonValue) Object() (JsonValue_Field_List, error) {
	if s.Struct.Uint16(0) != 5 {
		panic("Which() != object")
	}
	p, err := s.Struct.Ptr(0)
	return JsonValue_Field_List{List: p.List()}, err
}

func (s JsonValue) HasObject() bool {
	if s.Struct.Uint16(0) != 5 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s JsonValue) SetObject(v JsonValue_Field_List) error {
	s.Struct.SetUint16(0, 5)
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewObject sets the object field to a newly
// allocated JsonValue_Field_List, preferring placement in s's segment.
func (s JsonValue) NewObject(n int32) (JsonValue_Field_List, error) {
	s.Struct.SetUint16(0, 5)
	l, err := NewJsonValue_Field_List(s.Struct.Segment(), n)
	if err != nil {
		return JsonValue_Field_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s JsonValue) Call() (JsonValue_Call, error) {
	if s.Struct.Uint16(0) != 6 {
		panic("Which() != call")
	}
	p, err := s.Struct.Ptr(0)
	return JsonValue_Call{Struct: p.Struct()}, err
}

func (s JsonValue) HasCall() bool {
	if s.Struct.Uint16(0) != 6 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s JsonValue) SetCall(v JsonValue_Call) error {
	s.Struct.SetUint16(0, 6)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewCall sets the call field to a newly
// allocated JsonValue_Call struct, preferring placement in s's segment.
func (s JsonValue) NewCall() (JsonValue_Call, error) {
	s.Struct.SetUint16(0, 6)
	ss, err := NewJsonValue_Call(s.Struct.Segment())
	if err != nil {
		return JsonValue_Call{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// JsonValue_List is a list of JsonValue.
type JsonValue_List struct{ capnp.List }

// NewJsonValue creates a new list of JsonValue.
func NewJsonValue_List(s *capnp.Segment, sz int32) (JsonValue_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1}, sz)
	return JsonValue_List{l}, err
}

func (s JsonValue_List) At(i int) JsonValue { return JsonValue{s.List.Struct(i)} }

func (s JsonValue_List) Set(i int, v JsonValue) error { return s.List.SetStruct(i, v.Struct) }

func (s JsonValue_List) String() string {
	str, _ := text.MarshalList(0x8825ffaa852cda72, s.List)
	return str
}

// JsonValue_Promise is a wrapper for a JsonValue promised by a client call.
type JsonValue_Promise struct{ *capnp.Pipeline }

func (p JsonValue_Promise) Struct() (JsonValue, error) {
	s, err := p.Pipeline.Struct()
	return JsonValue{s}, err
}

func (p JsonValue_Promise) Call() JsonValue_Call_Promise {
	return JsonValue_Call_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type JsonValue_Field struct{ capnp.Struct }

// JsonValue_Field_TypeID is the unique identifier for the type JsonValue_Field.
const JsonValue_Field_TypeID = 0xc27855d853a937cc

func NewJsonValue_Field(s *capnp.Segment) (JsonValue_Field, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return JsonValue_Field{st}, err
}

func NewRootJsonValue_Field(s *capnp.Segment) (JsonValue_Field, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return JsonValue_Field{st}, err
}

func ReadRootJsonValue_Field(msg *capnp.Message) (JsonValue_Field, error) {
	root, err := msg.RootPtr()
	return JsonValue_Field{root.Struct()}, err
}

func (s JsonValue_Field) String() string {
	str, _ := text.Marshal(0xc27855d853a937cc, s.Struct)
	return str
}

func (s JsonValue_Field) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s JsonValue_Field) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s JsonValue_Field) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s JsonValue_Field) SetName(v string) error {
	return s.Struct.SetText(0, v)
}

func (s JsonValue_Field) Value() (JsonValue, error) {
	p, err := s.Struct.Ptr(1)
	return JsonValue{Struct: p.Struct()}, err
}

func (s JsonValue_Field) HasValue() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s JsonValue_Field) SetValue(v JsonValue) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewValue sets the value field to a newly
// allocated JsonValue struct, preferring placement in s's segment.
func (s JsonValue_Field) NewValue() (JsonValue, error) {
	ss, err := NewJsonValue(s.Struct.Segment())
	if err != nil {
		return JsonValue{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

// JsonValue_Field_List is a list of JsonValue_Field.
type JsonValue_Field_List struct{ capnp.List }

// NewJsonValue_Field creates a new list of JsonValue_Field.
func NewJsonValue_Field_List(s *capnp.Segment, sz int32) (JsonValue_Field_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return JsonValue_Field_List{l}, err
}

func (s JsonValue_Field_List) At(i int) JsonValue_Field { return JsonValue_Field{s.List.Struct(i)} }

func (s JsonValue_Field_List) Set(i int, v JsonValue_Field) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s JsonValue_Field_List) String() string {
	str, _ := text.MarshalList(0xc27855d853a937cc, s.List)
	return str
}

// JsonValue_Field_Promise is a wrapper for a JsonValue_Field promised by a client call.
type JsonValue_Field_Promise struct{ *capnp.Pipeline }

func (p JsonValue_Field_Promise) Struct() (JsonValue_Field, error) {
	s, err := p.Pipeline.Struct()
	return JsonValue_Field{s}, err
}

func (p JsonValue_Field_Promise) Value() JsonValue_Promise {
	return JsonValue_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type JsonValue_Call struct{ capnp.Struct }

// JsonValue_Call_TypeID is the unique identifier for the type JsonValue_Call.
const JsonValue_Call_TypeID = 0x9bbf84153dd4bb60

func NewJsonValue_Call(s *capnp.Segment) (JsonValue_Call, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return JsonValue_Call{st}, err
}

func NewRootJsonValue_Call(s *capnp.Segment) (JsonValue_Call, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return JsonValue_Call{st}, err
}

func ReadRootJsonValue_Call(msg *capnp.Message) (JsonValue_Call, error) {
	root, err := msg.RootPtr()
	return JsonValue_Call{root.Struct()}, err
}

func (s JsonValue_Call) String() string {
	str, _ := text.Marshal(0x9bbf84153dd4bb60, s.Struct)
	return str
}

func (s JsonValue_Call) Function() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s JsonValue_Call) HasFunction() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s JsonValue_Call) FunctionBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s JsonValue_Call) SetFunction(v string) error {
	return s.Struct.SetText(0, v)
}

func (s JsonValue_Call) Params() (JsonValue_List, error) {
	p, err := s.Struct.Ptr(1)
	return JsonValue_List{List: p.List()}, err
}

func (s JsonValue_Call) HasParams() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s JsonValue_Call) SetParams(v JsonValue_List) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewParams sets the params field to a newly
// allocated JsonValue_List, preferring placement in s's segment.
func (s JsonValue_Call) NewParams(n int32) (JsonValue_List, error) {
	l, err := NewJsonValue_List(s.Struct.Segment(), n)
	if err != nil {
		return JsonValue_List{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

// JsonValue_Call_List is a list of JsonValue_Call.
type JsonValue_Call_List struct{ capnp.List }

// NewJsonValue_Call creates a new list of JsonValue_Call.
func NewJsonValue_Call_List(s *capnp.Segment, sz int32) (JsonValue_Call_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return JsonValue_Call_List{l}, err
}

func (s JsonValue_Call_List) At(i int) JsonValue_Call { return JsonValue_Call{s.List.Struct(i)} }

func (s JsonValue_Call_List) Set(i int, v JsonValue_Call) error { return s.List.SetStruct(i, v.Struct) }

func (s JsonValue_Call_List) String() string {
	str, _ := text.MarshalList(0x9bbf84153dd4bb60, s.List)
	return str
}

// JsonValue_Call_Promise is a wrapper for a JsonValue_Call promised by a client call.
type JsonValue_Call_Promise struct{ *capnp.Pipeline }

func (p JsonValue_Call_Promise) Struct() (JsonValue_Call, error) {
	s, err := p.Pipeline.Struct()
	return JsonValue_Call{s}, err
}

const schema_8ef99297a43a5e34 = "x\xdat\x92AHTQ\x14\x86\xff\xff\xde7\xa3\xe2" +
	"L\xf3^\xf3\xa4\x16\x89\x9b\xa2\x1235!\x18\x88)" +
	"-\x09\x17\xe1m\xa8e\xf9\x1c_1r}O\xde8" +
	"Y+7\x05\xb5\xa8(\x88\x16\xad\x826\xbajQ\x90" +
	"\x94\x94\xd2\xb2U\xb4\x08W-\xda\x04m\xdaee7" +
	"n\x92O$w\x87s\xce=\xff\xcf\xf7\xdf\x9e\x93<" +
	".z33\x12P\x072Y\x93\xact\xdd\x987\xfb" +
	"nB\xb5R\x98\xfe\x0b\xa5'\x0f\xef\xaf\xde\xc1)6" +
	"5\x01\xc5\xc7\x9c/\xceq?pd\x91w\x09\x9a\xd1" +
	"\x97\x1f\x8e\xb5]\x7f\xfd\x08^\x1b\xd3\xb7\x19a\x97o" +
	"\xc9\xf7\xc5\x07\xd2V\xf7\xe4\x0ch\xde\x1d\x9d\xab|<" +
	"wu\xf9\x7f\xbbtV\x8ay\xc7V-\xce\x0c\xce\x98" +
	"\x89z\x1cuW\x83)FS\xa5\xe1z\x1c\x9d/\x04" +
	"\xba\x11\xaafn>\xd3\xd2\xb7I?\xd3\xd91T\x0b" +
	"\xf5xa0\xd0Z\xed\x91N\xce\x18\x87\x80\xf7\xbc\x13" +
	"PO%\xd5+\xc1v\xfe6\xaeO\xdb^\x18\x00\xd4" +
	"3I\xf5F\xb0]\xac\x19\xfa\x14\x80\xb7X\x02\xd4\x0b" +
	"I\xf5V0/\x7f\x19\x9f\x12\xf0\x96J\xdeR\x87\xfa" +
	"$\xa9\xbe\x0a\xe6\x9d\x9f\xc6\xa7\x03x_\xfa\x00\xf5Y" +
	"\xf2,\x05\xf3\x99\x1f\xc6g\x06\xf0\xd6\xec\x89\xef\x92\x15" +
	"\xdf\xb6\xb3\xab\xc6g\x16(z\xec\x04*9JVv" +
	"S\xb0\x105\xb4Fvv,\x8eu\x18D$\x04\x09" +
	"\x96\xa3\xc6\xe4X\x98\xb0\x15\x82\xad`\xb9>\x9d\xd4\xa2" +
	"\xcb\xca\xa10\xdfn\x1f\xde\xb5sta\x19\xca\x11<" +
	"\xe1\x929\xc0\xe3\xc0\xec\xfa\xcaE\x809\x08\xe6\xc0\x8e" +
	" I\x82k\xdc\x01\x8eH\xd2MY\x83\xb6Y\x8e\xc7" +
	"&\xc2\xeat:\xdf \xba>/T\x03\xad\xe9\xa6l" +
	"A\xba\xe0F&\xe2_&6\x92\xee\xc1@S\x8f\x90" +
	"\xaaY:\xc0_\xe2\x07\x87\xedg\x92T\xfd\x82\x1e\xb9" +
	"\xce\xbb\xd7R\xe9\x92T\xa7\x05\xcd\xa5FT\x9d\xae\xc5" +
	"\x11R\xd3\xe5\xa9 \x09&\xeb\xdb\xba\xdeF~\xa8\x16" +
	"J=\xbeE\xdf&\xbeWR\xf5l\xd2?\xd4\x97\x9a" +
	"*D\xc1d\xb8A\xeb\x8a=\xb4E\xd0\x05\xff\x04\x00" +
	"\x00\xff\xff\x9f\xdb\xc7\x93"

func init() {
	schemas.Register(schema_8ef99297a43a5e34,
		0x8825ffaa852cda72,
		0x9bbf84153dd4bb60,
		0xc27855d853a937cc)
}
