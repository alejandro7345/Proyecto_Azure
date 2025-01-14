// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.2
// source: proto/instrumento.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type InstrumentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *InstrumentRequest) Reset() {
	*x = InstrumentRequest{}
	mi := &file_proto_instrumento_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InstrumentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstrumentRequest) ProtoMessage() {}

func (x *InstrumentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_instrumento_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstrumentRequest.ProtoReflect.Descriptor instead.
func (*InstrumentRequest) Descriptor() ([]byte, []int) {
	return file_proto_instrumento_proto_rawDescGZIP(), []int{0}
}

func (x *InstrumentRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type InstrumentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type   string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Origin string `protobuf:"bytes,3,opt,name=origin,proto3" json:"origin,omitempty"`
}

func (x *InstrumentResponse) Reset() {
	*x = InstrumentResponse{}
	mi := &file_proto_instrumento_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InstrumentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstrumentResponse) ProtoMessage() {}

func (x *InstrumentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_instrumento_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstrumentResponse.ProtoReflect.Descriptor instead.
func (*InstrumentResponse) Descriptor() ([]byte, []int) {
	return file_proto_instrumento_proto_rawDescGZIP(), []int{1}
}

func (x *InstrumentResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InstrumentResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *InstrumentResponse) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

var File_proto_instrumento_proto protoreflect.FileDescriptor

var file_proto_instrumento_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x6f, 0x22, 0x27, 0x0a, 0x11, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x54, 0x0a, 0x12, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x32, 0x69, 0x0a, 0x11, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x66, 0x6f, 0x12,
	0x1e, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x6f, 0x2e, 0x49, 0x6e,
	0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1f, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x6f, 0x2e, 0x49, 0x6e,
	0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x1a, 0x5a, 0x18, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x6f, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_instrumento_proto_rawDescOnce sync.Once
	file_proto_instrumento_proto_rawDescData = file_proto_instrumento_proto_rawDesc
)

func file_proto_instrumento_proto_rawDescGZIP() []byte {
	file_proto_instrumento_proto_rawDescOnce.Do(func() {
		file_proto_instrumento_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_instrumento_proto_rawDescData)
	})
	return file_proto_instrumento_proto_rawDescData
}

var file_proto_instrumento_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_instrumento_proto_goTypes = []any{
	(*InstrumentRequest)(nil),  // 0: instrumento.InstrumentRequest
	(*InstrumentResponse)(nil), // 1: instrumento.InstrumentResponse
}
var file_proto_instrumento_proto_depIdxs = []int32{
	0, // 0: instrumento.InstrumentService.GetInstrumentinfo:input_type -> instrumento.InstrumentRequest
	1, // 1: instrumento.InstrumentService.GetInstrumentinfo:output_type -> instrumento.InstrumentResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_instrumento_proto_init() }
func file_proto_instrumento_proto_init() {
	if File_proto_instrumento_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_instrumento_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_instrumento_proto_goTypes,
		DependencyIndexes: file_proto_instrumento_proto_depIdxs,
		MessageInfos:      file_proto_instrumento_proto_msgTypes,
	}.Build()
	File_proto_instrumento_proto = out.File
	file_proto_instrumento_proto_rawDesc = nil
	file_proto_instrumento_proto_goTypes = nil
	file_proto_instrumento_proto_depIdxs = nil
}
