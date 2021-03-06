// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stream.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	stream.proto

It has these top-level messages:
	StreamPoint
	StreamRequest
	StreamResponse
*/
package pb

import (
	fmt "fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

	context "golang.org/x/net/context"

	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StreamPoint struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value int32  `protobuf:"varint,2,opt,name=value" json:"value,omitempty"`
}

func (m *StreamPoint) Reset()                    { *m = StreamPoint{} }
func (m *StreamPoint) String() string            { return proto.CompactTextString(m) }
func (*StreamPoint) ProtoMessage()               {}
func (*StreamPoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StreamPoint) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StreamPoint) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type StreamRequest struct {
	Pt *StreamPoint `protobuf:"bytes,1,opt,name=pt" json:"pt,omitempty"`
}

func (m *StreamRequest) Reset()                    { *m = StreamRequest{} }
func (m *StreamRequest) String() string            { return proto.CompactTextString(m) }
func (*StreamRequest) ProtoMessage()               {}
func (*StreamRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *StreamRequest) GetPt() *StreamPoint {
	if m != nil {
		return m.Pt
	}
	return nil
}

type StreamResponse struct {
	Pt *StreamPoint `protobuf:"bytes,1,opt,name=pt" json:"pt,omitempty"`
}

func (m *StreamResponse) Reset()                    { *m = StreamResponse{} }
func (m *StreamResponse) String() string            { return proto.CompactTextString(m) }
func (*StreamResponse) ProtoMessage()               {}
func (*StreamResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *StreamResponse) GetPt() *StreamPoint {
	if m != nil {
		return m.Pt
	}
	return nil
}

func init() {
	proto.RegisterType((*StreamPoint)(nil), "pb.StreamPoint")
	proto.RegisterType((*StreamRequest)(nil), "pb.StreamRequest")
	proto.RegisterType((*StreamResponse)(nil), "pb.StreamResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for StreamService service

type StreamServiceClient interface {
	List(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (StreamService_ListClient, error)
	Record(ctx context.Context, opts ...grpc.CallOption) (StreamService_RecordClient, error)
	Route(ctx context.Context, opts ...grpc.CallOption) (StreamService_RouteClient, error)
}

type streamServiceClient struct {
	cc *grpc.ClientConn
}

func NewStreamServiceClient(cc *grpc.ClientConn) StreamServiceClient {
	return &streamServiceClient{cc}
}

func (c *streamServiceClient) List(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (StreamService_ListClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_StreamService_serviceDesc.Streams[0], c.cc, "/pb.StreamService/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamService_ListClient interface {
	Recv() (*StreamResponse, error)
	grpc.ClientStream
}

type streamServiceListClient struct {
	grpc.ClientStream
}

func (x *streamServiceListClient) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamServiceClient) Record(ctx context.Context, opts ...grpc.CallOption) (StreamService_RecordClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_StreamService_serviceDesc.Streams[1], c.cc, "/pb.StreamService/Record", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceRecordClient{stream}
	return x, nil
}

type StreamService_RecordClient interface {
	Send(*StreamRequest) error
	CloseAndRecv() (*StreamResponse, error)
	grpc.ClientStream
}

type streamServiceRecordClient struct {
	grpc.ClientStream
}

func (x *streamServiceRecordClient) Send(m *StreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamServiceRecordClient) CloseAndRecv() (*StreamResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamServiceClient) Route(ctx context.Context, opts ...grpc.CallOption) (StreamService_RouteClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_StreamService_serviceDesc.Streams[2], c.cc, "/pb.StreamService/Route", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceRouteClient{stream}
	return x, nil
}

type StreamService_RouteClient interface {
	Send(*StreamRequest) error
	Recv() (*StreamResponse, error)
	grpc.ClientStream
}

type streamServiceRouteClient struct {
	grpc.ClientStream
}

func (x *streamServiceRouteClient) Send(m *StreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamServiceRouteClient) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for StreamService service

type StreamServiceServer interface {
	List(*StreamRequest, StreamService_ListServer) error
	Record(StreamService_RecordServer) error
	Route(StreamService_RouteServer) error
}

func RegisterStreamServiceServer(s *grpc.Server, srv StreamServiceServer) {
	s.RegisterService(&_StreamService_serviceDesc, srv)
}

func _StreamService_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamServiceServer).List(m, &streamServiceListServer{stream})
}

type StreamService_ListServer interface {
	Send(*StreamResponse) error
	grpc.ServerStream
}

type streamServiceListServer struct {
	grpc.ServerStream
}

func (x *streamServiceListServer) Send(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _StreamService_Record_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServiceServer).Record(&streamServiceRecordServer{stream})
}

