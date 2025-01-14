// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: pkg/ingress/get-instance-logs.proto

package ingress

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetWorkflowInstanceLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstanceId *string `protobuf:"bytes,1,opt,name=instanceId,proto3,oneof" json:"instanceId,omitempty"`
	Offset     *int32  `protobuf:"varint,2,opt,name=offset,proto3,oneof" json:"offset,omitempty"`
	Limit      *int32  `protobuf:"varint,3,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
}

func (x *GetWorkflowInstanceLogsRequest) Reset() {
	*x = GetWorkflowInstanceLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_ingress_get_instance_logs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWorkflowInstanceLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWorkflowInstanceLogsRequest) ProtoMessage() {}

func (x *GetWorkflowInstanceLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_ingress_get_instance_logs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWorkflowInstanceLogsRequest.ProtoReflect.Descriptor instead.
func (*GetWorkflowInstanceLogsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_ingress_get_instance_logs_proto_rawDescGZIP(), []int{0}
}

func (x *GetWorkflowInstanceLogsRequest) GetInstanceId() string {
	if x != nil && x.InstanceId != nil {
		return *x.InstanceId
	}
	return ""
}

func (x *GetWorkflowInstanceLogsRequest) GetOffset() int32 {
	if x != nil && x.Offset != nil {
		return *x.Offset
	}
	return 0
}

func (x *GetWorkflowInstanceLogsRequest) GetLimit() int32 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

type GetWorkflowInstanceLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkflowInstanceLogs []*GetWorkflowInstanceLogsResponse_WorkflowInstanceLog `protobuf:"bytes,1,rep,name=workflowInstanceLogs,proto3" json:"workflowInstanceLogs,omitempty"`
	Offset               *int32                                                 `protobuf:"varint,2,opt,name=offset,proto3,oneof" json:"offset,omitempty"`
	Limit                *int32                                                 `protobuf:"varint,3,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
}

func (x *GetWorkflowInstanceLogsResponse) Reset() {
	*x = GetWorkflowInstanceLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_ingress_get_instance_logs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWorkflowInstanceLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWorkflowInstanceLogsResponse) ProtoMessage() {}

func (x *GetWorkflowInstanceLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_ingress_get_instance_logs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWorkflowInstanceLogsResponse.ProtoReflect.Descriptor instead.
func (*GetWorkflowInstanceLogsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_ingress_get_instance_logs_proto_rawDescGZIP(), []int{1}
}

func (x *GetWorkflowInstanceLogsResponse) GetWorkflowInstanceLogs() []*GetWorkflowInstanceLogsResponse_WorkflowInstanceLog {
	if x != nil {
		return x.WorkflowInstanceLogs
	}
	return nil
}

func (x *GetWorkflowInstanceLogsResponse) GetOffset() int32 {
	if x != nil && x.Offset != nil {
		return *x.Offset
	}
	return 0
}

func (x *GetWorkflowInstanceLogsResponse) GetLimit() int32 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

type GetWorkflowInstanceLogsResponse_WorkflowInstanceLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        *string                `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	Level     *string                `protobuf:"bytes,2,opt,name=level,proto3,oneof" json:"level,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3,oneof" json:"timestamp,omitempty"`
	Message   *string                `protobuf:"bytes,4,opt,name=message,proto3,oneof" json:"message,omitempty"`
	Context   map[string]string      `protobuf:"bytes,5,rep,name=context,proto3" json:"context,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetWorkflowInstanceLogsResponse_WorkflowInstanceLog) Reset() {
	*x = GetWorkflowInstanceLogsResponse_WorkflowInstanceLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_ingress_get_instance_logs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWorkflowInstanceLogsResponse_WorkflowInstanceLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWorkflowInstanceLogsResponse_WorkflowInstanceLog) ProtoMessage() {}

func (x *GetWorkflowInstanceLogsResponse_WorkflowInstanceLog) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_ingress_get_instance_logs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWorkflowInstanceLogsResponse_WorkflowInstanceLog.ProtoReflect.Descriptor instead.
func (*GetWorkflowInstanceLogsResponse_WorkflowInstanceLog) Descriptor() ([]byte, []int) {
	return file_pkg_ingress_get_instance_logs_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetWorkflowInstanceLogsResponse_WorkflowInstanceLog) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *GetWorkflowInstanceLogsResponse_WorkflowInstanceLog) GetLevel() string {
	if x != nil && x.Level != nil {
		return *x.Level
	}
	return ""
}

func (x *GetWorkflowInstanceLogsResponse_WorkflowInstanceLog) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *GetWorkflowInstanceLogsResponse_WorkflowInstanceLog) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

