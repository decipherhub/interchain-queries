// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc_query/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_go "github.com/confio/ics23/go"
	types "github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgSubmitCrossChainQuery
type MsgSubmitCrossChainQuery struct {
	Path               string       `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	LocalTimeoutHeight types.Height `protobuf:"bytes,2,opt,name=local_timeout_height,json=localTimeoutHeight,proto3" json:"local_timeout_height" yaml:"timeout_height"`
	LocalTimeoutStamp  uint64       `protobuf:"varint,3,opt,name=local_timeout_stamp,json=localTimeoutStamp,proto3" json:"local_timeout_stamp,omitempty"`
	QueryHeight        uint64       `protobuf:"varint,4,opt,name=query_height,json=queryHeight,proto3" json:"query_height,omitempty"`
	ChainId            string       `protobuf:"bytes,5,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	Sender             string       `protobuf:"bytes,6,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (m *MsgSubmitCrossChainQuery) Reset()         { *m = MsgSubmitCrossChainQuery{} }
func (m *MsgSubmitCrossChainQuery) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitCrossChainQuery) ProtoMessage()    {}
func (*MsgSubmitCrossChainQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f4f91fbb08a4ad7, []int{0}
}
func (m *MsgSubmitCrossChainQuery) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitCrossChainQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitCrossChainQuery.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitCrossChainQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitCrossChainQuery.Merge(m, src)
}
func (m *MsgSubmitCrossChainQuery) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitCrossChainQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitCrossChainQuery.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitCrossChainQuery proto.InternalMessageInfo

// MsgSubmitCrossChainQueryResponse
type MsgSubmitCrossChainQueryResponse struct {
	QueryId string `protobuf:"bytes,1,opt,name=query_id,json=queryId,proto3" json:"query_id,omitempty"`
	CapKey  uint64 `protobuf:"varint,2,opt,name=cap_key,json=capKey,proto3" json:"cap_key,omitempty"`
}

func (m *MsgSubmitCrossChainQueryResponse) Reset()         { *m = MsgSubmitCrossChainQueryResponse{} }
func (m *MsgSubmitCrossChainQueryResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitCrossChainQueryResponse) ProtoMessage()    {}
func (*MsgSubmitCrossChainQueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f4f91fbb08a4ad7, []int{1}
}
func (m *MsgSubmitCrossChainQueryResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitCrossChainQueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitCrossChainQueryResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitCrossChainQueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitCrossChainQueryResponse.Merge(m, src)
}
func (m *MsgSubmitCrossChainQueryResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitCrossChainQueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitCrossChainQueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitCrossChainQueryResponse proto.InternalMessageInfo

func (m *MsgSubmitCrossChainQueryResponse) GetQueryId() string {
	if m != nil {
		return m.QueryId
	}
	return ""
}

func (m *MsgSubmitCrossChainQueryResponse) GetCapKey() uint64 {
	if m != nil {
		return m.CapKey
	}
	return 0
}

// MsgSubmitCrossChainQueryResult
type MsgSubmitCrossChainQueryResult struct {
	Id          string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Path        string           `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	QueryHeight uint64           `protobuf:"varint,3,opt,name=query_height,json=queryHeight,proto3" json:"query_height,omitempty"`
	Result      QueryResult      `protobuf:"varint,4,opt,name=result,proto3,enum=ibc_query.v1.QueryResult" json:"result,omitempty"`
	Data        []byte           `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	Sender      string           `protobuf:"bytes,6,opt,name=sender,proto3" json:"sender,omitempty"`
	ProofSpecs  []*_go.ProofSpec `protobuf:"bytes,7,rep,name=proof_specs,json=proofSpecs,proto3" json:"proof_specs,omitempty"`
}

func (m *MsgSubmitCrossChainQueryResult) Reset()         { *m = MsgSubmitCrossChainQueryResult{} }
func (m *MsgSubmitCrossChainQueryResult) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitCrossChainQueryResult) ProtoMessage()    {}
func (*MsgSubmitCrossChainQueryResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f4f91fbb08a4ad7, []int{2}
}
func (m *MsgSubmitCrossChainQueryResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitCrossChainQueryResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitCrossChainQueryResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitCrossChainQueryResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitCrossChainQueryResult.Merge(m, src)
}
func (m *MsgSubmitCrossChainQueryResult) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitCrossChainQueryResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitCrossChainQueryResult.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitCrossChainQueryResult proto.InternalMessageInfo

