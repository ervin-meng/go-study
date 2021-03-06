// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.13.0
// source: study.proto

package protobuf

import (
	context "context"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Sys int32

const (
	Sys_Unknown Sys = 0
	Sys_IOS     Sys = 1
	Sys_Android Sys = 2
)

// Enum value maps for Sys.
var (
	Sys_name = map[int32]string{
		0: "Unknown",
		1: "IOS",
		2: "Android",
	}
	Sys_value = map[string]int32{
		"Unknown": 0,
		"IOS":     1,
		"Android": 2,
	}
)

func (x Sys) Enum() *Sys {
	p := new(Sys)
	*p = x
	return p
}

func (x Sys) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Sys) Descriptor() protoreflect.EnumDescriptor {
	return file_study_proto_enumTypes[0].Descriptor()
}

func (Sys) Type() protoreflect.EnumType {
	return &file_study_proto_enumTypes[0]
}

func (x Sys) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Sys.Descriptor instead.
func (Sys) EnumDescriptor() ([]byte, []int) {
	return file_study_proto_rawDescGZIP(), []int{0}
}

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_study_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_study_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_study_proto_rawDescGZIP(), []int{0}
}

func (x *Location) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Location) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type StudyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` //1 ???????????????
	Location *Location   `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	Path     []*Location `protobuf:"bytes,3,rep,name=path,proto3" json:"path,omitempty"`
	Sys      Sys         `protobuf:"varint,4,opt,name=sys,proto3,enum=Sys" json:"sys,omitempty"`
}

func (x *StudyRequest) Reset() {
	*x = StudyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_study_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StudyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StudyRequest) ProtoMessage() {}

func (x *StudyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_study_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StudyRequest.ProtoReflect.Descriptor instead.
func (*StudyRequest) Descriptor() ([]byte, []int) {
	return file_study_proto_rawDescGZIP(), []int{1}
}

func (x *StudyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StudyRequest) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *StudyRequest) GetPath() []*Location {
	if x != nil {
		return x.Path
	}
	return nil
}

func (x *StudyRequest) GetSys() Sys {
	if x != nil {
		return x.Sys
	}
	return Sys_Unknown
}

type StudyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sys int32 `protobuf:"varint,1,opt,name=sys,proto3" json:"sys,omitempty"`
	//map <string, string> mp = 2; //?????????string??? key ?????? ?????????string???value?????????????????????????????????????????????????????????map
	AddTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=addTime,proto3" json:"addTime,omitempty"`
}

func (x *StudyResponse) Reset() {
	*x = StudyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_study_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StudyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StudyResponse) ProtoMessage() {}

func (x *StudyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_study_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StudyResponse.ProtoReflect.Descriptor instead.
func (*StudyResponse) Descriptor() ([]byte, []int) {
	return file_study_proto_rawDescGZIP(), []int{2}
}

func (x *StudyResponse) GetSys() int32 {
	if x != nil {
		return x.Sys
	}
	return 0
}

func (x *StudyResponse) GetAddTime() *timestamp.Timestamp {
	if x != nil {
		return x.AddTime
	}
	return nil
}

var File_study_proto protoreflect.FileDescriptor

var file_study_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x74, 0x75, 0x64, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44,
	0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61,
	0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61,
	0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x22, 0x80, 0x01, 0x0a, 0x0c, 0x53, 0x74, 0x75, 0x64, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1d, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09,
	0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12,
	0x16, 0x0a, 0x03, 0x73, 0x79, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x04, 0x2e, 0x53,
	0x79, 0x73, 0x52, 0x03, 0x73, 0x79, 0x73, 0x22, 0x57, 0x0a, 0x0d, 0x53, 0x74, 0x75, 0x64, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x79, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x73, 0x79, 0x73, 0x12, 0x34, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x61, 0x64, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x2a, 0x28, 0x0a, 0x03, 0x53, 0x79, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f,
	0x77, 0x6e, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x49, 0x4f, 0x53, 0x10, 0x01, 0x12, 0x0b, 0x0a,
	0x07, 0x41, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x10, 0x02, 0x32, 0x34, 0x0a, 0x0c, 0x53, 0x74,
	0x75, 0x64, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x03, 0x47, 0x65,
	0x74, 0x12, 0x0d, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0e, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_study_proto_rawDescOnce sync.Once
	file_study_proto_rawDescData = file_study_proto_rawDesc
)

func file_study_proto_rawDescGZIP() []byte {
	file_study_proto_rawDescOnce.Do(func() {
		file_study_proto_rawDescData = protoimpl.X.CompressGZIP(file_study_proto_rawDescData)
	})
	return file_study_proto_rawDescData
}

var file_study_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_study_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_study_proto_goTypes = []interface{}{
	(Sys)(0),                    // 0: Sys
	(*Location)(nil),            // 1: Location
	(*StudyRequest)(nil),        // 2: StudyRequest
	(*StudyResponse)(nil),       // 3: StudyResponse
	(*timestamp.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_study_proto_depIdxs = []int32{
	1, // 0: StudyRequest.location:type_name -> Location
	1, // 1: StudyRequest.path:type_name -> Location
	0, // 2: StudyRequest.sys:type_name -> Sys
	4, // 3: StudyResponse.addTime:type_name -> google.protobuf.Timestamp
	2, // 4: StudyService.Get:input_type -> StudyRequest
	3, // 5: StudyService.Get:output_type -> StudyResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_study_proto_init() }
func file_study_proto_init() {
	if File_study_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_study_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_study_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StudyRequest); i {
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
		file_study_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StudyResponse); i {
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
			RawDescriptor: file_study_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_study_proto_goTypes,
		DependencyIndexes: file_study_proto_depIdxs,
		EnumInfos:         file_study_proto_enumTypes,
		MessageInfos:      file_study_proto_msgTypes,
	}.Build()
	File_study_proto = out.File
	file_study_proto_rawDesc = nil
	file_study_proto_goTypes = nil
	file_study_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StudyServiceClient is the client API for StudyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StudyServiceClient interface {
	Get(ctx context.Context, in *StudyRequest, opts ...grpc.CallOption) (*StudyResponse, error)
}

type studyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStudyServiceClient(cc grpc.ClientConnInterface) StudyServiceClient {
	return &studyServiceClient{cc}
}

func (c *studyServiceClient) Get(ctx context.Context, in *StudyRequest, opts ...grpc.CallOption) (*StudyResponse, error) {
	out := new(StudyResponse)
	err := c.cc.Invoke(ctx, "/StudyService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StudyServiceServer is the server API for StudyService service.
type StudyServiceServer interface {
	Get(context.Context, *StudyRequest) (*StudyResponse, error)
}

// UnimplementedStudyServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStudyServiceServer struct {
}

func (*UnimplementedStudyServiceServer) Get(context.Context, *StudyRequest) (*StudyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterStudyServiceServer(s *grpc.Server, srv StudyServiceServer) {
	s.RegisterService(&_StudyService_serviceDesc, srv)
}

func _StudyService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StudyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudyServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/StudyService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudyServiceServer).Get(ctx, req.(*StudyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StudyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "StudyService",
	HandlerType: (*StudyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _StudyService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "study.proto",
}
