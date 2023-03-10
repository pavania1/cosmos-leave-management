// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.3
// source: proto/cosmos/lms/v1beta1/quer.proto

package types

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

type GetstudentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetstudentsRequest) Reset() {
	*x = GetstudentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cosmos_lms_v1beta1_quer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetstudentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetstudentsRequest) ProtoMessage() {}

func (x *GetstudentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cosmos_lms_v1beta1_quer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetstudentsRequest.ProtoReflect.Descriptor instead.
func (*GetstudentsRequest) Descriptor() ([]byte, []int) {
	return file_proto_cosmos_lms_v1beta1_quer_proto_rawDescGZIP(), []int{0}
}

type GetstudentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Students []*Student `protobuf:"bytes,1,rep,name=students,proto3" json:"students,omitempty"`
}

func (x *GetstudentsResponse) Reset() {
	*x = GetstudentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cosmos_lms_v1beta1_quer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetstudentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetstudentsResponse) ProtoMessage() {}

func (x *GetstudentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cosmos_lms_v1beta1_quer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetstudentsResponse.ProtoReflect.Descriptor instead.
func (*GetstudentsResponse) Descriptor() ([]byte, []int) {
	return file_proto_cosmos_lms_v1beta1_quer_proto_rawDescGZIP(), []int{1}
}

func (x *GetstudentsResponse) GetStudents() []*Student {
	if x != nil {
		return x.Students
	}
	return nil
}

type GetLeaveRequestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetLeaveRequestRequest) Reset() {
	*x = GetLeaveRequestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cosmos_lms_v1beta1_quer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLeaveRequestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLeaveRequestRequest) ProtoMessage() {}

func (x *GetLeaveRequestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cosmos_lms_v1beta1_quer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLeaveRequestRequest.ProtoReflect.Descriptor instead.
func (*GetLeaveRequestRequest) Descriptor() ([]byte, []int) {
	return file_proto_cosmos_lms_v1beta1_quer_proto_rawDescGZIP(), []int{2}
}

type GetLeaveRequestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Leaverequest []*ApplyLeaveRequest `protobuf:"bytes,1,rep,name=leaverequest,proto3" json:"leaverequest,omitempty"`
}

func (x *GetLeaveRequestResponse) Reset() {
	*x = GetLeaveRequestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cosmos_lms_v1beta1_quer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLeaveRequestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLeaveRequestResponse) ProtoMessage() {}

func (x *GetLeaveRequestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cosmos_lms_v1beta1_quer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLeaveRequestResponse.ProtoReflect.Descriptor instead.
func (*GetLeaveRequestResponse) Descriptor() ([]byte, []int) {
	return file_proto_cosmos_lms_v1beta1_quer_proto_rawDescGZIP(), []int{3}
}

func (x *GetLeaveRequestResponse) GetLeaverequest() []*ApplyLeaveRequest {
	if x != nil {
		return x.Leaverequest
	}
	return nil
}

type GetLeaveApproveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetLeaveApproveRequest) Reset() {
	*x = GetLeaveApproveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cosmos_lms_v1beta1_quer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLeaveApproveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLeaveApproveRequest) ProtoMessage() {}

func (x *GetLeaveApproveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cosmos_lms_v1beta1_quer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLeaveApproveRequest.ProtoReflect.Descriptor instead.
func (*GetLeaveApproveRequest) Descriptor() ([]byte, []int) {
	return file_proto_cosmos_lms_v1beta1_quer_proto_rawDescGZIP(), []int{4}
}

type GetLeaveApproveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Approveleaverequest []*AcceptLeaveRequest `protobuf:"bytes,1,rep,name=Approveleaverequest,proto3" json:"Approveleaverequest,omitempty"`
}

func (x *GetLeaveApproveResponse) Reset() {
	*x = GetLeaveApproveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cosmos_lms_v1beta1_quer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLeaveApproveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLeaveApproveResponse) ProtoMessage() {}

func (x *GetLeaveApproveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cosmos_lms_v1beta1_quer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLeaveApproveResponse.ProtoReflect.Descriptor instead.
func (*GetLeaveApproveResponse) Descriptor() ([]byte, []int) {
	return file_proto_cosmos_lms_v1beta1_quer_proto_rawDescGZIP(), []int{5}
}

func (x *GetLeaveApproveResponse) GetApproveleaverequest() []*AcceptLeaveRequest {
	if x != nil {
		return x.Approveleaverequest
	}
	return nil
}

type GetRegisterAdminRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetRegisterAdminRequest) Reset() {
	*x = GetRegisterAdminRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cosmos_lms_v1beta1_quer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRegisterAdminRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRegisterAdminRequest) ProtoMessage() {}