func (x *GetWorkflowInstanceLogsResponse_WorkflowInstanceLog) GetContext() map[string]string {
	if x != nil {
		return x.Context
	}
	return nil
}

var File_pkg_ingress_get_instance_logs_proto protoreflect.FileDescriptor

var file_pkg_ingress_get_instance_logs_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x6b, 0x67, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x67, 0x65,
	0x74, 0x2d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2d, 0x6c, 0x6f, 0x67, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xa1, 0x01, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x23, 0x0a, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x48, 0x02, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x88, 0x01, 0x01, 0x42,
	0x0d, 0x0a, 0x0b, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x42, 0x09,
	0x0a, 0x07, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x22, 0xd2, 0x04, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x70, 0x0a, 0x14, 0x77, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x4c, 0x6f, 0x67, 0x52, 0x14, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x1b, 0x0a, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x88, 0x01,
	0x01, 0x1a, 0xef, 0x02, 0x0a, 0x13, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x19,
	0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52,
	0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x3d, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x02, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x63, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x49, 0x2e, 0x69, 0x6e, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x1a, 0x3a, 0x0a, 0x0c,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x42, 0x08,
	0x0a, 0x06, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x6f, 0x72, 0x74, 0x65, 0x69, 0x6c, 0x2f, 0x64,
	0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x69, 0x6e, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_ingress_get_instance_logs_proto_rawDescOnce sync.Once
	file_pkg_ingress_get_instance_logs_proto_rawDescData = file_pkg_ingress_get_instance_logs_proto_rawDesc
)

func file_pkg_ingress_get_instance_logs_proto_rawDescGZIP() []byte {
	file_pkg_ingress_get_instance_logs_proto_rawDescOnce.Do(func() {
		file_pkg_ingress_get_instance_logs_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_ingress_get_instance_logs_proto_rawDescData)
	})
	return file_pkg_ingress_get_instance_logs_proto_rawDescData
}

var file_pkg_ingress_get_instance_logs_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_ingress_get_instance_logs_proto_goTypes = []interface{}{
	(*GetWorkflowInstanceLogsRequest)(nil),                      // 0: ingress.GetWorkflowInstanceLogsRequest
	(*GetWorkflowInstanceLogsResponse)(nil),                     // 1: ingress.GetWorkflowInstanceLogsResponse
	(*GetWorkflowInstanceLogsResponse_WorkflowInstanceLog)(nil), // 2: ingress.GetWorkflowInstanceLogsResponse.WorkflowInstanceLog
	nil,                           // 3: ingress.GetWorkflowInstanceLogsResponse.WorkflowInstanceLog.ContextEntry
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_pkg_ingress_get_instance_logs_proto_depIdxs = []int32{
	2, // 0: ingress.GetWorkflowInstanceLogsResponse.workflowInstanceLogs:type_name -> ingress.GetWorkflowInstanceLogsResponse.WorkflowInstanceLog
	4, // 1: ingress.GetWorkflowInstanceLogsResponse.WorkflowInstanceLog.timestamp:type_name -> google.protobuf.Timestamp
	3, // 2: ingress.GetWorkflowInstanceLogsResponse.WorkflowInstanceLog.context:type_name -> ingress.GetWorkflowInstanceLogsResponse.WorkflowInstanceLog.ContextEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_ingress_get_instance_logs_proto_init() }
func file_pkg_ingress_get_instance_logs_proto_init() {
	if File_pkg_ingress_get_instance_logs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_ingress_get_instance_logs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWorkflowInstanceLogsRequest); i {
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
		file_pkg_ingress_get_instance_logs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWorkflowInstanceLogsResponse); i {
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
		file_pkg_ingress_get_instance_logs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWorkflowInstanceLogsResponse_WorkflowInstanceLog); i {
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
	file_pkg_ingress_get_instance_logs_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_pkg_ingress_get_instance_logs_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_pkg_ingress_get_instance_logs_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_ingress_get_instance_logs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_ingress_get_instance_logs_proto_goTypes,
		DependencyIndexes: file_pkg_ingress_get_instance_logs_proto_depIdxs,
		MessageInfos:      file_pkg_ingress_get_instance_logs_proto_msgTypes,
	}.Build()
	File_pkg_ingress_get_instance_logs_proto = out.File
	file_pkg_ingress_get_instance_logs_proto_rawDesc = nil
	file_pkg_ingress_get_instance_logs_proto_goTypes = nil
	file_pkg_ingress_get_instance_logs_proto_depIdxs = nil
}
