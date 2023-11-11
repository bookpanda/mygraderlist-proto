// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.24.4
// source: emoji.proto

package emoji

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

type Emoji struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Emoji     string `protobuf:"bytes,2,opt,name=emoji,proto3" json:"emoji,omitempty"`
	ProblemId string `protobuf:"bytes,3,opt,name=problemId,proto3" json:"problemId,omitempty"`
	UserId    string `protobuf:"bytes,4,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *Emoji) Reset() {
	*x = Emoji{}
	if protoimpl.UnsafeEnabled {
		mi := &file_emoji_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Emoji) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Emoji) ProtoMessage() {}

func (x *Emoji) ProtoReflect() protoreflect.Message {
	mi := &file_emoji_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Emoji.ProtoReflect.Descriptor instead.
func (*Emoji) Descriptor() ([]byte, []int) {
	return file_emoji_proto_rawDescGZIP(), []int{0}
}

func (x *Emoji) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Emoji) GetEmoji() string {
	if x != nil {
		return x.Emoji
	}
	return ""
}

func (x *Emoji) GetProblemId() string {
	if x != nil {
		return x.ProblemId
	}
	return ""
}

func (x *Emoji) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// FindAll
type FindAllEmojiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FindAllEmojiRequest) Reset() {
	*x = FindAllEmojiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_emoji_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllEmojiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllEmojiRequest) ProtoMessage() {}

func (x *FindAllEmojiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_emoji_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllEmojiRequest.ProtoReflect.Descriptor instead.
func (*FindAllEmojiRequest) Descriptor() ([]byte, []int) {
	return file_emoji_proto_rawDescGZIP(), []int{1}
}

type FindAllEmojiResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Emojis []*Emoji `protobuf:"bytes,1,rep,name=emojis,proto3" json:"emojis,omitempty"`
}

func (x *FindAllEmojiResponse) Reset() {
	*x = FindAllEmojiResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_emoji_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllEmojiResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllEmojiResponse) ProtoMessage() {}

func (x *FindAllEmojiResponse) ProtoReflect() protoreflect.Message {
	mi := &file_emoji_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllEmojiResponse.ProtoReflect.Descriptor instead.
func (*FindAllEmojiResponse) Descriptor() ([]byte, []int) {
	return file_emoji_proto_rawDescGZIP(), []int{2}
}

func (x *FindAllEmojiResponse) GetEmojis() []*Emoji {
	if x != nil {
		return x.Emojis
	}
	return nil
}

// FindByUserId
type FindByUserIdEmojiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *FindByUserIdEmojiRequest) Reset() {
	*x = FindByUserIdEmojiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_emoji_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindByUserIdEmojiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindByUserIdEmojiRequest) ProtoMessage() {}

func (x *FindByUserIdEmojiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_emoji_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindByUserIdEmojiRequest.ProtoReflect.Descriptor instead.
func (*FindByUserIdEmojiRequest) Descriptor() ([]byte, []int) {
	return file_emoji_proto_rawDescGZIP(), []int{3}
}

func (x *FindByUserIdEmojiRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type FindByUserIdEmojiResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Emojis []*Emoji `protobuf:"bytes,1,rep,name=emojis,proto3" json:"emojis,omitempty"`
}

func (x *FindByUserIdEmojiResponse) Reset() {
	*x = FindByUserIdEmojiResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_emoji_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindByUserIdEmojiResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindByUserIdEmojiResponse) ProtoMessage() {}

func (x *FindByUserIdEmojiResponse) ProtoReflect() protoreflect.Message {
	mi := &file_emoji_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindByUserIdEmojiResponse.ProtoReflect.Descriptor instead.
func (*FindByUserIdEmojiResponse) Descriptor() ([]byte, []int) {
	return file_emoji_proto_rawDescGZIP(), []int{4}
}

func (x *FindByUserIdEmojiResponse) GetEmojis() []*Emoji {
	if x != nil {
		return x.Emojis
	}
	return nil
}

// Create
type CreateEmojiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Emoji *Emoji `protobuf:"bytes,1,opt,name=emoji,proto3" json:"emoji,omitempty"`
}

func (x *CreateEmojiRequest) Reset() {
	*x = CreateEmojiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_emoji_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEmojiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEmojiRequest) ProtoMessage() {}

func (x *CreateEmojiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_emoji_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEmojiRequest.ProtoReflect.Descriptor instead.
func (*CreateEmojiRequest) Descriptor() ([]byte, []int) {
	return file_emoji_proto_rawDescGZIP(), []int{5}
}

func (x *CreateEmojiRequest) GetEmoji() *Emoji {
	if x != nil {
		return x.Emoji
	}
	return nil
}

type CreateEmojiResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Emoji *Emoji `protobuf:"bytes,1,opt,name=emoji,proto3" json:"emoji,omitempty"`
}

func (x *CreateEmojiResponse) Reset() {
	*x = CreateEmojiResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_emoji_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEmojiResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEmojiResponse) ProtoMessage() {}

func (x *CreateEmojiResponse) ProtoReflect() protoreflect.Message {
	mi := &file_emoji_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEmojiResponse.ProtoReflect.Descriptor instead.
func (*CreateEmojiResponse) Descriptor() ([]byte, []int) {
	return file_emoji_proto_rawDescGZIP(), []int{6}
}

func (x *CreateEmojiResponse) GetEmoji() *Emoji {
	if x != nil {
		return x.Emoji
	}
	return nil
}

// Delete
type DeleteEmojiRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteEmojiRequest) Reset() {
	*x = DeleteEmojiRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_emoji_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEmojiRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEmojiRequest) ProtoMessage() {}

func (x *DeleteEmojiRequest) ProtoReflect() protoreflect.Message {
	mi := &file_emoji_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEmojiRequest.ProtoReflect.Descriptor instead.
func (*DeleteEmojiRequest) Descriptor() ([]byte, []int) {
	return file_emoji_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteEmojiRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteEmojiResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteEmojiResponse) Reset() {
	*x = DeleteEmojiResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_emoji_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEmojiResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEmojiResponse) ProtoMessage() {}

func (x *DeleteEmojiResponse) ProtoReflect() protoreflect.Message {
	mi := &file_emoji_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEmojiResponse.ProtoReflect.Descriptor instead.
func (*DeleteEmojiResponse) Descriptor() ([]byte, []int) {
	return file_emoji_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteEmojiResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_emoji_proto protoreflect.FileDescriptor

var file_emoji_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65,
	0x6d, 0x6f, 0x6a, 0x69, 0x22, 0x63, 0x0a, 0x05, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x6f, 0x6a, 0x69, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x46, 0x69, 0x6e,
	0x64, 0x41, 0x6c, 0x6c, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x3c, 0x0a, 0x14, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45, 0x6d, 0x6f, 0x6a, 0x69,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x65, 0x6d, 0x6f, 0x6a,
	0x69, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x65, 0x6d, 0x6f, 0x6a, 0x69,
	0x2e, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x52, 0x06, 0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x73, 0x22, 0x32,
	0x0a, 0x18, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x45, 0x6d,
	0x6f, 0x6a, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x41, 0x0a, 0x19, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x24, 0x0a, 0x06, 0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x2e, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x52, 0x06, 0x65,
	0x6d, 0x6f, 0x6a, 0x69, 0x73, 0x22, 0x38, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45,
	0x6d, 0x6f, 0x6a, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x05, 0x65,
	0x6d, 0x6f, 0x6a, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x65, 0x6d, 0x6f,
	0x6a, 0x69, 0x2e, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x52, 0x05, 0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x22,
	0x39, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x2e, 0x45, 0x6d,
	0x6f, 0x6a, 0x69, 0x52, 0x05, 0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x22, 0x24, 0x0a, 0x12, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x2f, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x32, 0xb9, 0x02, 0x0a, 0x0c, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x4e, 0x0a, 0x07, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x12, 0x1f, 0x2e,
	0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20,
	0x2e, 0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x53, 0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1f, 0x2e, 0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x42,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x2e, 0x46, 0x69, 0x6e, 0x64,
	0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x19, 0x2e, 0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x65,
	0x6d, 0x6f, 0x6a, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x6f, 0x6a, 0x69,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x06, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6d,
	0x6f, 0x6a, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1c, 0x5a,
	0x1a, 0x4d, 0x79, 0x47, 0x72, 0x61, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x2f, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_emoji_proto_rawDescOnce sync.Once
	file_emoji_proto_rawDescData = file_emoji_proto_rawDesc
)

func file_emoji_proto_rawDescGZIP() []byte {
	file_emoji_proto_rawDescOnce.Do(func() {
		file_emoji_proto_rawDescData = protoimpl.X.CompressGZIP(file_emoji_proto_rawDescData)
	})
	return file_emoji_proto_rawDescData
}

var file_emoji_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_emoji_proto_goTypes = []interface{}{
	(*Emoji)(nil),                     // 0: emoji.Emoji
	(*FindAllEmojiRequest)(nil),       // 1: emoji.FindAllEmojiRequest
	(*FindAllEmojiResponse)(nil),      // 2: emoji.FindAllEmojiResponse
	(*FindByUserIdEmojiRequest)(nil),  // 3: emoji.FindByUserIdEmojiRequest
	(*FindByUserIdEmojiResponse)(nil), // 4: emoji.FindByUserIdEmojiResponse
	(*CreateEmojiRequest)(nil),        // 5: emoji.CreateEmojiRequest
	(*CreateEmojiResponse)(nil),       // 6: emoji.CreateEmojiResponse
	(*DeleteEmojiRequest)(nil),        // 7: emoji.DeleteEmojiRequest
	(*DeleteEmojiResponse)(nil),       // 8: emoji.DeleteEmojiResponse
}
var file_emoji_proto_depIdxs = []int32{
	0, // 0: emoji.FindAllEmojiResponse.emojis:type_name -> emoji.Emoji
	0, // 1: emoji.FindByUserIdEmojiResponse.emojis:type_name -> emoji.Emoji
	0, // 2: emoji.CreateEmojiRequest.emoji:type_name -> emoji.Emoji
	0, // 3: emoji.CreateEmojiResponse.emoji:type_name -> emoji.Emoji
	3, // 4: emoji.EmojiService.FindAll:input_type -> emoji.FindByUserIdEmojiRequest
	3, // 5: emoji.EmojiService.FindByUserId:input_type -> emoji.FindByUserIdEmojiRequest
	5, // 6: emoji.EmojiService.Create:input_type -> emoji.CreateEmojiRequest
	7, // 7: emoji.EmojiService.Delete:input_type -> emoji.DeleteEmojiRequest
	4, // 8: emoji.EmojiService.FindAll:output_type -> emoji.FindByUserIdEmojiResponse
	4, // 9: emoji.EmojiService.FindByUserId:output_type -> emoji.FindByUserIdEmojiResponse
	6, // 10: emoji.EmojiService.Create:output_type -> emoji.CreateEmojiResponse
	8, // 11: emoji.EmojiService.Delete:output_type -> emoji.DeleteEmojiResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_emoji_proto_init() }
func file_emoji_proto_init() {
	if File_emoji_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_emoji_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Emoji); i {
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
		file_emoji_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllEmojiRequest); i {
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
		file_emoji_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllEmojiResponse); i {
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
		file_emoji_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindByUserIdEmojiRequest); i {
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
		file_emoji_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindByUserIdEmojiResponse); i {
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
		file_emoji_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEmojiRequest); i {
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
		file_emoji_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEmojiResponse); i {
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
		file_emoji_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEmojiRequest); i {
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
		file_emoji_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEmojiResponse); i {
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
			RawDescriptor: file_emoji_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_emoji_proto_goTypes,
		DependencyIndexes: file_emoji_proto_depIdxs,
		MessageInfos:      file_emoji_proto_msgTypes,
	}.Build()
	File_emoji_proto = out.File
	file_emoji_proto_rawDesc = nil
	file_emoji_proto_goTypes = nil
	file_emoji_proto_depIdxs = nil
}
