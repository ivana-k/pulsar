// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: profiles.proto

package __

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

type SeccompProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace    string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Application  string `protobuf:"bytes,2,opt,name=application,proto3" json:"application,omitempty"`
	Name         string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Version      string `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	Architecture string `protobuf:"bytes,5,opt,name=architecture,proto3" json:"architecture,omitempty"`
}

func (x *SeccompProfile) Reset() {
	*x = SeccompProfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profiles_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeccompProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeccompProfile) ProtoMessage() {}

func (x *SeccompProfile) ProtoReflect() protoreflect.Message {
	mi := &file_profiles_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeccompProfile.ProtoReflect.Descriptor instead.
func (*SeccompProfile) Descriptor() ([]byte, []int) {
	return file_profiles_proto_rawDescGZIP(), []int{0}
}

func (x *SeccompProfile) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *SeccompProfile) GetApplication() string {
	if x != nil {
		return x.Application
	}
	return ""
}

func (x *SeccompProfile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SeccompProfile) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *SeccompProfile) GetArchitecture() string {
	if x != nil {
		return x.Architecture
	}
	return ""
}

type Syscalls struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Names  []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
	Action string   `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
}

func (x *Syscalls) Reset() {
	*x = Syscalls{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profiles_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Syscalls) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Syscalls) ProtoMessage() {}

func (x *Syscalls) ProtoReflect() protoreflect.Message {
	mi := &file_profiles_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Syscalls.ProtoReflect.Descriptor instead.
func (*Syscalls) Descriptor() ([]byte, []int) {
	return file_profiles_proto_rawDescGZIP(), []int{1}
}

func (x *Syscalls) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

func (x *Syscalls) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

type SeccompProfileDefinition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DefaultAction string      `protobuf:"bytes,1,opt,name=defaultAction,proto3" json:"defaultAction,omitempty"`
	Architectures []string    `protobuf:"bytes,2,rep,name=architectures,proto3" json:"architectures,omitempty"`
	Syscalls      []*Syscalls `protobuf:"bytes,3,rep,name=syscalls,proto3" json:"syscalls,omitempty"`
}

func (x *SeccompProfileDefinition) Reset() {
	*x = SeccompProfileDefinition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profiles_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeccompProfileDefinition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeccompProfileDefinition) ProtoMessage() {}

