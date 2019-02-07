// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helper/storagepacker/types.proto

package storagepacker

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Item represents a entry that gets inserted into the storage packer
type Item struct {
	// ID is the UUID to identify the item
	ID string `sentinel:"" protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// message is the contents of the item
	Message              *any.Any `sentinel:"" protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0e98c66c4f51b7f, []int{0}
}

func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Item) GetMessage() *any.Any {
	if m != nil {
		return m.Message
	}
	return nil
}

// Bucket is a construct to hold multiple items within itself. This
// abstraction contains multiple buckets of the same kind within itself and
// shares amont them the items that get inserted. When the bucket as a whole
// gets too big to hold more items, the contained buckets gets pushed out only
// to become independent buckets. Hence, this can grow infinitely in terms of
// storage space for items that get inserted.
type Bucket struct {
	// Key is the storage path where the bucket gets stored
	Key string `sentinel:"" protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Items holds the items contained within this bucket. Used by v1.
	Items []*Item `sentinel:"" protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	// ItemMap stores a mapping of item ID to message. Used by v2.
	ItemMap map[string]*any.Any `sentinel:"" protobuf:"bytes,3,rep,name=item_map,json=itemMap,proto3" json:"item_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Buckets are the buckets contained within this bucket
	Buckets              map[string]*Bucket `sentinel:"" protobuf:"bytes,4,rep,name=buckets,proto3" json:"buckets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Bucket) Reset()         { *m = Bucket{} }
func (m *Bucket) String() string { return proto.CompactTextString(m) }
func (*Bucket) ProtoMessage()    {}
func (*Bucket) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0e98c66c4f51b7f, []int{1}
}

func (m *Bucket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bucket.Unmarshal(m, b)
}
func (m *Bucket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bucket.Marshal(b, m, deterministic)
}
func (m *Bucket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bucket.Merge(m, src)
}
func (m *Bucket) XXX_Size() int {
	return xxx_messageInfo_Bucket.Size(m)
}
func (m *Bucket) XXX_DiscardUnknown() {
	xxx_messageInfo_Bucket.DiscardUnknown(m)
}

var xxx_messageInfo_Bucket proto.InternalMessageInfo

func (m *Bucket) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Bucket) GetItems() []*Item {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *Bucket) GetItemMap() map[string]*any.Any {
	if m != nil {
		return m.ItemMap
	}
	return nil
}

func (m *Bucket) GetBuckets() map[string]*Bucket {
	if m != nil {
		return m.Buckets
	}
	return nil
}

func init() {
	proto.RegisterType((*Item)(nil), "storagepacker.Item")
	proto.RegisterType((*Bucket)(nil), "storagepacker.Bucket")
	proto.RegisterMapType((map[string]*Bucket)(nil), "storagepacker.Bucket.BucketsEntry")
	proto.RegisterMapType((map[string]*any.Any)(nil), "storagepacker.Bucket.ItemMapEntry")
}

func init() { proto.RegisterFile("helper/storagepacker/types.proto", fileDescriptor_c0e98c66c4f51b7f) }

var fileDescriptor_c0e98c66c4f51b7f = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x4b, 0xfb, 0x30,
	0x18, 0xc6, 0x69, 0xf7, 0xeb, 0xfb, 0x7d, 0x37, 0x45, 0xa2, 0x42, 0xdd, 0xa9, 0xec, 0x34, 0x15,
	0x12, 0x9c, 0x17, 0x11, 0x3d, 0x38, 0x50, 0xf0, 0x20, 0x68, 0x8f, 0x5e, 0x24, 0xed, 0x5e, 0xdb,
	0xb0, 0x76, 0x09, 0x49, 0x3a, 0xe8, 0x5f, 0xec, 0xbf, 0x21, 0x6b, 0x36, 0x58, 0xa5, 0xec, 0xd4,
	0xb7, 0x3c, 0xcf, 0xf3, 0x79, 0x9f, 0x84, 0x40, 0x98, 0x61, 0xae, 0x50, 0x33, 0x63, 0xa5, 0xe6,
	0x29, 0x2a, 0x9e, 0x2c, 0x51, 0x33, 0x5b, 0x29, 0x34, 0x54, 0x69, 0x69, 0x25, 0x39, 0x6a, 0x48,
	0xe3, 0x8b, 0x54, 0xca, 0x34, 0x47, 0x56, 0x8b, 0x71, 0xf9, 0xcd, 0xf8, 0xaa, 0x72, 0xce, 0xc9,
	0x0b, 0x74, 0x5f, 0x2d, 0x16, 0xe4, 0x18, 0x7c, 0xb1, 0x08, 0xbc, 0xd0, 0x9b, 0xfe, 0x8f, 0x7c,
	0xb1, 0x20, 0x14, 0x06, 0x05, 0x1a, 0xc3, 0x53, 0x0c, 0xfc, 0xd0, 0x9b, 0x0e, 0x67, 0x67, 0xd4,
	0x41, 0xe8, 0x0e, 0x42, 0x9f, 0x56, 0x55, 0xb4, 0x33, 0x4d, 0x7e, 0x7c, 0xe8, 0xcf, 0xcb, 0x64,
	0x89, 0x96, 0x9c, 0x40, 0x67, 0x89, 0xd5, 0x96, 0xb5, 0x19, 0xc9, 0x25, 0xf4, 0x84, 0xc5, 0xc2,
	0x04, 0x7e, 0xd8, 0x99, 0x0e, 0x67, 0xa7, 0xb4, 0x51, 0x8f, 0x6e, 0x0a, 0x44, 0xce, 0x41, 0x1e,
	0xe1, 0xdf, 0x66, 0xf8, 0x2a, 0xb8, 0x0a, 0x3a, 0xb5, 0x7b, 0xf2, 0xc7, 0xed, 0xb6, 0xd4, 0xa1,
	0x37, 0xae, 0x9e, 0x57, 0x56, 0x57, 0xd1, 0x40, 0xb8, 0x3f, 0xf2, 0x00, 0x83, 0xb8, 0xd6, 0x4d,
	0xd0, 0x3d, 0x94, 0x76, 0x1f, 0xb3, 0x4d, 0x6f, 0x23, 0xe3, 0x77, 0x18, 0xed, 0x63, 0x5b, 0x4e,
	0x72, 0x05, 0xbd, 0x35, 0xcf, 0xcb, 0xc3, 0x97, 0xe2, 0x2c, 0xf7, 0xfe, 0x9d, 0x37, 0xfe, 0x80,
	0xd1, 0xfe, 0xaa, 0x16, 0xe2, 0x75, 0x93, 0x78, 0xde, 0xda, 0x77, 0x0f, 0x39, 0xbf, 0xf9, 0x64,
	0xa9, 0xb0, 0x59, 0x19, 0xd3, 0x44, 0x16, 0x2c, 0xe3, 0x26, 0x13, 0x89, 0xd4, 0x8a, 0xad, 0x79,
	0x99, 0x5b, 0xd6, 0xf6, 0x34, 0xe2, 0x7e, 0x5d, 0xef, 0xf6, 0x37, 0x00, 0x00, 0xff, 0xff, 0x7b,
	0x21, 0xba, 0x2f, 0x39, 0x02, 0x00, 0x00,
}
