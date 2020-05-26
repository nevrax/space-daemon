// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.4
// source: space.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type PathInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *PathInfoRequest) Reset() {
	*x = PathInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_space_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PathInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PathInfoRequest) ProtoMessage() {}

func (x *PathInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_space_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PathInfoRequest.ProtoReflect.Descriptor instead.
func (*PathInfoRequest) Descriptor() ([]byte, []int) {
	return file_space_proto_rawDescGZIP(), []int{0}
}

func (x *PathInfoRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type PathInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path     string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	IpfsHash string `protobuf:"bytes,2,opt,name=ipfsHash,proto3" json:"ipfsHash,omitempty"`
	IsDir    bool   `protobuf:"varint,3,opt,name=isDir,proto3" json:"isDir,omitempty"`
}

func (x *PathInfoResponse) Reset() {
	*x = PathInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_space_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PathInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PathInfoResponse) ProtoMessage() {}

func (x *PathInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_space_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PathInfoResponse.ProtoReflect.Descriptor instead.
func (*PathInfoResponse) Descriptor() ([]byte, []int) {
	return file_space_proto_rawDescGZIP(), []int{1}
}

func (x *PathInfoResponse) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *PathInfoResponse) GetIpfsHash() string {
	if x != nil {
		return x.IpfsHash
	}
	return ""
}

func (x *PathInfoResponse) GetIsDir() bool {
	if x != nil {
		return x.IsDir
	}
	return false
}

type GenerateKeyPairRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GenerateKeyPairRequest) Reset() {
	*x = GenerateKeyPairRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_space_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateKeyPairRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateKeyPairRequest) ProtoMessage() {}

