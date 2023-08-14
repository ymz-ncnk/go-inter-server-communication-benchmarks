// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.10
// source: protobuf/data.proto

package protobuf

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

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bool    bool    `protobuf:"varint,1,opt,name=Bool,proto3" json:"Bool,omitempty"`
	Int64   int64   `protobuf:"varint,2,opt,name=Int64,proto3" json:"Int64,omitempty"`
	String_ string  `protobuf:"bytes,3,opt,name=String,proto3" json:"String,omitempty"`
	Float64 float64 `protobuf:"fixed64,4,opt,name=Float64,proto3" json:"Float64,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_protobuf_data_proto_rawDescGZIP(), []int{0}
}

func (x *Data) GetBool() bool {
	if x != nil {
		return x.Bool
	}
	return false
}

func (x *Data) GetInt64() int64 {
	if x != nil {
		return x.Int64
	}
	return 0
}

func (x *Data) GetString_() string {
	if x != nil {
		return x.String_
	}
	return ""
}

func (x *Data) GetFloat64() float64 {
	if x != nil {
		return x.Float64
	}
	return 0
}

var File_protobuf_data_proto protoreflect.FileDescriptor

var file_protobuf_data_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22,
	0x62, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x49,
	0x6e, 0x74, 0x36, 0x34, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x49, 0x6e, 0x74, 0x36,
	0x34, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x46, 0x6c, 0x6f,
	0x61, 0x74, 0x36, 0x34, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x46, 0x6c, 0x6f, 0x61,
	0x74, 0x36, 0x34, 0x42, 0x4b, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x79, 0x6d, 0x7a, 0x2d, 0x6e, 0x63, 0x6e, 0x6b, 0x2f, 0x67, 0x6f, 0x2d, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61,
	0x72, 0x6b, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_data_proto_rawDescOnce sync.Once
	file_protobuf_data_proto_rawDescData = file_protobuf_data_proto_rawDesc
)

func file_protobuf_data_proto_rawDescGZIP() []byte {
	file_protobuf_data_proto_rawDescOnce.Do(func() {
		file_protobuf_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_data_proto_rawDescData)
	})
	return file_protobuf_data_proto_rawDescData
}

var file_protobuf_data_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protobuf_data_proto_goTypes = []interface{}{
	(*Data)(nil), // 0: protobuf.Data
}
var file_protobuf_data_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobuf_data_proto_init() }
func file_protobuf_data_proto_init() {
	if File_protobuf_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protobuf_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_data_proto_goTypes,
		DependencyIndexes: file_protobuf_data_proto_depIdxs,
		MessageInfos:      file_protobuf_data_proto_msgTypes,
	}.Build()
	File_protobuf_data_proto = out.File
	file_protobuf_data_proto_rawDesc = nil
	file_protobuf_data_proto_goTypes = nil
	file_protobuf_data_proto_depIdxs = nil
}