// MsgSubmitCrossChainQueryResultResponse
type MsgSubmitCrossChainQueryResultResponse struct {
}

func (m *MsgSubmitCrossChainQueryResultResponse) Reset() {
	*m = MsgSubmitCrossChainQueryResultResponse{}
}
func (m *MsgSubmitCrossChainQueryResultResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitCrossChainQueryResultResponse) ProtoMessage()    {}
func (*MsgSubmitCrossChainQueryResultResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f4f91fbb08a4ad7, []int{3}
}
func (m *MsgSubmitCrossChainQueryResultResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitCrossChainQueryResultResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitCrossChainQueryResultResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitCrossChainQueryResultResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitCrossChainQueryResultResponse.Merge(m, src)
}
func (m *MsgSubmitCrossChainQueryResultResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitCrossChainQueryResultResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitCrossChainQueryResultResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitCrossChainQueryResultResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSubmitCrossChainQuery)(nil), "ibc_query.v1.MsgSubmitCrossChainQuery")
	proto.RegisterType((*MsgSubmitCrossChainQueryResponse)(nil), "ibc_query.v1.MsgSubmitCrossChainQueryResponse")
	proto.RegisterType((*MsgSubmitCrossChainQueryResult)(nil), "ibc_query.v1.MsgSubmitCrossChainQueryResult")
	proto.RegisterType((*MsgSubmitCrossChainQueryResultResponse)(nil), "ibc_query.v1.MsgSubmitCrossChainQueryResultResponse")
}

func init() { proto.RegisterFile("ibc_query/v1/tx.proto", fileDescriptor_4f4f91fbb08a4ad7) }