func (x *SeccompProfileDefinition) ProtoReflect() protoreflect.Message {
	mi := &file_profiles_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeccompProfileDefinition.ProtoReflect.Descriptor instead.
func (*SeccompProfileDefinition) Descriptor() ([]byte, []int) {
	return file_profiles_proto_rawDescGZIP(), []int{2}
}

func (x *SeccompProfileDefinition) GetDefaultAction() string {
	if x != nil {
		return x.DefaultAction
	}
	return ""
}

func (x *SeccompProfileDefinition) GetArchitectures() []string {
	if x != nil {
		return x.Architectures
	}
	return nil
}

func (x *SeccompProfileDefinition) GetSyscalls() []*Syscalls {
	if x != nil {
		return x.Syscalls
	}
	return nil
}

type SeccompProfileDefinitionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profile    *SeccompProfile           `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	Definition *SeccompProfileDefinition `protobuf:"bytes,2,opt,name=definition,proto3" json:"definition,omitempty"`
}

func (x *SeccompProfileDefinitionRequest) Reset() {
	*x = SeccompProfileDefinitionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profiles_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeccompProfileDefinitionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeccompProfileDefinitionRequest) ProtoMessage() {}

func (x *SeccompProfileDefinitionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_profiles_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeccompProfileDefinitionRequest.ProtoReflect.Descriptor instead.
func (*SeccompProfileDefinitionRequest) Descriptor() ([]byte, []int) {
	return file_profiles_proto_rawDescGZIP(), []int{3}
}

func (x *SeccompProfileDefinitionRequest) GetProfile() *SeccompProfile {
	if x != nil {
		return x.Profile
	}
	return nil
}

func (x *SeccompProfileDefinitionRequest) GetDefinition() *SeccompProfileDefinition {
	if x != nil {
		return x.Definition
	}
	return nil
}

type ExtendSeccompProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExtendProfile *SeccompProfile `protobuf:"bytes,1,opt,name=extendProfile,proto3" json:"extendProfile,omitempty"`
	DefineProfile *SeccompProfile `protobuf:"bytes,2,opt,name=defineProfile,proto3" json:"defineProfile,omitempty"`
	Syscalls      []*Syscalls     `protobuf:"bytes,3,rep,name=syscalls,proto3" json:"syscalls,omitempty"`
}

func (x *ExtendSeccompProfileRequest) Reset() {
	*x = ExtendSeccompProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profiles_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtendSeccompProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtendSeccompProfileRequest) ProtoMessage() {}

func (x *ExtendSeccompProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_profiles_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtendSeccompProfileRequest.ProtoReflect.Descriptor instead.
func (*ExtendSeccompProfileRequest) Descriptor() ([]byte, []int) {
	return file_profiles_proto_rawDescGZIP(), []int{4}
}

func (x *ExtendSeccompProfileRequest) GetExtendProfile() *SeccompProfile {
	if x != nil {
		return x.ExtendProfile
	}
	return nil
}

func (x *ExtendSeccompProfileRequest) GetDefineProfile() *SeccompProfile {
	if x != nil {
		return x.DefineProfile
	}
	return nil
}

func (x *ExtendSeccompProfileRequest) GetSyscalls() []*Syscalls {
	if x != nil {
		return x.Syscalls
	}
	return nil
}

type GetSeccompProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Profile string `protobuf:"bytes,2,opt,name=profile,proto3" json:"profile,omitempty"`
}

func (x *GetSeccompProfileResponse) Reset() {
	*x = GetSeccompProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profiles_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSeccompProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSeccompProfileResponse) ProtoMessage() {}

func (x *GetSeccompProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_profiles_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSeccompProfileResponse.ProtoReflect.Descriptor instead.
func (*GetSeccompProfileResponse) Descriptor() ([]byte, []int) {
	return file_profiles_proto_rawDescGZIP(), []int{5}
}

func (x *GetSeccompProfileResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetSeccompProfileResponse) GetProfile() string {
	if x != nil {
		return x.Profile
	}
	return ""
}

type BasicResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *BasicResponse) Reset() {
	*x = BasicResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profiles_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BasicResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasicResponse) ProtoMessage() {}

func (x *BasicResponse) ProtoReflect() protoreflect.Message {
	mi := &file_profiles_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BasicResponse.ProtoReflect.Descriptor instead.
func (*BasicResponse) Descriptor() ([]byte, []int) {
	return file_profiles_proto_rawDescGZIP(), []int{6}
}

func (x *BasicResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *BasicResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetAllDescendantProfilesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profiles []*SeccompProfileDefinitionRequest `protobuf:"bytes,1,rep,name=profiles,proto3" json:"profiles,omitempty"`
}

func (x *GetAllDescendantProfilesResponse) Reset() {
	*x = GetAllDescendantProfilesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profiles_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllDescendantProfilesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllDescendantProfilesResponse) ProtoMessage() {}

func (x *GetAllDescendantProfilesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_profiles_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllDescendantProfilesResponse.ProtoReflect.Descriptor instead.
func (*GetAllDescendantProfilesResponse) Descriptor() ([]byte, []int) {
	return file_profiles_proto_rawDescGZIP(), []int{7}
}

func (x *GetAllDescendantProfilesResponse) GetProfiles() []*SeccompProfileDefinitionRequest {
	if x != nil {
		return x.Profiles
	}
	return nil
}

var File_profiles_proto protoreflect.FileDescriptor

var file_profiles_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x22, 0xa2, 0x01, 0x0a,
	0x0e, 0x53, 0x65, 0x63, 0x63, 0x6f, 0x6d, 0x70, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a,
	0x0c, 0x61, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x22, 0x38, 0x0a, 0x08, 0x53, 0x79, 0x73, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x98, 0x01, 0x0a, 0x18,
	0x53, 0x65, 0x63, 0x63, 0x6f, 0x6d, 0x70, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x65,
	0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24,
	0x0a, 0x0d, 0x61, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x08, 0x73, 0x79, 0x73, 0x63, 0x61, 0x6c, 0x6c, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x2e, 0x53, 0x79, 0x73, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x52, 0x08, 0x73, 0x79,
	0x73, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x22, 0x9d, 0x01, 0x0a, 0x1f, 0x53, 0x65, 0x63, 0x63, 0x6f,
	0x6d, 0x70, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x07, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x53, 0x65, 0x63, 0x63, 0x6f, 0x6d, 0x70,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x12, 0x44, 0x0a, 0x0a, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c,
	0x64, 0x2e, 0x53, 0x65, 0x63, 0x63, 0x6f, 0x6d, 0x70, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x64, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xd3, 0x01, 0x0a, 0x1b, 0x45, 0x78, 0x74, 0x65, 0x6e,
	0x64, 0x53, 0x65, 0x63, 0x63, 0x6f, 0x6d, 0x70, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x0d, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x53, 0x65, 0x63, 0x63, 0x6f,
	0x6d, 0x70, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x0d, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x64, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x40, 0x0a, 0x0d, 0x64, 0x65, 0x66, 0x69,
	0x6e, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x53, 0x65, 0x63,
	0x63, 0x6f, 0x6d, 0x70, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x0d, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x73, 0x79,
	0x73, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x53, 0x79, 0x73, 0x63, 0x61, 0x6c,
	0x6c, 0x73, 0x52, 0x08, 0x73, 0x79, 0x73, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x22, 0x4f, 0x0a, 0x19,
	0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x63, 0x6f, 0x6d, 0x70, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x43, 0x0a,
	0x0d, 0x42, 0x61, 0x73, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x6b, 0x0a, 0x20, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x65, 0x73, 0x63,
	0x65, 0x6e, 0x64, 0x61, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x53, 0x65, 0x63, 0x63, 0x6f, 0x6d, 0x70, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x32,
	0xfb, 0x03, 0x0a, 0x0e, 0x53, 0x65, 0x63, 0x63, 0x6f, 0x6d, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x60, 0x0a, 0x14, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x63, 0x63,
	0x6f, 0x6d, 0x70, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x2b, 0x2e, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x53, 0x65, 0x63, 0x63, 0x6f, 0x6d, 0x70, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x63, 0x6f,
	0x6d, 0x70, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1a, 0x2e, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x53, 0x65, 0x63, 0x63, 0x6f, 0x6d, 0x70, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x1a, 0x25, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x63, 0x6f, 0x6d, 0x70, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5c,
	0x0a, 0x14, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x53, 0x65, 0x63, 0x63, 0x6f, 0x6d, 0x70, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x27, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x53, 0x65, 0x63, 0x63, 0x6f, 0x6d,
	0x70, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x42, 0x61, 0x73,
	0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x66, 0x0a, 0x18,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x65, 0x73, 0x63, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x74,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x1a, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x53, 0x65, 0x63, 0x63, 0x6f, 0x6d, 0x70, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x1a, 0x2c, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c,
	0x64, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x65, 0x73, 0x63, 0x65, 0x6e, 0x64, 0x61,
	0x6e, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x67, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x63, 0x6f,
	0x6d, 0x70, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x42, 0x79, 0x50, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x12, 0x1a, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x53,
	0x65, 0x63, 0x63, 0x6f, 0x6d, 0x70, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x1a, 0x2c, 0x2e,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x44, 0x65, 0x73, 0x63, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x03, 0x5a,
	0x01, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_profiles_proto_rawDescOnce sync.Once
	file_profiles_proto_rawDescData = file_profiles_proto_rawDesc
)

func file_profiles_proto_rawDescGZIP() []byte {
	file_profiles_proto_rawDescOnce.Do(func() {
		file_profiles_proto_rawDescData = protoimpl.X.CompressGZIP(file_profiles_proto_rawDescData)
	})
	return file_profiles_proto_rawDescData
}

var file_profiles_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_profiles_proto_goTypes = []interface{}{
	(*SeccompProfile)(nil),                   // 0: helloworld.SeccompProfile
	(*Syscalls)(nil),                         // 1: helloworld.Syscalls
	(*SeccompProfileDefinition)(nil),         // 2: helloworld.SeccompProfileDefinition
	(*SeccompProfileDefinitionRequest)(nil),  // 3: helloworld.SeccompProfileDefinitionRequest
	(*ExtendSeccompProfileRequest)(nil),      // 4: helloworld.ExtendSeccompProfileRequest
	(*GetSeccompProfileResponse)(nil),        // 5: helloworld.GetSeccompProfileResponse
	(*BasicResponse)(nil),                    // 6: helloworld.BasicResponse
	(*GetAllDescendantProfilesResponse)(nil), // 7: helloworld.GetAllDescendantProfilesResponse
}
var file_profiles_proto_depIdxs = []int32{
	1,  // 0: helloworld.SeccompProfileDefinition.syscalls:type_name -> helloworld.Syscalls
	0,  // 1: helloworld.SeccompProfileDefinitionRequest.profile:type_name -> helloworld.SeccompProfile
	2,  // 2: helloworld.SeccompProfileDefinitionRequest.definition:type_name -> helloworld.SeccompProfileDefinition
	0,  // 3: helloworld.ExtendSeccompProfileRequest.extendProfile:type_name -> helloworld.SeccompProfile
	0,  // 4: helloworld.ExtendSeccompProfileRequest.defineProfile:type_name -> helloworld.SeccompProfile
	1,  // 5: helloworld.ExtendSeccompProfileRequest.syscalls:type_name -> helloworld.Syscalls
	3,  // 6: helloworld.GetAllDescendantProfilesResponse.profiles:type_name -> helloworld.SeccompProfileDefinitionRequest
	3,  // 7: helloworld.SeccompService.DefineSeccompProfile:input_type -> helloworld.SeccompProfileDefinitionRequest
	0,  // 8: helloworld.SeccompService.GetSeccompProfile:input_type -> helloworld.SeccompProfile
	4,  // 9: helloworld.SeccompService.ExtendSeccompProfile:input_type -> helloworld.ExtendSeccompProfileRequest
	0,  // 10: helloworld.SeccompService.GetAllDescendantProfiles:input_type -> helloworld.SeccompProfile
	0,  // 11: helloworld.SeccompService.GetSeccompProfileByPrefix:input_type -> helloworld.SeccompProfile
	6,  // 12: helloworld.SeccompService.DefineSeccompProfile:output_type -> helloworld.BasicResponse
	5,  // 13: helloworld.SeccompService.GetSeccompProfile:output_type -> helloworld.GetSeccompProfileResponse
	6,  // 14: helloworld.SeccompService.ExtendSeccompProfile:output_type -> helloworld.BasicResponse
	7,  // 15: helloworld.SeccompService.GetAllDescendantProfiles:output_type -> helloworld.GetAllDescendantProfilesResponse
	7,  // 16: helloworld.SeccompService.GetSeccompProfileByPrefix:output_type -> helloworld.GetAllDescendantProfilesResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_profiles_proto_init() }
func file_profiles_proto_init() {
	if File_profiles_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_profiles_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeccompProfile); i {
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
		file_profiles_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Syscalls); i {
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
		file_profiles_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeccompProfileDefinition); i {
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
		file_profiles_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeccompProfileDefinitionRequest); i {
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
		file_profiles_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtendSeccompProfileRequest); i {
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
		file_profiles_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSeccompProfileResponse); i {
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
		file_profiles_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BasicResponse); i {
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
		file_profiles_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllDescendantProfilesResponse); i {
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
			RawDescriptor: file_profiles_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_profiles_proto_goTypes,
		DependencyIndexes: file_profiles_proto_depIdxs,
		MessageInfos:      file_profiles_proto_msgTypes,
	}.Build()
	File_profiles_proto = out.File
	file_profiles_proto_rawDesc = nil
	file_profiles_proto_goTypes = nil
	file_profiles_proto_depIdxs = nil
}
