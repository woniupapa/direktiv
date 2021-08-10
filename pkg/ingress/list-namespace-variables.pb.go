// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: pkg/ingress/list-namespace-variables.proto

package ingress

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

type ListNamespaceVariablesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace *string `protobuf:"bytes,1,opt,name=namespace,proto3,oneof" json:"namespace,omitempty"`
}

func (x *ListNamespaceVariablesRequest) Reset() {
	*x = ListNamespaceVariablesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_ingress_list_namespace_variables_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNamespaceVariablesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNamespaceVariablesRequest) ProtoMessage() {}

func (x *ListNamespaceVariablesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_ingress_list_namespace_variables_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNamespaceVariablesRequest.ProtoReflect.Descriptor instead.
func (*ListNamespaceVariablesRequest) Descriptor() ([]byte, []int) {
	return file_pkg_ingress_list_namespace_variables_proto_rawDescGZIP(), []int{0}
}

func (x *ListNamespaceVariablesRequest) GetNamespace() string {
	if x != nil && x.Namespace != nil {
		return *x.Namespace
	}
	return ""
}

type ListNamespaceVariablesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Variables []*ListNamespaceVariablesResponse_Variable `protobuf:"bytes,1,rep,name=variables,proto3" json:"variables,omitempty"`
}

func (x *ListNamespaceVariablesResponse) Reset() {
	*x = ListNamespaceVariablesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_ingress_list_namespace_variables_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNamespaceVariablesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNamespaceVariablesResponse) ProtoMessage() {}

func (x *ListNamespaceVariablesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_ingress_list_namespace_variables_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNamespaceVariablesResponse.ProtoReflect.Descriptor instead.
func (*ListNamespaceVariablesResponse) Descriptor() ([]byte, []int) {
	return file_pkg_ingress_list_namespace_variables_proto_rawDescGZIP(), []int{1}
}

func (x *ListNamespaceVariablesResponse) GetVariables() []*ListNamespaceVariablesResponse_Variable {
	if x != nil {
		return x.Variables
	}
	return nil
}

type ListNamespaceVariablesResponse_Variable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name *string `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Size *int64  `protobuf:"varint,2,opt,name=size,proto3,oneof" json:"size,omitempty"`
}

func (x *ListNamespaceVariablesResponse_Variable) Reset() {
	*x = ListNamespaceVariablesResponse_Variable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_ingress_list_namespace_variables_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNamespaceVariablesResponse_Variable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNamespaceVariablesResponse_Variable) ProtoMessage() {}

func (x *ListNamespaceVariablesResponse_Variable) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_ingress_list_namespace_variables_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNamespaceVariablesResponse_Variable.ProtoReflect.Descriptor instead.
func (*ListNamespaceVariablesResponse_Variable) Descriptor() ([]byte, []int) {
	return file_pkg_ingress_list_namespace_variables_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ListNamespaceVariablesResponse_Variable) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *ListNamespaceVariablesResponse_Variable) GetSize() int64 {
	if x != nil && x.Size != nil {
		return *x.Size
	}
	return 0
}

var File_pkg_ingress_list_namespace_variables_proto protoreflect.FileDescriptor

var file_pkg_ingress_list_namespace_variables_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x70, 0x6b, 0x67, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x6c, 0x69,
	0x73, 0x74, 0x2d, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2d, 0x76, 0x61, 0x72,
	0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x69, 0x6e,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x22, 0x50, 0x0a, 0x1d, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0xc0, 0x01, 0x0a, 0x1e, 0x4c, 0x69, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x09, 0x76, 0x61,
	0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e,
	0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x52,
	0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x1a, 0x4e, 0x0a, 0x08, 0x56, 0x61,
	0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x17, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x6f, 0x72, 0x74, 0x65, 0x69, 0x6c,
	0x2f, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x69, 0x6e,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_ingress_list_namespace_variables_proto_rawDescOnce sync.Once
	file_pkg_ingress_list_namespace_variables_proto_rawDescData = file_pkg_ingress_list_namespace_variables_proto_rawDesc
)

func file_pkg_ingress_list_namespace_variables_proto_rawDescGZIP() []byte {
	file_pkg_ingress_list_namespace_variables_proto_rawDescOnce.Do(func() {
		file_pkg_ingress_list_namespace_variables_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_ingress_list_namespace_variables_proto_rawDescData)
	})
	return file_pkg_ingress_list_namespace_variables_proto_rawDescData
}

var file_pkg_ingress_list_namespace_variables_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_ingress_list_namespace_variables_proto_goTypes = []interface{}{
	(*ListNamespaceVariablesRequest)(nil),           // 0: ingress.ListNamespaceVariablesRequest
	(*ListNamespaceVariablesResponse)(nil),          // 1: ingress.ListNamespaceVariablesResponse
	(*ListNamespaceVariablesResponse_Variable)(nil), // 2: ingress.ListNamespaceVariablesResponse.Variable
}
var file_pkg_ingress_list_namespace_variables_proto_depIdxs = []int32{
	2, // 0: ingress.ListNamespaceVariablesResponse.variables:type_name -> ingress.ListNamespaceVariablesResponse.Variable
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_ingress_list_namespace_variables_proto_init() }
func file_pkg_ingress_list_namespace_variables_proto_init() {
	if File_pkg_ingress_list_namespace_variables_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_ingress_list_namespace_variables_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNamespaceVariablesRequest); i {
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
		file_pkg_ingress_list_namespace_variables_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNamespaceVariablesResponse); i {
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
		file_pkg_ingress_list_namespace_variables_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNamespaceVariablesResponse_Variable); i {
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
	file_pkg_ingress_list_namespace_variables_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_pkg_ingress_list_namespace_variables_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_ingress_list_namespace_variables_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_ingress_list_namespace_variables_proto_goTypes,
		DependencyIndexes: file_pkg_ingress_list_namespace_variables_proto_depIdxs,
		MessageInfos:      file_pkg_ingress_list_namespace_variables_proto_msgTypes,
	}.Build()
	File_pkg_ingress_list_namespace_variables_proto = out.File
	file_pkg_ingress_list_namespace_variables_proto_rawDesc = nil
	file_pkg_ingress_list_namespace_variables_proto_goTypes = nil
	file_pkg_ingress_list_namespace_variables_proto_depIdxs = nil
}
