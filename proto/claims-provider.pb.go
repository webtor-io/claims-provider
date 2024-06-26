// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: claims-provider.proto

package claims_provider

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_claims_provider_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_claims_provider_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_claims_provider_proto_rawDescGZIP(), []int{0}
}

func (x *GetRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Context *Context `protobuf:"bytes,1,opt,name=context,proto3" json:"context,omitempty"`
	Claims  *Claims  `protobuf:"bytes,2,opt,name=claims,proto3" json:"claims,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_claims_provider_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_claims_provider_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_claims_provider_proto_rawDescGZIP(), []int{1}
}

func (x *GetResponse) GetContext() *Context {
	if x != nil {
		return x.Context
	}
	return nil
}

func (x *GetResponse) GetClaims() *Claims {
	if x != nil {
		return x.Claims
	}
	return nil
}

type Context struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tier *Tier `protobuf:"bytes,1,opt,name=tier,proto3" json:"tier,omitempty"`
}

func (x *Context) Reset() {
	*x = Context{}
	if protoimpl.UnsafeEnabled {
		mi := &file_claims_provider_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Context) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Context) ProtoMessage() {}

func (x *Context) ProtoReflect() protoreflect.Message {
	mi := &file_claims_provider_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Context.ProtoReflect.Descriptor instead.
func (*Context) Descriptor() ([]byte, []int) {
	return file_claims_provider_proto_rawDescGZIP(), []int{2}
}

func (x *Context) GetTier() *Tier {
	if x != nil {
		return x.Tier
	}
	return nil
}

type Tier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Tier) Reset() {
	*x = Tier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_claims_provider_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tier) ProtoMessage() {}

func (x *Tier) ProtoReflect() protoreflect.Message {
	mi := &file_claims_provider_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tier.ProtoReflect.Descriptor instead.
func (*Tier) Descriptor() ([]byte, []int) {
	return file_claims_provider_proto_rawDescGZIP(), []int{3}
}

func (x *Tier) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Tier) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Claims struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Connection *Connection `protobuf:"bytes,1,opt,name=connection,proto3" json:"connection,omitempty"`
	Embed      *Embed      `protobuf:"bytes,2,opt,name=embed,proto3" json:"embed,omitempty"`
	Site       *Site       `protobuf:"bytes,3,opt,name=site,proto3" json:"site,omitempty"`
}

func (x *Claims) Reset() {
	*x = Claims{}
	if protoimpl.UnsafeEnabled {
		mi := &file_claims_provider_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Claims) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Claims) ProtoMessage() {}

func (x *Claims) ProtoReflect() protoreflect.Message {
	mi := &file_claims_provider_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Claims.ProtoReflect.Descriptor instead.
func (*Claims) Descriptor() ([]byte, []int) {
	return file_claims_provider_proto_rawDescGZIP(), []int{4}
}

func (x *Claims) GetConnection() *Connection {
	if x != nil {
		return x.Connection
	}
	return nil
}

func (x *Claims) GetEmbed() *Embed {
	if x != nil {
		return x.Embed
	}
	return nil
}

func (x *Claims) GetSite() *Site {
	if x != nil {
		return x.Site
	}
	return nil
}

type Connection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rate uint64 `protobuf:"varint,1,opt,name=rate,proto3" json:"rate,omitempty"`
}

func (x *Connection) Reset() {
	*x = Connection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_claims_provider_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Connection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Connection) ProtoMessage() {}

func (x *Connection) ProtoReflect() protoreflect.Message {
	mi := &file_claims_provider_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Connection.ProtoReflect.Descriptor instead.
func (*Connection) Descriptor() ([]byte, []int) {
	return file_claims_provider_proto_rawDescGZIP(), []int{5}
}

func (x *Connection) GetRate() uint64 {
	if x != nil {
		return x.Rate
	}
	return 0
}

type Embed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NoAds bool `protobuf:"varint,1,opt,name=no_ads,json=noAds,proto3" json:"no_ads,omitempty"`
}

func (x *Embed) Reset() {
	*x = Embed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_claims_provider_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Embed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Embed) ProtoMessage() {}

func (x *Embed) ProtoReflect() protoreflect.Message {
	mi := &file_claims_provider_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Embed.ProtoReflect.Descriptor instead.
func (*Embed) Descriptor() ([]byte, []int) {
	return file_claims_provider_proto_rawDescGZIP(), []int{6}
}

func (x *Embed) GetNoAds() bool {
	if x != nil {
		return x.NoAds
	}
	return false
}

type Site struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NoAds bool `protobuf:"varint,1,opt,name=no_ads,json=noAds,proto3" json:"no_ads,omitempty"`
}

func (x *Site) Reset() {
	*x = Site{}
	if protoimpl.UnsafeEnabled {
		mi := &file_claims_provider_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Site) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Site) ProtoMessage() {}

func (x *Site) ProtoReflect() protoreflect.Message {
	mi := &file_claims_provider_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Site.ProtoReflect.Descriptor instead.
func (*Site) Descriptor() ([]byte, []int) {
	return file_claims_provider_proto_rawDescGZIP(), []int{7}
}

func (x *Site) GetNoAds() bool {
	if x != nil {
		return x.NoAds
	}
	return false
}

var File_claims_provider_proto protoreflect.FileDescriptor

var file_claims_provider_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x52, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1f,
	0x0a, 0x06, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07,
	0x2e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x52, 0x06, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x22,
	0x24, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x19, 0x0a, 0x04, 0x74, 0x69,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x54, 0x69, 0x65, 0x72, 0x52,
	0x04, 0x74, 0x69, 0x65, 0x72, 0x22, 0x2a, 0x0a, 0x04, 0x54, 0x69, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x6e, 0x0a, 0x06, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x12, 0x2b, 0x0a, 0x0a, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x05, 0x65, 0x6d, 0x62, 0x65,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x52,
	0x05, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x12, 0x19, 0x0a, 0x04, 0x73, 0x69, 0x74, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x53, 0x69, 0x74, 0x65, 0x52, 0x04, 0x73, 0x69, 0x74,
	0x65, 0x22, 0x20, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x72,
	0x61, 0x74, 0x65, 0x22, 0x1e, 0x0a, 0x05, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x12, 0x15, 0x0a, 0x06,
	0x6e, 0x6f, 0x5f, 0x61, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x6e, 0x6f,
	0x41, 0x64, 0x73, 0x22, 0x1d, 0x0a, 0x04, 0x53, 0x69, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x6e,
	0x6f, 0x5f, 0x61, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x6e, 0x6f, 0x41,
	0x64, 0x73, 0x32, 0x32, 0x0a, 0x0e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0b, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_claims_provider_proto_rawDescOnce sync.Once
	file_claims_provider_proto_rawDescData = file_claims_provider_proto_rawDesc
)

func file_claims_provider_proto_rawDescGZIP() []byte {
	file_claims_provider_proto_rawDescOnce.Do(func() {
		file_claims_provider_proto_rawDescData = protoimpl.X.CompressGZIP(file_claims_provider_proto_rawDescData)
	})
	return file_claims_provider_proto_rawDescData
}

var file_claims_provider_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_claims_provider_proto_goTypes = []interface{}{
	(*GetRequest)(nil),  // 0: GetRequest
	(*GetResponse)(nil), // 1: GetResponse
	(*Context)(nil),     // 2: Context
	(*Tier)(nil),        // 3: Tier
	(*Claims)(nil),      // 4: Claims
	(*Connection)(nil),  // 5: Connection
	(*Embed)(nil),       // 6: Embed
	(*Site)(nil),        // 7: Site
}
var file_claims_provider_proto_depIdxs = []int32{
	2, // 0: GetResponse.context:type_name -> Context
	4, // 1: GetResponse.claims:type_name -> Claims
	3, // 2: Context.tier:type_name -> Tier
	5, // 3: Claims.connection:type_name -> Connection
	6, // 4: Claims.embed:type_name -> Embed
	7, // 5: Claims.site:type_name -> Site
	0, // 6: ClaimsProvider.Get:input_type -> GetRequest
	1, // 7: ClaimsProvider.Get:output_type -> GetResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_claims_provider_proto_init() }
func file_claims_provider_proto_init() {
	if File_claims_provider_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_claims_provider_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_claims_provider_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_claims_provider_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Context); i {
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
		file_claims_provider_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tier); i {
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
		file_claims_provider_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Claims); i {
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
		file_claims_provider_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Connection); i {
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
		file_claims_provider_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Embed); i {
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
		file_claims_provider_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Site); i {
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
			RawDescriptor: file_claims_provider_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_claims_provider_proto_goTypes,
		DependencyIndexes: file_claims_provider_proto_depIdxs,
		MessageInfos:      file_claims_provider_proto_msgTypes,
	}.Build()
	File_claims_provider_proto = out.File
	file_claims_provider_proto_rawDesc = nil
	file_claims_provider_proto_goTypes = nil
	file_claims_provider_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ClaimsProviderClient is the client API for ClaimsProvider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClaimsProviderClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type claimsProviderClient struct {
	cc grpc.ClientConnInterface
}

func NewClaimsProviderClient(cc grpc.ClientConnInterface) ClaimsProviderClient {
	return &claimsProviderClient{cc}
}

func (c *claimsProviderClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/ClaimsProvider/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClaimsProviderServer is the server API for ClaimsProvider service.
type ClaimsProviderServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
}

// UnimplementedClaimsProviderServer can be embedded to have forward compatible implementations.
type UnimplementedClaimsProviderServer struct {
}

func (*UnimplementedClaimsProviderServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterClaimsProviderServer(s *grpc.Server, srv ClaimsProviderServer) {
	s.RegisterService(&_ClaimsProvider_serviceDesc, srv)
}

func _ClaimsProvider_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClaimsProviderServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ClaimsProvider/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClaimsProviderServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClaimsProvider_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ClaimsProvider",
	HandlerType: (*ClaimsProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _ClaimsProvider_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "claims-provider.proto",
}
