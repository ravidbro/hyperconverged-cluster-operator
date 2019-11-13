// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/user_location_view_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Request message for [UserLocationViewService.GetUserLocationView][google.ads.googleads.v2.services.UserLocationViewService.GetUserLocationView].
type GetUserLocationViewRequest struct {
	// The resource name of the user location view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserLocationViewRequest) Reset()         { *m = GetUserLocationViewRequest{} }
func (m *GetUserLocationViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserLocationViewRequest) ProtoMessage()    {}
func (*GetUserLocationViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ceb1a60c4f1c0c60, []int{0}
}

func (m *GetUserLocationViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserLocationViewRequest.Unmarshal(m, b)
}
func (m *GetUserLocationViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserLocationViewRequest.Marshal(b, m, deterministic)
}
func (m *GetUserLocationViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserLocationViewRequest.Merge(m, src)
}
func (m *GetUserLocationViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserLocationViewRequest.Size(m)
}
func (m *GetUserLocationViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserLocationViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserLocationViewRequest proto.InternalMessageInfo

func (m *GetUserLocationViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetUserLocationViewRequest)(nil), "google.ads.googleads.v2.services.GetUserLocationViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/user_location_view_service.proto", fileDescriptor_ceb1a60c4f1c0c60)
}

var fileDescriptor_ceb1a60c4f1c0c60 = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x3f, 0xeb, 0xd3, 0x40,
	0x18, 0x26, 0x11, 0x04, 0x83, 0x2e, 0x71, 0x68, 0x89, 0x1d, 0x4a, 0xed, 0x20, 0x1d, 0xee, 0x20,
	0xc5, 0xe5, 0xaa, 0xc3, 0x75, 0xa9, 0x83, 0x48, 0xa9, 0x98, 0x41, 0x02, 0x21, 0x26, 0x2f, 0x21,
	0x90, 0xe4, 0xea, 0xbd, 0x49, 0x3a, 0x88, 0x8b, 0x7e, 0x04, 0xbf, 0x81, 0xa3, 0xdf, 0xc3, 0xa5,
	0xab, 0x5f, 0xc1, 0xc9, 0x0f, 0x21, 0x3f, 0x92, 0xcb, 0xa5, 0xfd, 0x95, 0x86, 0x6e, 0x0f, 0xf7,
	0x3e, 0x7f, 0xee, 0x7d, 0xee, 0x2c, 0x9e, 0x08, 0x91, 0x64, 0x40, 0xc3, 0x18, 0xa9, 0x82, 0x0d,
	0xaa, 0x5d, 0x8a, 0x20, 0xeb, 0x34, 0x02, 0xa4, 0x15, 0x82, 0x0c, 0x32, 0x11, 0x85, 0x65, 0x2a,
	0x8a, 0xa0, 0x4e, 0xe1, 0x10, 0x74, 0x33, 0xb2, 0x97, 0xa2, 0x14, 0xf6, 0x54, 0xe9, 0x48, 0x18,
	0x23, 0xe9, 0x2d, 0x48, 0xed, 0x12, 0x6d, 0xe1, 0xb0, 0xa1, 0x10, 0x09, 0x28, 0x2a, 0x79, 0x3d,
	0x45, 0xb9, 0x3b, 0x13, 0xad, 0xdd, 0xa7, 0x34, 0x2c, 0x0a, 0x51, 0xb6, 0x0c, 0xec, 0xa6, 0xa3,
	0xb3, 0x69, 0x94, 0xa5, 0x50, 0x94, 0x6a, 0x30, 0xe3, 0x96, 0xb3, 0x81, 0xf2, 0x03, 0x82, 0x7c,
	0xdb, 0x99, 0x7a, 0x29, 0x1c, 0x76, 0xf0, 0xb9, 0x02, 0x2c, 0xed, 0xe7, 0xd6, 0x13, 0x1d, 0x1d,
	0x14, 0x61, 0x0e, 0x63, 0x63, 0x6a, 0xbc, 0x78, 0xb4, 0x7b, 0xac, 0x0f, 0xdf, 0x85, 0x39, 0xb8,
	0xff, 0x0d, 0x6b, 0x74, 0x69, 0xf0, 0x5e, 0xad, 0x64, 0xff, 0x36, 0xac, 0xa7, 0x57, 0xfc, 0xed,
	0x57, 0xe4, 0x56, 0x19, 0x64, 0xf8, 0x5a, 0xce, 0x72, 0x50, 0xdd, 0x17, 0x45, 0x2e, 0xb5, 0xb3,
	0xd5, 0xb7, 0x3f, 0x7f, 0x7f, 0x98, 0x2f, 0xed, 0x65, 0x53, 0xe8, 0x97, 0x7b, 0x6b, 0xbd, 0x8e,
	0x2a, 0x2c, 0x45, 0x0e, 0x12, 0xe9, 0xa2, 0x6d, 0xf8, 0x5c, 0x88, 0x74, 0xf1, 0xd5, 0x79, 0x76,
	0xe4, 0xe3, 0x53, 0x50, 0x87, 0xf6, 0x29, 0x92, 0x48, 0xe4, 0xeb, 0xef, 0xa6, 0x35, 0x8f, 0x44,
	0x7e, 0x73, 0xa5, 0xf5, 0x64, 0xa0, 0xa6, 0x6d, 0xf3, 0x14, 0x5b, 0xe3, 0xe3, 0x9b, 0xce, 0x21,
	0x11, 0x59, 0x58, 0x24, 0x44, 0xc8, 0x84, 0x26, 0x50, 0xb4, 0x0f, 0x45, 0x4f, 0x99, 0xc3, 0x7f,
	0x70, 0xa5, 0xc1, 0x4f, 0xf3, 0xc1, 0x86, 0xf3, 0x5f, 0xe6, 0x74, 0xa3, 0x0c, 0x79, 0x8c, 0x44,
	0xc1, 0x06, 0x79, 0x2e, 0xe9, 0x82, 0xf1, 0xa8, 0x29, 0x3e, 0x8f, 0xd1, 0xef, 0x29, 0xbe, 0xe7,
	0xfa, 0x9a, 0xf2, 0xcf, 0x9c, 0xab, 0x73, 0xc6, 0x78, 0x8c, 0x8c, 0xf5, 0x24, 0xc6, 0x3c, 0x97,
	0x31, 0x4d, 0xfb, 0xf4, 0xb0, 0xbd, 0xe7, 0xf2, 0x2e, 0x00, 0x00, 0xff, 0xff, 0x47, 0xd9, 0x3d,
	0x8d, 0x2a, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserLocationViewServiceClient is the client API for UserLocationViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserLocationViewServiceClient interface {
	// Returns the requested user location view in full detail.
	GetUserLocationView(ctx context.Context, in *GetUserLocationViewRequest, opts ...grpc.CallOption) (*resources.UserLocationView, error)
}

type userLocationViewServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserLocationViewServiceClient(cc *grpc.ClientConn) UserLocationViewServiceClient {
	return &userLocationViewServiceClient{cc}
}

func (c *userLocationViewServiceClient) GetUserLocationView(ctx context.Context, in *GetUserLocationViewRequest, opts ...grpc.CallOption) (*resources.UserLocationView, error) {
	out := new(resources.UserLocationView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.UserLocationViewService/GetUserLocationView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserLocationViewServiceServer is the server API for UserLocationViewService service.
type UserLocationViewServiceServer interface {
	// Returns the requested user location view in full detail.
	GetUserLocationView(context.Context, *GetUserLocationViewRequest) (*resources.UserLocationView, error)
}

// UnimplementedUserLocationViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserLocationViewServiceServer struct {
}

func (*UnimplementedUserLocationViewServiceServer) GetUserLocationView(ctx context.Context, req *GetUserLocationViewRequest) (*resources.UserLocationView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserLocationView not implemented")
}

func RegisterUserLocationViewServiceServer(s *grpc.Server, srv UserLocationViewServiceServer) {
	s.RegisterService(&_UserLocationViewService_serviceDesc, srv)
}

func _UserLocationViewService_GetUserLocationView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserLocationViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLocationViewServiceServer).GetUserLocationView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.UserLocationViewService/GetUserLocationView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLocationViewServiceServer).GetUserLocationView(ctx, req.(*GetUserLocationViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserLocationViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.UserLocationViewService",
	HandlerType: (*UserLocationViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserLocationView",
			Handler:    _UserLocationViewService_GetUserLocationView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/user_location_view_service.proto",
}