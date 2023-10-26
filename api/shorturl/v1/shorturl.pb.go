// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: shorturl/v1/shorturl.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ExpandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shorten string `protobuf:"bytes,1,opt,name=shorten,proto3" json:"shorten,omitempty"`
}

func (x *ExpandRequest) Reset() {
	*x = ExpandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shorturl_v1_shorturl_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpandRequest) ProtoMessage() {}

func (x *ExpandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shorturl_v1_shorturl_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpandRequest.ProtoReflect.Descriptor instead.
func (*ExpandRequest) Descriptor() ([]byte, []int) {
	return file_shorturl_v1_shorturl_proto_rawDescGZIP(), []int{0}
}

func (x *ExpandRequest) GetShorten() string {
	if x != nil {
		return x.Shorten
	}
	return ""
}

type ExpandReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *ExpandReply) Reset() {
	*x = ExpandReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shorturl_v1_shorturl_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpandReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpandReply) ProtoMessage() {}

func (x *ExpandReply) ProtoReflect() protoreflect.Message {
	mi := &file_shorturl_v1_shorturl_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpandReply.ProtoReflect.Descriptor instead.
func (*ExpandReply) Descriptor() ([]byte, []int) {
	return file_shorturl_v1_shorturl_proto_rawDescGZIP(), []int{1}
}

func (x *ExpandReply) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type ShortenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *ShortenRequest) Reset() {
	*x = ShortenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shorturl_v1_shorturl_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenRequest) ProtoMessage() {}

func (x *ShortenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shorturl_v1_shorturl_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenRequest.ProtoReflect.Descriptor instead.
func (*ShortenRequest) Descriptor() ([]byte, []int) {
	return file_shorturl_v1_shorturl_proto_rawDescGZIP(), []int{2}
}

func (x *ShortenRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type ShortenReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shorten string `protobuf:"bytes,1,opt,name=shorten,proto3" json:"shorten,omitempty"`
}

func (x *ShortenReply) Reset() {
	*x = ShortenReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shorturl_v1_shorturl_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortenReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenReply) ProtoMessage() {}

func (x *ShortenReply) ProtoReflect() protoreflect.Message {
	mi := &file_shorturl_v1_shorturl_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenReply.ProtoReflect.Descriptor instead.
func (*ShortenReply) Descriptor() ([]byte, []int) {
	return file_shorturl_v1_shorturl_proto_rawDescGZIP(), []int{3}
}

func (x *ShortenReply) GetShorten() string {
	if x != nil {
		return x.Shorten
	}
	return ""
}

var File_shorturl_v1_shorturl_proto protoreflect.FileDescriptor

var file_shorturl_v1_shorturl_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a, 0x0d, 0x45,
	0x78, 0x70, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x22, 0x1f, 0x0a, 0x0b, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x22, 0x0a, 0x0e, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x28, 0x0a, 0x0c, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x65, 0x6e, 0x32, 0xd3, 0x01, 0x0a, 0x08, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55,
	0x72, 0x6c, 0x12, 0x61, 0x0a, 0x06, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x12, 0x1e, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x78, 0x70, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x78, 0x70, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x13, 0x12, 0x11, 0x2f, 0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x2f, 0x7b, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x65, 0x6e, 0x7d, 0x12, 0x64, 0x0a, 0x07, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e,
	0x12, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x22, 0x0e, 0x2f, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x65, 0x6e, 0x2f, 0x7b, 0x75, 0x72, 0x6c, 0x7d, 0x42, 0x38, 0x0a, 0x0f, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x2e, 0x76, 0x31, 0x50, 0x01,
	0x5a, 0x23, 0x6e, 0x65, 0x78, 0x74, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6c, 0x61, 0x79, 0x6f,
	0x75, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x75, 0x72, 0x6c, 0x2f,
	0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shorturl_v1_shorturl_proto_rawDescOnce sync.Once
	file_shorturl_v1_shorturl_proto_rawDescData = file_shorturl_v1_shorturl_proto_rawDesc
)

func file_shorturl_v1_shorturl_proto_rawDescGZIP() []byte {
	file_shorturl_v1_shorturl_proto_rawDescOnce.Do(func() {
		file_shorturl_v1_shorturl_proto_rawDescData = protoimpl.X.CompressGZIP(file_shorturl_v1_shorturl_proto_rawDescData)
	})
	return file_shorturl_v1_shorturl_proto_rawDescData
}

var file_shorturl_v1_shorturl_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_shorturl_v1_shorturl_proto_goTypes = []interface{}{
	(*ExpandRequest)(nil),  // 0: api.shorturl.v1.ExpandRequest
	(*ExpandReply)(nil),    // 1: api.shorturl.v1.ExpandReply
	(*ShortenRequest)(nil), // 2: api.shorturl.v1.ShortenRequest
	(*ShortenReply)(nil),   // 3: api.shorturl.v1.ShortenReply
}
var file_shorturl_v1_shorturl_proto_depIdxs = []int32{
	0, // 0: api.shorturl.v1.ShortUrl.Expand:input_type -> api.shorturl.v1.ExpandRequest
	2, // 1: api.shorturl.v1.ShortUrl.Shorten:input_type -> api.shorturl.v1.ShortenRequest
	1, // 2: api.shorturl.v1.ShortUrl.Expand:output_type -> api.shorturl.v1.ExpandReply
	3, // 3: api.shorturl.v1.ShortUrl.Shorten:output_type -> api.shorturl.v1.ShortenReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_shorturl_v1_shorturl_proto_init() }
func file_shorturl_v1_shorturl_proto_init() {
	if File_shorturl_v1_shorturl_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shorturl_v1_shorturl_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpandRequest); i {
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
		file_shorturl_v1_shorturl_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpandReply); i {
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
		file_shorturl_v1_shorturl_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortenRequest); i {
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
		file_shorturl_v1_shorturl_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortenReply); i {
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
			RawDescriptor: file_shorturl_v1_shorturl_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shorturl_v1_shorturl_proto_goTypes,
		DependencyIndexes: file_shorturl_v1_shorturl_proto_depIdxs,
		MessageInfos:      file_shorturl_v1_shorturl_proto_msgTypes,
	}.Build()
	File_shorturl_v1_shorturl_proto = out.File
	file_shorturl_v1_shorturl_proto_rawDesc = nil
	file_shorturl_v1_shorturl_proto_goTypes = nil
	file_shorturl_v1_shorturl_proto_depIdxs = nil
}
