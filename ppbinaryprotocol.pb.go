// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.0
// source: ppbinaryprotocol.proto

package BinaryProtocolService

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type DecodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data      []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	DeviceEUI string `protobuf:"bytes,2,opt,name=deviceEUI,proto3" json:"deviceEUI,omitempty"`
}

func (x *DecodeRequest) Reset() {
	*x = DecodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ppbinaryprotocol_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecodeRequest) ProtoMessage() {}

func (x *DecodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ppbinaryprotocol_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecodeRequest.ProtoReflect.Descriptor instead.
func (*DecodeRequest) Descriptor() ([]byte, []int) {
	return file_ppbinaryprotocol_proto_rawDescGZIP(), []int{0}
}

func (x *DecodeRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *DecodeRequest) GetDeviceEUI() string {
	if x != nil {
		return x.DeviceEUI
	}
	return ""
}

type DecodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []*PowerpilotProtoMessage `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *DecodeResponse) Reset() {
	*x = DecodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ppbinaryprotocol_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecodeResponse) ProtoMessage() {}

func (x *DecodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ppbinaryprotocol_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecodeResponse.ProtoReflect.Descriptor instead.
func (*DecodeResponse) Descriptor() ([]byte, []int) {
	return file_ppbinaryprotocol_proto_rawDescGZIP(), []int{1}
}

func (x *DecodeResponse) GetMessages() []*PowerpilotProtoMessage {
	if x != nil {
		return x.Messages
	}
	return nil
}

type PowerpilotProtoMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *PowerpilotProtoMessage) Reset() {
	*x = PowerpilotProtoMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ppbinaryprotocol_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PowerpilotProtoMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PowerpilotProtoMessage) ProtoMessage() {}

func (x *PowerpilotProtoMessage) ProtoReflect() protoreflect.Message {
	mi := &file_ppbinaryprotocol_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PowerpilotProtoMessage.ProtoReflect.Descriptor instead.
func (*PowerpilotProtoMessage) Descriptor() ([]byte, []int) {
	return file_ppbinaryprotocol_proto_rawDescGZIP(), []int{2}
}

func (x *PowerpilotProtoMessage) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *PowerpilotProtoMessage) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type EncodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *EncodeRequest) Reset() {
	*x = EncodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ppbinaryprotocol_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncodeRequest) ProtoMessage() {}

func (x *EncodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ppbinaryprotocol_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncodeRequest.ProtoReflect.Descriptor instead.
func (*EncodeRequest) Descriptor() ([]byte, []int) {
	return file_ppbinaryprotocol_proto_rawDescGZIP(), []int{3}
}

func (x *EncodeRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *EncodeRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type PowerPilotBinaryMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PowerPilotBinaryMessage) Reset() {
	*x = PowerPilotBinaryMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ppbinaryprotocol_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PowerPilotBinaryMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PowerPilotBinaryMessage) ProtoMessage() {}

func (x *PowerPilotBinaryMessage) ProtoReflect() protoreflect.Message {
	mi := &file_ppbinaryprotocol_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PowerPilotBinaryMessage.ProtoReflect.Descriptor instead.
func (*PowerPilotBinaryMessage) Descriptor() ([]byte, []int) {
	return file_ppbinaryprotocol_proto_rawDescGZIP(), []int{4}
}

func (x *PowerPilotBinaryMessage) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_ppbinaryprotocol_proto protoreflect.FileDescriptor

var file_ppbinaryprotocol_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x70, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x70, 0x70, 0x62, 0x69, 0x6e, 0x61,
	0x72, 0x79, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x41, 0x0a, 0x0d, 0x44, 0x65,
	0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x1c, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x45, 0x55, 0x49, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x45, 0x55, 0x49, 0x22, 0x56, 0x0a,
	0x0e, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x44, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x70, 0x70, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x70, 0x69, 0x6c, 0x6f, 0x74, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x42, 0x0a, 0x16, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x70, 0x69,
	0x6c, 0x6f, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x37, 0x0a, 0x0d, 0x45, 0x6e, 0x63,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x2d, 0x0a, 0x17, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x50, 0x69, 0x6c, 0x6f, 0x74,
	0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x32, 0xb3, 0x01, 0x0a, 0x0e, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x4b, 0x0a, 0x06, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1f,
	0x2e, 0x70, 0x70, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x20, 0x2e, 0x70, 0x70, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x54, 0x0a, 0x06, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1f, 0x2e, 0x70, 0x70,
	0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x45,
	0x6e, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x70,
	0x70, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x50, 0x6f, 0x77, 0x65, 0x72, 0x50, 0x69, 0x6c, 0x6f, 0x74, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x31, 0x5a, 0x17, 0x2e, 0x3b, 0x42, 0x69, 0x6e,
	0x61, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0xaa, 0x02, 0x15, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_ppbinaryprotocol_proto_rawDescOnce sync.Once
	file_ppbinaryprotocol_proto_rawDescData = file_ppbinaryprotocol_proto_rawDesc
)

func file_ppbinaryprotocol_proto_rawDescGZIP() []byte {
	file_ppbinaryprotocol_proto_rawDescOnce.Do(func() {
		file_ppbinaryprotocol_proto_rawDescData = protoimpl.X.CompressGZIP(file_ppbinaryprotocol_proto_rawDescData)
	})
	return file_ppbinaryprotocol_proto_rawDescData
}

var file_ppbinaryprotocol_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_ppbinaryprotocol_proto_goTypes = []interface{}{
	(*DecodeRequest)(nil),           // 0: ppbinaryprotocol.DecodeRequest
	(*DecodeResponse)(nil),          // 1: ppbinaryprotocol.DecodeResponse
	(*PowerpilotProtoMessage)(nil),  // 2: ppbinaryprotocol.PowerpilotProtoMessage
	(*EncodeRequest)(nil),           // 3: ppbinaryprotocol.EncodeRequest
	(*PowerPilotBinaryMessage)(nil), // 4: ppbinaryprotocol.PowerPilotBinaryMessage
}
var file_ppbinaryprotocol_proto_depIdxs = []int32{
	2, // 0: ppbinaryprotocol.DecodeResponse.messages:type_name -> ppbinaryprotocol.PowerpilotProtoMessage
	0, // 1: ppbinaryprotocol.BinaryProtocol.Decode:input_type -> ppbinaryprotocol.DecodeRequest
	3, // 2: ppbinaryprotocol.BinaryProtocol.Encode:input_type -> ppbinaryprotocol.EncodeRequest
	1, // 3: ppbinaryprotocol.BinaryProtocol.Decode:output_type -> ppbinaryprotocol.DecodeResponse
	4, // 4: ppbinaryprotocol.BinaryProtocol.Encode:output_type -> ppbinaryprotocol.PowerPilotBinaryMessage
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ppbinaryprotocol_proto_init() }
func file_ppbinaryprotocol_proto_init() {
	if File_ppbinaryprotocol_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ppbinaryprotocol_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecodeRequest); i {
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
		file_ppbinaryprotocol_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecodeResponse); i {
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
		file_ppbinaryprotocol_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PowerpilotProtoMessage); i {
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
		file_ppbinaryprotocol_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncodeRequest); i {
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
		file_ppbinaryprotocol_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PowerPilotBinaryMessage); i {
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
			RawDescriptor: file_ppbinaryprotocol_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ppbinaryprotocol_proto_goTypes,
		DependencyIndexes: file_ppbinaryprotocol_proto_depIdxs,
		MessageInfos:      file_ppbinaryprotocol_proto_msgTypes,
	}.Build()
	File_ppbinaryprotocol_proto = out.File
	file_ppbinaryprotocol_proto_rawDesc = nil
	file_ppbinaryprotocol_proto_goTypes = nil
	file_ppbinaryprotocol_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BinaryProtocolClient is the client API for BinaryProtocol service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BinaryProtocolClient interface {
	Decode(ctx context.Context, in *DecodeRequest, opts ...grpc.CallOption) (*DecodeResponse, error)
	Encode(ctx context.Context, in *EncodeRequest, opts ...grpc.CallOption) (*PowerPilotBinaryMessage, error)
}

type binaryProtocolClient struct {
	cc grpc.ClientConnInterface
}

func NewBinaryProtocolClient(cc grpc.ClientConnInterface) BinaryProtocolClient {
	return &binaryProtocolClient{cc}
}

func (c *binaryProtocolClient) Decode(ctx context.Context, in *DecodeRequest, opts ...grpc.CallOption) (*DecodeResponse, error) {
	out := new(DecodeResponse)
	err := c.cc.Invoke(ctx, "/ppbinaryprotocol.BinaryProtocol/Decode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *binaryProtocolClient) Encode(ctx context.Context, in *EncodeRequest, opts ...grpc.CallOption) (*PowerPilotBinaryMessage, error) {
	out := new(PowerPilotBinaryMessage)
	err := c.cc.Invoke(ctx, "/ppbinaryprotocol.BinaryProtocol/Encode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BinaryProtocolServer is the server API for BinaryProtocol service.
type BinaryProtocolServer interface {
	Decode(context.Context, *DecodeRequest) (*DecodeResponse, error)
	Encode(context.Context, *EncodeRequest) (*PowerPilotBinaryMessage, error)
}

// UnimplementedBinaryProtocolServer can be embedded to have forward compatible implementations.
type UnimplementedBinaryProtocolServer struct {
}

func (*UnimplementedBinaryProtocolServer) Decode(context.Context, *DecodeRequest) (*DecodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Decode not implemented")
}
func (*UnimplementedBinaryProtocolServer) Encode(context.Context, *EncodeRequest) (*PowerPilotBinaryMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Encode not implemented")
}

func RegisterBinaryProtocolServer(s *grpc.Server, srv BinaryProtocolServer) {
	s.RegisterService(&_BinaryProtocol_serviceDesc, srv)
}

func _BinaryProtocol_Decode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BinaryProtocolServer).Decode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ppbinaryprotocol.BinaryProtocol/Decode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BinaryProtocolServer).Decode(ctx, req.(*DecodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BinaryProtocol_Encode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EncodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BinaryProtocolServer).Encode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ppbinaryprotocol.BinaryProtocol/Encode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BinaryProtocolServer).Encode(ctx, req.(*EncodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BinaryProtocol_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ppbinaryprotocol.BinaryProtocol",
	HandlerType: (*BinaryProtocolServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Decode",
			Handler:    _BinaryProtocol_Decode_Handler,
		},
		{
			MethodName: "Encode",
			Handler:    _BinaryProtocol_Encode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ppbinaryprotocol.proto",
}