var fileDescriptor_4f4f91fbb08a4ad7 = []byte{
	// 578 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xb6, 0x9d, 0x90, 0x94, 0x4b, 0x54, 0xc1, 0xd1, 0x82, 0x1b, 0x84, 0x13, 0x3c, 0x54, 0x19,
	0xe0, 0xac, 0xa4, 0x88, 0xa1, 0x63, 0xbb, 0x50, 0xa1, 0x4a, 0xc5, 0x45, 0x0c, 0x2c, 0x96, 0x7d,
	0x3e, 0x9c, 0x53, 0xed, 0x9c, 0xf1, 0x9d, 0xa3, 0x66, 0x44, 0x2c, 0x8c, 0xfc, 0x84, 0xee, 0xfc,
	0x91, 0x8e, 0x1d, 0x99, 0x2a, 0x94, 0x2c, 0xcc, 0x8c, 0x4c, 0xe8, 0xce, 0x4e, 0x48, 0x80, 0x44,
	0x65, 0x7b, 0xef, 0xde, 0xf7, 0xbe, 0xf7, 0xde, 0xf7, 0x9e, 0x0e, 0x6c, 0xd3, 0x00, 0x7b, 0xef,
	0x73, 0x92, 0x8d, 0x9d, 0x51, 0xcf, 0x11, 0xe7, 0x28, 0xcd, 0x98, 0x60, 0xb0, 0x39, 0x7f, 0x46,
	0xa3, 0x5e, 0x6b, 0x2b, 0x62, 0x11, 0x53, 0x01, 0x47, 0x5a, 0x05, 0xa6, 0xd5, 0xa6, 0x01, 0x76,
	0x30, 0xcb, 0x88, 0x83, 0x63, 0x4a, 0x86, 0x42, 0x12, 0x14, 0x56, 0x09, 0xb0, 0x97, 0xb8, 0x71,
	0xc6, 0x38, 0xc7, 0x03, 0x9f, 0x0e, 0x0b, 0xde, 0x02, 0xd3, 0x4c, 0x33, 0xc6, 0xde, 0xf1, 0xc2,
	0xb3, 0xbf, 0x18, 0xc0, 0x3c, 0xe6, 0xd1, 0x69, 0x1e, 0x24, 0x54, 0x1c, 0xca, 0x84, 0x43, 0x99,
	0xf0, 0x4a, 0x26, 0x40, 0x08, 0xaa, 0xa9, 0x2f, 0x06, 0xa6, 0xde, 0xd1, 0xbb, 0xb7, 0x5d, 0x65,
	0xc3, 0x33, 0xb0, 0x15, 0x33, 0xec, 0xc7, 0x9e, 0xa0, 0x09, 0x61, 0xb9, 0xf0, 0x06, 0x84, 0x46,
	0x03, 0x61, 0x1a, 0x1d, 0xbd, 0xdb, 0xe8, 0xb7, 0x10, 0x0d, 0x30, 0x92, 0x2d, 0xa2, 0xb2, 0xb1,
	0x51, 0x0f, 0xbd, 0x50, 0x88, 0x83, 0x47, 0x97, 0xd7, 0x6d, 0xed, 0xc7, 0x75, 0x7b, 0x7b, 0xec,
	0x27, 0xf1, 0xbe, 0xbd, 0x9c, 0x6f, 0xbb, 0x50, 0xd1, 0xbe, 0x2e, 0x5e, 0x8b, 0x14, 0x88, 0xc0,
	0xbd, 0xe5, 0x62, 0x5c, 0xf8, 0x49, 0x6a, 0x56, 0x3a, 0x7a, 0xb7, 0xea, 0xde, 0x5d, 0x4c, 0x38,
	0x95, 0x01, 0xf8, 0x18, 0x34, 0xd5, 0xa8, 0xb3, 0xa6, 0xaa, 0x0a, 0xd8, 0x50, 0x6f, 0x25, 0xe5,
	0x0e, 0xd8, 0x50, 0x92, 0x78, 0x34, 0x34, 0x6f, 0xa9, 0xb9, 0xea, 0xca, 0x3f, 0x0a, 0xe1, 0x7d,
	0x50, 0xe3, 0x64, 0x18, 0x92, 0xcc, 0xac, 0xa9, 0x40, 0xe9, 0xed, 0x6f, 0x7c, 0xba, 0x68, 0x6b,
	0xdf, 0x2f, 0xda, 0x9a, 0xfd, 0x06, 0x74, 0x56, 0x89, 0xe5, 0x12, 0x9e, 0xb2, 0x21, 0x27, 0xb2,
	0x40, 0xd1, 0x03, 0x0d, 0x4b, 0xe1, 0xea, 0xca, 0x3f, 0x0a, 0xe1, 0x03, 0x50, 0xc7, 0x7e, 0xea,
	0x9d, 0x91, 0xb1, 0x92, 0xab, 0xea, 0xd6, 0xb0, 0x9f, 0xbe, 0x24, 0x63, 0xfb, 0xa3, 0x01, 0xac,
	0x35, 0xc4, 0x79, 0x2c, 0xe0, 0x26, 0x30, 0xe6, 0x84, 0x06, 0x0d, 0xe7, 0xbb, 0x31, 0x16, 0x76,
	0xf3, 0xe7, 0xf8, 0x95, 0xbf, 0xc7, 0xef, 0x81, 0x5a, 0xa6, 0x08, 0x95, 0x36, 0x9b, 0xfd, 0x1d,
	0xb4, 0x78, 0x77, 0x68, 0xa1, 0xa2, 0x5b, 0x02, 0x65, 0xa5, 0xd0, 0x17, 0xbe, 0x52, 0xab, 0xe9,
	0x2a, 0x7b, 0x95, 0x54, 0xb0, 0x07, 0x1a, 0xea, 0xbc, 0x3c, 0x9e, 0x12, 0xcc, 0xcd, 0x7a, 0xa7,
	0xd2, 0x6d, 0xf4, 0xef, 0x20, 0x8a, 0x79, 0x7f, 0x0f, 0x9d, 0xc8, 0xc8, 0x69, 0x4a, 0xb0, 0x0b,
	0xd2, 0x99, 0xc9, 0x17, 0xd4, 0xed, 0x82, 0xdd, 0xf5, 0x22, 0xcc, 0x34, 0xee, 0xff, 0xd4, 0x41,
	0xe5, 0x98, 0x47, 0x90, 0x81, 0xed, 0x7f, 0x5f, 0xee, 0xee, 0xf2, 0x58, 0xab, 0x68, 0x5b, 0xe8,
	0x66, 0xb8, 0xf9, 0x72, 0x3f, 0xe8, 0xe0, 0xe1, 0xba, 0x2d, 0x3d, 0xb9, 0x31, 0x5f, 0x1e, 0x8b,
	0xd6, 0xb3, 0xff, 0x41, 0xcf, 0x7a, 0x38, 0x38, 0xb9, 0x9c, 0x58, 0xfa, 0xd5, 0xc4, 0xd2, 0xbf,
	0x4d, 0x2c, 0xfd, 0xf3, 0xd4, 0xd2, 0xae, 0xa6, 0x96, 0xf6, 0x75, 0x6a, 0x69, 0x6f, 0x9f, 0x47,
	0x54, 0x0c, 0xf2, 0x00, 0x61, 0x96, 0x38, 0x98, 0xf1, 0x84, 0x71, 0x87, 0x0e, 0x05, 0xc9, 0xd4,
	0x81, 0x3f, 0x95, 0x75, 0x28, 0xe1, 0xce, 0xb9, 0xf3, 0xfb, 0x97, 0x10, 0xe3, 0x94, 0xf0, 0xa0,
	0xa6, 0xfe, 0x82, 0xbd, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x03, 0x51, 0x3a, 0x32, 0x9b, 0x04,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// submit query request
	SubmitCrossChainQuery(ctx context.Context, in *MsgSubmitCrossChainQuery, opts ...grpc.CallOption) (*MsgSubmitCrossChainQueryResponse, error)
	// submit query result
	SubmitCrossChainQueryResult(ctx context.Context, in *MsgSubmitCrossChainQueryResult, opts ...grpc.CallOption) (*MsgSubmitCrossChainQueryResultResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SubmitCrossChainQuery(ctx context.Context, in *MsgSubmitCrossChainQuery, opts ...grpc.CallOption) (*MsgSubmitCrossChainQueryResponse, error) {
	out := new(MsgSubmitCrossChainQueryResponse)
	err := c.cc.Invoke(ctx, "/ibc_query.v1.Msg/SubmitCrossChainQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SubmitCrossChainQueryResult(ctx context.Context, in *MsgSubmitCrossChainQueryResult, opts ...grpc.CallOption) (*MsgSubmitCrossChainQueryResultResponse, error) {
	out := new(MsgSubmitCrossChainQueryResultResponse)
	err := c.cc.Invoke(ctx, "/ibc_query.v1.Msg/SubmitCrossChainQueryResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// submit query request
	SubmitCrossChainQuery(context.Context, *MsgSubmitCrossChainQuery) (*MsgSubmitCrossChainQueryResponse, error)
	// submit query result
	SubmitCrossChainQueryResult(context.Context, *MsgSubmitCrossChainQueryResult) (*MsgSubmitCrossChainQueryResultResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SubmitCrossChainQuery(ctx context.Context, req *MsgSubmitCrossChainQuery) (*MsgSubmitCrossChainQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitCrossChainQuery not implemented")
}
func (*UnimplementedMsgServer) SubmitCrossChainQueryResult(ctx context.Context, req *MsgSubmitCrossChainQueryResult) (*MsgSubmitCrossChainQueryResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitCrossChainQueryResult not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SubmitCrossChainQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitCrossChainQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitCrossChainQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc_query.v1.Msg/SubmitCrossChainQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitCrossChainQuery(ctx, req.(*MsgSubmitCrossChainQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SubmitCrossChainQueryResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitCrossChainQueryResult)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitCrossChainQueryResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc_query.v1.Msg/SubmitCrossChainQueryResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitCrossChainQueryResult(ctx, req.(*MsgSubmitCrossChainQueryResult))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ibc_query.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitCrossChainQuery",
			Handler:    _Msg_SubmitCrossChainQuery_Handler,
		},
		{
			MethodName: "SubmitCrossChainQueryResult",
			Handler:    _Msg_SubmitCrossChainQueryResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ibc_query/v1/tx.proto",
}

func (m *MsgSubmitCrossChainQuery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitCrossChainQuery) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitCrossChainQuery) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0x2a
	}
	if m.QueryHeight != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.QueryHeight))
		i--
		dAtA[i] = 0x20
	}
	if m.LocalTimeoutStamp != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.LocalTimeoutStamp))
		i--
		dAtA[i] = 0x18
	}
	{
		size, err := m.LocalTimeoutHeight.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Path) > 0 {
		i -= len(m.Path)
		copy(dAtA[i:], m.Path)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Path)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSubmitCrossChainQueryResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitCrossChainQueryResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitCrossChainQueryResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CapKey != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.CapKey))
		i--
		dAtA[i] = 0x10
	}
	if len(m.QueryId) > 0 {
		i -= len(m.QueryId)
		copy(dAtA[i:], m.QueryId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.QueryId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSubmitCrossChainQueryResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitCrossChainQueryResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitCrossChainQueryResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ProofSpecs) > 0 {
		for iNdEx := len(m.ProofSpecs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ProofSpecs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Result != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Result))
		i--
		dAtA[i] = 0x20
	}
	if m.QueryHeight != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.QueryHeight))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Path) > 0 {
		i -= len(m.Path)
		copy(dAtA[i:], m.Path)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Path)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSubmitCrossChainQueryResultResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitCrossChainQueryResultResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitCrossChainQueryResultResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSubmitCrossChainQuery) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.LocalTimeoutHeight.Size()
	n += 1 + l + sovTx(uint64(l))
	if m.LocalTimeoutStamp != 0 {
		n += 1 + sovTx(uint64(m.LocalTimeoutStamp))
	}
	if m.QueryHeight != 0 {
		n += 1 + sovTx(uint64(m.QueryHeight))
	}
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSubmitCrossChainQueryResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.QueryId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.CapKey != 0 {
		n += 1 + sovTx(uint64(m.CapKey))
	}
	return n
}

func (m *MsgSubmitCrossChainQueryResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.QueryHeight != 0 {
		n += 1 + sovTx(uint64(m.QueryHeight))
	}
	if m.Result != 0 {
		n += 1 + sovTx(uint64(m.Result))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.ProofSpecs) > 0 {
		for _, e := range m.ProofSpecs {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgSubmitCrossChainQueryResultResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSubmitCrossChainQuery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSubmitCrossChainQuery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitCrossChainQuery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LocalTimeoutHeight", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LocalTimeoutHeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LocalTimeoutStamp", wireType)
			}
			m.LocalTimeoutStamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LocalTimeoutStamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryHeight", wireType)
			}
			m.QueryHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.QueryHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSubmitCrossChainQueryResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSubmitCrossChainQueryResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitCrossChainQueryResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QueryId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CapKey", wireType)
			}
			m.CapKey = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CapKey |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSubmitCrossChainQueryResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSubmitCrossChainQueryResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitCrossChainQueryResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryHeight", wireType)
			}
			m.QueryHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.QueryHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			m.Result = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Result |= QueryResult(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProofSpecs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProofSpecs = append(m.ProofSpecs, &_go.ProofSpec{})
			if err := m.ProofSpecs[len(m.ProofSpecs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSubmitCrossChainQueryResultResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSubmitCrossChainQueryResultResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitCrossChainQueryResultResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
