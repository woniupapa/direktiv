// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: pkg/functions/grpc/functions-list.proto

package grpc

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type FunctionsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info        *BaseInfo    `protobuf:"bytes,1,opt,name=info,proto3,oneof" json:"info,omitempty"`
	Status      *string      `protobuf:"bytes,2,opt,name=status,proto3,oneof" json:"status,omitempty"`
	Conditions  []*Condition `protobuf:"bytes,3,rep,name=conditions,proto3" json:"conditions,omitempty"`
	ServiceName *string      `protobuf:"bytes,4,opt,name=serviceName,proto3,oneof" json:"serviceName,omitempty"`
}

func (x *FunctionsInfo) Reset() {
	*x = FunctionsInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_functions_grpc_functions_list_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionsInfo) ProtoMessage() {}

func (x *FunctionsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_functions_grpc_functions_list_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionsInfo.ProtoReflect.Descriptor instead.
func (*FunctionsInfo) Descriptor() ([]byte, []int) {
	return file_pkg_functions_grpc_functions_list_proto_rawDescGZIP(), []int{0}
}

func (x *FunctionsInfo) GetInfo() *BaseInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *FunctionsInfo) GetStatus() string {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ""
}

func (x *FunctionsInfo) GetConditions() []*Condition {
	if x != nil {
		return x.Conditions
	}
	return nil
}

func (x *FunctionsInfo) GetServiceName() string {
	if x != nil && x.ServiceName != nil {
		return *x.ServiceName
	}
	return ""
}

type ListFunctionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config    *FunctionsConfig `protobuf:"bytes,1,opt,name=config,proto3,oneof" json:"config,omitempty"`
	Functions []*FunctionsInfo `protobuf:"bytes,2,rep,name=functions,proto3" json:"functions,omitempty"`
}

func (x *ListFunctionsResponse) Reset() {
	*x = ListFunctionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_functions_grpc_functions_list_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFunctionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFunctionsResponse) ProtoMessage() {}

func (x *ListFunctionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_functions_grpc_functions_list_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFunctionsResponse.ProtoReflect.Descriptor instead.
func (*ListFunctionsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_functions_grpc_functions_list_proto_rawDescGZIP(), []int{1}
}

func (x *ListFunctionsResponse) GetConfig() *FunctionsConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *ListFunctionsResponse) GetFunctions() []*FunctionsInfo {
	if x != nil {
		return x.Functions
	}
	return nil
}

type ListFunctionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Annotations map[string]string `protobuf:"bytes,1,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ListFunctionsRequest) Reset() {
	*x = ListFunctionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_functions_grpc_functions_list_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFunctionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFunctionsRequest) ProtoMessage() {}

func (x *ListFunctionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_functions_grpc_functions_list_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFunctionsRequest.ProtoReflect.Descriptor instead.
func (*ListFunctionsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_functions_grpc_functions_list_proto_rawDescGZIP(), []int{2}
}

func (x *ListFunctionsRequest) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

var File_pkg_functions_grpc_functions_list_proto protoreflect.FileDescriptor

var file_pkg_functions_grpc_functions_list_proto_rawDesc = []byte{
	0x0a, 0x27, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x6c,
	0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x1a,
	0x29, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x01, 0x0a, 0x0d, 0x46,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x27, 0x0a, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88,
	0x01, 0x01, 0x12, 0x2f, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x25, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0e,
	0x0a, 0x0c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x89,
	0x01, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48,
	0x00, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x88, 0x01, 0x01, 0x12, 0x31, 0x0a, 0x09,
	0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42,
	0x09, 0x0a, 0x07, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0xa5, 0x01, 0x0a, 0x14, 0x4c,
	0x69, 0x73, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x4d, 0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x1a, 0x3e, 0x0a, 0x10, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x76, 0x6f, 0x72, 0x74, 0x65, 0x69, 0x6c, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69,
	0x76, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_functions_grpc_functions_list_proto_rawDescOnce sync.Once
	file_pkg_functions_grpc_functions_list_proto_rawDescData = file_pkg_functions_grpc_functions_list_proto_rawDesc
)

func file_pkg_functions_grpc_functions_list_proto_rawDescGZIP() []byte {
	file_pkg_functions_grpc_functions_list_proto_rawDescOnce.Do(func() {
		file_pkg_functions_grpc_functions_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_functions_grpc_functions_list_proto_rawDescData)
	})
	return file_pkg_functions_grpc_functions_list_proto_rawDescData
}

var file_pkg_functions_grpc_functions_list_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_functions_grpc_functions_list_proto_goTypes = []interface{}{
	(*FunctionsInfo)(nil),         // 0: grpc.FunctionsInfo
	(*ListFunctionsResponse)(nil), // 1: grpc.ListFunctionsResponse
	(*ListFunctionsRequest)(nil),  // 2: grpc.ListFunctionsRequest
	nil,                           // 3: grpc.ListFunctionsRequest.AnnotationsEntry
	(*BaseInfo)(nil),              // 4: grpc.BaseInfo
	(*Condition)(nil),             // 5: grpc.Condition
	(*FunctionsConfig)(nil),       // 6: grpc.FunctionsConfig
}
var file_pkg_functions_grpc_functions_list_proto_depIdxs = []int32{
	4, // 0: grpc.FunctionsInfo.info:type_name -> grpc.BaseInfo
	5, // 1: grpc.FunctionsInfo.conditions:type_name -> grpc.Condition
	6, // 2: grpc.ListFunctionsResponse.config:type_name -> grpc.FunctionsConfig
	0, // 3: grpc.ListFunctionsResponse.functions:type_name -> grpc.FunctionsInfo
	3, // 4: grpc.ListFunctionsRequest.annotations:type_name -> grpc.ListFunctionsRequest.AnnotationsEntry
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_pkg_functions_grpc_functions_list_proto_init() }
func file_pkg_functions_grpc_functions_list_proto_init() {
	if File_pkg_functions_grpc_functions_list_proto != nil {
		return
	}
	file_pkg_functions_grpc_functions_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pkg_functions_grpc_functions_list_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FunctionsInfo); i {
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
		file_pkg_functions_grpc_functions_list_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFunctionsResponse); i {
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
		file_pkg_functions_grpc_functions_list_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFunctionsRequest); i {
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
	file_pkg_functions_grpc_functions_list_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_pkg_functions_grpc_functions_list_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_functions_grpc_functions_list_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_functions_grpc_functions_list_proto_goTypes,
		DependencyIndexes: file_pkg_functions_grpc_functions_list_proto_depIdxs,
		MessageInfos:      file_pkg_functions_grpc_functions_list_proto_msgTypes,
	}.Build()
	File_pkg_functions_grpc_functions_list_proto = out.File
	file_pkg_functions_grpc_functions_list_proto_rawDesc = nil
	file_pkg_functions_grpc_functions_list_proto_goTypes = nil
	file_pkg_functions_grpc_functions_list_proto_depIdxs = nil
}