func (x *GenerateKeyPairRequest) ProtoReflect() protoreflect.Message {
	mi := &file_space_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateKeyPairRequest.ProtoReflect.Descriptor instead.
func (*GenerateKeyPairRequest) Descriptor() ([]byte, []int) {
	return file_space_proto_rawDescGZIP(), []int{2}
}

type GenerateKeyPairResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublicKey  string `protobuf:"bytes,1,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	PrivateKey string `protobuf:"bytes,2,opt,name=privateKey,proto3" json:"privateKey,omitempty"`
}

func (x *GenerateKeyPairResponse) Reset() {
	*x = GenerateKeyPairResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_space_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateKeyPairResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateKeyPairResponse) ProtoMessage() {}

func (x *GenerateKeyPairResponse) ProtoReflect() protoreflect.Message {
	mi := &file_space_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateKeyPairResponse.ProtoReflect.Descriptor instead.
func (*GenerateKeyPairResponse) Descriptor() ([]byte, []int) {
	return file_space_proto_rawDescGZIP(), []int{3}
}

func (x *GenerateKeyPairResponse) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *GenerateKeyPairResponse) GetPrivateKey() string {
	if x != nil {
		return x.PrivateKey
	}
	return ""
}

type FileEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *FileEventResponse) Reset() {
	*x = FileEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_space_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileEventResponse) ProtoMessage() {}

func (x *FileEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_space_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileEventResponse.ProtoReflect.Descriptor instead.
func (*FileEventResponse) Descriptor() ([]byte, []int) {
	return file_space_proto_rawDescGZIP(), []int{4}
}

func (x *FileEventResponse) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

var File_space_proto protoreflect.FileDescriptor

var file_space_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x25, 0x0a, 0x0f, 0x50, 0x61, 0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x58, 0x0a, 0x10, 0x50, 0x61, 0x74, 0x68,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x12, 0x1a, 0x0a, 0x08, 0x69, 0x70, 0x66, 0x73, 0x48, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x69, 0x70, 0x66, 0x73, 0x48, 0x61, 0x73, 0x68, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x73, 0x44, 0x69, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x44,
	0x69, 0x72, 0x22, 0x18, 0x0a, 0x16, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4b, 0x65,
	0x79, 0x50, 0x61, 0x69, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x57, 0x0a, 0x17,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x50, 0x61, 0x69, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x4b, 0x65, 0x79, 0x22, 0x27, 0x0a, 0x11, 0x46, 0x69, 0x6c, 0x65, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x32, 0xc0,
	0x02, 0x0a, 0x08, 0x53, 0x70, 0x61, 0x63, 0x65, 0x41, 0x70, 0x69, 0x12, 0x40, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x50, 0x61, 0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x2e, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a,
	0x0f, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x50, 0x61, 0x69, 0x72,
	0x12, 0x1d, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x4b, 0x65, 0x79, 0x50, 0x61, 0x69, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x4b, 0x65, 0x79, 0x50, 0x61, 0x69, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x5b, 0x0a, 0x18, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79,
	0x50, 0x61, 0x69, 0x72, 0x57, 0x69, 0x74, 0x68, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x12, 0x1d, 0x2e,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4b, 0x65,
	0x79, 0x50, 0x61, 0x69, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79,
	0x50, 0x61, 0x69, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41,
	0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30,
	0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_space_proto_rawDescOnce sync.Once
	file_space_proto_rawDescData = file_space_proto_rawDesc
)

func file_space_proto_rawDescGZIP() []byte {
	file_space_proto_rawDescOnce.Do(func() {
		file_space_proto_rawDescData = protoimpl.X.CompressGZIP(file_space_proto_rawDescData)
	})
	return file_space_proto_rawDescData
}

var file_space_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_space_proto_goTypes = []interface{}{
	(*PathInfoRequest)(nil),         // 0: space.PathInfoRequest
	(*PathInfoResponse)(nil),        // 1: space.PathInfoResponse
	(*GenerateKeyPairRequest)(nil),  // 2: space.GenerateKeyPairRequest
	(*GenerateKeyPairResponse)(nil), // 3: space.GenerateKeyPairResponse
	(*FileEventResponse)(nil),       // 4: space.FileEventResponse
	(*empty.Empty)(nil),             // 5: google.protobuf.Empty
}
var file_space_proto_depIdxs = []int32{
	0, // 0: space.SpaceApi.GetPathInfo:input_type -> space.PathInfoRequest
	2, // 1: space.SpaceApi.GenerateKeyPair:input_type -> space.GenerateKeyPairRequest
	2, // 2: space.SpaceApi.GenerateKeyPairWithForce:input_type -> space.GenerateKeyPairRequest
	5, // 3: space.SpaceApi.Subscribe:input_type -> google.protobuf.Empty
	1, // 4: space.SpaceApi.GetPathInfo:output_type -> space.PathInfoResponse
	3, // 5: space.SpaceApi.GenerateKeyPair:output_type -> space.GenerateKeyPairResponse
	3, // 6: space.SpaceApi.GenerateKeyPairWithForce:output_type -> space.GenerateKeyPairResponse
	4, // 7: space.SpaceApi.Subscribe:output_type -> space.FileEventResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_space_proto_init() }
func file_space_proto_init() {
	if File_space_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_space_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PathInfoRequest); i {
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
		file_space_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PathInfoResponse); i {
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
		file_space_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateKeyPairRequest); i {
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
		file_space_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateKeyPairResponse); i {
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
		file_space_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileEventResponse); i {
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
			RawDescriptor: file_space_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_space_proto_goTypes,
		DependencyIndexes: file_space_proto_depIdxs,
		MessageInfos:      file_space_proto_msgTypes,
	}.Build()
	File_space_proto = out.File
	file_space_proto_rawDesc = nil
	file_space_proto_goTypes = nil
	file_space_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SpaceApiClient is the client API for SpaceApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SpaceApiClient interface {
	GetPathInfo(ctx context.Context, in *PathInfoRequest, opts ...grpc.CallOption) (*PathInfoResponse, error)
	GenerateKeyPair(ctx context.Context, in *GenerateKeyPairRequest, opts ...grpc.CallOption) (*GenerateKeyPairResponse, error)
	GenerateKeyPairWithForce(ctx context.Context, in *GenerateKeyPairRequest, opts ...grpc.CallOption) (*GenerateKeyPairResponse, error)
	Subscribe(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (SpaceApi_SubscribeClient, error)
}

type spaceApiClient struct {
	cc grpc.ClientConnInterface
}

func NewSpaceApiClient(cc grpc.ClientConnInterface) SpaceApiClient {
	return &spaceApiClient{cc}
}

func (c *spaceApiClient) GetPathInfo(ctx context.Context, in *PathInfoRequest, opts ...grpc.CallOption) (*PathInfoResponse, error) {
	out := new(PathInfoResponse)
	err := c.cc.Invoke(ctx, "/space.SpaceApi/GetPathInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceApiClient) GenerateKeyPair(ctx context.Context, in *GenerateKeyPairRequest, opts ...grpc.CallOption) (*GenerateKeyPairResponse, error) {
	out := new(GenerateKeyPairResponse)
	err := c.cc.Invoke(ctx, "/space.SpaceApi/GenerateKeyPair", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceApiClient) GenerateKeyPairWithForce(ctx context.Context, in *GenerateKeyPairRequest, opts ...grpc.CallOption) (*GenerateKeyPairResponse, error) {
	out := new(GenerateKeyPairResponse)
	err := c.cc.Invoke(ctx, "/space.SpaceApi/GenerateKeyPairWithForce", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceApiClient) Subscribe(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (SpaceApi_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SpaceApi_serviceDesc.Streams[0], "/space.SpaceApi/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &spaceApiSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SpaceApi_SubscribeClient interface {
	Recv() (*FileEventResponse, error)
	grpc.ClientStream
}

type spaceApiSubscribeClient struct {
	grpc.ClientStream
}

func (x *spaceApiSubscribeClient) Recv() (*FileEventResponse, error) {
	m := new(FileEventResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SpaceApiServer is the server API for SpaceApi service.
type SpaceApiServer interface {
	GetPathInfo(context.Context, *PathInfoRequest) (*PathInfoResponse, error)
	GenerateKeyPair(context.Context, *GenerateKeyPairRequest) (*GenerateKeyPairResponse, error)
	GenerateKeyPairWithForce(context.Context, *GenerateKeyPairRequest) (*GenerateKeyPairResponse, error)
	Subscribe(*empty.Empty, SpaceApi_SubscribeServer) error
}

// UnimplementedSpaceApiServer can be embedded to have forward compatible implementations.
type UnimplementedSpaceApiServer struct {
}

func (*UnimplementedSpaceApiServer) GetPathInfo(context.Context, *PathInfoRequest) (*PathInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPathInfo not implemented")
}
func (*UnimplementedSpaceApiServer) GenerateKeyPair(context.Context, *GenerateKeyPairRequest) (*GenerateKeyPairResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateKeyPair not implemented")
}
func (*UnimplementedSpaceApiServer) GenerateKeyPairWithForce(context.Context, *GenerateKeyPairRequest) (*GenerateKeyPairResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateKeyPairWithForce not implemented")
}
func (*UnimplementedSpaceApiServer) Subscribe(*empty.Empty, SpaceApi_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}

func RegisterSpaceApiServer(s *grpc.Server, srv SpaceApiServer) {
	s.RegisterService(&_SpaceApi_serviceDesc, srv)
}

func _SpaceApi_GetPathInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PathInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceApiServer).GetPathInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/space.SpaceApi/GetPathInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceApiServer).GetPathInfo(ctx, req.(*PathInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpaceApi_GenerateKeyPair_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateKeyPairRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceApiServer).GenerateKeyPair(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/space.SpaceApi/GenerateKeyPair",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceApiServer).GenerateKeyPair(ctx, req.(*GenerateKeyPairRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpaceApi_GenerateKeyPairWithForce_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateKeyPairRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceApiServer).GenerateKeyPairWithForce(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/space.SpaceApi/GenerateKeyPairWithForce",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceApiServer).GenerateKeyPairWithForce(ctx, req.(*GenerateKeyPairRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpaceApi_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SpaceApiServer).Subscribe(m, &spaceApiSubscribeServer{stream})
}

type SpaceApi_SubscribeServer interface {
	Send(*FileEventResponse) error
	grpc.ServerStream
}

type spaceApiSubscribeServer struct {
	grpc.ServerStream
}

func (x *spaceApiSubscribeServer) Send(m *FileEventResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _SpaceApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "space.SpaceApi",
	HandlerType: (*SpaceApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPathInfo",
			Handler:    _SpaceApi_GetPathInfo_Handler,
		},
		{
			MethodName: "GenerateKeyPair",
			Handler:    _SpaceApi_GenerateKeyPair_Handler,
		},
		{
			MethodName: "GenerateKeyPairWithForce",
			Handler:    _SpaceApi_GenerateKeyPairWithForce_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _SpaceApi_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "space.proto",
}
