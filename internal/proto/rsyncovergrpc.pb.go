// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.1
// source: rsyncovergrpc.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TransferRequest struct {
	state                       protoimpl.MessageState `protogen:"open.v1"`
	Args                        []string               `protobuf:"bytes,1,rep,name=args" json:"args,omitempty"`
	MinSupportedProtocolVersion *int32                 `protobuf:"varint,2,opt,name=min_supported_protocol_version,json=minSupportedProtocolVersion" json:"min_supported_protocol_version,omitempty"`
	MaxSupportedProtocolVersion *int32                 `protobuf:"varint,3,opt,name=max_supported_protocol_version,json=maxSupportedProtocolVersion" json:"max_supported_protocol_version,omitempty"`
	unknownFields               protoimpl.UnknownFields
	sizeCache                   protoimpl.SizeCache
}

func (x *TransferRequest) Reset() {
	*x = TransferRequest{}
	mi := &file_rsyncovergrpc_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransferRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferRequest) ProtoMessage() {}

func (x *TransferRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rsyncovergrpc_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferRequest.ProtoReflect.Descriptor instead.
func (*TransferRequest) Descriptor() ([]byte, []int) {
	return file_rsyncovergrpc_proto_rawDescGZIP(), []int{0}
}

func (x *TransferRequest) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *TransferRequest) GetMinSupportedProtocolVersion() int32 {
	if x != nil && x.MinSupportedProtocolVersion != nil {
		return *x.MinSupportedProtocolVersion
	}
	return 0
}

func (x *TransferRequest) GetMaxSupportedProtocolVersion() int32 {
	if x != nil && x.MaxSupportedProtocolVersion != nil {
		return *x.MaxSupportedProtocolVersion
	}
	return 0
}

type TransferChunk struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// request must be set on the first request of the stream,
	// and is ignored on all other requests.
	Request       *TransferRequest `protobuf:"bytes,1,opt,name=request" json:"request,omitempty"`
	Chunk         []byte           `protobuf:"bytes,2,opt,name=chunk" json:"chunk,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TransferChunk) Reset() {
	*x = TransferChunk{}
	mi := &file_rsyncovergrpc_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransferChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferChunk) ProtoMessage() {}

func (x *TransferChunk) ProtoReflect() protoreflect.Message {
	mi := &file_rsyncovergrpc_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferChunk.ProtoReflect.Descriptor instead.
func (*TransferChunk) Descriptor() ([]byte, []int) {
	return file_rsyncovergrpc_proto_rawDescGZIP(), []int{1}
}

func (x *TransferChunk) GetRequest() *TransferRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *TransferChunk) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

var File_rsyncovergrpc_proto protoreflect.FileDescriptor

var file_rsyncovergrpc_proto_rawDesc = string([]byte{
	0x0a, 0x13, 0x72, 0x73, 0x79, 0x6e, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x72, 0x73, 0x79, 0x6e, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x67, 0x72, 0x70, 0x63, 0x22, 0xaf, 0x01, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x12, 0x43, 0x0a, 0x1e,
	0x6d, 0x69, 0x6e, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x1b, 0x6d, 0x69, 0x6e, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x43, 0x0a, 0x1e, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x1b, 0x6d, 0x61, 0x78, 0x53, 0x75,
	0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x5f, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x38, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x73, 0x79, 0x6e, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x32, 0x52, 0x0a, 0x05, 0x52, 0x73, 0x79, 0x6e, 0x63,
	0x12, 0x49, 0x0a, 0x05, 0x52, 0x73, 0x79, 0x6e, 0x63, 0x12, 0x1c, 0x2e, 0x72, 0x73, 0x79, 0x6e,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x1a, 0x1c, 0x2e, 0x72, 0x73, 0x79, 0x6e, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x43, 0x68, 0x75, 0x6e, 0x6b, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x36, 0x5a, 0x34, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x61, 0x70, 0x65, 0x6c,
	0x62, 0x65, 0x72, 0x67, 0x2f, 0x72, 0x73, 0x79, 0x6e, 0x63, 0x2d, 0x6f, 0x76, 0x65, 0x72, 0x2d,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x08, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x70, 0xe8, 0x07,
})

var (
	file_rsyncovergrpc_proto_rawDescOnce sync.Once
	file_rsyncovergrpc_proto_rawDescData []byte
)

func file_rsyncovergrpc_proto_rawDescGZIP() []byte {
	file_rsyncovergrpc_proto_rawDescOnce.Do(func() {
		file_rsyncovergrpc_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_rsyncovergrpc_proto_rawDesc), len(file_rsyncovergrpc_proto_rawDesc)))
	})
	return file_rsyncovergrpc_proto_rawDescData
}

var file_rsyncovergrpc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rsyncovergrpc_proto_goTypes = []any{
	(*TransferRequest)(nil), // 0: rsyncovergrpc.TransferRequest
	(*TransferChunk)(nil),   // 1: rsyncovergrpc.TransferChunk
}
var file_rsyncovergrpc_proto_depIdxs = []int32{
	0, // 0: rsyncovergrpc.TransferChunk.request:type_name -> rsyncovergrpc.TransferRequest
	1, // 1: rsyncovergrpc.Rsync.Rsync:input_type -> rsyncovergrpc.TransferChunk
	1, // 2: rsyncovergrpc.Rsync.Rsync:output_type -> rsyncovergrpc.TransferChunk
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rsyncovergrpc_proto_init() }
func file_rsyncovergrpc_proto_init() {
	if File_rsyncovergrpc_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_rsyncovergrpc_proto_rawDesc), len(file_rsyncovergrpc_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rsyncovergrpc_proto_goTypes,
		DependencyIndexes: file_rsyncovergrpc_proto_depIdxs,
		MessageInfos:      file_rsyncovergrpc_proto_msgTypes,
	}.Build()
	File_rsyncovergrpc_proto = out.File
	file_rsyncovergrpc_proto_goTypes = nil
	file_rsyncovergrpc_proto_depIdxs = nil
}
