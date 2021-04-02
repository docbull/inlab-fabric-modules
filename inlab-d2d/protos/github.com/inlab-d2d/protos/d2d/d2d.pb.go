// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: d2d.proto

package d2d

import (
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
	return file_d2d_proto_enumTypes[0].Descriptor()
}

func (StatusCode) Type() protoreflect.EnumType {
	return &file_d2d_proto_enumTypes[0]
}

func (x StatusCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusCode.Descriptor instead.
func (StatusCode) EnumDescriptor() ([]byte, []int) {
	return file_d2d_proto_rawDescGZIP(), []int{0}
}

// Envelope contains a marshalled
// GossipMessage and a signature over it.
// It may also contain a SecretEnvelope
// which is a marshalled Secret
type Envelope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload        []byte          `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Signature      []byte          `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	SecretEnvelope *SecretEnvelope `protobuf:"bytes,3,opt,name=secret_envelope,json=secretEnvelope,proto3" json:"secret_envelope,omitempty"`
}

func (x *Envelope) Reset() {
	*x = Envelope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_d2d_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Envelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Envelope) ProtoMessage() {}

func (x *Envelope) ProtoReflect() protoreflect.Message {
	mi := &file_d2d_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Envelope.ProtoReflect.Descriptor instead.
func (*Envelope) Descriptor() ([]byte, []int) {
	return file_d2d_proto_rawDescGZIP(), []int{0}
}

func (x *Envelope) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *Envelope) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *Envelope) GetSecretEnvelope() *SecretEnvelope {
	if x != nil {
		return x.SecretEnvelope
	}
	return nil
}

// SecretEnvelope is a marshalled Secret
// and a signature over it.
// The signature should be validated by the peer
// that signed the Envelope the SecretEnvelope
// came with
type SecretEnvelope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload   []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *SecretEnvelope) Reset() {
	*x = SecretEnvelope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_d2d_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretEnvelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretEnvelope) ProtoMessage() {}

func (x *SecretEnvelope) ProtoReflect() protoreflect.Message {
	mi := &file_d2d_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretEnvelope.ProtoReflect.Descriptor instead.
func (*SecretEnvelope) Descriptor() ([]byte, []int) {
	return file_d2d_proto_rawDescGZIP(), []int{1}
}

func (x *SecretEnvelope) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *SecretEnvelope) GetSignature() []byte {
	if x != nil {
		return x.Signature
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
		mi := &file_d2d_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_d2d_proto_msgTypes[2]
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
	return file_d2d_proto_rawDescGZIP(), []int{2}
}

func (x *Status) GetCode() StatusCode {
	if x != nil {
		return x.Code
	}
	return StatusCode_Unknown
}

var File_d2d_proto protoreflect.FileDescriptor

var file_d2d_proto_rawDesc = []byte{
	0x0a, 0x09, 0x64, 0x32, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x70, 0x65, 0x65,
	0x72, 0x22, 0x81, 0x01, 0x0a, 0x08, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x3d, 0x0a, 0x0f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x5f, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x45, 0x6e, 0x76,
	0x65, 0x6c, 0x6f, 0x70, 0x65, 0x52, 0x0e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x45, 0x6e, 0x76,
	0x65, 0x6c, 0x6f, 0x70, 0x65, 0x22, 0x48, 0x0a, 0x0e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x45,
	0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22,
	0x2e, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x24, 0x0a, 0x04, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x2a,
	0x2d, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x6b,
	0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x02, 0x32, 0x3e,
	0x0a, 0x0a, 0x44, 0x32, 0x44, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x0e,
	0x44, 0x32, 0x44, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x0e,
	0x2e, 0x70, 0x65, 0x65, 0x72, 0x2e, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x1a, 0x0c,
	0x2e, 0x70, 0x65, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x42, 0x21,
	0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x6c,
	0x61, 0x62, 0x2d, 0x64, 0x32, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x64, 0x32,
	0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_d2d_proto_rawDescOnce sync.Once
	file_d2d_proto_rawDescData = file_d2d_proto_rawDesc
)

func file_d2d_proto_rawDescGZIP() []byte {
	file_d2d_proto_rawDescOnce.Do(func() {
		file_d2d_proto_rawDescData = protoimpl.X.CompressGZIP(file_d2d_proto_rawDescData)
	})
	return file_d2d_proto_rawDescData
}

var file_d2d_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_d2d_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_d2d_proto_goTypes = []interface{}{
	(StatusCode)(0),        // 0: peer.StatusCode
	(*Envelope)(nil),       // 1: peer.Envelope
	(*SecretEnvelope)(nil), // 2: peer.SecretEnvelope
	(*Status)(nil),         // 3: peer.Status
}
var file_d2d_proto_depIdxs = []int32{
	2, // 0: peer.Envelope.secret_envelope:type_name -> peer.SecretEnvelope
	0, // 1: peer.Status.Code:type_name -> peer.StatusCode
	1, // 2: peer.D2DService.D2DBlockStream:input_type -> peer.Envelope
	3, // 3: peer.D2DService.D2DBlockStream:output_type -> peer.Status
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_d2d_proto_init() }
func file_d2d_proto_init() {
	if File_d2d_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_d2d_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Envelope); i {
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
		file_d2d_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretEnvelope); i {
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
		file_d2d_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_d2d_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_d2d_proto_goTypes,
		DependencyIndexes: file_d2d_proto_depIdxs,
		EnumInfos:         file_d2d_proto_enumTypes,
		MessageInfos:      file_d2d_proto_msgTypes,
	}.Build()
	File_d2d_proto = out.File
	file_d2d_proto_rawDesc = nil
	file_d2d_proto_goTypes = nil
	file_d2d_proto_depIdxs = nil
}
