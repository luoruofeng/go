// Code generated by protoc-gen-go. DO NOT EDIT.
// source: addr.proto

package tutorial

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Person_PhotoType int32

const (
	Person_MOBILE Person_PhotoType = 0
	Person_HOME   Person_PhotoType = 1
	Person_WORK   Person_PhotoType = 2
)

var Person_PhotoType_name = map[int32]string{
	0: "MOBILE",
	1: "HOME",
	2: "WORK",
}

var Person_PhotoType_value = map[string]int32{
	"MOBILE": 0,
	"HOME":   1,
	"WORK":   2,
}

func (x Person_PhotoType) String() string {
	return proto.EnumName(Person_PhotoType_name, int32(x))
}

func (Person_PhotoType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eaa67da8ec6ef8e1, []int{0, 0}
}

//每个元素上的“ = 1”，“ = 2”标记标识该字段在二进制编码中使用的唯一“标记”。标签编号1至15与较高的编号相比，编码所需的字节减少了一个字节，因此，为了进行优化，您可以决定将这些标签用于常用或重复的元素，而将标签16和更高的标签用于较少使用的可选元素。
type Person struct {
	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id    int32  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	//如果一个字段为repeated，则该字段可以重复任意次（包括零次）。重复值的顺序将保留在协议缓冲区中。将重复字段视为动态大小的数组。
	Phones               []*Person_PhoneNumber `protobuf:"bytes,4,rep,name=phones,proto3" json:"phones,omitempty"`
	LastUpdated          *timestamp.Timestamp  `protobuf:"bytes,5,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_eaa67da8ec6ef8e1, []int{0}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Person) GetPhones() []*Person_PhoneNumber {
	if m != nil {
		return m.Phones
	}
	return nil
}

func (m *Person) GetLastUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

type Person_PhoneNumber struct {
	Number               string           `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Type                 Person_PhotoType `protobuf:"varint,2,opt,name=type,proto3,enum=tutorial.Person_PhotoType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Person_PhoneNumber) Reset()         { *m = Person_PhoneNumber{} }
func (m *Person_PhoneNumber) String() string { return proto.CompactTextString(m) }
func (*Person_PhoneNumber) ProtoMessage()    {}
func (*Person_PhoneNumber) Descriptor() ([]byte, []int) {
	return fileDescriptor_eaa67da8ec6ef8e1, []int{0, 0}
}

func (m *Person_PhoneNumber) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person_PhoneNumber.Unmarshal(m, b)
}
func (m *Person_PhoneNumber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person_PhoneNumber.Marshal(b, m, deterministic)
}
func (m *Person_PhoneNumber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person_PhoneNumber.Merge(m, src)
}
func (m *Person_PhoneNumber) XXX_Size() int {
	return xxx_messageInfo_Person_PhoneNumber.Size(m)
}
func (m *Person_PhoneNumber) XXX_DiscardUnknown() {
	xxx_messageInfo_Person_PhoneNumber.DiscardUnknown(m)
}

var xxx_messageInfo_Person_PhoneNumber proto.InternalMessageInfo

func (m *Person_PhoneNumber) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *Person_PhoneNumber) GetType() Person_PhotoType {
	if m != nil {
		return m.Type
	}
	return Person_MOBILE
}

// Our address book file is just one of these.
type AddressBook struct {
	People               []*Person `protobuf:"bytes,1,rep,name=people,proto3" json:"people,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AddressBook) Reset()         { *m = AddressBook{} }
func (m *AddressBook) String() string { return proto.CompactTextString(m) }
func (*AddressBook) ProtoMessage()    {}
func (*AddressBook) Descriptor() ([]byte, []int) {
	return fileDescriptor_eaa67da8ec6ef8e1, []int{1}
}

func (m *AddressBook) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressBook.Unmarshal(m, b)
}
func (m *AddressBook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressBook.Marshal(b, m, deterministic)
}
func (m *AddressBook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressBook.Merge(m, src)
}
func (m *AddressBook) XXX_Size() int {
	return xxx_messageInfo_AddressBook.Size(m)
}
func (m *AddressBook) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressBook.DiscardUnknown(m)
}

var xxx_messageInfo_AddressBook proto.InternalMessageInfo

func (m *AddressBook) GetPeople() []*Person {
	if m != nil {
		return m.People
	}
	return nil
}

