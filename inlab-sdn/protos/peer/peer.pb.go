// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: peer.proto

package peer

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

type StatusCode int32

const (
	StatusCode_Unknown StatusCode = 0
	StatusCode_Ok      StatusCode = 1
	StatusCode_Failed  StatusCode = 2
)

// Enum value maps for StatusCode.
var (
	StatusCode_name = map[int32]string{
		0: "Unknown",
		1: "Ok",
		2: "Failed",
	}
	StatusCode_value = map[string]int32{
		"Unknown": 0,
		"Ok":      1,
		"Failed":  2,
	}
)

func (x StatusCode) Enum() *StatusCode {
	p := new(StatusCode)
	*p = x
	return p
}

func (x StatusCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusCode) Descriptor() protoreflect.EnumDescriptor {
	return file_peer_proto_enumTypes[0].Descriptor()
}

func (StatusCode) Type() protoreflect.EnumType {
	return &file_peer_proto_enumTypes[0]
}

func (x StatusCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusCode.Descriptor instead.
func (StatusCode) EnumDescriptor() ([]byte, []int) {
	return file_peer_proto_rawDescGZIP(), []int{0}
}

// Indivisual information of the peers
type PeerConnection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mobile *PeerInfo `protobuf:"bytes,1,opt,name=Mobile,proto3" json:"Mobile,omitempty"`
	D2D    *PeerInfo `protobuf:"bytes,2,opt,name=D2D,proto3" json:"D2D,omitempty"`
}

func (x *PeerConnection) Reset() {
	*x = PeerConnection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerConnection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerConnection) ProtoMessage() {}

func (x *PeerConnection) ProtoReflect() protoreflect.Message {
	mi := &file_peer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerConnection.ProtoReflect.Descriptor instead.
func (*PeerConnection) Descriptor() ([]byte, []int) {
	return file_peer_proto_rawDescGZIP(), []int{0}
}

func (x *PeerConnection) GetMobile() *PeerInfo {
	if x != nil {
		return x.Mobile
	}
	return nil
}

func (x *PeerConnection) GetD2D() *PeerInfo {
	if x != nil {
		return x.D2D
	}
	return nil
}

type PeerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoint string       `protobuf:"bytes,1,opt,name=Endpoint,proto3" json:"Endpoint,omitempty"`
	NetInfo  *NetworkInfo `protobuf:"bytes,2,opt,name=NetInfo,proto3" json:"NetInfo,omitempty"`
}

func (x *PeerInfo) Reset() {
	*x = PeerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerInfo) ProtoMessage() {}

func (x *PeerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_peer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerInfo.ProtoReflect.Descriptor instead.
func (*PeerInfo) Descriptor() ([]byte, []int) {
	return file_peer_proto_rawDescGZIP(), []int{1}
}

func (x *PeerInfo) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *PeerInfo) GetNetInfo() *NetworkInfo {
	if x != nil {
		return x.NetInfo
	}
	return nil
}

type NetworkInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetworkInterface string  `protobuf:"bytes,1,opt,name=NetworkInterface,proto3" json:"NetworkInterface,omitempty"`
	NetworkStrength  float32 `protobuf:"fixed32,2,opt,name=NetworkStrength,proto3" json:"NetworkStrength,omitempty"`
}

func (x *NetworkInfo) Reset() {
	*x = NetworkInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkInfo) ProtoMessage() {}

func (x *NetworkInfo) ProtoReflect() protoreflect.Message {
	mi := &file_peer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkInfo.ProtoReflect.Descriptor instead.
func (*NetworkInfo) Descriptor() ([]byte, []int) {
	return file_peer_proto_rawDescGZIP(), []int{2}
}

func (x *NetworkInfo) GetNetworkInterface() string {
	if x != nil {
		return x.NetworkInterface
	}
	return ""
}

func (x *NetworkInfo) GetNetworkStrength() float32 {
	if x != nil {
		return x.NetworkStrength
	}
	return 0
}

type OverlayStructure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoint string   `protobuf:"bytes,1,opt,name=Endpoint,proto3" json:"Endpoint,omitempty"`
	SubPeers []string `protobuf:"bytes,2,rep,name=SubPeers,proto3" json:"SubPeers,omitempty"`
}

