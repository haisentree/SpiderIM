// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: msg_gateway.proto

package pbMsgGateway

import (
	context "context"
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

type SingleMsgReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SendID  uint64 `protobuf:"varint,1,opt,name=sendID,proto3" json:"sendID,omitempty"`
	RecvID  uint64 `protobuf:"varint,2,opt,name=recvID,proto3" json:"recvID,omitempty"`
	MsgType uint32 `protobuf:"varint,3,opt,name=msgType,proto3" json:"msgType,omitempty"`
	Content string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *SingleMsgReq) Reset() {
	*x = SingleMsgReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_gateway_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SingleMsgReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SingleMsgReq) ProtoMessage() {}

func (x *SingleMsgReq) ProtoReflect() protoreflect.Message {
	mi := &file_msg_gateway_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SingleMsgReq.ProtoReflect.Descriptor instead.
func (*SingleMsgReq) Descriptor() ([]byte, []int) {
	return file_msg_gateway_proto_rawDescGZIP(), []int{0}
}

func (x *SingleMsgReq) GetSendID() uint64 {
	if x != nil {
		return x.SendID
	}
	return 0
}

func (x *SingleMsgReq) GetRecvID() uint64 {
	if x != nil {
		return x.RecvID
	}
	return 0
}

func (x *SingleMsgReq) GetMsgType() uint32 {
	if x != nil {
		return x.MsgType
	}
	return 0
}

func (x *SingleMsgReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type SingleMsgResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SingleMsgResp) Reset() {
	*x = SingleMsgResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_gateway_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SingleMsgResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SingleMsgResp) ProtoMessage() {}

func (x *SingleMsgResp) ProtoReflect() protoreflect.Message {
	mi := &file_msg_gateway_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SingleMsgResp.ProtoReflect.Descriptor instead.
func (*SingleMsgResp) Descriptor() ([]byte, []int) {
	return file_msg_gateway_proto_rawDescGZIP(), []int{1}
}

func (x *SingleMsgResp) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SingleMsgResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ListMsgReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SendID  uint64   `protobuf:"varint,1,opt,name=sendID,proto3" json:"sendID,omitempty"`
	RecvID  []uint64 `protobuf:"varint,2,rep,packed,name=recvID,proto3" json:"recvID,omitempty"`
	MsgType uint32   `protobuf:"varint,3,opt,name=msgType,proto3" json:"msgType,omitempty"`
	SeqID   uint64   `protobuf:"varint,4,opt,name=seqID,proto3" json:"seqID,omitempty"`
	Content string   `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *ListMsgReq) Reset() {
	*x = ListMsgReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_gateway_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMsgReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMsgReq) ProtoMessage() {}

func (x *ListMsgReq) ProtoReflect() protoreflect.Message {
	mi := &file_msg_gateway_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMsgReq.ProtoReflect.Descriptor instead.
func (*ListMsgReq) Descriptor() ([]byte, []int) {
	return file_msg_gateway_proto_rawDescGZIP(), []int{2}
}

func (x *ListMsgReq) GetSendID() uint64 {
	if x != nil {
		return x.SendID
	}
	return 0
}

func (x *ListMsgReq) GetRecvID() []uint64 {
	if x != nil {
		return x.RecvID
	}
	return nil
}

func (x *ListMsgReq) GetMsgType() uint32 {
	if x != nil {
		return x.MsgType
	}
	return 0
}

func (x *ListMsgReq) GetSeqID() uint64 {
	if x != nil {
		return x.SeqID
	}
	return 0
}

func (x *ListMsgReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type ListMsgResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ListMsgResp) Reset() {
	*x = ListMsgResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_gateway_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMsgResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMsgResp) ProtoMessage() {}

func (x *ListMsgResp) ProtoReflect() protoreflect.Message {
	mi := &file_msg_gateway_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMsgResp.ProtoReflect.Descriptor instead.
func (*ListMsgResp) Descriptor() ([]byte, []int) {
	return file_msg_gateway_proto_rawDescGZIP(), []int{3}
}

func (x *ListMsgResp) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ListMsgResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_msg_gateway_proto protoreflect.FileDescriptor

var file_msg_gateway_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x73, 0x67, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x62, 0x4d, 0x73, 0x67, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x22, 0x72, 0x0a, 0x0c, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x4d, 0x73, 0x67, 0x52, 0x65,
	0x71, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x63,
	0x76, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x72, 0x65, 0x63, 0x76, 0x49,
	0x44, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x3d, 0x0a, 0x0d, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x4d,
	0x73, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x86, 0x01, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x73, 0x67,
	0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x63, 0x76, 0x49, 0x44, 0x18, 0x02, 0x20, 0x03, 0x28, 0x04, 0x52, 0x06, 0x72, 0x65, 0x63,
	0x76, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x65, 0x71, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x73, 0x65,
	0x71, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x3b, 0x0a,
	0x0b, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xa0, 0x01, 0x0a, 0x0a, 0x4d,
	0x73, 0x67, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x4b, 0x0a, 0x10, 0x52, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x4d, 0x73, 0x67, 0x12, 0x1a, 0x2e,
	0x70, 0x62, 0x4d, 0x73, 0x67, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x53, 0x69, 0x6e,
	0x67, 0x6c, 0x65, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x70, 0x62, 0x4d, 0x73,
	0x67, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x4d,
	0x73, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x45, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x4d, 0x73, 0x67,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x52,
	0x65, 0x71, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x4d, 0x73, 0x67, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x70, 0x42, 0x10, 0x5a,
	0x0e, 0x2e, 0x3b, 0x70, 0x62, 0x4d, 0x73, 0x67, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_msg_gateway_proto_rawDescOnce sync.Once
	file_msg_gateway_proto_rawDescData = file_msg_gateway_proto_rawDesc
)

func file_msg_gateway_proto_rawDescGZIP() []byte {
	file_msg_gateway_proto_rawDescOnce.Do(func() {
		file_msg_gateway_proto_rawDescData = protoimpl.X.CompressGZIP(file_msg_gateway_proto_rawDescData)
	})
	return file_msg_gateway_proto_rawDescData
}

var file_msg_gateway_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_msg_gateway_proto_goTypes = []interface{}{
	(*SingleMsgReq)(nil),  // 0: pbMsgGateway.SingleMsgReq
	(*SingleMsgResp)(nil), // 1: pbMsgGateway.SingleMsgResp
	(*ListMsgReq)(nil),    // 2: pbMsgGateway.ListMsgReq
	(*ListMsgResp)(nil),   // 3: pbMsgGateway.ListMsgResp
}
var file_msg_gateway_proto_depIdxs = []int32{
	0, // 0: pbMsgGateway.MsgGateway.ReceiveSingleMsg:input_type -> pbMsgGateway.SingleMsgReq
	2, // 1: pbMsgGateway.MsgGateway.ReceiveListMsg:input_type -> pbMsgGateway.ListMsgReq
	1, // 2: pbMsgGateway.MsgGateway.ReceiveSingleMsg:output_type -> pbMsgGateway.SingleMsgResp
	3, // 3: pbMsgGateway.MsgGateway.ReceiveListMsg:output_type -> pbMsgGateway.ListMsgResp
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_msg_gateway_proto_init() }
func file_msg_gateway_proto_init() {
	if File_msg_gateway_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_msg_gateway_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SingleMsgReq); i {
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
		file_msg_gateway_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SingleMsgResp); i {
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
		file_msg_gateway_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMsgReq); i {
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
		file_msg_gateway_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMsgResp); i {
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
			RawDescriptor: file_msg_gateway_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_msg_gateway_proto_goTypes,
		DependencyIndexes: file_msg_gateway_proto_depIdxs,
		MessageInfos:      file_msg_gateway_proto_msgTypes,
	}.Build()
	File_msg_gateway_proto = out.File
	file_msg_gateway_proto_rawDesc = nil
	file_msg_gateway_proto_goTypes = nil
	file_msg_gateway_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MsgGatewayClient is the client API for MsgGateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgGatewayClient interface {
	ReceiveSingleMsg(ctx context.Context, in *SingleMsgReq, opts ...grpc.CallOption) (*SingleMsgResp, error)
	ReceiveListMsg(ctx context.Context, in *ListMsgReq, opts ...grpc.CallOption) (*ListMsgResp, error)
}

type msgGatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgGatewayClient(cc grpc.ClientConnInterface) MsgGatewayClient {
	return &msgGatewayClient{cc}
}

func (c *msgGatewayClient) ReceiveSingleMsg(ctx context.Context, in *SingleMsgReq, opts ...grpc.CallOption) (*SingleMsgResp, error) {
	out := new(SingleMsgResp)
	err := c.cc.Invoke(ctx, "/pbMsgGateway.MsgGateway/ReceiveSingleMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgGatewayClient) ReceiveListMsg(ctx context.Context, in *ListMsgReq, opts ...grpc.CallOption) (*ListMsgResp, error) {
	out := new(ListMsgResp)
	err := c.cc.Invoke(ctx, "/pbMsgGateway.MsgGateway/ReceiveListMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgGatewayServer is the server API for MsgGateway service.
type MsgGatewayServer interface {
	ReceiveSingleMsg(context.Context, *SingleMsgReq) (*SingleMsgResp, error)
	ReceiveListMsg(context.Context, *ListMsgReq) (*ListMsgResp, error)
}

// UnimplementedMsgGatewayServer can be embedded to have forward compatible implementations.
type UnimplementedMsgGatewayServer struct {
}

func (*UnimplementedMsgGatewayServer) ReceiveSingleMsg(context.Context, *SingleMsgReq) (*SingleMsgResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveSingleMsg not implemented")
}
func (*UnimplementedMsgGatewayServer) ReceiveListMsg(context.Context, *ListMsgReq) (*ListMsgResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveListMsg not implemented")
}

func RegisterMsgGatewayServer(s *grpc.Server, srv MsgGatewayServer) {
	s.RegisterService(&_MsgGateway_serviceDesc, srv)
}

func _MsgGateway_ReceiveSingleMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgGatewayServer).ReceiveSingleMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbMsgGateway.MsgGateway/ReceiveSingleMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgGatewayServer).ReceiveSingleMsg(ctx, req.(*SingleMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgGateway_ReceiveListMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgGatewayServer).ReceiveListMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbMsgGateway.MsgGateway/ReceiveListMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgGatewayServer).ReceiveListMsg(ctx, req.(*ListMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _MsgGateway_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pbMsgGateway.MsgGateway",
	HandlerType: (*MsgGatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReceiveSingleMsg",
			Handler:    _MsgGateway_ReceiveSingleMsg_Handler,
		},
		{
			MethodName: "ReceiveListMsg",
			Handler:    _MsgGateway_ReceiveListMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "msg_gateway.proto",
}
