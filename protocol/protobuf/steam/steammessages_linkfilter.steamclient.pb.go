// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: steammessages_linkfilter.steamclient.proto

package steam

import (
	proto "github.com/golang/protobuf/proto"
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

type CCommunity_GetLinkFilterHashPrefixes_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HitType *uint32 `protobuf:"varint,1,opt,name=hit_type,json=hitType" json:"hit_type,omitempty"`
	Count   *uint32 `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
	Start   *uint64 `protobuf:"varint,3,opt,name=start" json:"start,omitempty"`
}

func (x *CCommunity_GetLinkFilterHashPrefixes_Request) Reset() {
	*x = CCommunity_GetLinkFilterHashPrefixes_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_linkfilter_steamclient_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CCommunity_GetLinkFilterHashPrefixes_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CCommunity_GetLinkFilterHashPrefixes_Request) ProtoMessage() {}

func (x *CCommunity_GetLinkFilterHashPrefixes_Request) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_linkfilter_steamclient_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CCommunity_GetLinkFilterHashPrefixes_Request.ProtoReflect.Descriptor instead.
func (*CCommunity_GetLinkFilterHashPrefixes_Request) Descriptor() ([]byte, []int) {
	return file_steammessages_linkfilter_steamclient_proto_rawDescGZIP(), []int{0}
}

func (x *CCommunity_GetLinkFilterHashPrefixes_Request) GetHitType() uint32 {
	if x != nil && x.HitType != nil {
		return *x.HitType
	}
	return 0
}

func (x *CCommunity_GetLinkFilterHashPrefixes_Request) GetCount() uint32 {
	if x != nil && x.Count != nil {
		return *x.Count
	}
	return 0
}

func (x *CCommunity_GetLinkFilterHashPrefixes_Request) GetStart() uint64 {
	if x != nil && x.Start != nil {
		return *x.Start
	}
	return 0
}

type CCommunity_GetLinkFilterHashPrefixes_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HashPrefixes []uint32 `protobuf:"varint,1,rep,name=hash_prefixes,json=hashPrefixes" json:"hash_prefixes,omitempty"`
}

func (x *CCommunity_GetLinkFilterHashPrefixes_Response) Reset() {
	*x = CCommunity_GetLinkFilterHashPrefixes_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_linkfilter_steamclient_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CCommunity_GetLinkFilterHashPrefixes_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CCommunity_GetLinkFilterHashPrefixes_Response) ProtoMessage() {}

func (x *CCommunity_GetLinkFilterHashPrefixes_Response) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_linkfilter_steamclient_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CCommunity_GetLinkFilterHashPrefixes_Response.ProtoReflect.Descriptor instead.
func (*CCommunity_GetLinkFilterHashPrefixes_Response) Descriptor() ([]byte, []int) {
	return file_steammessages_linkfilter_steamclient_proto_rawDescGZIP(), []int{1}
}

func (x *CCommunity_GetLinkFilterHashPrefixes_Response) GetHashPrefixes() []uint32 {
	if x != nil {
		return x.HashPrefixes
	}
	return nil
}

type CCommunity_GetLinkFilterHashes_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HitType *uint32 `protobuf:"varint,1,opt,name=hit_type,json=hitType" json:"hit_type,omitempty"`
	Count   *uint32 `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
	Start   *uint64 `protobuf:"varint,3,opt,name=start" json:"start,omitempty"`
}

func (x *CCommunity_GetLinkFilterHashes_Request) Reset() {
	*x = CCommunity_GetLinkFilterHashes_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_linkfilter_steamclient_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CCommunity_GetLinkFilterHashes_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CCommunity_GetLinkFilterHashes_Request) ProtoMessage() {}

func (x *CCommunity_GetLinkFilterHashes_Request) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_linkfilter_steamclient_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CCommunity_GetLinkFilterHashes_Request.ProtoReflect.Descriptor instead.
func (*CCommunity_GetLinkFilterHashes_Request) Descriptor() ([]byte, []int) {
	return file_steammessages_linkfilter_steamclient_proto_rawDescGZIP(), []int{2}
}

func (x *CCommunity_GetLinkFilterHashes_Request) GetHitType() uint32 {
	if x != nil && x.HitType != nil {
		return *x.HitType
	}
	return 0
}

func (x *CCommunity_GetLinkFilterHashes_Request) GetCount() uint32 {
	if x != nil && x.Count != nil {
		return *x.Count
	}
	return 0
}