func (x *GetRegisterAdminRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cosmos_lms_v1beta1_quer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRegisterAdminRequest.ProtoReflect.Descriptor instead.
func (*GetRegisterAdminRequest) Descriptor() ([]byte, []int) {
	return file_proto_cosmos_lms_v1beta1_quer_proto_rawDescGZIP(), []int{6}
}

type GetRegisterAdminResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegisterAdminrequest []*RegisterAdminRequest `protobuf:"bytes,1,rep,name=registerAdminrequest,proto3" json:"registerAdminrequest,omitempty"`
}

func (x *GetRegisterAdminResponse) Reset() {
	*x = GetRegisterAdminResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cosmos_lms_v1beta1_quer_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRegisterAdminResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRegisterAdminResponse) ProtoMessage() {}

func (x *GetRegisterAdminResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cosmos_lms_v1beta1_quer_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRegisterAdminResponse.ProtoReflect.Descriptor instead.
func (*GetRegisterAdminResponse) Descriptor() ([]byte, []int) {
	return file_proto_cosmos_lms_v1beta1_quer_proto_rawDescGZIP(), []int{7}
}

func (x *GetRegisterAdminResponse) GetRegisterAdminrequest() []*RegisterAdminRequest {
	if x != nil {
		return x.RegisterAdminrequest
	}
	return nil
}

var File_proto_cosmos_lms_v1beta1_quer_proto protoreflect.FileDescriptor

var file_proto_cosmos_lms_v1beta1_quer_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x6c,
	0x6d, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6c, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x1a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2f, 0x6c, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x74, 0x78, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x73, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x47, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6c, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x73, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x73, 0x22, 0x18, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x61, 0x76, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x5d,
	0x0a, 0x17, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0c, 0x6c, 0x65, 0x61,
	0x76, 0x65, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x6c, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x41, 0x70,
	0x70, 0x6c, 0x79, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x0c, 0x6c, 0x65, 0x61, 0x76, 0x65, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x18, 0x0a,
	0x16, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x6c, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4c, 0x65,
	0x61, 0x76, 0x65, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x51, 0x0a, 0x13, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x6c, 0x65, 0x61,
	0x76, 0x65, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x6c, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x41, 0x63,
	0x63, 0x65, 0x70, 0x74, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x13, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x6c, 0x65, 0x61, 0x76, 0x65, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x19, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x71, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x14,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6c, 0x6d, 0x73,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x14, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x32, 0xeb, 0x02, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x58, 0x0a,
	0x0d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x21,
	0x2e, 0x6c, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x22, 0x2e, 0x6c, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x53, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x2e, 0x6c, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6c, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x0a, 0x41, 0x70,
	0x70, 0x6c, 0x79, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x12, 0x23, 0x2e, 0x6c, 0x6d, 0x73, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e,
	0x6c, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c,
	0x65, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x4c,
	0x65, 0x61, 0x76, 0x65, 0x12, 0x23, 0x2e, 0x6c, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x41, 0x70, 0x70, 0x72, 0x6f,
	0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6c, 0x6d, 0x73, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x61, 0x76, 0x65,
	0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x78, 0x2f, 0x6c, 0x6d, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_cosmos_lms_v1beta1_quer_proto_rawDescOnce sync.Once
	file_proto_cosmos_lms_v1beta1_quer_proto_rawDescData = file_proto_cosmos_lms_v1beta1_quer_proto_rawDesc
)

func file_proto_cosmos_lms_v1beta1_quer_proto_rawDescGZIP() []byte {
	file_proto_cosmos_lms_v1beta1_quer_proto_rawDescOnce.Do(func() {
		file_proto_cosmos_lms_v1beta1_quer_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_cosmos_lms_v1beta1_quer_proto_rawDescData)
	})
	return file_proto_cosmos_lms_v1beta1_quer_proto_rawDescData
}

