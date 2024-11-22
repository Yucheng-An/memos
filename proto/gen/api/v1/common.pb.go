// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: api/v1/common.proto

package apiv1

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

type RowStatus int32

const (
	RowStatus_ROW_STATUS_UNSPECIFIED RowStatus = 0
	RowStatus_ACTIVE                 RowStatus = 1
	RowStatus_ARCHIVED               RowStatus = 2
)

// Enum value maps for RowStatus.
var (
	RowStatus_name = map[int32]string{
		0: "ROW_STATUS_UNSPECIFIED",
		1: "ACTIVE",
		2: "ARCHIVED",
	}
	RowStatus_value = map[string]int32{
		"ROW_STATUS_UNSPECIFIED": 0,
		"ACTIVE":                 1,
		"ARCHIVED":               2,
	}
)

func (x RowStatus) Enum() *RowStatus {
	p := new(RowStatus)
	*p = x
	return p
}

func (x RowStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RowStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v1_common_proto_enumTypes[0].Descriptor()
}

func (RowStatus) Type() protoreflect.EnumType {
	return &file_api_v1_common_proto_enumTypes[0]
}

func (x RowStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RowStatus.Descriptor instead.
func (RowStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_v1_common_proto_rawDescGZIP(), []int{0}
}

// Used internally for obfuscating the page token.
type PageToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *PageToken) Reset() {
	*x = PageToken{}
	mi := &file_api_v1_common_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PageToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageToken) ProtoMessage() {}

func (x *PageToken) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_common_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageToken.ProtoReflect.Descriptor instead.
func (*PageToken) Descriptor() ([]byte, []int) {
	return file_api_v1_common_proto_rawDescGZIP(), []int{0}
}

func (x *PageToken) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *PageToken) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

var File_api_v1_common_proto protoreflect.FileDescriptor

var file_api_v1_common_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x22, 0x39, 0x0a, 0x09, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x2a, 0x41,
	0x0a, 0x09, 0x52, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x16, 0x52,
	0x4f, 0x57, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56,
	0x45, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x52, 0x43, 0x48, 0x49, 0x56, 0x45, 0x44, 0x10,
	0x02, 0x42, 0xa3, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x75, 0x73, 0x65, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2f, 0x6d, 0x65, 0x6d, 0x6f, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x3b, 0x61, 0x70, 0x69, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4d, 0x41, 0x58, 0xaa, 0x02, 0x0c,
	0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0c, 0x4d,
	0x65, 0x6d, 0x6f, 0x73, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x18, 0x4d, 0x65,
	0x6d, 0x6f, 0x73, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x3a, 0x3a,
	0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_common_proto_rawDescOnce sync.Once
	file_api_v1_common_proto_rawDescData = file_api_v1_common_proto_rawDesc
)

func file_api_v1_common_proto_rawDescGZIP() []byte {
	file_api_v1_common_proto_rawDescOnce.Do(func() {
		file_api_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_common_proto_rawDescData)
	})
	return file_api_v1_common_proto_rawDescData
}

var file_api_v1_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_v1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_v1_common_proto_goTypes = []any{
	(RowStatus)(0),    // 0: memos.api.v1.RowStatus
	(*PageToken)(nil), // 1: memos.api.v1.PageToken
}
var file_api_v1_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_v1_common_proto_init() }
func file_api_v1_common_proto_init() {
	if File_api_v1_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1_common_proto_goTypes,
		DependencyIndexes: file_api_v1_common_proto_depIdxs,
		EnumInfos:         file_api_v1_common_proto_enumTypes,
		MessageInfos:      file_api_v1_common_proto_msgTypes,
	}.Build()
	File_api_v1_common_proto = out.File
	file_api_v1_common_proto_rawDesc = nil
	file_api_v1_common_proto_goTypes = nil
	file_api_v1_common_proto_depIdxs = nil
}
