// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/testprotos/messageset/msetextpb/msetextpb_hybrid/msetextpb.hybrid.proto

//go:build protoopaque

package msetextpb_hybrid

import (
	messagesetpb_hybrid "google.golang.org/protobuf/internal/testprotos/messageset/messagesetpb/messagesetpb_hybrid"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/gofeaturespb"
	reflect "reflect"
)

type Ext1 struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Ext1Field1  int32                  `protobuf:"varint,1,opt,name=ext1_field1,json=ext1Field1"`
	xxx_hidden_Ext1Field2  int32                  `protobuf:"varint,2,opt,name=ext1_field2,json=ext1Field2"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *Ext1) Reset() {
	*x = Ext1{}
	mi := &file_internal_testprotos_messageset_msetextpb_msetextpb_hybrid_msetextpb_hybrid_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Ext1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ext1) ProtoMessage() {}

func (x *Ext1) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_messageset_msetextpb_msetextpb_hybrid_msetextpb_hybrid_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Ext1) GetExt1Field1() int32 {
	if x != nil {
		return x.xxx_hidden_Ext1Field1
	}
	return 0
}

func (x *Ext1) GetExt1Field2() int32 {
	if x != nil {
		return x.xxx_hidden_Ext1Field2
	}
	return 0
}

func (x *Ext1) SetExt1Field1(v int32) {
	x.xxx_hidden_Ext1Field1 = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 2)
}

func (x *Ext1) SetExt1Field2(v int32) {
	x.xxx_hidden_Ext1Field2 = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 1, 2)
}

func (x *Ext1) HasExt1Field1() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *Ext1) HasExt1Field2() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 1)
}

func (x *Ext1) ClearExt1Field1() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_Ext1Field1 = 0
}

func (x *Ext1) ClearExt1Field2() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 1)
	x.xxx_hidden_Ext1Field2 = 0
}

type Ext1_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Ext1Field1 *int32
	Ext1Field2 *int32
}

func (b0 Ext1_builder) Build() *Ext1 {
	m0 := &Ext1{}
	b, x := &b0, m0
	_, _ = b, x
	if b.Ext1Field1 != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 2)
		x.xxx_hidden_Ext1Field1 = *b.Ext1Field1
	}
	if b.Ext1Field2 != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 1, 2)
		x.xxx_hidden_Ext1Field2 = *b.Ext1Field2
	}
	return m0
}

type Ext2 struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Ext2Field1  int32                  `protobuf:"varint,1,opt,name=ext2_field1,json=ext2Field1"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *Ext2) Reset() {
	*x = Ext2{}
	mi := &file_internal_testprotos_messageset_msetextpb_msetextpb_hybrid_msetextpb_hybrid_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Ext2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ext2) ProtoMessage() {}

func (x *Ext2) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_messageset_msetextpb_msetextpb_hybrid_msetextpb_hybrid_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Ext2) GetExt2Field1() int32 {
	if x != nil {
		return x.xxx_hidden_Ext2Field1
	}
	return 0
}

func (x *Ext2) SetExt2Field1(v int32) {
	x.xxx_hidden_Ext2Field1 = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 1)
}

func (x *Ext2) HasExt2Field1() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *Ext2) ClearExt2Field1() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_Ext2Field1 = 0
}

type Ext2_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Ext2Field1 *int32
}

func (b0 Ext2_builder) Build() *Ext2 {
	m0 := &Ext2{}
	b, x := &b0, m0
	_, _ = b, x
	if b.Ext2Field1 != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 1)
		x.xxx_hidden_Ext2Field1 = *b.Ext2Field1
	}
	return m0
}

type ExtRequired struct {
	state                     protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_RequiredField1 int32                  `protobuf:"varint,1,req,name=required_field1,json=requiredField1"`
	XXX_raceDetectHookData    protoimpl.RaceDetectHookData
	XXX_presence              [1]uint32
	unknownFields             protoimpl.UnknownFields
	sizeCache                 protoimpl.SizeCache
}

func (x *ExtRequired) Reset() {
	*x = ExtRequired{}
	mi := &file_internal_testprotos_messageset_msetextpb_msetextpb_hybrid_msetextpb_hybrid_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExtRequired) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtRequired) ProtoMessage() {}

func (x *ExtRequired) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_messageset_msetextpb_msetextpb_hybrid_msetextpb_hybrid_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *ExtRequired) GetRequiredField1() int32 {
	if x != nil {
		return x.xxx_hidden_RequiredField1
	}
	return 0
}

func (x *ExtRequired) SetRequiredField1(v int32) {
	x.xxx_hidden_RequiredField1 = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 1)
}