type StreamService_RecordServer interface {
	SendAndClose(*StreamResponse) error
	Recv() (*StreamRequest, error)
	grpc.ServerStream
}

type streamServiceRecordServer struct {
	grpc.ServerStream
}

func (x *streamServiceRecordServer) SendAndClose(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamServiceRecordServer) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _StreamService_Route_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServiceServer).Route(&streamServiceRouteServer{stream})
}

type StreamService_RouteServer interface {
	Send(*StreamResponse) error
	Recv() (*StreamRequest, error)
	grpc.ServerStream
}

type streamServiceRouteServer struct {
	grpc.ServerStream
}

func (x *streamServiceRouteServer) Send(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamServiceRouteServer) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _StreamService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.StreamService",
	HandlerType: (*StreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _StreamService_List_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Record",
			Handler:       _StreamService_Record_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Route",
			Handler:       _StreamService_Route_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "stream.proto",
}

func init() { proto.RegisterFile("stream.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0xd0, 0x4b, 0x4e, 0x80, 0x30,
	0x10, 0x06, 0x60, 0xdb, 0x00, 0x89, 0x83, 0x8f, 0x38, 0x71, 0x41, 0xdc, 0x48, 0xba, 0x62, 0x45,
	0x78, 0x98, 0x78, 0x09, 0x17, 0xa6, 0x9c, 0x00, 0x70, 0x16, 0x24, 0x42, 0x6b, 0x3b, 0x70, 0x2b,
	0xef, 0x68, 0x52, 0x24, 0xb2, 0x53, 0x77, 0xed, 0x9f, 0x7e, 0xfd, 0x27, 0x03, 0x57, 0x9e, 0x1d,
	0xf5, 0x73, 0x69, 0x9d, 0x61, 0x83, 0xd2, 0x0e, 0xea, 0x19, 0xd2, 0x2e, 0x64, 0xaf, 0x66, 0x5a,
	0x18, 0x11, 0xa2, 0xa5, 0x9f, 0x29, 0x13, 0xb9, 0x28, 0x2e, 0x75, 0x38, 0xe3, 0x3d, 0xc4, 0x5b,
	0xff, 0xbe, 0x52, 0x26, 0x73, 0x51, 0xc4, 0x7a, 0xbf, 0xa8, 0x0a, 0xae, 0x77, 0xa8, 0xe9, 0x63,
	0x25, 0xcf, 0xf8, 0x08, 0xd2, 0x72, 0x80, 0x69, 0x73, 0x5b, 0xda, 0xa1, 0x3c, 0xfd, 0xab, 0xa5,
	0x65, 0x55, 0xc3, 0xcd, 0x21, 0xbc, 0x35, 0x8b, 0xa7, 0x5f, 0x49, 0xf3, 0x29, 0x8e, 0x96, 0x8e,
	0xdc, 0x36, 0x8d, 0x84, 0x35, 0x44, 0x2f, 0x93, 0x67, 0xbc, 0xfb, 0x79, 0xfe, 0x3d, 0xc0, 0x03,
	0x9e, 0xa3, 0xbd, 0x41, 0x5d, 0x54, 0x02, 0x5b, 0x48, 0x34, 0x8d, 0xc6, 0xbd, 0xfd, 0x19, 0x15,
	0x02, 0x9f, 0x20, 0xd6, 0x66, 0x65, 0xfa, 0x87, 0xa9, 0xc4, 0x90, 0x84, 0xc5, 0xb6, 0x5f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x7d, 0xfd, 0xfd, 0xbc, 0x68, 0x01, 0x00, 0x00,
}