func (x *OverlayStructure) Reset() {
	*x = OverlayStructure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OverlayStructure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OverlayStructure) ProtoMessage() {}

func (x *OverlayStructure) ProtoReflect() protoreflect.Message {
	mi := &file_peer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OverlayStructure.ProtoReflect.Descriptor instead.
func (*OverlayStructure) Descriptor() ([]byte, []int) {
	return file_peer_proto_rawDescGZIP(), []int{3}
}

func (x *OverlayStructure) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *OverlayStructure) GetSubPeers() []string {
	if x != nil {
		return x.SubPeers
	}
	return nil
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code StatusCode `protobuf:"varint,1,opt,name=Code,proto3,enum=peer.StatusCode" json:"Code,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_peer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_peer_proto_rawDescGZIP(), []int{4}
}

func (x *Status) GetCode() StatusCode {
	if x != nil {
		return x.Code
	}
	return StatusCode_Unknown
}

var File_peer_proto protoreflect.FileDescriptor

var file_peer_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x70, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x70, 0x65,
	0x65, 0x72, 0x22, 0x5a, 0x0a, 0x0e, 0x50, 0x65, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x06, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x2e, 0x50, 0x65, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x03,
	0x44, 0x32, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x65, 0x65, 0x72,
	0x2e, 0x50, 0x65, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x03, 0x44, 0x32, 0x44, 0x22, 0x53,
	0x0a, 0x08, 0x50, 0x65, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x07, 0x4e, 0x65, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x2e, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x4e, 0x65, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x63, 0x0a, 0x0b, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x2a, 0x0a, 0x10, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x28,
	0x0a, 0x0f, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x74, 0x72, 0x65, 0x6e, 0x67, 0x74,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0f, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x53, 0x74, 0x72, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x22, 0x4a, 0x0a, 0x10, 0x4f, 0x76, 0x65, 0x72,
	0x6c, 0x61, 0x79, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x75, 0x62, 0x50,
	0x65, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x53, 0x75, 0x62, 0x50,
	0x65, 0x65, 0x72, 0x73, 0x22, 0x2e, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x24,
	0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x70,
	0x65, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04,
	0x43, 0x6f, 0x64, 0x65, 0x2a, 0x2d, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12,
	0x06, 0x0a, 0x02, 0x4f, 0x6b, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x61, 0x69, 0x6c, 0x65,
	0x64, 0x10, 0x02, 0x32, 0x8e, 0x01, 0x0a, 0x0a, 0x53, 0x44, 0x4e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3c, 0x0a, 0x14, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x65, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x2e, 0x70, 0x65, 0x65,
	0x72, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x0c, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00,
	0x12, 0x42, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x76, 0x65, 0x72, 0x6c, 0x61,
	0x79, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x0e, 0x2e, 0x70, 0x65, 0x65,
	0x72, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x16, 0x2e, 0x70, 0x65, 0x65,
	0x72, 0x2e, 0x4f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x79, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x69, 0x6e, 0x6c, 0x61, 0x62, 0x2d, 0x73, 0x64,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_peer_proto_rawDescOnce sync.Once
	file_peer_proto_rawDescData = file_peer_proto_rawDesc
)

func file_peer_proto_rawDescGZIP() []byte {
	file_peer_proto_rawDescOnce.Do(func() {
		file_peer_proto_rawDescData = protoimpl.X.CompressGZIP(file_peer_proto_rawDescData)
	})
	return file_peer_proto_rawDescData
}

var file_peer_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_peer_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_peer_proto_goTypes = []interface{}{
	(StatusCode)(0),          // 0: peer.StatusCode
	(*PeerConnection)(nil),   // 1: peer.PeerConnection
	(*PeerInfo)(nil),         // 2: peer.PeerInfo
	(*NetworkInfo)(nil),      // 3: peer.NetworkInfo
	(*OverlayStructure)(nil), // 4: peer.OverlayStructure
	(*Status)(nil),           // 5: peer.Status
}
var file_peer_proto_depIdxs = []int32{
	2, // 0: peer.PeerConnection.Mobile:type_name -> peer.PeerInfo
	2, // 1: peer.PeerConnection.D2D:type_name -> peer.PeerInfo
	3, // 2: peer.PeerInfo.NetInfo:type_name -> peer.NetworkInfo
	0, // 3: peer.Status.Code:type_name -> peer.StatusCode
	1, // 4: peer.SDNService.StorePeerInformation:input_type -> peer.PeerConnection
	2, // 5: peer.SDNService.UpdateOverlayStructure:input_type -> peer.PeerInfo
	5, // 6: peer.SDNService.StorePeerInformation:output_type -> peer.Status
	4, // 7: peer.SDNService.UpdateOverlayStructure:output_type -> peer.OverlayStructure
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_peer_proto_init() }
func file_peer_proto_init() {
	if File_peer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_peer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerConnection); i {
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
		file_peer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerInfo); i {
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
		file_peer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkInfo); i {
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
		file_peer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OverlayStructure); i {
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
		file_peer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
			RawDescriptor: file_peer_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_peer_proto_goTypes,
		DependencyIndexes: file_peer_proto_depIdxs,
		EnumInfos:         file_peer_proto_enumTypes,
		MessageInfos:      file_peer_proto_msgTypes,
	}.Build()
	File_peer_proto = out.File
	file_peer_proto_rawDesc = nil
	file_peer_proto_goTypes = nil
	file_peer_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SDNServiceClient is the client API for SDNService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SDNServiceClient interface {
	StorePeerInformation(ctx context.Context, in *PeerConnection, opts ...grpc.CallOption) (*Status, error)
	UpdateOverlayStructure(ctx context.Context, in *PeerInfo, opts ...grpc.CallOption) (*OverlayStructure, error)
}

type sDNServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSDNServiceClient(cc grpc.ClientConnInterface) SDNServiceClient {
	return &sDNServiceClient{cc}
}

func (c *sDNServiceClient) StorePeerInformation(ctx context.Context, in *PeerConnection, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/peer.SDNService/StorePeerInformation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sDNServiceClient) UpdateOverlayStructure(ctx context.Context, in *PeerInfo, opts ...grpc.CallOption) (*OverlayStructure, error) {
	out := new(OverlayStructure)
	err := c.cc.Invoke(ctx, "/peer.SDNService/UpdateOverlayStructure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SDNServiceServer is the server API for SDNService service.
type SDNServiceServer interface {
	StorePeerInformation(context.Context, *PeerConnection) (*Status, error)
	UpdateOverlayStructure(context.Context, *PeerInfo) (*OverlayStructure, error)
}

// UnimplementedSDNServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSDNServiceServer struct {
}

func (*UnimplementedSDNServiceServer) StorePeerInformation(context.Context, *PeerConnection) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StorePeerInformation not implemented")
}
func (*UnimplementedSDNServiceServer) UpdateOverlayStructure(context.Context, *PeerInfo) (*OverlayStructure, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOverlayStructure not implemented")
}

func RegisterSDNServiceServer(s *grpc.Server, srv SDNServiceServer) {
	s.RegisterService(&_SDNService_serviceDesc, srv)
}

func _SDNService_StorePeerInformation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeerConnection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SDNServiceServer).StorePeerInformation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/peer.SDNService/StorePeerInformation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SDNServiceServer).StorePeerInformation(ctx, req.(*PeerConnection))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDNService_UpdateOverlayStructure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeerInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SDNServiceServer).UpdateOverlayStructure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/peer.SDNService/UpdateOverlayStructure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SDNServiceServer).UpdateOverlayStructure(ctx, req.(*PeerInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _SDNService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "peer.SDNService",
	HandlerType: (*SDNServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StorePeerInformation",
			Handler:    _SDNService_StorePeerInformation_Handler,
		},
		{
			MethodName: "UpdateOverlayStructure",
			Handler:    _SDNService_UpdateOverlayStructure_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "peer.proto",
}