func (x *CCommunity_GetLinkFilterHashes_Request) GetStart() uint64 {
	if x != nil && x.Start != nil {
		return *x.Start
	}
	return 0
}

type CCommunity_GetLinkFilterHashes_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hashes [][]byte `protobuf:"bytes,1,rep,name=hashes" json:"hashes,omitempty"`
}

func (x *CCommunity_GetLinkFilterHashes_Response) Reset() {
	*x = CCommunity_GetLinkFilterHashes_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_linkfilter_steamclient_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CCommunity_GetLinkFilterHashes_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CCommunity_GetLinkFilterHashes_Response) ProtoMessage() {}

func (x *CCommunity_GetLinkFilterHashes_Response) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_linkfilter_steamclient_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CCommunity_GetLinkFilterHashes_Response.ProtoReflect.Descriptor instead.
func (*CCommunity_GetLinkFilterHashes_Response) Descriptor() ([]byte, []int) {
	return file_steammessages_linkfilter_steamclient_proto_rawDescGZIP(), []int{3}
}

func (x *CCommunity_GetLinkFilterHashes_Response) GetHashes() [][]byte {
	if x != nil {
		return x.Hashes
	}
	return nil
}

type CCommunity_GetLinkFilterListVersion_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HitType *uint32 `protobuf:"varint,1,opt,name=hit_type,json=hitType" json:"hit_type,omitempty"`
}

func (x *CCommunity_GetLinkFilterListVersion_Request) Reset() {
	*x = CCommunity_GetLinkFilterListVersion_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_linkfilter_steamclient_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CCommunity_GetLinkFilterListVersion_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CCommunity_GetLinkFilterListVersion_Request) ProtoMessage() {}

func (x *CCommunity_GetLinkFilterListVersion_Request) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_linkfilter_steamclient_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CCommunity_GetLinkFilterListVersion_Request.ProtoReflect.Descriptor instead.
func (*CCommunity_GetLinkFilterListVersion_Request) Descriptor() ([]byte, []int) {
	return file_steammessages_linkfilter_steamclient_proto_rawDescGZIP(), []int{4}
}

func (x *CCommunity_GetLinkFilterListVersion_Request) GetHitType() uint32 {
	if x != nil && x.HitType != nil {
		return *x.HitType
	}
	return 0
}

type CCommunity_GetLinkFilterListVersion_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version *string `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	Count   *uint64 `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
}

func (x *CCommunity_GetLinkFilterListVersion_Response) Reset() {
	*x = CCommunity_GetLinkFilterListVersion_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_linkfilter_steamclient_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CCommunity_GetLinkFilterListVersion_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CCommunity_GetLinkFilterListVersion_Response) ProtoMessage() {}

func (x *CCommunity_GetLinkFilterListVersion_Response) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_linkfilter_steamclient_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CCommunity_GetLinkFilterListVersion_Response.ProtoReflect.Descriptor instead.
func (*CCommunity_GetLinkFilterListVersion_Response) Descriptor() ([]byte, []int) {
	return file_steammessages_linkfilter_steamclient_proto_rawDescGZIP(), []int{5}
}

func (x *CCommunity_GetLinkFilterListVersion_Response) GetVersion() string {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return ""
}

func (x *CCommunity_GetLinkFilterListVersion_Response) GetCount() uint64 {
	if x != nil && x.Count != nil {
		return *x.Count
	}
	return 0
}

var File_steammessages_linkfilter_steamclient_proto protoreflect.FileDescriptor

