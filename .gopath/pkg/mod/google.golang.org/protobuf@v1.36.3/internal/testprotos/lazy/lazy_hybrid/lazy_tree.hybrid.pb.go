// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/testprotos/lazy/lazy_hybrid/lazy_tree.hybrid.proto

//go:build !protoopaque

package lazy_hybrid

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/gofeaturespb"
	reflect "reflect"
)

type Node struct {
	state         protoimpl.MessageState `protogen:"hybrid.v1"`
	Nested        *Node                  `protobuf:"bytes,99,opt,name=nested" json:"nested,omitempty"`
	Int32         *int32                 `protobuf:"varint,1,opt,name=int32" json:"int32,omitempty"`
	Int64         *int64                 `protobuf:"varint,2,opt,name=int64" json:"int64,omitempty"`
	Uint32        *uint32                `protobuf:"varint,3,opt,name=uint32" json:"uint32,omitempty"`
	Uint64        *uint64                `protobuf:"varint,4,opt,name=uint64" json:"uint64,omitempty"`
	Sint32        *int32                 `protobuf:"zigzag32,5,opt,name=sint32" json:"sint32,omitempty"`
	Sint64        *int64                 `protobuf:"zigzag64,6,opt,name=sint64" json:"sint64,omitempty"`
	Fixed32       *uint32                `protobuf:"fixed32,7,opt,name=fixed32" json:"fixed32,omitempty"`
	Fixed64       *uint64                `protobuf:"fixed64,8,opt,name=fixed64" json:"fixed64,omitempty"`
	Sfixed32      *int32                 `protobuf:"fixed32,9,opt,name=sfixed32" json:"sfixed32,omitempty"`
	Sfixed64      *int64                 `protobuf:"fixed64,10,opt,name=sfixed64" json:"sfixed64,omitempty"`
	Float         *float32               `protobuf:"fixed32,11,opt,name=float" json:"float,omitempty"`
	Double        *float64               `protobuf:"fixed64,12,opt,name=double" json:"double,omitempty"`
	Bool          *bool                  `protobuf:"varint,13,opt,name=bool" json:"bool,omitempty"`
	String_       *string                `protobuf:"bytes,14,opt,name=string" json:"string,omitempty"`
	Bytes         []byte                 `protobuf:"bytes,15,opt,name=bytes" json:"bytes,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Node) Reset() {
	*x = Node{}
	mi := &file_internal_testprotos_lazy_lazy_hybrid_lazy_tree_hybrid_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_lazy_lazy_hybrid_lazy_tree_hybrid_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Node) GetNested() *Node {
	if x != nil {
		return x.Nested
	}
	return nil
}

func (x *Node) GetInt32() int32 {
	if x != nil && x.Int32 != nil {
		return *x.Int32
	}
	return 0
}

func (x *Node) GetInt64() int64 {
	if x != nil && x.Int64 != nil {
		return *x.Int64
	}
	return 0
}

func (x *Node) GetUint32() uint32 {
	if x != nil && x.Uint32 != nil {
		return *x.Uint32
	}
	return 0
}

func (x *Node) GetUint64() uint64 {
	if x != nil && x.Uint64 != nil {
		return *x.Uint64
	}
	return 0
}

func (x *Node) GetSint32() int32 {
	if x != nil && x.Sint32 != nil {
		return *x.Sint32
	}
	return 0
}

func (x *Node) GetSint64() int64 {
	if x != nil && x.Sint64 != nil {
		return *x.Sint64
	}
	return 0
}

func (x *Node) GetFixed32() uint32 {
	if x != nil && x.Fixed32 != nil {
		return *x.Fixed32
	}
	return 0
}

func (x *Node) GetFixed64() uint64 {
	if x != nil && x.Fixed64 != nil {
		return *x.Fixed64
	}
	return 0
}

func (x *Node) GetSfixed32() int32 {
	if x != nil && x.Sfixed32 != nil {
		return *x.Sfixed32
	}
	return 0
}

func (x *Node) GetSfixed64() int64 {
	if x != nil && x.Sfixed64 != nil {
		return *x.Sfixed64
	}
	return 0
}

func (x *Node) GetFloat() float32 {
	if x != nil && x.Float != nil {
		return *x.Float
	}
	return 0
}

func (x *Node) GetDouble() float64 {
	if x != nil && x.Double != nil {
		return *x.Double
	}
	return 0
}

func (x *Node) GetBool() bool {
	if x != nil && x.Bool != nil {
		return *x.Bool
	}
	return false
}

func (x *Node) GetString() string {
	if x != nil && x.String_ != nil {
		return *x.String_
	}
	return ""
}

// Deprecated: Use GetString instead.
func (x *Node) GetString_() string {
	return x.GetString()
}

func (x *Node) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

func (x *Node) SetNested(v *Node) {
	x.Nested = v
}

func (x *Node) SetInt32(v int32) {
	x.Int32 = &v
}

func (x *Node) SetInt64(v int64) {
	x.Int64 = &v
}

func (x *Node) SetUint32(v uint32) {
	x.Uint32 = &v
}

func (x *Node) SetUint64(v uint64) {
	x.Uint64 = &v
}

func (x *Node) SetSint32(v int32) {
	x.Sint32 = &v
}

func (x *Node) SetSint64(v int64) {
	x.Sint64 = &v
}

func (x *Node) SetFixed32(v uint32) {
	x.Fixed32 = &v
}

func (x *Node) SetFixed64(v uint64) {
	x.Fixed64 = &v
}

func (x *Node) SetSfixed32(v int32) {
	x.Sfixed32 = &v
}

func (x *Node) SetSfixed64(v int64) {
	x.Sfixed64 = &v
}

func (x *Node) SetFloat(v float32) {
	x.Float = &v
}

func (x *Node) SetDouble(v float64) {
	x.Double = &v
}

func (x *Node) SetBool(v bool) {
	x.Bool = &v
}

func (x *Node) SetString(v string) {
	x.String_ = &v
}

func (x *Node) SetBytes(v []byte) {
	if v == nil {
		v = []byte{}
	}
	x.Bytes = v
}

func (x *Node) HasNested() bool {
	if x == nil {
		return false
	}
	return x.Nested != nil
}

func (x *Node) HasInt32() bool {
	if x == nil {
		return false
	}
	return x.Int32 != nil
}

func (x *Node) HasInt64() bool {
	if x == nil {
		return false
	}
	return x.Int64 != nil
}

func (x *Node) HasUint32() bool {
	if x == nil {
		return false
	}
	return x.Uint32 != nil
}

func (x *Node) HasUint64() bool {
	if x == nil {
		return false
	}
	return x.Uint64 != nil
}

func (x *Node) HasSint32() bool {
	if x == nil {
		return false
	}
	return x.Sint32 != nil
}

func (x *Node) HasSint64() bool {
	if x == nil {
		return false
	}
	return x.Sint64 != nil
}

func (x *Node) HasFixed32() bool {
	if x == nil {
		return false
	}
	return x.Fixed32 != nil
}

func (x *Node) HasFixed64() bool {
	if x == nil {
		return false
	}
	return x.Fixed64 != nil
}

func (x *Node) HasSfixed32() bool {
	if x == nil {
		return false
	}
	return x.Sfixed32 != nil
}

func (x *Node) HasSfixed64() bool {
	if x == nil {
		return false
	}
	return x.Sfixed64 != nil
}

func (x *Node) HasFloat() bool {
	if x == nil {
		return false
	}
	return x.Float != nil
}

func (x *Node) HasDouble() bool {
	if x == nil {
		return false
	}
	return x.Double != nil
}

func (x *Node) HasBool() bool {
	if x == nil {
		return false
	}
	return x.Bool != nil
}

func (x *Node) HasString() bool {
	if x == nil {
		return false
	}
	return x.String_ != nil
}

func (x *Node) HasBytes() bool {
	if x == nil {
		return false
	}
	return x.Bytes != nil
}

func (x *Node) ClearNested() {
	x.Nested = nil
}

func (x *Node) ClearInt32() {
	x.Int32 = nil
}

func (x *Node) ClearInt64() {
	x.Int64 = nil
}

func (x *Node) ClearUint32() {
	x.Uint32 = nil
}

func (x *Node) ClearUint64() {
	x.Uint64 = nil
}

func (x *Node) ClearSint32() {
	x.Sint32 = nil
}

func (x *Node) ClearSint64() {
	x.Sint64 = nil
}

func (x *Node) ClearFixed32() {
	x.Fixed32 = nil
}

func (x *Node) ClearFixed64() {
	x.Fixed64 = nil
}

func (x *Node) ClearSfixed32() {
	x.Sfixed32 = nil
}

func (x *Node) ClearSfixed64() {
	x.Sfixed64 = nil
}

func (x *Node) ClearFloat() {
	x.Float = nil
}

func (x *Node) ClearDouble() {
	x.Double = nil
}

func (x *Node) ClearBool() {
	x.Bool = nil
}

func (x *Node) ClearString() {
	x.String_ = nil
}

func (x *Node) ClearBytes() {
	x.Bytes = nil
}

type Node_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Nested   *Node
	Int32    *int32
	Int64    *int64
	Uint32   *uint32
	Uint64   *uint64
	Sint32   *int32
	Sint64   *int64
	Fixed32  *uint32
	Fixed64  *uint64
	Sfixed32 *int32
	Sfixed64 *int64
	Float    *float32
	Double   *float64
	Bool     *bool
	String   *string
	Bytes    []byte
}

func (b0 Node_builder) Build() *Node {
	m0 := &Node{}
	b, x := &b0, m0
	_, _ = b, x
	x.Nested = b.Nested
	x.Int32 = b.Int32
	x.Int64 = b.Int64
	x.Uint32 = b.Uint32
	x.Uint64 = b.Uint64
	x.Sint32 = b.Sint32
	x.Sint64 = b.Sint64
	x.Fixed32 = b.Fixed32
	x.Fixed64 = b.Fixed64
	x.Sfixed32 = b.Sfixed32
	x.Sfixed64 = b.Sfixed64
	x.Float = b.Float
	x.Double = b.Double
	x.Bool = b.Bool
	x.String_ = b.String
	x.Bytes = b.Bytes
	return m0
}

var File_internal_testprotos_lazy_lazy_hybrid_lazy_tree_hybrid_proto protoreflect.FileDescriptor

var file_internal_testprotos_lazy_lazy_hybrid_lazy_tree_hybrid_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6c, 0x61, 0x7a, 0x79, 0x2f, 0x6c, 0x61, 0x7a, 0x79, 0x5f,
	0x68, 0x79, 0x62, 0x72, 0x69, 0x64, 0x2f, 0x6c, 0x61, 0x7a, 0x79, 0x5f, 0x74, 0x72, 0x65, 0x65,
	0x2e, 0x68, 0x79, 0x62, 0x72, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x68,
	0x79, 0x62, 0x72, 0x69, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x5f, 0x74, 0x72, 0x65, 0x65, 0x1a,
	0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x67, 0x6f, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa2, 0x03, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x6e,
	0x65, 0x73, 0x74, 0x65, 0x64, 0x18, 0x63, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x68, 0x79,
	0x62, 0x72, 0x69, 0x64, 0x2e, 0x6c, 0x61, 0x7a, 0x79, 0x5f, 0x74, 0x72, 0x65, 0x65, 0x2e, 0x4e,
	0x6f, 0x64, 0x65, 0x42, 0x02, 0x28, 0x01, 0x52, 0x06, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x69, 0x6e,
	0x74, 0x33, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x05, 0x20, 0x01, 0x28, 0x11, 0x52, 0x06, 0x73, 0x69, 0x6e,
	0x74, 0x33, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x12, 0x52, 0x06, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x18, 0x0a, 0x07, 0x66,
	0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x18, 0x07, 0x20, 0x01, 0x28, 0x07, 0x52, 0x07, 0x66, 0x69,
	0x78, 0x65, 0x64, 0x33, 0x32, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x06, 0x52, 0x07, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x12,
	0x1a, 0x0a, 0x08, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0f, 0x52, 0x08, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x10, 0x52, 0x08, 0x73,
	0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x64,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6c, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x42, 0x49, 0x5a, 0x3f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6c, 0x61, 0x7a, 0x79, 0x2f, 0x6c,
	0x61, 0x7a, 0x79, 0x5f, 0x68, 0x79, 0x62, 0x72, 0x69, 0x64, 0x92, 0x03, 0x05, 0xd2, 0x3e, 0x02,
	0x10, 0x02, 0x62, 0x08, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x70, 0xe8, 0x07,
}

var file_internal_testprotos_lazy_lazy_hybrid_lazy_tree_hybrid_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internal_testprotos_lazy_lazy_hybrid_lazy_tree_hybrid_proto_goTypes = []any{
	(*Node)(nil), // 0: hybrid.lazy_tree.Node
}
var file_internal_testprotos_lazy_lazy_hybrid_lazy_tree_hybrid_proto_depIdxs = []int32{
	0, // 0: hybrid.lazy_tree.Node.nested:type_name -> hybrid.lazy_tree.Node
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_testprotos_lazy_lazy_hybrid_lazy_tree_hybrid_proto_init() }
func file_internal_testprotos_lazy_lazy_hybrid_lazy_tree_hybrid_proto_init() {
	if File_internal_testprotos_lazy_lazy_hybrid_lazy_tree_hybrid_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_testprotos_lazy_lazy_hybrid_lazy_tree_hybrid_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_testprotos_lazy_lazy_hybrid_lazy_tree_hybrid_proto_goTypes,
		DependencyIndexes: file_internal_testprotos_lazy_lazy_hybrid_lazy_tree_hybrid_proto_depIdxs,
		MessageInfos:      file_internal_testprotos_lazy_lazy_hybrid_lazy_tree_hybrid_proto_msgTypes,
	}.Build()
	File_internal_testprotos_lazy_lazy_hybrid_lazy_tree_hybrid_proto = out.File
	file_internal_testprotos_lazy_lazy_hybrid_lazy_tree_hybrid_proto_rawDesc = nil
	file_internal_testprotos_lazy_lazy_hybrid_lazy_tree_hybrid_proto_goTypes = nil
	file_internal_testprotos_lazy_lazy_hybrid_lazy_tree_hybrid_proto_depIdxs = nil
}