func (x *ExtRequired) HasRequiredField1() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *ExtRequired) ClearRequiredField1() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_RequiredField1 = 0
}

type ExtRequired_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	RequiredField1 *int32
}

func (b0 ExtRequired_builder) Build() *ExtRequired {
	m0 := &ExtRequired{}
	b, x := &b0, m0
	_, _ = b, x
	if b.RequiredField1 != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 1)
		x.xxx_hidden_RequiredField1 = *b.RequiredField1
	}
	return m0
}

type ExtLargeNumber struct {
	state         protoimpl.MessageState `protogen:"opaque.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExtLargeNumber) Reset() {
	*x = ExtLargeNumber{}
	mi := &file_internal_testprotos_messageset_msetextpb_msetextpb_hybrid_msetextpb_hybrid_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExtLargeNumber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtLargeNumber) ProtoMessage() {}

func (x *ExtLargeNumber) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_messageset_msetextpb_msetextpb_hybrid_msetextpb_hybrid_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

type ExtLargeNumber_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

}

func (b0 ExtLargeNumber_builder) Build() *ExtLargeNumber {
	m0 := &ExtLargeNumber{}
	b, x := &b0, m0
	_, _ = b, x
	return m0
}

var file_internal_testprotos_messageset_msetextpb_msetextpb_hybrid_msetextpb_hybrid_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*messagesetpb_hybrid.MessageSet)(nil),
		ExtensionType: (*Ext1)(nil),
		Field:         1000,
		Name:          "hybrid.goproto.proto.messageset.Ext1.message_set_ext1",
		Tag:           "bytes,1000,opt,name=message_set_ext1",
		Filename:      "internal/testprotos/messageset/msetextpb/msetextpb_hybrid/msetextpb.hybrid.proto",
	},
	{
		ExtendedType:  (*messagesetpb_hybrid.MessageSet)(nil),
		ExtensionType: (*Ext2)(nil),
		Field:         1001,
		Name:          "hybrid.goproto.proto.messageset.Ext2.message_set_ext2",
		Tag:           "bytes,1001,opt,name=message_set_ext2",
		Filename:      "internal/testprotos/messageset/msetextpb/msetextpb_hybrid/msetextpb.hybrid.proto",
	},
	{
		ExtendedType:  (*messagesetpb_hybrid.MessageSet)(nil),
		ExtensionType: (*ExtRequired)(nil),
		Field:         1002,
		Name:          "hybrid.goproto.proto.messageset.ExtRequired.message_set_extrequired",
		Tag:           "bytes,1002,opt,name=message_set_extrequired",
		Filename:      "internal/testprotos/messageset/msetextpb/msetextpb_hybrid/msetextpb.hybrid.proto",
	},
	{
		ExtendedType:  (*messagesetpb_hybrid.MessageSet)(nil),
		ExtensionType: (*ExtLargeNumber)(nil),
		Field:         536870912,
		Name:          "hybrid.goproto.proto.messageset.ExtLargeNumber.message_set_extlarge",
		Tag:           "bytes,536870912,opt,name=message_set_extlarge",
		Filename:      "internal/testprotos/messageset/msetextpb/msetextpb_hybrid/msetextpb.hybrid.proto",
	},
}

// Extension fields to messagesetpb_hybrid.MessageSet.
var (
	// optional hybrid.goproto.proto.messageset.Ext1 message_set_ext1 = 1000;
	E_Ext1_MessageSetExt1 = &file_internal_testprotos_messageset_msetextpb_msetextpb_hybrid_msetextpb_hybrid_proto_extTypes[0]
	// optional hybrid.goproto.proto.messageset.Ext2 message_set_ext2 = 1001;
	E_Ext2_MessageSetExt2 = &file_internal_testprotos_messageset_msetextpb_msetextpb_hybrid_msetextpb_hybrid_proto_extTypes[1]
	// optional hybrid.goproto.proto.messageset.ExtRequired message_set_extrequired = 1002;
	E_ExtRequired_MessageSetExtrequired = &file_internal_testprotos_messageset_msetextpb_msetextpb_hybrid_msetextpb_hybrid_proto_extTypes[2]
	// optional hybrid.goproto.proto.messageset.ExtLargeNumber message_set_extlarge = 536870912;
	E_ExtLargeNumber_MessageSetExtlarge = &file_internal_testprotos_messageset_msetextpb_msetextpb_hybrid_msetextpb_hybrid_proto_extTypes[3] // 1<<29
)

var File_internal_testprotos_messageset_msetextpb_msetextpb_hybrid_msetextpb_hybrid_proto protoreflect.FileDescriptor

var file_internal_testprotos_messageset_msetextpb_msetextpb_hybrid_msetextpb_hybrid_proto_rawDesc = []byte{
	0x0a, 0x50, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x65, 0x74,
	0x2f, 0x6d, 0x73, 0x65, 0x74, 0x65, 0x78, 0x74, 0x70, 0x62, 0x2f, 0x6d, 0x73, 0x65, 0x74, 0x65,
	0x78, 0x74, 0x70, 0x62, 0x5f, 0x68, 0x79, 0x62, 0x72, 0x69, 0x64, 0x2f, 0x6d, 0x73, 0x65, 0x74,
	0x65, 0x78, 0x74, 0x70, 0x62, 0x2e, 0x68, 0x79, 0x62, 0x72, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1f, 0x68, 0x79, 0x62, 0x72, 0x69, 0x64, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x65, 0x74, 0x1a, 0x58, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65,
	0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x65, 0x74, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x65, 0x74, 0x70, 0x62,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x65, 0x74, 0x70, 0x62, 0x5f, 0x68, 0x79,
	0x62, 0x72, 0x69, 0x64, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x74,
	0x2e, 0x68, 0x79, 0x62, 0x72, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67,
	0x6f, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xc7, 0x01, 0x0a, 0x04, 0x45, 0x78, 0x74, 0x31, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x74,
	0x31, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x65, 0x78, 0x74, 0x31, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78,
	0x74, 0x31, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x65, 0x78, 0x74, 0x31, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x32, 0x7d, 0x0a, 0x10, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x65, 0x78, 0x74, 0x31, 0x12,
	0x2b, 0x2e, 0x68, 0x79, 0x62, 0x72, 0x69, 0x64, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x65,
	0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x74, 0x18, 0xe8, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x68, 0x79, 0x62, 0x72, 0x69, 0x64, 0x2e, 0x67, 0x6f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x65, 0x74, 0x2e, 0x45, 0x78, 0x74, 0x31, 0x52, 0x0e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x53, 0x65, 0x74, 0x45, 0x78, 0x74, 0x31, 0x22, 0xa6, 0x01, 0x0a, 0x04, 0x45,
	0x78, 0x74, 0x32, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x32, 0x5f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x32, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x31, 0x32, 0x7d, 0x0a, 0x10, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x65, 0x74, 0x5f, 0x65, 0x78, 0x74, 0x32, 0x12, 0x2b, 0x2e, 0x68, 0x79, 0x62, 0x72, 0x69,
	0x64, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x65, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x53, 0x65, 0x74, 0x18, 0xe9, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x68,
	0x79, 0x62, 0x72, 0x69, 0x64, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x65, 0x74, 0x2e, 0x45,
	0x78, 0x74, 0x32, 0x52, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x74, 0x45,
	0x78, 0x74, 0x32, 0x22, 0xd2, 0x01, 0x0a, 0x0b, 0x45, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x64, 0x12, 0x2e, 0x0a, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x5f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x05, 0xaa, 0x01,
	0x02, 0x08, 0x03, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x31, 0x32, 0x92, 0x01, 0x0a, 0x17, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x65, 0x74, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12,
	0x2b, 0x2e, 0x68, 0x79, 0x62, 0x72, 0x69, 0x64, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x65,
	0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x74, 0x18, 0xea, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x68, 0x79, 0x62, 0x72, 0x69, 0x64, 0x2e, 0x67, 0x6f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x65, 0x74, 0x2e, 0x45, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x64, 0x52, 0x15, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x74, 0x45, 0x78, 0x74,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0xa5, 0x01, 0x0a, 0x0e, 0x45, 0x78, 0x74,
	0x4c, 0x61, 0x72, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x32, 0x92, 0x01, 0x0a, 0x14,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x65, 0x78, 0x74, 0x6c,
	0x61, 0x72, 0x67, 0x65, 0x12, 0x2b, 0x2e, 0x68, 0x79, 0x62, 0x72, 0x69, 0x64, 0x2e, 0x67, 0x6f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x65, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65,
	0x74, 0x18, 0x80, 0x80, 0x80, 0x80, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x68, 0x79,
	0x62, 0x72, 0x69, 0x64, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x65, 0x74, 0x2e, 0x45, 0x78,
	0x74, 0x4c, 0x61, 0x72, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x12, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x74, 0x45, 0x78, 0x74, 0x6c, 0x61, 0x72, 0x67, 0x65,
	0x42, 0x5e, 0x5a, 0x54, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x65, 0x74, 0x2f, 0x6d, 0x73,
	0x65, 0x74, 0x65, 0x78, 0x74, 0x70, 0x62, 0x2f, 0x6d, 0x73, 0x65, 0x74, 0x65, 0x78, 0x74, 0x70,
	0x62, 0x5f, 0x68, 0x79, 0x62, 0x72, 0x69, 0x64, 0x92, 0x03, 0x05, 0xd2, 0x3e, 0x02, 0x10, 0x02,
	0x62, 0x08, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x70, 0xe8, 0x07,
}

var file_internal_testprotos_messageset_msetextpb_msetextpb_hybrid_msetextpb_hybrid_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_internal_testprotos_messageset_msetextpb_msetextpb_hybrid_msetextpb_hybrid_proto_goTypes = []any{
	(*Ext1)(nil),                           // 0: hybrid.goproto.proto.messageset.Ext1
	(*Ext2)(nil),                           // 1: hybrid.goproto.proto.messageset.Ext2
	(*ExtRequired)(nil),                    // 2: hybrid.goproto.proto.messageset.ExtRequired
	(*ExtLargeNumber)(nil),                 // 3: hybrid.goproto.proto.messageset.ExtLargeNumber
	(*messagesetpb_hybrid.MessageSet)(nil), // 4: hybrid.goproto.proto.messageset.MessageSet
}
var file_internal_testprotos_messageset_msetextpb_msetextpb_hybrid_msetextpb_hybrid_proto_depIdxs = []int32{
	4, // 0: hybrid.goproto.proto.messageset.Ext1.message_set_ext1:extendee -> hybrid.goproto.proto.messageset.MessageSet
	4, // 1: hybrid.goproto.proto.messageset.Ext2.message_set_ext2:extendee -> hybrid.goproto.proto.messageset.MessageSet
	4, // 2: hybrid.goproto.proto.messageset.ExtRequired.message_set_extrequired:extendee -> hybrid.goproto.proto.messageset.MessageSet
	4, // 3: hybrid.goproto.proto.messageset.ExtLargeNumber.message_set_extlarge:extendee -> hybrid.goproto.proto.messageset.MessageSet
	0, // 4: hybrid.goproto.proto.messageset.Ext1.message_set_ext1:type_name -> hybrid.goproto.proto.messageset.Ext1
	1, // 5: hybrid.goproto.proto.messageset.Ext2.message_set_ext2:type_name -> hybrid.goproto.proto.messageset.Ext2
	2, // 6: hybrid.goproto.proto.messageset.ExtRequired.message_set_extrequired:type_name -> hybrid.goproto.proto.messageset.ExtRequired
	3, // 7: hybrid.goproto.proto.messageset.ExtLargeNumber.message_set_extlarge:type_name -> hybrid.goproto.proto.messageset.ExtLargeNumber
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	4, // [4:8] is the sub-list for extension type_name
	0, // [0:4] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_internal_testprotos_messageset_msetextpb_msetextpb_hybrid_msetextpb_hybrid_proto_init()
}
func file_internal_testprotos_messageset_msetextpb_msetextpb_hybrid_msetextpb_hybrid_proto_init() {
	if File_internal_testprotos_messageset_msetextpb_msetextpb_hybrid_msetextpb_hybrid_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_testprotos_messageset_msetextpb_msetextpb_hybrid_msetextpb_hybrid_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 4,
			NumServices:   0,
		},
		GoTypes:           file_internal_testprotos_messageset_msetextpb_msetextpb_hybrid_msetextpb_hybrid_proto_goTypes,
		DependencyIndexes: file_internal_testprotos_messageset_msetextpb_msetextpb_hybrid_msetextpb_hybrid_proto_depIdxs,
		MessageInfos:      file_internal_testprotos_messageset_msetextpb_msetextpb_hybrid_msetextpb_hybrid_proto_msgTypes,
		ExtensionInfos:    file_internal_testprotos_messageset_msetextpb_msetextpb_hybrid_msetextpb_hybrid_proto_extTypes,
	}.Build()
	File_internal_testprotos_messageset_msetextpb_msetextpb_hybrid_msetextpb_hybrid_proto = out.File
	file_internal_testprotos_messageset_msetextpb_msetextpb_hybrid_msetextpb_hybrid_proto_rawDesc = nil
	file_internal_testprotos_messageset_msetextpb_msetextpb_hybrid_msetextpb_hybrid_proto_goTypes = nil
	file_internal_testprotos_messageset_msetextpb_msetextpb_hybrid_msetextpb_hybrid_proto_depIdxs = nil
}