var file_steammessages_linkfilter_steamclient_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5f,
	0x6c, 0x69, 0x6e, 0x6b, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x73, 0x74,
	0x65, 0x61, 0x6d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x75, 0x6e, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x02, 0x0a, 0x2c, 0x43,
	0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x5f, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6e,
	0x6b, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x48, 0x61, 0x73, 0x68, 0x50, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x65, 0x73, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x50, 0x0a, 0x08, 0x68,
	0x69, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x35, 0x82,
	0xb5, 0x18, 0x31, 0x54, 0x68, 0x65, 0x20, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x64,
	0x20, 0x68, 0x69, 0x74, 0x73, 0x20, 0x77, 0x69, 0x6c, 0x6c, 0x20, 0x62, 0x65, 0x20, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x52, 0x07, 0x68, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x63, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x4d, 0x82, 0xb5,
	0x18, 0x49, 0x54, 0x68, 0x65, 0x20, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x20, 0x6f, 0x66, 0x20,
	0x68, 0x69, 0x74, 0x73, 0x20, 0x74, 0x6f, 0x20, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65,
	0x20, 0x69, 0x6e, 0x20, 0x61, 0x20, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x20, 0x62, 0x61, 0x74,
	0x63, 0x68, 0x2e, 0x20, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x79, 0x20, 0x30, 0x20, 0x66, 0x6f,
	0x72, 0x20, 0x6e, 0x6f, 0x20, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x4d, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x42, 0x37, 0x82, 0xb5, 0x18, 0x33, 0x54, 0x68, 0x65, 0x20, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x20, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x72,
	0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x69, 0x6e, 0x67, 0x20, 0x68, 0x69, 0x74, 0x73, 0x20, 0x69,
	0x6e, 0x20, 0x62, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x2e, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x22, 0x95, 0x01, 0x0a, 0x2d, 0x43, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79,
	0x5f, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x48, 0x61,
	0x73, 0x68, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x73, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x64, 0x0a, 0x0d, 0x68, 0x61, 0x73, 0x68, 0x5f, 0x70, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x42, 0x3f, 0x82, 0xb5, 0x18, 0x3b,
	0x54, 0x68, 0x65, 0x20, 0x66, 0x69, 0x72, 0x73, 0x74, 0x20, 0x33, 0x32, 0x20, 0x62, 0x69, 0x74,
	0x73, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x53, 0x48, 0x41, 0x31, 0x20, 0x68, 0x61,
	0x73, 0x68, 0x65, 0x73, 0x20, 0x6f, 0x66, 0x20, 0x65, 0x61, 0x63, 0x68, 0x20, 0x68, 0x69, 0x74,
	0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x2e, 0x52, 0x0c, 0x68, 0x61, 0x73,
	0x68, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x73, 0x22, 0xae, 0x02, 0x0a, 0x26, 0x43, 0x43,
	0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x5f, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x5f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x50, 0x0a, 0x08, 0x68, 0x69, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x35, 0x82, 0xb5, 0x18, 0x31, 0x54, 0x68, 0x65, 0x20,
	0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x64, 0x20, 0x68, 0x69, 0x74, 0x73, 0x20, 0x77,
	0x69, 0x6c, 0x6c, 0x20, 0x62, 0x65, 0x20, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x65, 0x64, 0x20,
	0x74, 0x6f, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x52, 0x07, 0x68,
	0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x63, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x4d, 0x82, 0xb5, 0x18, 0x49, 0x54, 0x68, 0x65, 0x20, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x20, 0x6f, 0x66, 0x20, 0x68, 0x69, 0x74, 0x73, 0x20, 0x74, 0x6f,
	0x20, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x20, 0x69, 0x6e, 0x20, 0x61, 0x20, 0x73,
	0x69, 0x6e, 0x67, 0x6c, 0x65, 0x20, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x20, 0x53, 0x70, 0x65,
	0x63, 0x69, 0x66, 0x79, 0x20, 0x30, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x6e, 0x6f, 0x20, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x2e, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x4d, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x42, 0x37, 0x82, 0xb5, 0x18, 0x33,
	0x54, 0x68, 0x65, 0x20, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x20, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x69,
	0x6e, 0x67, 0x20, 0x68, 0x69, 0x74, 0x73, 0x20, 0x69, 0x6e, 0x20, 0x62, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x73, 0x2e, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x22, 0x72, 0x0a, 0x27, 0x43, 0x43,
	0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x5f, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x5f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x06, 0x68, 0x61, 0x73, 0x68, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0c, 0x42, 0x2f, 0x82, 0xb5, 0x18, 0x2b, 0x41, 0x20, 0x6c, 0x69, 0x73,
	0x74, 0x20, 0x6f, 0x66, 0x20, 0x68, 0x61, 0x73, 0x68, 0x65, 0x73, 0x20, 0x72, 0x65, 0x74, 0x75,
	0x72, 0x6e, 0x65, 0x64, 0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x6f, 0x75, 0x72, 0x20, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x52, 0x06, 0x68, 0x61, 0x73, 0x68, 0x65, 0x73, 0x22, 0x89,
	0x01, 0x0a, 0x2b, 0x43, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x5f, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x5a,
	0x0a, 0x08, 0x68, 0x69, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x42, 0x3f, 0x82, 0xb5, 0x18, 0x3b, 0x54, 0x68, 0x65, 0x20, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x20, 0x68, 0x61, 0x73, 0x68, 0x65, 0x73, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x74, 0x68, 0x69,
	0x73, 0x20, 0x74, 0x79, 0x70, 0x65, 0x27, 0x73, 0x20, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x20,
	0x77, 0x69, 0x6c, 0x6c, 0x20, 0x62, 0x65, 0x20, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64,
	0x2e, 0x52, 0x07, 0x68, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0xba, 0x01, 0x0a, 0x2c, 0x43,
	0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x5f, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6e,
	0x6b, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x35, 0x82, 0xb5,
	0x18, 0x31, 0x41, 0x20, 0x68, 0x61, 0x73, 0x68, 0x20, 0x62, 0x75, 0x69, 0x6c, 0x74, 0x20, 0x75,
	0x73, 0x69, 0x6e, 0x67, 0x20, 0x74, 0x68, 0x65, 0x20, 0x49, 0x44, 0x73, 0x20, 0x6f, 0x66, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x20, 0x68, 0x69,
	0x74, 0x73, 0x2e, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x23, 0x82, 0xb5, 0x18,
	0x1f, 0x54, 0x68, 0x65, 0x20, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x20, 0x6f, 0x66, 0x20, 0x68,
	0x69, 0x74, 0x73, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x74, 0x79, 0x70, 0x65,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xc6, 0x05, 0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x74, 0x79, 0x4c, 0x69, 0x6e, 0x6b, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12,
	0xd2, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x48, 0x61, 0x73, 0x68, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x73, 0x12, 0x2d, 0x2e,
	0x43, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x5f, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x6e, 0x6b, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x48, 0x61, 0x73, 0x68, 0x50, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x65, 0x73, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x43,
	0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x5f, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6e,
	0x6b, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x48, 0x61, 0x73, 0x68, 0x50, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x65, 0x73, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x56, 0x82, 0xb5,
	0x18, 0x52, 0x47, 0x65, 0x74, 0x20, 0x61, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20,
	0x68, 0x61, 0x73, 0x68, 0x20, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x73, 0x20, 0x66, 0x6f,
	0x72, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x20,
	0x68, 0x69, 0x74, 0x20, 0x74, 0x79, 0x70, 0x65, 0x2c, 0x20, 0x74, 0x6f, 0x20, 0x75, 0x73, 0x65,
	0x20, 0x66, 0x6f, 0x72, 0x20, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x20, 0x63, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x67, 0x2e, 0x12, 0xb9, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x12, 0x27, 0x2e, 0x43,
	0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x5f, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6e,
	0x6b, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x5f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x43, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x5f, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x4f, 0x82, 0xb5, 0x18, 0x4b, 0x47, 0x65, 0x74, 0x20, 0x61, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20,
	0x6f, 0x66, 0x20, 0x68, 0x61, 0x73, 0x68, 0x65, 0x73, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x74, 0x68,
	0x65, 0x20, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x20, 0x68, 0x69, 0x74, 0x20,
	0x74, 0x79, 0x70, 0x65, 0x2c, 0x20, 0x74, 0x6f, 0x20, 0x75, 0x73, 0x65, 0x20, 0x66, 0x6f, 0x72,
	0x20, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x20, 0x63, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2e,
	0x12, 0xcd, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x2e,
	0x43, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x5f, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x6e, 0x6b, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x43, 0x43,
	0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x5f, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x54, 0x82, 0xb5, 0x18, 0x50,
	0x47, 0x65, 0x74, 0x20, 0x61, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x68, 0x61,
	0x73, 0x68, 0x65, 0x73, 0x20, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x69, 0x6e, 0x67, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x6f, 0x66, 0x20, 0x65,
	0x61, 0x63, 0x68, 0x20, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x20, 0x73, 0x69, 0x7a, 0x65, 0x2e,
	0x1a, 0x4e, 0x82, 0xb5, 0x18, 0x4a, 0x41, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x20,
	0x66, 0x6f, 0x72, 0x20, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x20, 0x64, 0x61,
	0x74, 0x61, 0x20, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x20, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x20, 0x43,
	0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x20, 0x70, 0x68, 0x69, 0x73, 0x68, 0x69, 0x6e,
	0x67, 0x20, 0x6c, 0x69, 0x6e, 0x6b, 0x20, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67,
	0x42, 0x1e, 0x5a, 0x19, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x80, 0x01, 0x01,
}

