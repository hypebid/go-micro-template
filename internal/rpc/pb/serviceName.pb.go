// protoc --go_out=. --go_opt=paths=source_relative \
// --go-grpc_out=. --go-grpc_opt=paths=source_relative \
// serviceName.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: serviceName.proto

package pb

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

type HealthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HealthRequest) Reset() {
	*x = HealthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serviceName_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthRequest) ProtoMessage() {}

func (x *HealthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_serviceName_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthRequest.ProtoReflect.Descriptor instead.
func (*HealthRequest) Descriptor() ([]byte, []int) {
	return file_serviceName_proto_rawDescGZIP(), []int{0}
}

func (x *HealthRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type HealthStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId  string `protobuf:"bytes,1,opt,name=transactionId,proto3" json:"transactionId,omitempty"`
	ServiceName    string `protobuf:"bytes,2,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
	ReleaseDate    string `protobuf:"bytes,3,opt,name=releaseDate,proto3" json:"releaseDate,omitempty"`
	ReleaseSlug    string `protobuf:"bytes,4,opt,name=releaseSlug,proto3" json:"releaseSlug,omitempty"`
	ReleaseVersion string `protobuf:"bytes,5,opt,name=releaseVersion,proto3" json:"releaseVersion,omitempty"`
	DatabaseOnline bool   `protobuf:"varint,6,opt,name=databaseOnline,proto3" json:"databaseOnline,omitempty"`
	Message        string `protobuf:"bytes,7,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HealthStatus) Reset() {
	*x = HealthStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serviceName_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthStatus) ProtoMessage() {}

func (x *HealthStatus) ProtoReflect() protoreflect.Message {
	mi := &file_serviceName_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthStatus.ProtoReflect.Descriptor instead.
func (*HealthStatus) Descriptor() ([]byte, []int) {
	return file_serviceName_proto_rawDescGZIP(), []int{1}
}

func (x *HealthStatus) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *HealthStatus) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *HealthStatus) GetReleaseDate() string {
	if x != nil {
		return x.ReleaseDate
	}
	return ""
}

func (x *HealthStatus) GetReleaseSlug() string {
	if x != nil {
		return x.ReleaseSlug
	}
	return ""
}

func (x *HealthStatus) GetReleaseVersion() string {
	if x != nil {
		return x.ReleaseVersion
	}
	return ""
}

func (x *HealthStatus) GetDatabaseOnline() bool {
	if x != nil {
		return x.DatabaseOnline
	}
	return false
}

func (x *HealthStatus) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_serviceName_proto protoreflect.FileDescriptor

var file_serviceName_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x68, 0x79, 0x70, 0x65, 0x62, 0x69, 0x64, 0x22, 0x29, 0x0a, 0x0d,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x84, 0x02, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x20,
	0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44, 0x61, 0x74, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x53, 0x6c, 0x75,
	0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x53, 0x6c, 0x75, 0x67, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0e,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x4f, 0x6e,
	0x6c, 0x69, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x4d,
	0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3e, 0x0a,
	0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x16, 0x2e, 0x68,
	0x79, 0x70, 0x65, 0x62, 0x69, 0x64, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x68, 0x79, 0x70, 0x65, 0x62, 0x69, 0x64, 0x2e, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x42, 0x37, 0x5a,
	0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x79, 0x70, 0x65,
	0x62, 0x69, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2d, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_serviceName_proto_rawDescOnce sync.Once
	file_serviceName_proto_rawDescData = file_serviceName_proto_rawDesc
)

func file_serviceName_proto_rawDescGZIP() []byte {
	file_serviceName_proto_rawDescOnce.Do(func() {
		file_serviceName_proto_rawDescData = protoimpl.X.CompressGZIP(file_serviceName_proto_rawDescData)
	})
	return file_serviceName_proto_rawDescData
}

var file_serviceName_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_serviceName_proto_goTypes = []interface{}{
	(*HealthRequest)(nil), // 0: hypebid.HealthRequest
	(*HealthStatus)(nil),  // 1: hypebid.HealthStatus
}
var file_serviceName_proto_depIdxs = []int32{
	0, // 0: hypebid.ServiceName.HealthCheck:input_type -> hypebid.HealthRequest
	1, // 1: hypebid.ServiceName.HealthCheck:output_type -> hypebid.HealthStatus
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_serviceName_proto_init() }
func file_serviceName_proto_init() {
	if File_serviceName_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_serviceName_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthRequest); i {
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
		file_serviceName_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthStatus); i {
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
			RawDescriptor: file_serviceName_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_serviceName_proto_goTypes,
		DependencyIndexes: file_serviceName_proto_depIdxs,
		MessageInfos:      file_serviceName_proto_msgTypes,
	}.Build()
	File_serviceName_proto = out.File
	file_serviceName_proto_rawDesc = nil
	file_serviceName_proto_goTypes = nil
	file_serviceName_proto_depIdxs = nil
}