var file_proto_cosmos_lms_v1beta1_quer_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_cosmos_lms_v1beta1_quer_proto_goTypes = []interface{}{
	(*GetstudentsRequest)(nil),       // 0: lms.v1beta1.GetstudentsRequest
	(*GetstudentsResponse)(nil),      // 1: lms.v1beta1.GetstudentsResponse
	(*GetLeaveRequestRequest)(nil),   // 2: lms.v1beta1.GetLeaveRequestRequest
	(*GetLeaveRequestResponse)(nil),  // 3: lms.v1beta1.GetLeaveRequestResponse
	(*GetLeaveApproveRequest)(nil),   // 4: lms.v1beta1.GetLeaveApproveRequest
	(*GetLeaveApproveResponse)(nil),  // 5: lms.v1beta1.GetLeaveApproveResponse
	(*GetRegisterAdminRequest)(nil),  // 6: lms.v1beta1.GetRegisterAdminRequest
	(*GetRegisterAdminResponse)(nil), // 7: lms.v1beta1.GetRegisterAdminResponse
	(*Student)(nil),                  // 8: lms.v1beta1.Student
	(*ApplyLeaveRequest)(nil),        // 9: lms.v1beta1.ApplyLeaveRequest
	(*AcceptLeaveRequest)(nil),       // 10: lms.v1beta1.AcceptLeaveRequest
	(*RegisterAdminRequest)(nil),     // 11: lms.v1beta1.RegisterAdminRequest
	(*RegisterAdminResponse)(nil),    // 12: lms.v1beta1.RegisterAdminResponse
}
var file_proto_cosmos_lms_v1beta1_quer_proto_depIdxs = []int32{
	8,  // 0: lms.v1beta1.GetstudentsResponse.students:type_name -> lms.v1beta1.Student
	9,  // 1: lms.v1beta1.GetLeaveRequestResponse.leaverequest:type_name -> lms.v1beta1.ApplyLeaveRequest
	10, // 2: lms.v1beta1.GetLeaveApproveResponse.Approveleaverequest:type_name -> lms.v1beta1.AcceptLeaveRequest
	11, // 3: lms.v1beta1.GetRegisterAdminResponse.registerAdminrequest:type_name -> lms.v1beta1.RegisterAdminRequest
	11, // 4: lms.v1beta1.Query.RegisterAdmin:input_type -> lms.v1beta1.RegisterAdminRequest
	0,  // 5: lms.v1beta1.Query.AddStudent:input_type -> lms.v1beta1.GetstudentsRequest
	2,  // 6: lms.v1beta1.Query.ApplyLeave:input_type -> lms.v1beta1.GetLeaveRequestRequest
	4,  // 7: lms.v1beta1.Query.AcceptLeave:input_type -> lms.v1beta1.GetLeaveApproveRequest
	12, // 8: lms.v1beta1.Query.RegisterAdmin:output_type -> lms.v1beta1.RegisterAdminResponse
	1,  // 9: lms.v1beta1.Query.AddStudent:output_type -> lms.v1beta1.GetstudentsResponse
	3,  // 10: lms.v1beta1.Query.ApplyLeave:output_type -> lms.v1beta1.GetLeaveRequestResponse
	5,  // 11: lms.v1beta1.Query.AcceptLeave:output_type -> lms.v1beta1.GetLeaveApproveResponse
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_proto_cosmos_lms_v1beta1_quer_proto_init() }
func file_proto_cosmos_lms_v1beta1_quer_proto_init() {
	if File_proto_cosmos_lms_v1beta1_quer_proto != nil {
		return
	}
	file_proto_cosmos_lms_v1beta1_tx_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_cosmos_lms_v1beta1_quer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetstudentsRequest); i {
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
		file_proto_cosmos_lms_v1beta1_quer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetstudentsResponse); i {
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
		file_proto_cosmos_lms_v1beta1_quer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLeaveRequestRequest); i {
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
		file_proto_cosmos_lms_v1beta1_quer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLeaveRequestResponse); i {
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
		file_proto_cosmos_lms_v1beta1_quer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLeaveApproveRequest); i {
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
		file_proto_cosmos_lms_v1beta1_quer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLeaveApproveResponse); i {
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
		file_proto_cosmos_lms_v1beta1_quer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRegisterAdminRequest); i {
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
		file_proto_cosmos_lms_v1beta1_quer_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRegisterAdminResponse); i {
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
			RawDescriptor: file_proto_cosmos_lms_v1beta1_quer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_cosmos_lms_v1beta1_quer_proto_goTypes,
		DependencyIndexes: file_proto_cosmos_lms_v1beta1_quer_proto_depIdxs,
		MessageInfos:      file_proto_cosmos_lms_v1beta1_quer_proto_msgTypes,
	}.Build()
	File_proto_cosmos_lms_v1beta1_quer_proto = out.File
	file_proto_cosmos_lms_v1beta1_quer_proto_rawDesc = nil
	file_proto_cosmos_lms_v1beta1_quer_proto_goTypes = nil
	file_proto_cosmos_lms_v1beta1_quer_proto_depIdxs = nil
}