var (
	file_steammessages_linkfilter_steamclient_proto_rawDescOnce sync.Once
	file_steammessages_linkfilter_steamclient_proto_rawDescData = file_steammessages_linkfilter_steamclient_proto_rawDesc
)

func file_steammessages_linkfilter_steamclient_proto_rawDescGZIP() []byte {
	file_steammessages_linkfilter_steamclient_proto_rawDescOnce.Do(func() {
		file_steammessages_linkfilter_steamclient_proto_rawDescData = protoimpl.X.CompressGZIP(file_steammessages_linkfilter_steamclient_proto_rawDescData)
	})
	return file_steammessages_linkfilter_steamclient_proto_rawDescData
}

var file_steammessages_linkfilter_steamclient_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_steammessages_linkfilter_steamclient_proto_goTypes = []interface{}{
	(*CCommunity_GetLinkFilterHashPrefixes_Request)(nil),  // 0: CCommunity_GetLinkFilterHashPrefixes_Request
	(*CCommunity_GetLinkFilterHashPrefixes_Response)(nil), // 1: CCommunity_GetLinkFilterHashPrefixes_Response
	(*CCommunity_GetLinkFilterHashes_Request)(nil),        // 2: CCommunity_GetLinkFilterHashes_Request
	(*CCommunity_GetLinkFilterHashes_Response)(nil),       // 3: CCommunity_GetLinkFilterHashes_Response
	(*CCommunity_GetLinkFilterListVersion_Request)(nil),   // 4: CCommunity_GetLinkFilterListVersion_Request
	(*CCommunity_GetLinkFilterListVersion_Response)(nil),  // 5: CCommunity_GetLinkFilterListVersion_Response
}
var file_steammessages_linkfilter_steamclient_proto_depIdxs = []int32{
	0, // 0: CommunityLinkFilter.GetLinkFilterHashPrefixes:input_type -> CCommunity_GetLinkFilterHashPrefixes_Request
	2, // 1: CommunityLinkFilter.GetLinkFilterHashes:input_type -> CCommunity_GetLinkFilterHashes_Request
	4, // 2: CommunityLinkFilter.GetLinkFilterListVersion:input_type -> CCommunity_GetLinkFilterListVersion_Request
	1, // 3: CommunityLinkFilter.GetLinkFilterHashPrefixes:output_type -> CCommunity_GetLinkFilterHashPrefixes_Response
	3, // 4: CommunityLinkFilter.GetLinkFilterHashes:output_type -> CCommunity_GetLinkFilterHashes_Response
	5, // 5: CommunityLinkFilter.GetLinkFilterListVersion:output_type -> CCommunity_GetLinkFilterListVersion_Response
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_steammessages_linkfilter_steamclient_proto_init() }
func file_steammessages_linkfilter_steamclient_proto_init() {
	if File_steammessages_linkfilter_steamclient_proto != nil {
		return
	}
	file_steammessages_unified_base_steamclient_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_steammessages_linkfilter_steamclient_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CCommunity_GetLinkFilterHashPrefixes_Request); i {
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
		file_steammessages_linkfilter_steamclient_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CCommunity_GetLinkFilterHashPrefixes_Response); i {
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
		file_steammessages_linkfilter_steamclient_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CCommunity_GetLinkFilterHashes_Request); i {
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
		file_steammessages_linkfilter_steamclient_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CCommunity_GetLinkFilterHashes_Response); i {
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
		file_steammessages_linkfilter_steamclient_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CCommunity_GetLinkFilterListVersion_Request); i {
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
		file_steammessages_linkfilter_steamclient_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CCommunity_GetLinkFilterListVersion_Response); i {
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
			RawDescriptor: file_steammessages_linkfilter_steamclient_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_steammessages_linkfilter_steamclient_proto_goTypes,
		DependencyIndexes: file_steammessages_linkfilter_steamclient_proto_depIdxs,
		MessageInfos:      file_steammessages_linkfilter_steamclient_proto_msgTypes,
	}.Build()
	File_steammessages_linkfilter_steamclient_proto = out.File
	file_steammessages_linkfilter_steamclient_proto_rawDesc = nil
	file_steammessages_linkfilter_steamclient_proto_goTypes = nil
	file_steammessages_linkfilter_steamclient_proto_depIdxs = nil
}
