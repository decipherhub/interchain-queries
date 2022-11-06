// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/applications/ibc_query/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// QueryCrossChainQuery
type QueryCrossChainQuery struct {
	// query id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryCrossChainQuery) Reset()         { *m = QueryCrossChainQuery{} }
func (m *QueryCrossChainQuery) String() string { return proto.CompactTextString(m) }
func (*QueryCrossChainQuery) ProtoMessage()    {}
func (*QueryCrossChainQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_41a30807902c5755, []int{0}
}
func (m *QueryCrossChainQuery) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCrossChainQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCrossChainQuery.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCrossChainQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCrossChainQuery.Merge(m, src)
}
func (m *QueryCrossChainQuery) XXX_Size() int {
	return m.Size()
}
func (m *QueryCrossChainQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCrossChainQuery.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCrossChainQuery proto.InternalMessageInfo

func (m *QueryCrossChainQuery) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// QueryCrossChainQueryResponse
type QueryCrossChainQueryResponse struct {
	Result *CrossChainQuery `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (m *QueryCrossChainQueryResponse) Reset()         { *m = QueryCrossChainQueryResponse{} }
func (m *QueryCrossChainQueryResponse) String() string { return proto.CompactTextString(m) }
func (*QueryCrossChainQueryResponse) ProtoMessage()    {}
func (*QueryCrossChainQueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_41a30807902c5755, []int{1}
}
func (m *QueryCrossChainQueryResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCrossChainQueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCrossChainQueryResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCrossChainQueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCrossChainQueryResponse.Merge(m, src)
}
func (m *QueryCrossChainQueryResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryCrossChainQueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCrossChainQueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCrossChainQueryResponse proto.InternalMessageInfo

func (m *QueryCrossChainQueryResponse) GetResult() *CrossChainQuery {
	if m != nil {
		return m.Result
	}
	return nil
}

// QueryCrossChainQueryResult
type QueryCrossChainQueryResult struct {
	// query id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryCrossChainQueryResult) Reset()         { *m = QueryCrossChainQueryResult{} }
func (m *QueryCrossChainQueryResult) String() string { return proto.CompactTextString(m) }
func (*QueryCrossChainQueryResult) ProtoMessage()    {}
func (*QueryCrossChainQueryResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_41a30807902c5755, []int{2}
}
func (m *QueryCrossChainQueryResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCrossChainQueryResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCrossChainQueryResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCrossChainQueryResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCrossChainQueryResult.Merge(m, src)
}
func (m *QueryCrossChainQueryResult) XXX_Size() int {
	return m.Size()
}
func (m *QueryCrossChainQueryResult) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCrossChainQueryResult.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCrossChainQueryResult proto.InternalMessageInfo

func (m *QueryCrossChainQueryResult) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// QueryCrossChainQueryResultResponse
type QueryCrossChainQueryResultResponse struct {
	Id     string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Result QueryResult `protobuf:"varint,2,opt,name=result,proto3,enum=ibc.applications.ibc_query.v1.QueryResult" json:"result,omitempty"`
	Data   []byte      `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Sender string      `protobuf:"bytes,4,opt,name=sender,proto3" json:"sender,omitempty"`
}

func (m *QueryCrossChainQueryResultResponse) Reset()         { *m = QueryCrossChainQueryResultResponse{} }
func (m *QueryCrossChainQueryResultResponse) String() string { return proto.CompactTextString(m) }
func (*QueryCrossChainQueryResultResponse) ProtoMessage()    {}
func (*QueryCrossChainQueryResultResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_41a30807902c5755, []int{3}
}
func (m *QueryCrossChainQueryResultResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCrossChainQueryResultResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCrossChainQueryResultResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCrossChainQueryResultResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCrossChainQueryResultResponse.Merge(m, src)
}
func (m *QueryCrossChainQueryResultResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryCrossChainQueryResultResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCrossChainQueryResultResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCrossChainQueryResultResponse proto.InternalMessageInfo

func (m *QueryCrossChainQueryResultResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *QueryCrossChainQueryResultResponse) GetResult() QueryResult {
	if m != nil {
		return m.Result
	}
	return QueryResult_QUERY_RESULT_UNSPECIFIED
}

func (m *QueryCrossChainQueryResultResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *QueryCrossChainQueryResultResponse) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func init() {
	proto.RegisterType((*QueryCrossChainQuery)(nil), "ibc.applications.ibc_query.v1.QueryCrossChainQuery")
	proto.RegisterType((*QueryCrossChainQueryResponse)(nil), "ibc.applications.ibc_query.v1.QueryCrossChainQueryResponse")
	proto.RegisterType((*QueryCrossChainQueryResult)(nil), "ibc.applications.ibc_query.v1.QueryCrossChainQueryResult")
	proto.RegisterType((*QueryCrossChainQueryResultResponse)(nil), "ibc.applications.ibc_query.v1.QueryCrossChainQueryResultResponse")
}

func init() {
	proto.RegisterFile("ibc/applications/ibc_query/v1/query.proto", fileDescriptor_41a30807902c5755)
}

var fileDescriptor_41a30807902c5755 = []byte{
	// 417 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x4d, 0x8b, 0xd3, 0x40,
	0x18, 0xee, 0xc4, 0x5a, 0x70, 0x94, 0x0a, 0x83, 0x4a, 0x89, 0x6d, 0x28, 0x11, 0x24, 0x16, 0x9d,
	0xb1, 0x2d, 0x08, 0xe2, 0xc9, 0x16, 0x3c, 0x6b, 0x8f, 0x5e, 0x4a, 0x32, 0x19, 0xdb, 0x81, 0x34,
	0x13, 0x33, 0x93, 0x62, 0x11, 0x2f, 0xfe, 0x02, 0xc1, 0xbf, 0xb0, 0x87, 0xfd, 0x27, 0xbb, 0xc7,
	0xc2, 0x5e, 0xf6, 0xb8, 0xb4, 0xfb, 0x33, 0xf6, 0xb0, 0x64, 0x9a, 0x0d, 0x21, 0xdb, 0x76, 0x97,
	0xde, 0xde, 0x97, 0x79, 0x9e, 0xf7, 0x79, 0xe6, 0xfd, 0x80, 0x6f, 0xb8, 0x47, 0x89, 0x1b, 0x45,
	0x01, 0xa7, 0xae, 0xe2, 0x22, 0x94, 0x84, 0x7b, 0x74, 0xfc, 0x33, 0x61, 0xf1, 0x82, 0xcc, 0xbb,
	0x44, 0x07, 0x38, 0x8a, 0x85, 0x12, 0xa8, 0xc5, 0x3d, 0x8a, 0x8b, 0x50, 0x9c, 0x43, 0xf1, 0xbc,
	0x6b, 0x36, 0x27, 0x42, 0x4c, 0x02, 0x46, 0xdc, 0x88, 0x13, 0x37, 0x0c, 0x85, 0xca, 0x40, 0x9a,
	0x6c, 0xf6, 0xf7, 0xeb, 0xd0, 0x58, 0x48, 0x49, 0xa7, 0x2e, 0x0f, 0x0b, 0x8a, 0xf6, 0x6b, 0xf8,
	0xec, 0x5b, 0x9a, 0x0e, 0xd3, 0xd7, 0x61, 0xfa, 0xaa, 0x53, 0x54, 0x87, 0x06, 0xf7, 0x1b, 0xa0,
	0x0d, 0x9c, 0x47, 0x23, 0x83, 0xfb, 0xf6, 0x0f, 0xd8, 0xdc, 0x86, 0x1b, 0x31, 0x19, 0x89, 0x50,
	0x32, 0xf4, 0x05, 0xd6, 0x62, 0x26, 0x93, 0x40, 0x69, 0xce, 0xe3, 0x1e, 0xc6, 0x7b, 0xbf, 0x82,
	0xcb, 0x75, 0x32, 0xb6, 0xfd, 0x16, 0x9a, 0x3b, 0x74, 0x92, 0x40, 0xdd, 0x72, 0x75, 0x04, 0xa0,
	0xbd, 0x1b, 0x9e, 0x9b, 0x2b, 0xd1, 0xd0, 0x20, 0x37, 0x6b, 0xb4, 0x81, 0x53, 0xef, 0x75, 0xee,
	0x30, 0x5b, 0xac, 0x99, 0x31, 0x11, 0x82, 0x55, 0xdf, 0x55, 0x6e, 0xe3, 0x41, 0x1b, 0x38, 0x4f,
	0x46, 0x3a, 0x46, 0x2f, 0x60, 0x4d, 0xb2, 0xd0, 0x67, 0x71, 0xa3, 0xaa, 0xb5, 0xb2, 0xac, 0x77,
	0x65, 0xc0, 0x87, 0x9b, 0xb6, 0x1e, 0x03, 0xf8, 0xb4, 0xdc, 0xea, 0xfe, 0x7d, 0xd4, 0x4b, 0x24,
	0xf3, 0xd3, 0x01, 0xa4, 0x9b, 0x7e, 0xd8, 0xaf, 0xfe, 0x9e, 0x5d, 0xfe, 0x37, 0x5a, 0xe8, 0x25,
	0xc9, 0x56, 0xa6, 0xb4, 0x2a, 0xbf, 0xb9, 0xff, 0x07, 0x9d, 0x00, 0xf8, 0x7c, 0xfb, 0x14, 0x3e,
	0x1e, 0xa6, 0x9d, 0x04, 0xca, 0xfc, 0x7c, 0x30, 0x35, 0x37, 0xff, 0x5e, 0x9b, 0xef, 0x20, 0x67,
	0x87, 0x79, 0x1d, 0x8c, 0x37, 0x53, 0xd2, 0x3f, 0x19, 0x7c, 0x3d, 0x5d, 0x59, 0x60, 0xb9, 0xb2,
	0xc0, 0xc5, 0xca, 0x02, 0xff, 0xd6, 0x56, 0x65, 0xb9, 0xb6, 0x2a, 0xe7, 0x6b, 0xab, 0xf2, 0xfd,
	0xc3, 0x84, 0xab, 0x69, 0xe2, 0x61, 0x2a, 0x66, 0x84, 0x0a, 0x39, 0x13, 0x92, 0xf0, 0x50, 0xb1,
	0x58, 0x9f, 0xc9, 0xbb, 0xb4, 0x12, 0x67, 0x92, 0xfc, 0x2a, 0x08, 0xa8, 0x45, 0xc4, 0xa4, 0x57,
	0xd3, 0xc7, 0xd3, 0xbf, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x5a, 0x24, 0xae, 0xf8, 0xdb, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// query CrossChainQuery
	CrossChainQuery(ctx context.Context, in *QueryCrossChainQuery, opts ...grpc.CallOption) (*QueryCrossChainQueryResponse, error)
	// query CrossChainQueryResult
	CrossChainQueryResult(ctx context.Context, in *QueryCrossChainQueryResult, opts ...grpc.CallOption) (*QueryCrossChainQueryResultResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) CrossChainQuery(ctx context.Context, in *QueryCrossChainQuery, opts ...grpc.CallOption) (*QueryCrossChainQueryResponse, error) {
	out := new(QueryCrossChainQueryResponse)
	err := c.cc.Invoke(ctx, "/ibc.applications.ibc_query.v1.Query/CrossChainQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CrossChainQueryResult(ctx context.Context, in *QueryCrossChainQueryResult, opts ...grpc.CallOption) (*QueryCrossChainQueryResultResponse, error) {
	out := new(QueryCrossChainQueryResultResponse)
	err := c.cc.Invoke(ctx, "/ibc.applications.ibc_query.v1.Query/CrossChainQueryResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// query CrossChainQuery
	CrossChainQuery(context.Context, *QueryCrossChainQuery) (*QueryCrossChainQueryResponse, error)
	// query CrossChainQueryResult
	CrossChainQueryResult(context.Context, *QueryCrossChainQueryResult) (*QueryCrossChainQueryResultResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) CrossChainQuery(ctx context.Context, req *QueryCrossChainQuery) (*QueryCrossChainQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CrossChainQuery not implemented")
}
func (*UnimplementedQueryServer) CrossChainQueryResult(ctx context.Context, req *QueryCrossChainQueryResult) (*QueryCrossChainQueryResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CrossChainQueryResult not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_CrossChainQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCrossChainQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CrossChainQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.applications.ibc_query.v1.Query/CrossChainQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CrossChainQuery(ctx, req.(*QueryCrossChainQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CrossChainQueryResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCrossChainQueryResult)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CrossChainQueryResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.applications.ibc_query.v1.Query/CrossChainQueryResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CrossChainQueryResult(ctx, req.(*QueryCrossChainQueryResult))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ibc.applications.ibc_query.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CrossChainQuery",
			Handler:    _Query_CrossChainQuery_Handler,
		},
		{
			MethodName: "CrossChainQueryResult",
			Handler:    _Query_CrossChainQueryResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ibc/applications/ibc_query/v1/query.proto",
}

func (m *QueryCrossChainQuery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCrossChainQuery) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCrossChainQuery) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryCrossChainQueryResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCrossChainQueryResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCrossChainQueryResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Result != nil {
		{
			size, err := m.Result.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryCrossChainQueryResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCrossChainQueryResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCrossChainQueryResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryCrossChainQueryResultResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCrossChainQueryResultResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCrossChainQueryResultResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Result != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Result))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryCrossChainQuery) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryCrossChainQueryResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Result != nil {
		l = m.Result.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryCrossChainQueryResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryCrossChainQueryResultResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Result != 0 {
		n += 1 + sovQuery(uint64(m.Result))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryCrossChainQuery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryCrossChainQuery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCrossChainQuery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryCrossChainQueryResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryCrossChainQueryResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCrossChainQueryResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Result == nil {
				m.Result = &CrossChainQuery{}
			}
			if err := m.Result.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryCrossChainQueryResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryCrossChainQueryResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCrossChainQueryResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryCrossChainQueryResultResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryCrossChainQueryResultResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCrossChainQueryResultResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			m.Result = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