func init() {
	proto.RegisterEnum("tutorial.Person_PhotoType", Person_PhotoType_name, Person_PhotoType_value)
	proto.RegisterType((*Person)(nil), "tutorial.Person")
	proto.RegisterType((*Person_PhoneNumber)(nil), "tutorial.Person.PhoneNumber")
	proto.RegisterType((*AddressBook)(nil), "tutorial.AddressBook")
}

func init() { proto.RegisterFile("addr.proto", fileDescriptor_eaa67da8ec6ef8e1) }

var fileDescriptor_eaa67da8ec6ef8e1 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x51, 0x4b, 0xc3, 0x30,
	0x10, 0xc7, 0x6d, 0xd7, 0x95, 0xed, 0x2a, 0xa3, 0x04, 0x91, 0x52, 0x04, 0xcb, 0x9e, 0x0a, 0x42,
	0x06, 0x53, 0xf0, 0xc9, 0x07, 0x07, 0x03, 0x45, 0xe7, 0x46, 0xd8, 0xf0, 0x51, 0x32, 0x72, 0xce,
	0x62, 0xdb, 0x84, 0x26, 0x7d, 0xd8, 0x67, 0xf3, 0xcb, 0x49, 0xd3, 0x56, 0x44, 0x7c, 0xfb, 0xe7,
	0xee, 0xc7, 0xe5, 0x77, 0x07, 0xc0, 0x85, 0xa8, 0xa8, 0xaa, 0xa4, 0x91, 0x64, 0x64, 0x6a, 0x23,
	0xab, 0x8c, 0xe7, 0xf1, 0xe5, 0x41, 0xca, 0x43, 0x8e, 0x33, 0x5b, 0xdf, 0xd7, 0xef, 0x33, 0x93,
	0x15, 0xa8, 0x0d, 0x2f, 0x54, 0x8b, 0x4e, 0xbf, 0x5c, 0xf0, 0x37, 0x58, 0x69, 0x59, 0x12, 0x02,
	0x5e, 0xc9, 0x0b, 0x8c, 0x9c, 0xc4, 0x49, 0xc7, 0xcc, 0x66, 0x32, 0x01, 0x37, 0x13, 0x91, 0x9b,
	0x38, 0xe9, 0x90, 0xb9, 0x99, 0x20, 0x67, 0x30, 0xc4, 0x82, 0x67, 0x79, 0x34, 0xb0, 0x50, 0xfb,
	0x20, 0x37, 0xe0, 0xab, 0x0f, 0x59, 0xa2, 0x8e, 0xbc, 0x64, 0x90, 0x06, 0xf3, 0x0b, 0xda, 0x0b,
	0xd0, 0x76, 0x36, 0xdd, 0x34, 0xed, 0x97, 0xba, 0xd8, 0x63, 0xc5, 0x3a, 0x96, 0xdc, 0xc1, 0x69,
	0xce, 0xb5, 0x79, 0xab, 0x95, 0xe0, 0x06, 0x45, 0x34, 0x4c, 0x9c, 0x34, 0x98, 0xc7, 0xb4, 0x55,
	0xa6, 0xbd, 0x32, 0xdd, 0xf6, 0xca, 0x2c, 0x68, 0xf8, 0x5d, 0x8b, 0xc7, 0x3b, 0x08, 0x7e, 0x4d,
	0x25, 0xe7, 0xe0, 0x97, 0x36, 0x75, 0xfe, 0xdd, 0x8b, 0x50, 0xf0, 0xcc, 0x51, 0xa1, 0xdd, 0x61,
	0x32, 0x8f, 0xff, 0x33, 0x33, 0x72, 0x7b, 0x54, 0xc8, 0x2c, 0x37, 0xbd, 0x82, 0xf1, 0x4f, 0x89,
	0x00, 0xf8, 0xab, 0xf5, 0xe2, 0xf1, 0x79, 0x19, 0x9e, 0x90, 0x11, 0x78, 0x0f, 0xeb, 0xd5, 0x32,
	0x74, 0x9a, 0xf4, 0xba, 0x66, 0x4f, 0xa1, 0x3b, 0xbd, 0x85, 0xe0, 0x5e, 0x88, 0x0a, 0xb5, 0x5e,
	0x48, 0xf9, 0x49, 0x52, 0xf0, 0x15, 0x4a, 0x95, 0x37, 0x37, 0x6c, 0xee, 0x10, 0xfe, 0xfd, 0x8d,
	0x75, 0xfd, 0xbd, 0x6f, 0xb7, 0xbb, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x9c, 0xb8, 0xa9, 0xf1,
	0xb6, 0x01, 0x00, 0x00,
}
