// Code generated by protoc-gen-go. DO NOT EDIT.
// source: city.proto

package city

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

type CityRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Country              string   `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CityRequest) Reset()         { *m = CityRequest{} }
func (m *CityRequest) String() string { return proto.CompactTextString(m) }
func (*CityRequest) ProtoMessage()    {}
func (*CityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f63625ed15d5bd67, []int{0}
}

func (m *CityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CityRequest.Unmarshal(m, b)
}
func (m *CityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CityRequest.Marshal(b, m, deterministic)
}
func (m *CityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CityRequest.Merge(m, src)
}
func (m *CityRequest) XXX_Size() int {
	return xxx_messageInfo_CityRequest.Size(m)
}
func (m *CityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CityRequest proto.InternalMessageInfo

func (m *CityRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CityRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CityRequest) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

type City struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	FindName             string   `protobuf:"bytes,3,opt,name=findName,proto3" json:"findName,omitempty"`
	Country              string   `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	Lat                  float64  `protobuf:"fixed64,5,opt,name=lat,proto3" json:"lat,omitempty"`
	Lon                  float64  `protobuf:"fixed64,6,opt,name=lon,proto3" json:"lon,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *City) Reset()         { *m = City{} }
func (m *City) String() string { return proto.CompactTextString(m) }
func (*City) ProtoMessage()    {}
func (*City) Descriptor() ([]byte, []int) {
	return fileDescriptor_f63625ed15d5bd67, []int{1}
}

func (m *City) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_City.Unmarshal(m, b)
}
func (m *City) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_City.Marshal(b, m, deterministic)
}
func (m *City) XXX_Merge(src proto.Message) {
	xxx_messageInfo_City.Merge(m, src)
}
func (m *City) XXX_Size() int {
	return xxx_messageInfo_City.Size(m)
}
func (m *City) XXX_DiscardUnknown() {
	xxx_messageInfo_City.DiscardUnknown(m)
}

var xxx_messageInfo_City proto.InternalMessageInfo

func (m *City) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *City) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *City) GetFindName() string {
	if m != nil {
		return m.FindName
	}
	return ""
}

func (m *City) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *City) GetLat() float64 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *City) GetLon() float64 {
	if m != nil {
		return m.Lon
	}
	return 0
}

func init() {
	proto.RegisterType((*CityRequest)(nil), "city.CityRequest")
	proto.RegisterType((*City)(nil), "city.City")
}

func init() {
	proto.RegisterFile("city.proto", fileDescriptor_f63625ed15d5bd67)
}

var fileDescriptor_f63625ed15d5bd67 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xce, 0x2c, 0xa9,
	0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0xbc, 0xb9, 0xb8, 0x9d, 0x33,
	0x4b, 0x2a, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24,
	0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0x98, 0x32, 0x53, 0x84, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73,
	0x53, 0x25, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x21, 0x09, 0x2e, 0xf6, 0xe4, 0xfc,
	0xd2, 0xbc, 0x92, 0xa2, 0x4a, 0x09, 0x66, 0xb0, 0x30, 0x8c, 0xab, 0xd4, 0xc2, 0xc8, 0xc5, 0x02,
	0x32, 0x0d, 0xc9, 0x18, 0x66, 0x9c, 0xc6, 0x48, 0x71, 0x71, 0xa4, 0x65, 0xe6, 0xa5, 0xf8, 0x81,
	0xc4, 0x21, 0xe6, 0xc0, 0xf9, 0xc8, 0x56, 0xb0, 0xa0, 0x58, 0x21, 0x24, 0xc0, 0xc5, 0x9c, 0x93,
	0x58, 0x22, 0xc1, 0xaa, 0xc0, 0xa8, 0xc1, 0x18, 0x04, 0x62, 0x82, 0x45, 0xf2, 0xf3, 0x24, 0xd8,
	0xa0, 0x22, 0xf9, 0x79, 0x46, 0x59, 0x10, 0x3f, 0x05, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a,
	0x69, 0x71, 0xb1, 0xbb, 0xa7, 0x96, 0x80, 0xdd, 0x25, 0xa8, 0x07, 0x0e, 0x00, 0x24, 0x1f, 0x4b,
	0x71, 0x21, 0x84, 0x94, 0x18, 0x84, 0xf4, 0xb9, 0xb8, 0x7c, 0x32, 0x8b, 0x41, 0x8a, 0x33, 0x53,
	0x8b, 0x09, 0x2a, 0x37, 0x60, 0x4c, 0x62, 0x03, 0x07, 0xa6, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff,
	0xdc, 0xfe, 0x53, 0x4e, 0x5a, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CityServiceClient is the client API for CityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CityServiceClient interface {
	GetCity(ctx context.Context, in *CityRequest, opts ...grpc.CallOption) (*City, error)
	ListCities(ctx context.Context, in *CityRequest, opts ...grpc.CallOption) (CityService_ListCitiesClient, error)
}

type cityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCityServiceClient(cc grpc.ClientConnInterface) CityServiceClient {
	return &cityServiceClient{cc}
}

func (c *cityServiceClient) GetCity(ctx context.Context, in *CityRequest, opts ...grpc.CallOption) (*City, error) {
	out := new(City)
	err := c.cc.Invoke(ctx, "/city.CityService/GetCity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cityServiceClient) ListCities(ctx context.Context, in *CityRequest, opts ...grpc.CallOption) (CityService_ListCitiesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CityService_serviceDesc.Streams[0], "/city.CityService/ListCities", opts...)
	if err != nil {
		return nil, err
	}
	x := &cityServiceListCitiesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CityService_ListCitiesClient interface {
	Recv() (*City, error)
	grpc.ClientStream
}

type cityServiceListCitiesClient struct {
	grpc.ClientStream
}

func (x *cityServiceListCitiesClient) Recv() (*City, error) {
	m := new(City)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CityServiceServer is the server API for CityService service.
type CityServiceServer interface {
	GetCity(context.Context, *CityRequest) (*City, error)
	ListCities(*CityRequest, CityService_ListCitiesServer) error
}

// UnimplementedCityServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCityServiceServer struct {
}

func (*UnimplementedCityServiceServer) GetCity(ctx context.Context, req *CityRequest) (*City, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCity not implemented")
}
func (*UnimplementedCityServiceServer) ListCities(req *CityRequest, srv CityService_ListCitiesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListCities not implemented")
}

func RegisterCityServiceServer(s *grpc.Server, srv CityServiceServer) {
	s.RegisterService(&_CityService_serviceDesc, srv)
}

func _CityService_GetCity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CityServiceServer).GetCity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/city.CityService/GetCity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CityServiceServer).GetCity(ctx, req.(*CityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CityService_ListCities_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CityRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CityServiceServer).ListCities(m, &cityServiceListCitiesServer{stream})
}

type CityService_ListCitiesServer interface {
	Send(*City) error
	grpc.ServerStream
}

type cityServiceListCitiesServer struct {
	grpc.ServerStream
}

func (x *cityServiceListCitiesServer) Send(m *City) error {
	return x.ServerStream.SendMsg(m)
}

var _CityService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "city.CityService",
	HandlerType: (*CityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCity",
			Handler:    _CityService_GetCity_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListCities",
			Handler:       _CityService_ListCities_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "city.proto",
}
