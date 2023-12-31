// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v4.25.1
// source: package.proto

package pb

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

type Package struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeUID    []byte `protobuf:"bytes,1,opt,name=NodeUID,proto3" json:"NodeUID,omitempty"`        // 消息接收节点
	MessageID  int32  `protobuf:"varint,2,opt,name=MessageID,proto3" json:"MessageID,omitempty"`   // 消息ID
	UserID     int32  `protobuf:"varint,3,opt,name=UserID,proto3" json:"UserID,omitempty"`         // Front消息必要字段
	Idempotent int64  `protobuf:"varint,4,opt,name=Idempotent,proto3" json:"Idempotent,omitempty"` // 幂等
	RealAddr   string `protobuf:"bytes,5,opt,name=RealAddr,proto3" json:"RealAddr,omitempty"`      // 客户端真实地址, Front请求时由Kraken赋值
	Token      []byte `protobuf:"bytes,6,opt,name=Token,proto3" json:"Token,omitempty"`            // 会话, Front消息必要字段
	Data       []byte `protobuf:"bytes,7,opt,name=Data,proto3" json:"Data,omitempty"`              // 消息内容
}

func (x *Package) Reset() {
	*x = Package{}
	if protoimpl.UnsafeEnabled {
		mi := &file_package_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Package) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Package) ProtoMessage() {}

func (x *Package) ProtoReflect() protoreflect.Message {
	mi := &file_package_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Package.ProtoReflect.Descriptor instead.
func (*Package) Descriptor() ([]byte, []int) {
	return file_package_proto_rawDescGZIP(), []int{0}
}

func (x *Package) GetNodeUID() []byte {
	if x != nil {
		return x.NodeUID
	}
	return nil
}

func (x *Package) GetMessageID() int32 {
	if x != nil {
		return x.MessageID
	}
	return 0
}

func (x *Package) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *Package) GetIdempotent() int64 {
	if x != nil {
		return x.Idempotent
	}
	return 0
}

func (x *Package) GetRealAddr() string {
	if x != nil {
		return x.RealAddr
	}
	return ""
}

func (x *Package) GetToken() []byte {
	if x != nil {
		return x.Token
	}
	return nil
}

func (x *Package) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_package_proto protoreflect.FileDescriptor

var file_package_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x22, 0xbf, 0x01, 0x0a, 0x07, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x4e, 0x6f, 0x64, 0x65, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x4e, 0x6f, 0x64, 0x65, 0x55, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x1e, 0x0a, 0x0a, 0x49, 0x64, 0x65, 0x6d, 0x70, 0x6f, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x49, 0x64, 0x65, 0x6d, 0x70, 0x6f, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x52, 0x65, 0x61, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x52, 0x65, 0x61, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x44, 0x61, 0x74, 0x61, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_package_proto_rawDescOnce sync.Once
	file_package_proto_rawDescData = file_package_proto_rawDesc
)

func file_package_proto_rawDescGZIP() []byte {
	file_package_proto_rawDescOnce.Do(func() {
		file_package_proto_rawDescData = protoimpl.X.CompressGZIP(file_package_proto_rawDescData)
	})
	return file_package_proto_rawDescData
}

var file_package_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_package_proto_goTypes = []interface{}{
	(*Package)(nil), // 0: pb.Package
}
var file_package_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_package_proto_init() }
func file_package_proto_init() {
	if File_package_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_package_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Package); i {
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
			RawDescriptor: file_package_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_package_proto_goTypes,
		DependencyIndexes: file_package_proto_depIdxs,
		MessageInfos:      file_package_proto_msgTypes,
	}.Build()
	File_package_proto = out.File
	file_package_proto_rawDesc = nil
	file_package_proto_goTypes = nil
	file_package_proto_depIdxs = nil
}
